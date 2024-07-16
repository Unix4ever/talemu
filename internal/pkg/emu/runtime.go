// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package emu

import (
	"context"

	"github.com/cosi-project/runtime/pkg/controller"
	"github.com/cosi-project/runtime/pkg/controller/runtime"
	"github.com/cosi-project/runtime/pkg/state"
	"go.uber.org/zap"

	"github.com/siderolabs/talemu/internal/pkg/emu/controllers"
	"github.com/siderolabs/talemu/internal/pkg/kubefactory"
)

// Runtime creates COSI runtime attached to the global state.
type Runtime struct {
	globalState state.State
	kubernetes  *kubefactory.Kubernetes
	runtime     *runtime.Runtime
	logger      *zap.Logger
}

// NewRuntime creates new runtime.
func NewRuntime(globalState state.State, kubernetes *kubefactory.Kubernetes, logger *zap.Logger) (*Runtime, error) {
	controllers := []controller.Controller{
		&controllers.ClusterCleanupController{
			Kubernetes: kubernetes,
		},
	}

	runtime, err := runtime.NewRuntime(globalState, logger)
	if err != nil {
		return nil, err
	}

	for _, ctrl := range controllers {
		if err = runtime.RegisterController(ctrl); err != nil {
			return nil, err
		}
	}

	return &Runtime{
		globalState: globalState,
		kubernetes:  kubernetes,
		runtime:     runtime,
		logger:      logger,
	}, nil
}

// Run starts the runtime.
func (rt *Runtime) Run(ctx context.Context) error {
	rt.logger.Info("starting global runtime")

	return rt.runtime.Run(ctx)
}
