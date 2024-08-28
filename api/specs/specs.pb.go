// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: specs/specs.proto

package specs

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClusterStatusSpec defines cluster status of the emulator.
type ClusterStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bootstrapped    bool     `protobuf:"varint,1,opt,name=bootstrapped,proto3" json:"bootstrapped,omitempty"`
	ControlPlanes   uint32   `protobuf:"varint,2,opt,name=control_planes,json=controlPlanes,proto3" json:"control_planes,omitempty"`
	Workers         uint32   `protobuf:"varint,3,opt,name=workers,proto3" json:"workers,omitempty"`
	DenyEtcdMembers []string `protobuf:"bytes,4,rep,name=deny_etcd_members,json=denyEtcdMembers,proto3" json:"deny_etcd_members,omitempty"`
	Kubeconfig      []byte   `protobuf:"bytes,5,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
}

func (x *ClusterStatusSpec) Reset() {
	*x = ClusterStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatusSpec) ProtoMessage() {}

func (x *ClusterStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStatusSpec.ProtoReflect.Descriptor instead.
func (*ClusterStatusSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterStatusSpec) GetBootstrapped() bool {
	if x != nil {
		return x.Bootstrapped
	}
	return false
}

func (x *ClusterStatusSpec) GetControlPlanes() uint32 {
	if x != nil {
		return x.ControlPlanes
	}
	return 0
}

func (x *ClusterStatusSpec) GetWorkers() uint32 {
	if x != nil {
		return x.Workers
	}
	return 0
}

func (x *ClusterStatusSpec) GetDenyEtcdMembers() []string {
	if x != nil {
		return x.DenyEtcdMembers
	}
	return nil
}

func (x *ClusterStatusSpec) GetKubeconfig() []byte {
	if x != nil {
		return x.Kubeconfig
	}
	return nil
}

// MachineStatusSpec is an emulated machine status.
type MachineStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses    []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	EtcdMemberId string   `protobuf:"bytes,2,opt,name=etcd_member_id,json=etcdMemberId,proto3" json:"etcd_member_id,omitempty"`
	Hostname     string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *MachineStatusSpec) Reset() {
	*x = MachineStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineStatusSpec) ProtoMessage() {}

func (x *MachineStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineStatusSpec.ProtoReflect.Descriptor instead.
func (*MachineStatusSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{1}
}

func (x *MachineStatusSpec) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *MachineStatusSpec) GetEtcdMemberId() string {
	if x != nil {
		return x.EtcdMemberId
	}
	return ""
}

func (x *MachineStatusSpec) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

// EventSinkStateSpec is defined per machine and resides in it's internal state
// describes which last version of a resource was reported to the events sink.
type EventSinkStateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions map[string]uint64 `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *EventSinkStateSpec) Reset() {
	*x = EventSinkStateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSinkStateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSinkStateSpec) ProtoMessage() {}

func (x *EventSinkStateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSinkStateSpec.ProtoReflect.Descriptor instead.
func (*EventSinkStateSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{2}
}

func (x *EventSinkStateSpec) GetVersions() map[string]uint64 {
	if x != nil {
		return x.Versions
	}
	return nil
}

// VersionSpec keeps current talos version of the emulated Talos node.
type VersionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Architecture string `protobuf:"bytes,2,opt,name=architecture,proto3" json:"architecture,omitempty"`
}

func (x *VersionSpec) Reset() {
	*x = VersionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionSpec) ProtoMessage() {}

func (x *VersionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionSpec.ProtoReflect.Descriptor instead.
func (*VersionSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{3}
}

func (x *VersionSpec) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *VersionSpec) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

// ImageSpec is the last image used in the upgrade request.
type ImageSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Schematic string `protobuf:"bytes,2,opt,name=schematic,proto3" json:"schematic,omitempty"`
}

func (x *ImageSpec) Reset() {
	*x = ImageSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageSpec) ProtoMessage() {}

func (x *ImageSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageSpec.ProtoReflect.Descriptor instead.
func (*ImageSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{4}
}

func (x *ImageSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ImageSpec) GetSchematic() string {
	if x != nil {
		return x.Schematic
	}
	return ""
}

// CachedImageSpec is the image pulled by the ImagePull API.
type CachedImageSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest string `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Size   int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *CachedImageSpec) Reset() {
	*x = CachedImageSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CachedImageSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CachedImageSpec) ProtoMessage() {}

func (x *CachedImageSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CachedImageSpec.ProtoReflect.Descriptor instead.
func (*CachedImageSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{5}
}

func (x *CachedImageSpec) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *CachedImageSpec) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// ServiceSpec is the fake service information.
type ServiceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State  string              `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Health *ServiceSpec_Health `protobuf:"bytes,3,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *ServiceSpec) Reset() {
	*x = ServiceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSpec) ProtoMessage() {}

func (x *ServiceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSpec.ProtoReflect.Descriptor instead.
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{6}
}

func (x *ServiceSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceSpec) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ServiceSpec) GetHealth() *ServiceSpec_Health {
	if x != nil {
		return x.Health
	}
	return nil
}

// RebootSpec keeps track of all reboots on the node.
type RebootSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Downtime *durationpb.Duration `protobuf:"bytes,1,opt,name=downtime,proto3" json:"downtime,omitempty"`
}

func (x *RebootSpec) Reset() {
	*x = RebootSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebootSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebootSpec) ProtoMessage() {}

func (x *RebootSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebootSpec.ProtoReflect.Descriptor instead.
func (*RebootSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{7}
}

func (x *RebootSpec) GetDowntime() *durationpb.Duration {
	if x != nil {
		return x.Downtime
	}
	return nil
}

// RebootStatusSpec is generated for each reboot spec.
type RebootStatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RebootStatusSpec) Reset() {
	*x = RebootStatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebootStatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebootStatusSpec) ProtoMessage() {}

func (x *RebootStatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebootStatusSpec.ProtoReflect.Descriptor instead.
func (*RebootStatusSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{8}
}

// MachineSpec is stored in Omni in the infra provisioner state.
type MachineSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot         int32  `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Uuid         string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Schematic    string `protobuf:"bytes,3,opt,name=schematic,proto3" json:"schematic,omitempty"`
	TalosVersion string `protobuf:"bytes,4,opt,name=talos_version,json=talosVersion,proto3" json:"talos_version,omitempty"`
}

func (x *MachineSpec) Reset() {
	*x = MachineSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineSpec) ProtoMessage() {}

func (x *MachineSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineSpec.ProtoReflect.Descriptor instead.
func (*MachineSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{9}
}

func (x *MachineSpec) GetSlot() int32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *MachineSpec) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MachineSpec) GetSchematic() string {
	if x != nil {
		return x.Schematic
	}
	return ""
}

func (x *MachineSpec) GetTalosVersion() string {
	if x != nil {
		return x.TalosVersion
	}
	return ""
}

// MachineTaskSpec is stored in the emulator state and the c.
type MachineTaskSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot           int32  `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Uuid           string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Schematic      string `protobuf:"bytes,3,opt,name=schematic,proto3" json:"schematic,omitempty"`
	TalosVersion   string `protobuf:"bytes,4,opt,name=talos_version,json=talosVersion,proto3" json:"talos_version,omitempty"`
	ConnectionArgs string `protobuf:"bytes,5,opt,name=connection_args,json=connectionArgs,proto3" json:"connection_args,omitempty"`
}

func (x *MachineTaskSpec) Reset() {
	*x = MachineTaskSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineTaskSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineTaskSpec) ProtoMessage() {}

func (x *MachineTaskSpec) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineTaskSpec.ProtoReflect.Descriptor instead.
func (*MachineTaskSpec) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{10}
}

func (x *MachineTaskSpec) GetSlot() int32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *MachineTaskSpec) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MachineTaskSpec) GetSchematic() string {
	if x != nil {
		return x.Schematic
	}
	return ""
}

func (x *MachineTaskSpec) GetTalosVersion() string {
	if x != nil {
		return x.TalosVersion
	}
	return ""
}

func (x *MachineTaskSpec) GetConnectionArgs() string {
	if x != nil {
		return x.ConnectionArgs
	}
	return ""
}

type ServiceSpec_Health struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unknown     bool                   `protobuf:"varint,1,opt,name=unknown,proto3" json:"unknown,omitempty"`
	Healthy     bool                   `protobuf:"varint,2,opt,name=healthy,proto3" json:"healthy,omitempty"`
	LastMessage string                 `protobuf:"bytes,3,opt,name=last_message,json=lastMessage,proto3" json:"last_message,omitempty"`
	LastChange  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_change,json=lastChange,proto3" json:"last_change,omitempty"`
}

func (x *ServiceSpec_Health) Reset() {
	*x = ServiceSpec_Health{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_specs_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSpec_Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSpec_Health) ProtoMessage() {}

func (x *ServiceSpec_Health) ProtoReflect() protoreflect.Message {
	mi := &file_specs_specs_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSpec_Health.ProtoReflect.Descriptor instead.
func (*ServiceSpec_Health) Descriptor() ([]byte, []int) {
	return file_specs_specs_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ServiceSpec_Health) GetUnknown() bool {
	if x != nil {
		return x.Unknown
	}
	return false
}

func (x *ServiceSpec_Health) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *ServiceSpec_Health) GetLastMessage() string {
	if x != nil {
		return x.LastMessage
	}
	return ""
}

func (x *ServiceSpec_Health) GetLastChange() *timestamppb.Timestamp {
	if x != nil {
		return x.LastChange
	}
	return nil
}

var File_specs_specs_proto protoreflect.FileDescriptor

var file_specs_specs_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6d, 0x75, 0x73, 0x70, 0x65, 0x63, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4,
	0x01, 0x0a, 0x11, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x6e,
	0x79, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x6e, 0x79, 0x45, 0x74, 0x63, 0x64, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x73, 0x0a, 0x11, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x74, 0x63, 0x64,
	0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x74, 0x63, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x12, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x46, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6d, 0x75, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22,
	0x43, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x22, 0x3d, 0x0a, 0x0f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6d, 0x75, 0x73,
	0x70, 0x65, 0x63, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a,
	0x9c, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x43,
	0x0a, 0x0a, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x35, 0x0a, 0x08,
	0x64, 0x6f, 0x77, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x70, 0x65, 0x63, 0x22, 0x78, 0x0a, 0x0b, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x61, 0x6c, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x61, 0x6c, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6d, 0x75, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_specs_specs_proto_rawDescOnce sync.Once
	file_specs_specs_proto_rawDescData = file_specs_specs_proto_rawDesc
)

func file_specs_specs_proto_rawDescGZIP() []byte {
	file_specs_specs_proto_rawDescOnce.Do(func() {
		file_specs_specs_proto_rawDescData = protoimpl.X.CompressGZIP(file_specs_specs_proto_rawDescData)
	})
	return file_specs_specs_proto_rawDescData
}

var file_specs_specs_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_specs_specs_proto_goTypes = []any{
	(*ClusterStatusSpec)(nil),     // 0: emuspecs.ClusterStatusSpec
	(*MachineStatusSpec)(nil),     // 1: emuspecs.MachineStatusSpec
	(*EventSinkStateSpec)(nil),    // 2: emuspecs.EventSinkStateSpec
	(*VersionSpec)(nil),           // 3: emuspecs.VersionSpec
	(*ImageSpec)(nil),             // 4: emuspecs.ImageSpec
	(*CachedImageSpec)(nil),       // 5: emuspecs.CachedImageSpec
	(*ServiceSpec)(nil),           // 6: emuspecs.ServiceSpec
	(*RebootSpec)(nil),            // 7: emuspecs.RebootSpec
	(*RebootStatusSpec)(nil),      // 8: emuspecs.RebootStatusSpec
	(*MachineSpec)(nil),           // 9: emuspecs.MachineSpec
	(*MachineTaskSpec)(nil),       // 10: emuspecs.MachineTaskSpec
	nil,                           // 11: emuspecs.EventSinkStateSpec.VersionsEntry
	(*ServiceSpec_Health)(nil),    // 12: emuspecs.ServiceSpec.Health
	(*durationpb.Duration)(nil),   // 13: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
}
var file_specs_specs_proto_depIdxs = []int32{
	11, // 0: emuspecs.EventSinkStateSpec.versions:type_name -> emuspecs.EventSinkStateSpec.VersionsEntry
	12, // 1: emuspecs.ServiceSpec.health:type_name -> emuspecs.ServiceSpec.Health
	13, // 2: emuspecs.RebootSpec.downtime:type_name -> google.protobuf.Duration
	14, // 3: emuspecs.ServiceSpec.Health.last_change:type_name -> google.protobuf.Timestamp
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_specs_specs_proto_init() }
func file_specs_specs_proto_init() {
	if File_specs_specs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_specs_specs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterStatusSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MachineStatusSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EventSinkStateSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VersionSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ImageSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CachedImageSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RebootSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RebootStatusSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*MachineSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*MachineTaskSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_specs_specs_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceSpec_Health); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_specs_specs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_specs_specs_proto_goTypes,
		DependencyIndexes: file_specs_specs_proto_depIdxs,
		MessageInfos:      file_specs_specs_proto_msgTypes,
	}.Build()
	File_specs_specs_proto = out.File
	file_specs_specs_proto_rawDesc = nil
	file_specs_specs_proto_goTypes = nil
	file_specs_specs_proto_depIdxs = nil
}
