// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: chat/server_v2_rpc/channel.proto

package server_v2_rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Behavior of moderation.
type ChannelModerationBehavior int32

const (
	// Behavior is unspecified.
	ChannelModerationBehavior_CHANNEL_MODERATION_BEHAVIOR_UNSPECIFIED ChannelModerationBehavior = 0
	// Flag a message.
	ChannelModerationBehavior_CHANNEL_MODERATION_BEHAVIOR_FLAG ChannelModerationBehavior = 1
	// Block a message.
	ChannelModerationBehavior_CHANNEL_MODERATION_BEHAVIOR_BLOCK ChannelModerationBehavior = 2
)

// Enum value maps for ChannelModerationBehavior.
var (
	ChannelModerationBehavior_name = map[int32]string{
		0: "CHANNEL_MODERATION_BEHAVIOR_UNSPECIFIED",
		1: "CHANNEL_MODERATION_BEHAVIOR_FLAG",
		2: "CHANNEL_MODERATION_BEHAVIOR_BLOCK",
	}
	ChannelModerationBehavior_value = map[string]int32{
		"CHANNEL_MODERATION_BEHAVIOR_UNSPECIFIED": 0,
		"CHANNEL_MODERATION_BEHAVIOR_FLAG":        1,
		"CHANNEL_MODERATION_BEHAVIOR_BLOCK":       2,
	}
)

func (x ChannelModerationBehavior) Enum() *ChannelModerationBehavior {
	p := new(ChannelModerationBehavior)
	*p = x
	return p
}

func (x ChannelModerationBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelModerationBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_server_v2_rpc_channel_proto_enumTypes[0].Descriptor()
}

func (ChannelModerationBehavior) Type() protoreflect.EnumType {
	return &file_chat_server_v2_rpc_channel_proto_enumTypes[0]
}

func (x ChannelModerationBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelModerationBehavior.Descriptor instead.
func (ChannelModerationBehavior) EnumDescriptor() ([]byte, []int) {
	return file_chat_server_v2_rpc_channel_proto_rawDescGZIP(), []int{0}
}

// Mode of automatic moderation.
type ChannelAutomodMode int32

const (
	// Mode is unspecified.
	ChannelAutomodMode_CHANNEL_AUTOMOD_MODE_UNSPECIFIED ChannelAutomodMode = 0
	// Moderation is disabled.
	ChannelAutomodMode_CHANNEL_AUTOMOD_MODE_DISABLED ChannelAutomodMode = 1
	// Simple moderation.
	ChannelAutomodMode_CHANNEL_AUTOMOD_MODE_SIMPLE ChannelAutomodMode = 2
	// AI moderation.
	ChannelAutomodMode_CHANNEL_AUTOMOD_MODE_AI ChannelAutomodMode = 3
)

// Enum value maps for ChannelAutomodMode.
var (
	ChannelAutomodMode_name = map[int32]string{
		0: "CHANNEL_AUTOMOD_MODE_UNSPECIFIED",
		1: "CHANNEL_AUTOMOD_MODE_DISABLED",
		2: "CHANNEL_AUTOMOD_MODE_SIMPLE",
		3: "CHANNEL_AUTOMOD_MODE_AI",
	}
	ChannelAutomodMode_value = map[string]int32{
		"CHANNEL_AUTOMOD_MODE_UNSPECIFIED": 0,
		"CHANNEL_AUTOMOD_MODE_DISABLED":    1,
		"CHANNEL_AUTOMOD_MODE_SIMPLE":      2,
		"CHANNEL_AUTOMOD_MODE_AI":          3,
	}
)

func (x ChannelAutomodMode) Enum() *ChannelAutomodMode {
	p := new(ChannelAutomodMode)
	*p = x
	return p
}

func (x ChannelAutomodMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelAutomodMode) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_server_v2_rpc_channel_proto_enumTypes[1].Descriptor()
}

func (ChannelAutomodMode) Type() protoreflect.EnumType {
	return &file_chat_server_v2_rpc_channel_proto_enumTypes[1]
}

func (x ChannelAutomodMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelAutomodMode.Descriptor instead.
func (ChannelAutomodMode) EnumDescriptor() ([]byte, []int) {
	return file_chat_server_v2_rpc_channel_proto_rawDescGZIP(), []int{1}
}

type ChannelConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether typing events are enabled.
	TypingEvents bool `protobuf:"varint,1,opt,name=typing_events,json=typingEvents,proto3" json:"typing_events,omitempty"`
	// Whether read events are enabled.
	ReadEvents bool `protobuf:"varint,2,opt,name=read_events,json=readEvents,proto3" json:"read_events,omitempty"`
	// Whether connect events are enabled.
	ConnectEvents bool `protobuf:"varint,3,opt,name=connect_events,json=connectEvents,proto3" json:"connect_events,omitempty"`
	// Whether custom events are enabled.
	CustomEvents bool `protobuf:"varint,4,opt,name=custom_events,json=customEvents,proto3" json:"custom_events,omitempty"`
	// Whether search feature is enabled.
	Search bool `protobuf:"varint,5,opt,name=search,proto3" json:"search,omitempty"`
	// Whether reactions are enabled.
	Reactions bool `protobuf:"varint,6,opt,name=reactions,proto3" json:"reactions,omitempty"`
	// Whether replies (threads) are enabled.
	Replies bool `protobuf:"varint,7,opt,name=replies,proto3" json:"replies,omitempty"`
	// Whether message quoting feature is enabled.
	Quotes bool `protobuf:"varint,8,opt,name=quotes,proto3" json:"quotes,omitempty"`
	// Whether mute feature is enabled.
	Mutes bool `protobuf:"varint,9,opt,name=mutes,proto3" json:"mutes,omitempty"`
	// Whether file uploads are enabled.
	Uploads bool `protobuf:"varint,10,opt,name=uploads,proto3" json:"uploads,omitempty"`
	// Whether URL enrichment feature is enabled.
	UrlEnrichment bool `protobuf:"varint,11,opt,name=url_enrichment,json=urlEnrichment,proto3" json:"url_enrichment,omitempty"`
	// Whether push notifications are enabled.
	PushNotifications bool `protobuf:"varint,12,opt,name=push_notifications,json=pushNotifications,proto3" json:"push_notifications,omitempty"`
	// Whether reminders feature is enabled.
	Reminders bool `protobuf:"varint,13,opt,name=reminders,proto3" json:"reminders,omitempty"`
	// Automatic moderation mode.
	Automod ChannelAutomodMode `protobuf:"varint,14,opt,name=automod,proto3,enum=stream.chat.server_v2_rpc.ChannelAutomodMode" json:"automod,omitempty"`
	// Automatic moderation behavior.
	AutomodBehavior ChannelModerationBehavior `protobuf:"varint,15,opt,name=automod_behavior,json=automodBehavior,proto3,enum=stream.chat.server_v2_rpc.ChannelModerationBehavior" json:"automod_behavior,omitempty"`
	// Automatic moderation thresholds.
	AutomodThresholds *ChannelModerationThresholds `protobuf:"bytes,16,opt,name=automod_thresholds,json=automodThresholds,proto3" json:"automod_thresholds,omitempty"`
	// Blocklist ID that is active in this channel.
	Blocklist string `protobuf:"bytes,17,opt,name=blocklist,proto3" json:"blocklist,omitempty"`
	// Behavior of the blocklist.
	BlocklistBehavior ChannelModerationBehavior `protobuf:"varint,18,opt,name=blocklist_behavior,json=blocklistBehavior,proto3,enum=stream.chat.server_v2_rpc.ChannelModerationBehavior" json:"blocklist_behavior,omitempty"`
}

func (x *ChannelConfig) Reset() {
	*x = ChannelConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_server_v2_rpc_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelConfig) ProtoMessage() {}

func (x *ChannelConfig) ProtoReflect() protoreflect.Message {
	mi := &file_chat_server_v2_rpc_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelConfig.ProtoReflect.Descriptor instead.
func (*ChannelConfig) Descriptor() ([]byte, []int) {
	return file_chat_server_v2_rpc_channel_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelConfig) GetTypingEvents() bool {
	if x != nil {
		return x.TypingEvents
	}
	return false
}

func (x *ChannelConfig) GetReadEvents() bool {
	if x != nil {
		return x.ReadEvents
	}
	return false
}

func (x *ChannelConfig) GetConnectEvents() bool {
	if x != nil {
		return x.ConnectEvents
	}
	return false
}

func (x *ChannelConfig) GetCustomEvents() bool {
	if x != nil {
		return x.CustomEvents
	}
	return false
}

func (x *ChannelConfig) GetSearch() bool {
	if x != nil {
		return x.Search
	}
	return false
}

func (x *ChannelConfig) GetReactions() bool {
	if x != nil {
		return x.Reactions
	}
	return false
}

func (x *ChannelConfig) GetReplies() bool {
	if x != nil {
		return x.Replies
	}
	return false
}

func (x *ChannelConfig) GetQuotes() bool {
	if x != nil {
		return x.Quotes
	}
	return false
}

func (x *ChannelConfig) GetMutes() bool {
	if x != nil {
		return x.Mutes
	}
	return false
}

func (x *ChannelConfig) GetUploads() bool {
	if x != nil {
		return x.Uploads
	}
	return false
}

func (x *ChannelConfig) GetUrlEnrichment() bool {
	if x != nil {
		return x.UrlEnrichment
	}
	return false
}

func (x *ChannelConfig) GetPushNotifications() bool {
	if x != nil {
		return x.PushNotifications
	}
	return false
}

func (x *ChannelConfig) GetReminders() bool {
	if x != nil {
		return x.Reminders
	}
	return false
}

func (x *ChannelConfig) GetAutomod() ChannelAutomodMode {
	if x != nil {
		return x.Automod
	}
	return ChannelAutomodMode_CHANNEL_AUTOMOD_MODE_UNSPECIFIED
}

func (x *ChannelConfig) GetAutomodBehavior() ChannelModerationBehavior {
	if x != nil {
		return x.AutomodBehavior
	}
	return ChannelModerationBehavior_CHANNEL_MODERATION_BEHAVIOR_UNSPECIFIED
}

func (x *ChannelConfig) GetAutomodThresholds() *ChannelModerationThresholds {
	if x != nil {
		return x.AutomodThresholds
	}
	return nil
}

func (x *ChannelConfig) GetBlocklist() string {
	if x != nil {
		return x.Blocklist
	}
	return ""
}

func (x *ChannelConfig) GetBlocklistBehavior() ChannelModerationBehavior {
	if x != nil {
		return x.BlocklistBehavior
	}
	return ChannelModerationBehavior_CHANNEL_MODERATION_BEHAVIOR_UNSPECIFIED
}

// Thresholds for automatic moderation.
type ChannelModerationThresholds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Thresholds for explicit content.
	Explicit *ChannelModerationThreshold `protobuf:"bytes,1,opt,name=explicit,proto3" json:"explicit,omitempty"`
	// Thresholds for spam.
	Spam *ChannelModerationThreshold `protobuf:"bytes,2,opt,name=spam,proto3" json:"spam,omitempty"`
	// Thresholds for toxic content.
	Toxic *ChannelModerationThreshold `protobuf:"bytes,3,opt,name=toxic,proto3" json:"toxic,omitempty"`
}

func (x *ChannelModerationThresholds) Reset() {
	*x = ChannelModerationThresholds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_server_v2_rpc_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelModerationThresholds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelModerationThresholds) ProtoMessage() {}

func (x *ChannelModerationThresholds) ProtoReflect() protoreflect.Message {
	mi := &file_chat_server_v2_rpc_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelModerationThresholds.ProtoReflect.Descriptor instead.
func (*ChannelModerationThresholds) Descriptor() ([]byte, []int) {
	return file_chat_server_v2_rpc_channel_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelModerationThresholds) GetExplicit() *ChannelModerationThreshold {
	if x != nil {
		return x.Explicit
	}
	return nil
}

func (x *ChannelModerationThresholds) GetSpam() *ChannelModerationThreshold {
	if x != nil {
		return x.Spam
	}
	return nil
}

func (x *ChannelModerationThresholds) GetToxic() *ChannelModerationThreshold {
	if x != nil {
		return x.Toxic
	}
	return nil
}

// Behavior thresholds for automatic moderaion.
type ChannelModerationThreshold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Threshold for flagging a message.
	Flag float64 `protobuf:"fixed64,1,opt,name=flag,proto3" json:"flag,omitempty"`
	// Threshold for blocking a message.
	Block float64 `protobuf:"fixed64,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *ChannelModerationThreshold) Reset() {
	*x = ChannelModerationThreshold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_server_v2_rpc_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelModerationThreshold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelModerationThreshold) ProtoMessage() {}

func (x *ChannelModerationThreshold) ProtoReflect() protoreflect.Message {
	mi := &file_chat_server_v2_rpc_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelModerationThreshold.ProtoReflect.Descriptor instead.
func (*ChannelModerationThreshold) Descriptor() ([]byte, []int) {
	return file_chat_server_v2_rpc_channel_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelModerationThreshold) GetFlag() float64 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *ChannelModerationThreshold) GetBlock() float64 {
	if x != nil {
		return x.Block
	}
	return 0
}

var File_chat_server_v2_rpc_channel_proto protoreflect.FileDescriptor

var file_chat_server_v2_rpc_channel_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32,
	0x5f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x22, 0xc1, 0x06,
	0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x72, 0x6c,
	0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x75, 0x72, 0x6c, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x2d, 0x0a, 0x12, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x75,
	0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a,
	0x07, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6f, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6f, 0x64, 0x12, 0x5f, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6f,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x34, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6f, 0x64, 0x42,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x65, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6f, 0x64, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x11, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6f, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x12,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x11,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x22, 0x88, 0x02, 0x0a, 0x1b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x73, 0x12, 0x51, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x12, 0x49, 0x0a, 0x04, 0x73, 0x70, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x04, 0x73, 0x70, 0x61, 0x6d, 0x12,
	0x4b, 0x0a, 0x05, 0x74, 0x6f, 0x78, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x05, 0x74, 0x6f, 0x78, 0x69, 0x63, 0x22, 0x46, 0x0a, 0x1a,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2a, 0x95, 0x01, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x27, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f,
	0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x24, 0x0a, 0x20, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f, 0x52, 0x5f, 0x46,
	0x4c, 0x41, 0x47, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x45, 0x48, 0x41,
	0x56, 0x49, 0x4f, 0x52, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x02, 0x2a, 0x9b, 0x01, 0x0a,
	0x12, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x6f, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x41,
	0x55, 0x54, 0x4f, 0x4d, 0x4f, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x4f, 0x44, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x4f, 0x44, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x1b, 0x0a,
	0x17, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x4f, 0x44,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x49, 0x10, 0x03, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_server_v2_rpc_channel_proto_rawDescOnce sync.Once
	file_chat_server_v2_rpc_channel_proto_rawDescData = file_chat_server_v2_rpc_channel_proto_rawDesc
)

func file_chat_server_v2_rpc_channel_proto_rawDescGZIP() []byte {
	file_chat_server_v2_rpc_channel_proto_rawDescOnce.Do(func() {
		file_chat_server_v2_rpc_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_server_v2_rpc_channel_proto_rawDescData)
	})
	return file_chat_server_v2_rpc_channel_proto_rawDescData
}

var file_chat_server_v2_rpc_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_chat_server_v2_rpc_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_server_v2_rpc_channel_proto_goTypes = []interface{}{
	(ChannelModerationBehavior)(0),      // 0: stream.chat.server_v2_rpc.ChannelModerationBehavior
	(ChannelAutomodMode)(0),             // 1: stream.chat.server_v2_rpc.ChannelAutomodMode
	(*ChannelConfig)(nil),               // 2: stream.chat.server_v2_rpc.ChannelConfig
	(*ChannelModerationThresholds)(nil), // 3: stream.chat.server_v2_rpc.ChannelModerationThresholds
	(*ChannelModerationThreshold)(nil),  // 4: stream.chat.server_v2_rpc.ChannelModerationThreshold
}
var file_chat_server_v2_rpc_channel_proto_depIdxs = []int32{
	1, // 0: stream.chat.server_v2_rpc.ChannelConfig.automod:type_name -> stream.chat.server_v2_rpc.ChannelAutomodMode
	0, // 1: stream.chat.server_v2_rpc.ChannelConfig.automod_behavior:type_name -> stream.chat.server_v2_rpc.ChannelModerationBehavior
	3, // 2: stream.chat.server_v2_rpc.ChannelConfig.automod_thresholds:type_name -> stream.chat.server_v2_rpc.ChannelModerationThresholds
	0, // 3: stream.chat.server_v2_rpc.ChannelConfig.blocklist_behavior:type_name -> stream.chat.server_v2_rpc.ChannelModerationBehavior
	4, // 4: stream.chat.server_v2_rpc.ChannelModerationThresholds.explicit:type_name -> stream.chat.server_v2_rpc.ChannelModerationThreshold
	4, // 5: stream.chat.server_v2_rpc.ChannelModerationThresholds.spam:type_name -> stream.chat.server_v2_rpc.ChannelModerationThreshold
	4, // 6: stream.chat.server_v2_rpc.ChannelModerationThresholds.toxic:type_name -> stream.chat.server_v2_rpc.ChannelModerationThreshold
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_chat_server_v2_rpc_channel_proto_init() }
func file_chat_server_v2_rpc_channel_proto_init() {
	if File_chat_server_v2_rpc_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_server_v2_rpc_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelConfig); i {
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
		file_chat_server_v2_rpc_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelModerationThresholds); i {
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
		file_chat_server_v2_rpc_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelModerationThreshold); i {
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
			RawDescriptor: file_chat_server_v2_rpc_channel_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_server_v2_rpc_channel_proto_goTypes,
		DependencyIndexes: file_chat_server_v2_rpc_channel_proto_depIdxs,
		EnumInfos:         file_chat_server_v2_rpc_channel_proto_enumTypes,
		MessageInfos:      file_chat_server_v2_rpc_channel_proto_msgTypes,
	}.Build()
	File_chat_server_v2_rpc_channel_proto = out.File
	file_chat_server_v2_rpc_channel_proto_rawDesc = nil
	file_chat_server_v2_rpc_channel_proto_goTypes = nil
	file_chat_server_v2_rpc_channel_proto_depIdxs = nil
}
