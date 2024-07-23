// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package runtime starts COSI state and runtime.
package runtime

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cosi-project/runtime/pkg/controller"
	"github.com/cosi-project/runtime/pkg/controller/runtime"
	"github.com/cosi-project/runtime/pkg/state"
	"go.uber.org/zap"

	"github.com/siderolabs/talemu/internal/pkg/kubefactory"
	"github.com/siderolabs/talemu/internal/pkg/machine/controllers"
	"github.com/siderolabs/talemu/internal/pkg/machine/logging"
	"github.com/siderolabs/talemu/internal/pkg/machine/runtime/resources/emu"
	"github.com/siderolabs/talemu/internal/pkg/machine/runtime/resources/talos"
	"github.com/siderolabs/talemu/internal/pkg/machine/services"
)

// Runtime handles COSI state setup and lifecycle.
type Runtime struct {
	state        state.State
	globalState  state.State
	runtime      *runtime.Runtime
	backingStore io.Closer
	id           string
}

// NewRuntime creates new runtime.
func NewRuntime(ctx context.Context, logger *zap.Logger, machineIndex int, id string, globalState state.State,
	kubernetes *kubefactory.Kubernetes, logSink *logging.ZapCore,
) (*Runtime, error) {
	stateDir := filepath.Join("_out/state/machines", strconv.FormatInt(int64(machineIndex), 10))

	err := os.MkdirAll(stateDir, 0o664)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return nil, err
	}

	st, backingStore, err := NewState(filepath.Join(stateDir, "state.db"), logger)
	if err != nil {
		return nil, err
	}

	if err = talos.Register(ctx, st); err != nil {
		return nil, err
	}

	machineStatus := emu.NewMachineStatus(emu.NamespaceName, id)
	if err = globalState.Create(ctx, machineStatus); err != nil && !state.IsConflictError(err) {
		return nil, err
	}

	qcontrollers := []controller.QController{
		controllers.NewRebootStatusController(),
	}

	controllers := []controller.Controller{
		&controllers.ManagerController{
			MachineIndex: machineIndex,
		},
		&controllers.LinkSpecController{},
		&controllers.LinkStatusController{},
		&controllers.APIDController{
			APID: services.NewAPID(id, st, globalState),
		},
		&controllers.AddressSpecController{},
		&controllers.GRPCTLSController{},
		&controllers.MachineTypeController{},
		&controllers.HostnameConfigController{
			MachineID: id,
		},
		&controllers.HostnameMergeController{},
		&controllers.HostnameSpecController{
			GlobalState: globalState,
			MachineID:   id,
		},
		&controllers.NodeAddressController{
			GlobalState: globalState,
			MachineID:   id,
		},
		&controllers.APICertSANsController{},
		controllers.NewRootOSController(),
		&controllers.ExtensionStatusController{},
		&controllers.MachineStatusController{State: st},
		&controllers.VersionController{},
		&controllers.NodeIdentityController{},
		&controllers.NodenameController{},
		&controllers.EtcdController{
			GlobalState: globalState,
			MachineID:   id,
		},
		&controllers.MountStatusController{},
		&controllers.LocalAffiliateController{},
		&controllers.MemberController{},
		controllers.NewClusterConfigController(),
		&controllers.AffiliateMergeController{},
		&controllers.DiscoveryServiceController{},
		&controllers.KubernetesSecretsController{},
		&controllers.KubernetesDynamicCertsController{},
		&controllers.KubernetesController{
			Kubernetes: kubernetes,
			MachineID:  id,
		},
		controllers.NewRootKubernetesController(),
		&controllers.KubernetesCertSANsController{},
		&controllers.RenderSecretsStaticPodController{
			MachineID: id,
		},
		&controllers.KubernetesNodeController{
			MachineID:   id,
			GlobalState: globalState,
		},
		&controllers.KubeconfigController{
			GlobalState: globalState,
		},
		&controllers.StaticPodController{
			MachineID: id,
		},
		&controllers.LogSinkController{
			LogSink: logSink,
		},
	}

	runtime, err := runtime.NewRuntime(st, logger)
	if err != nil {
		return nil, err
	}

	for _, ctrl := range controllers {
		if err = runtime.RegisterController(ctrl); err != nil {
			return nil, err
		}
	}

	for _, ctrl := range qcontrollers {
		if err = runtime.RegisterQController(ctrl); err != nil {
			return nil, err
		}
	}

	return &Runtime{
		state:        st,
		globalState:  globalState,
		runtime:      runtime,
		backingStore: backingStore,
		id:           id,
	}, nil
}

// Run starts COSI runtime.
func (r *Runtime) Run(ctx context.Context) error {
	defer r.backingStore.Close() //nolint:errcheck

	if err := r.runtime.Run(ctx); err != nil {
		return err
	}

	return nil
}

// State returns COSI state.
func (r *Runtime) State() state.State {
	return r.state
}
