# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/devcommon.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from evecommon import devmodelcommon_pb2 as evecommon_dot_devmodelcommon__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='config/devcommon.proto',
  package='org.lfedge.eve.config',
  syntax='proto3',
  serialized_options=b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x16\x63onfig/devcommon.proto\x12\x15org.lfedge.eve.config\x1a\x1e\x65vecommon/devmodelcommon.proto\"/\n\x0eUUIDandVersion\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\"F\n\x0c\x44\x65viceOpsCmd\x12\x0f\n\x07\x63ounter\x18\x02 \x01(\r\x12\x14\n\x0c\x64\x65siredState\x18\x03 \x01(\x08\x12\x0f\n\x07opsTime\x18\x04 \x01(\t\"(\n\nConfigItem\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"u\n\x07\x41\x64\x61pter\x12.\n\x04type\x18\x01 \x01(\x0e\x32 .org.lfedge.eve.common.PhyIoType\x12\x0c\n\x04name\x18\x02 \x01(\t\x12,\n\x06\x65th_vf\x18\x03 \x01(\x0b\x32\x1c.org.lfedge.eve.config.EthVF\"@\n\x05\x45thVF\x12\r\n\x05index\x18\x01 \x01(\r\x12\x0b\n\x03pci\x18\x02 \x01(\t\x12\x0b\n\x03mac\x18\x03 \x01(\t\x12\x0e\n\x06vlanId\x18\x04 \x01(\rB=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3'
  ,
  dependencies=[evecommon_dot_devmodelcommon__pb2.DESCRIPTOR,])




_UUIDANDVERSION = _descriptor.Descriptor(
  name='UUIDandVersion',
  full_name='org.lfedge.eve.config.UUIDandVersion',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='org.lfedge.eve.config.UUIDandVersion.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='version', full_name='org.lfedge.eve.config.UUIDandVersion.version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=81,
  serialized_end=128,
)


_DEVICEOPSCMD = _descriptor.Descriptor(
  name='DeviceOpsCmd',
  full_name='org.lfedge.eve.config.DeviceOpsCmd',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='counter', full_name='org.lfedge.eve.config.DeviceOpsCmd.counter', index=0,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='desiredState', full_name='org.lfedge.eve.config.DeviceOpsCmd.desiredState', index=1,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='opsTime', full_name='org.lfedge.eve.config.DeviceOpsCmd.opsTime', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=130,
  serialized_end=200,
)


_CONFIGITEM = _descriptor.Descriptor(
  name='ConfigItem',
  full_name='org.lfedge.eve.config.ConfigItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='org.lfedge.eve.config.ConfigItem.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='org.lfedge.eve.config.ConfigItem.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=202,
  serialized_end=242,
)


_ADAPTER = _descriptor.Descriptor(
  name='Adapter',
  full_name='org.lfedge.eve.config.Adapter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='org.lfedge.eve.config.Adapter.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='org.lfedge.eve.config.Adapter.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='eth_vf', full_name='org.lfedge.eve.config.Adapter.eth_vf', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=244,
  serialized_end=361,
)


_ETHVF = _descriptor.Descriptor(
  name='EthVF',
  full_name='org.lfedge.eve.config.EthVF',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='index', full_name='org.lfedge.eve.config.EthVF.index', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pci', full_name='org.lfedge.eve.config.EthVF.pci', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='mac', full_name='org.lfedge.eve.config.EthVF.mac', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='vlanId', full_name='org.lfedge.eve.config.EthVF.vlanId', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=363,
  serialized_end=427,
)

_ADAPTER.fields_by_name['type'].enum_type = evecommon_dot_devmodelcommon__pb2._PHYIOTYPE
_ADAPTER.fields_by_name['eth_vf'].message_type = _ETHVF
DESCRIPTOR.message_types_by_name['UUIDandVersion'] = _UUIDANDVERSION
DESCRIPTOR.message_types_by_name['DeviceOpsCmd'] = _DEVICEOPSCMD
DESCRIPTOR.message_types_by_name['ConfigItem'] = _CONFIGITEM
DESCRIPTOR.message_types_by_name['Adapter'] = _ADAPTER
DESCRIPTOR.message_types_by_name['EthVF'] = _ETHVF
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

UUIDandVersion = _reflection.GeneratedProtocolMessageType('UUIDandVersion', (_message.Message,), {
  'DESCRIPTOR' : _UUIDANDVERSION,
  '__module__' : 'config.devcommon_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.UUIDandVersion)
  })
_sym_db.RegisterMessage(UUIDandVersion)

DeviceOpsCmd = _reflection.GeneratedProtocolMessageType('DeviceOpsCmd', (_message.Message,), {
  'DESCRIPTOR' : _DEVICEOPSCMD,
  '__module__' : 'config.devcommon_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.DeviceOpsCmd)
  })
_sym_db.RegisterMessage(DeviceOpsCmd)

ConfigItem = _reflection.GeneratedProtocolMessageType('ConfigItem', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGITEM,
  '__module__' : 'config.devcommon_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.ConfigItem)
  })
_sym_db.RegisterMessage(ConfigItem)

Adapter = _reflection.GeneratedProtocolMessageType('Adapter', (_message.Message,), {
  'DESCRIPTOR' : _ADAPTER,
  '__module__' : 'config.devcommon_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.Adapter)
  })
_sym_db.RegisterMessage(Adapter)

EthVF = _reflection.GeneratedProtocolMessageType('EthVF', (_message.Message,), {
  'DESCRIPTOR' : _ETHVF,
  '__module__' : 'config.devcommon_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.EthVF)
  })
_sym_db.RegisterMessage(EthVF)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
