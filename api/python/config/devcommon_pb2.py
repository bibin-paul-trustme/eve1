# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: devcommon.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import devmodel_pb2 as devmodel__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='devcommon.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\037com.zededa.cloud.uservice.protoZ$github.com/lf-edge/eve/api/go/config'),
  serialized_pb=_b('\n\x0f\x64\x65vcommon.proto\x1a\x0e\x64\x65vmodel.proto\"/\n\x0eUUIDandVersion\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\"F\n\x0c\x44\x65viceOpsCmd\x12\x0f\n\x07\x63ounter\x18\x02 \x01(\r\x12\x14\n\x0c\x64\x65siredState\x18\x03 \x01(\x08\x12\x0f\n\x07opsTime\x18\x04 \x01(\t\"(\n\nConfigItem\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"1\n\x07\x41\x64\x61pter\x12\x18\n\x04type\x18\x01 \x01(\x0e\x32\n.PhyIoType\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xc1\x01\n\nCipherInfo\x12\n\n\x02id\x18\x64 \x01(\t\x12-\n\x11keyExchangeScheme\x18\x01 \x01(\x0e\x32\x12.KeyExchangeScheme\x12+\n\x10\x65ncryptionScheme\x18\x02 \x01(\x0e\x32\x11.EncryptionScheme\x12\x14\n\x0cinitialValue\x18\x03 \x01(\x0c\x12\x12\n\npublicCert\x18\x04 \x01(\x0c\x12\x0e\n\x06sha256\x18\x05 \x01(\t\x12\x11\n\tsignature\x18\x06 \x01(\x0c*/\n\x11KeyExchangeScheme\x12\x0c\n\x08KEA_NONE\x10\x00\x12\x0c\n\x08KEA_ECDH\x10\x01*3\n\x10\x45ncryptionScheme\x12\x0b\n\x07SA_NONE\x10\x00\x12\x12\n\x0eSA_AES_256_CFB\x10\x01\x42G\n\x1f\x63om.zededa.cloud.uservice.protoZ$github.com/lf-edge/eve/api/go/configb\x06proto3')
  ,
  dependencies=[devmodel__pb2.DESCRIPTOR,])

_KEYEXCHANGESCHEME = _descriptor.EnumDescriptor(
  name='KeyExchangeScheme',
  full_name='KeyExchangeScheme',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='KEA_NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='KEA_ECDH', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=445,
  serialized_end=492,
)
_sym_db.RegisterEnumDescriptor(_KEYEXCHANGESCHEME)

KeyExchangeScheme = enum_type_wrapper.EnumTypeWrapper(_KEYEXCHANGESCHEME)
_ENCRYPTIONSCHEME = _descriptor.EnumDescriptor(
  name='EncryptionScheme',
  full_name='EncryptionScheme',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SA_NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SA_AES_256_CFB', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=494,
  serialized_end=545,
)
_sym_db.RegisterEnumDescriptor(_ENCRYPTIONSCHEME)

EncryptionScheme = enum_type_wrapper.EnumTypeWrapper(_ENCRYPTIONSCHEME)
KEA_NONE = 0
KEA_ECDH = 1
SA_NONE = 0
SA_AES_256_CFB = 1



_UUIDANDVERSION = _descriptor.Descriptor(
  name='UUIDandVersion',
  full_name='UUIDandVersion',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='UUIDandVersion.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='UUIDandVersion.version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=35,
  serialized_end=82,
)


_DEVICEOPSCMD = _descriptor.Descriptor(
  name='DeviceOpsCmd',
  full_name='DeviceOpsCmd',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='counter', full_name='DeviceOpsCmd.counter', index=0,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='desiredState', full_name='DeviceOpsCmd.desiredState', index=1,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='opsTime', full_name='DeviceOpsCmd.opsTime', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=84,
  serialized_end=154,
)


_CONFIGITEM = _descriptor.Descriptor(
  name='ConfigItem',
  full_name='ConfigItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='ConfigItem.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='ConfigItem.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=156,
  serialized_end=196,
)


_ADAPTER = _descriptor.Descriptor(
  name='Adapter',
  full_name='Adapter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='Adapter.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='Adapter.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=198,
  serialized_end=247,
)


_CIPHERINFO = _descriptor.Descriptor(
  name='CipherInfo',
  full_name='CipherInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='CipherInfo.id', index=0,
      number=100, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='keyExchangeScheme', full_name='CipherInfo.keyExchangeScheme', index=1,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='encryptionScheme', full_name='CipherInfo.encryptionScheme', index=2,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='initialValue', full_name='CipherInfo.initialValue', index=3,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='publicCert', full_name='CipherInfo.publicCert', index=4,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sha256', full_name='CipherInfo.sha256', index=5,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='signature', full_name='CipherInfo.signature', index=6,
      number=6, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=250,
  serialized_end=443,
)

_ADAPTER.fields_by_name['type'].enum_type = devmodel__pb2._PHYIOTYPE
_CIPHERINFO.fields_by_name['keyExchangeScheme'].enum_type = _KEYEXCHANGESCHEME
_CIPHERINFO.fields_by_name['encryptionScheme'].enum_type = _ENCRYPTIONSCHEME
DESCRIPTOR.message_types_by_name['UUIDandVersion'] = _UUIDANDVERSION
DESCRIPTOR.message_types_by_name['DeviceOpsCmd'] = _DEVICEOPSCMD
DESCRIPTOR.message_types_by_name['ConfigItem'] = _CONFIGITEM
DESCRIPTOR.message_types_by_name['Adapter'] = _ADAPTER
DESCRIPTOR.message_types_by_name['CipherInfo'] = _CIPHERINFO
DESCRIPTOR.enum_types_by_name['KeyExchangeScheme'] = _KEYEXCHANGESCHEME
DESCRIPTOR.enum_types_by_name['EncryptionScheme'] = _ENCRYPTIONSCHEME
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

UUIDandVersion = _reflection.GeneratedProtocolMessageType('UUIDandVersion', (_message.Message,), dict(
  DESCRIPTOR = _UUIDANDVERSION,
  __module__ = 'devcommon_pb2'
  # @@protoc_insertion_point(class_scope:UUIDandVersion)
  ))
_sym_db.RegisterMessage(UUIDandVersion)

DeviceOpsCmd = _reflection.GeneratedProtocolMessageType('DeviceOpsCmd', (_message.Message,), dict(
  DESCRIPTOR = _DEVICEOPSCMD,
  __module__ = 'devcommon_pb2'
  # @@protoc_insertion_point(class_scope:DeviceOpsCmd)
  ))
_sym_db.RegisterMessage(DeviceOpsCmd)

ConfigItem = _reflection.GeneratedProtocolMessageType('ConfigItem', (_message.Message,), dict(
  DESCRIPTOR = _CONFIGITEM,
  __module__ = 'devcommon_pb2'
  # @@protoc_insertion_point(class_scope:ConfigItem)
  ))
_sym_db.RegisterMessage(ConfigItem)

Adapter = _reflection.GeneratedProtocolMessageType('Adapter', (_message.Message,), dict(
  DESCRIPTOR = _ADAPTER,
  __module__ = 'devcommon_pb2'
  # @@protoc_insertion_point(class_scope:Adapter)
  ))
_sym_db.RegisterMessage(Adapter)

CipherInfo = _reflection.GeneratedProtocolMessageType('CipherInfo', (_message.Message,), dict(
  DESCRIPTOR = _CIPHERINFO,
  __module__ = 'devcommon_pb2'
  # @@protoc_insertion_point(class_scope:CipherInfo)
  ))
_sym_db.RegisterMessage(CipherInfo)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
