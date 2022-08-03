# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/error_v2/error.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19\x63hat/error_v2/error.proto\x12\x14stream.chat.error_v2*\xbb\x08\n\tErrorCode\x12\x1a\n\x16\x45RROR_CODE_UNSPECIFIED\x10\x00\x12\x19\n\x15\x45RROR_CODE_VALIDATION\x10\x01\x12 \n\x13\x45RROR_CODE_INTERNAL\x10\xff\xff\xff\xff\xff\xff\xff\xff\xff\x01\x12\x19\n\x15\x45RROR_CODE_ACCESS_KEY\x10\x02\x12\x13\n\x0f\x45RROR_CODE_AUTH\x10\x05\x12!\n\x1d\x45RROR_CODE_DUPLICATE_USERNAME\x10\x06\x12\x1a\n\x16\x45RROR_CODE_RATELIMITED\x10\t\x12\x17\n\x13\x45RROR_CODE_NOTFOUND\x10\x10\x12\x1a\n\x16\x45RROR_CODE_NOT_ALLOWED\x10\x11\x12\"\n\x1e\x45RROR_CODE_EVENT_NOT_SUPPORTED\x10\x12\x12,\n(ERROR_CODE_CHANNEL_FEATURE_NOT_SUPPORTED\x10\x13\x12\x1f\n\x1b\x45RROR_CODE_MESSAGE_TOO_LONG\x10\x14\x12%\n!ERROR_CODE_MULTIPLE_NESTING_LEVEL\x10\x15\x12\x1e\n\x1a\x45RROR_CODE_PAYLOAD_TOO_BIG\x10\x16\x12\x1c\n\x18\x45RROR_CODE_TOKEN_EXPIRED\x10(\x12 \n\x1c\x45RROR_CODE_TOKEN_INVALID_YET\x10)\x12$\n ERROR_CODE_TOKEN_USED_BEFORE_IAT\x10*\x12&\n\"ERROR_CODE_TOKEN_INVALID_SIGNATURE\x10+\x12.\n*ERROR_CODE_CUSTOM_COMMAND_ENDPOINT_MISSING\x10,\x12\x31\n-ERROR_CODE_CUSTOM_COMMAND_ENDPOINT_CALL_ERROR\x10-\x12&\n\"ERROR_CODE_CONNECTION_ID_NOT_FOUND\x10.\x12&\n\"ERROR_CODE_COOLDOWN_PERIOD_NOT_MET\x10<\x12\x31\n-ERROR_CODE_QUERY_CHANNEL_PERMISSIONS_MISMATCH\x10\x46\x12#\n\x1f\x45RROR_CODE_TOO_MANY_CONNECTIONS\x10G\x12\'\n#ERROR_CODE_NOT_SUPPORTED_IN_PUSH_V1\x10H\x12 \n\x1c\x45RROR_CODE_MODERATION_FAILED\x10I\x12,\n(ERROR_CODE_VIDEO_PROVIDER_NOT_CONFIGURED\x10P\x12\x1e\n\x1a\x45RROR_CODE_INVALID_CALL_ID\x10Q\x12\'\n#ERROR_CODE_VIDEO_CREATE_CALL_FAILED\x10R\x12\x1c\n\x18\x45RROR_CODE_APP_SUSPENDED\x10\x63*`\n\x13ValidationErrorCode\x12%\n!VALIDATION_ERROR_CODE_UNSPECIFIED\x10\x00\x12\"\n\x1eVALIDATION_ERROR_CODE_REQUIRED\x10\x01\x42\x46ZDgithub.com/GetStream/stream-services/protobuf/chat/error_v2;error_v2b\x06proto3')

_ERRORCODE = DESCRIPTOR.enum_types_by_name['ErrorCode']
ErrorCode = enum_type_wrapper.EnumTypeWrapper(_ERRORCODE)
_VALIDATIONERRORCODE = DESCRIPTOR.enum_types_by_name['ValidationErrorCode']
ValidationErrorCode = enum_type_wrapper.EnumTypeWrapper(_VALIDATIONERRORCODE)
ERROR_CODE_UNSPECIFIED = 0
ERROR_CODE_VALIDATION = 1
ERROR_CODE_INTERNAL = -1
ERROR_CODE_ACCESS_KEY = 2
ERROR_CODE_AUTH = 5
ERROR_CODE_DUPLICATE_USERNAME = 6
ERROR_CODE_RATELIMITED = 9
ERROR_CODE_NOTFOUND = 16
ERROR_CODE_NOT_ALLOWED = 17
ERROR_CODE_EVENT_NOT_SUPPORTED = 18
ERROR_CODE_CHANNEL_FEATURE_NOT_SUPPORTED = 19
ERROR_CODE_MESSAGE_TOO_LONG = 20
ERROR_CODE_MULTIPLE_NESTING_LEVEL = 21
ERROR_CODE_PAYLOAD_TOO_BIG = 22
ERROR_CODE_TOKEN_EXPIRED = 40
ERROR_CODE_TOKEN_INVALID_YET = 41
ERROR_CODE_TOKEN_USED_BEFORE_IAT = 42
ERROR_CODE_TOKEN_INVALID_SIGNATURE = 43
ERROR_CODE_CUSTOM_COMMAND_ENDPOINT_MISSING = 44
ERROR_CODE_CUSTOM_COMMAND_ENDPOINT_CALL_ERROR = 45
ERROR_CODE_CONNECTION_ID_NOT_FOUND = 46
ERROR_CODE_COOLDOWN_PERIOD_NOT_MET = 60
ERROR_CODE_QUERY_CHANNEL_PERMISSIONS_MISMATCH = 70
ERROR_CODE_TOO_MANY_CONNECTIONS = 71
ERROR_CODE_NOT_SUPPORTED_IN_PUSH_V1 = 72
ERROR_CODE_MODERATION_FAILED = 73
ERROR_CODE_VIDEO_PROVIDER_NOT_CONFIGURED = 80
ERROR_CODE_INVALID_CALL_ID = 81
ERROR_CODE_VIDEO_CREATE_CALL_FAILED = 82
ERROR_CODE_APP_SUSPENDED = 99
VALIDATION_ERROR_CODE_UNSPECIFIED = 0
VALIDATION_ERROR_CODE_REQUIRED = 1


if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZDgithub.com/GetStream/stream-services/protobuf/chat/error_v2;error_v2'
  _ERRORCODE._serialized_start=52
  _ERRORCODE._serialized_end=1135
  _VALIDATIONERRORCODE._serialized_start=1137
  _VALIDATIONERRORCODE._serialized_end=1233
# @@protoc_insertion_point(module_scope)
