# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/devconfig.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from config import acipherinfo_pb2 as config_dot_acipherinfo__pb2
from config import appconfig_pb2 as config_dot_appconfig__pb2
from config import baseosconfig_pb2 as config_dot_baseosconfig__pb2
from config import devcommon_pb2 as config_dot_devcommon__pb2
from config import devmodel_pb2 as config_dot_devmodel__pb2
from config import netconfig_pb2 as config_dot_netconfig__pb2
from config import netinst_pb2 as config_dot_netinst__pb2
from config import storage_pb2 as config_dot_storage__pb2
from config import edgeview_pb2 as config_dot_edgeview__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='config/devconfig.proto',
  package='org.lfedge.eve.config',
  syntax='proto3',
  serialized_options=b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x16\x63onfig/devconfig.proto\x12\x15org.lfedge.eve.config\x1a\x18\x63onfig/acipherinfo.proto\x1a\x16\x63onfig/appconfig.proto\x1a\x19\x63onfig/baseosconfig.proto\x1a\x16\x63onfig/devcommon.proto\x1a\x15\x63onfig/devmodel.proto\x1a\x16\x63onfig/netconfig.proto\x1a\x14\x63onfig/netinst.proto\x1a\x14\x63onfig/storage.proto\x1a\x15\x63onfig/edgeview.proto\"\xe3\x0b\n\rEdgeDevConfig\x12\x31\n\x02id\x18\x01 \x01(\x0b\x32%.org.lfedge.eve.config.UUIDandVersion\x12\x36\n\x04\x61pps\x18\x04 \x03(\x0b\x32(.org.lfedge.eve.config.AppInstanceConfig\x12\x36\n\x08networks\x18\x05 \x03(\x0b\x32$.org.lfedge.eve.config.NetworkConfig\x12:\n\ndatastores\x18\x06 \x03(\x0b\x32&.org.lfedge.eve.config.DatastoreConfig\x12\x31\n\x04\x62\x61se\x18\x08 \x03(\x0b\x32#.org.lfedge.eve.config.BaseOSConfig\x12\x33\n\x06reboot\x18\t \x01(\x0b\x32#.org.lfedge.eve.config.DeviceOpsCmd\x12\x33\n\x06\x62\x61\x63kup\x18\n \x01(\x0b\x32#.org.lfedge.eve.config.DeviceOpsCmd\x12\x36\n\x0b\x63onfigItems\x18\x0b \x03(\x0b\x32!.org.lfedge.eve.config.ConfigItem\x12?\n\x11systemAdapterList\x18\x0c \x03(\x0b\x32$.org.lfedge.eve.config.SystemAdapter\x12\x37\n\x0c\x64\x65viceIoList\x18\r \x03(\x0b\x32!.org.lfedge.eve.config.PhysicalIO\x12\x14\n\x0cmanufacturer\x18\x0e \x01(\t\x12\x13\n\x0bproductName\x18\x0f \x01(\t\x12\x46\n\x10networkInstances\x18\x10 \x03(\x0b\x32,.org.lfedge.eve.config.NetworkInstanceConfig\x12<\n\x0e\x63ipherContexts\x18\x13 \x03(\x0b\x32$.org.lfedge.eve.config.CipherContext\x12\x37\n\x0b\x63ontentInfo\x18\x14 \x03(\x0b\x32\".org.lfedge.eve.config.ContentTree\x12.\n\x07volumes\x18\x15 \x03(\x0b\x32\x1d.org.lfedge.eve.config.Volume\x12!\n\x19\x63ontrollercert_confighash\x18\x16 \x01(\t\x12\x18\n\x10maintenance_mode\x18\x18 \x01(\x08\x12\x18\n\x10\x63ontroller_epoch\x18\x19 \x01(\x03\x12-\n\x06\x62\x61seos\x18\x1a \x01(\x0b\x32\x1d.org.lfedge.eve.config.BaseOS\x12\x16\n\x0eglobal_profile\x18\x1b \x01(\t\x12\x1c\n\x14local_profile_server\x18\x1c \x01(\t\x12\x1c\n\x14profile_server_token\x18\x1d \x01(\t\x12\x31\n\x05vlans\x18\x1e \x03(\x0b\x32\".org.lfedge.eve.config.VlanAdapter\x12\x31\n\x05\x62onds\x18\x1f \x03(\x0b\x32\".org.lfedge.eve.config.BondAdapter\x12\x37\n\x08\x65\x64geview\x18  \x01(\x0b\x32%.org.lfedge.eve.config.EdgeViewConfig\x12\x31\n\x05\x64isks\x18! \x01(\x0b\x32\".org.lfedge.eve.config.DisksConfig\x12\x35\n\x08shutdown\x18\" \x01(\x0b\x32#.org.lfedge.eve.config.DeviceOpsCmd\x12\x13\n\x0b\x64\x65vice_name\x18# \x01(\t\x12\x14\n\x0cproject_name\x18$ \x01(\t\x12\x12\n\nproject_id\x18% \x01(\t\x12\x17\n\x0f\x65nterprise_name\x18& \x01(\t\x12\x15\n\renterprise_id\x18\' \x01(\t\x12\x38\n\tsnapshots\x18( \x03(\x0b\x32%.org.lfedge.eve.config.SnapshotConfig\"<\n\rConfigRequest\x12\x12\n\nconfigHash\x18\x01 \x01(\t\x12\x17\n\x0fintegrity_token\x18\x02 \x01(\x0c\"Z\n\x0e\x43onfigResponse\x12\x34\n\x06\x63onfig\x18\x01 \x01(\x0b\x32$.org.lfedge.eve.config.EdgeDevConfig\x12\x12\n\nconfigHash\x18\x02 \x01(\tB=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3'
  ,
  dependencies=[config_dot_acipherinfo__pb2.DESCRIPTOR,config_dot_appconfig__pb2.DESCRIPTOR,config_dot_baseosconfig__pb2.DESCRIPTOR,config_dot_devcommon__pb2.DESCRIPTOR,config_dot_devmodel__pb2.DESCRIPTOR,config_dot_netconfig__pb2.DESCRIPTOR,config_dot_netinst__pb2.DESCRIPTOR,config_dot_storage__pb2.DESCRIPTOR,config_dot_edgeview__pb2.DESCRIPTOR,])




_EDGEDEVCONFIG = _descriptor.Descriptor(
  name='EdgeDevConfig',
  full_name='org.lfedge.eve.config.EdgeDevConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='org.lfedge.eve.config.EdgeDevConfig.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='apps', full_name='org.lfedge.eve.config.EdgeDevConfig.apps', index=1,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='networks', full_name='org.lfedge.eve.config.EdgeDevConfig.networks', index=2,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='datastores', full_name='org.lfedge.eve.config.EdgeDevConfig.datastores', index=3,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='base', full_name='org.lfedge.eve.config.EdgeDevConfig.base', index=4,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='reboot', full_name='org.lfedge.eve.config.EdgeDevConfig.reboot', index=5,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='backup', full_name='org.lfedge.eve.config.EdgeDevConfig.backup', index=6,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='configItems', full_name='org.lfedge.eve.config.EdgeDevConfig.configItems', index=7,
      number=11, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='systemAdapterList', full_name='org.lfedge.eve.config.EdgeDevConfig.systemAdapterList', index=8,
      number=12, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='deviceIoList', full_name='org.lfedge.eve.config.EdgeDevConfig.deviceIoList', index=9,
      number=13, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='manufacturer', full_name='org.lfedge.eve.config.EdgeDevConfig.manufacturer', index=10,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='productName', full_name='org.lfedge.eve.config.EdgeDevConfig.productName', index=11,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='networkInstances', full_name='org.lfedge.eve.config.EdgeDevConfig.networkInstances', index=12,
      number=16, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cipherContexts', full_name='org.lfedge.eve.config.EdgeDevConfig.cipherContexts', index=13,
      number=19, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='contentInfo', full_name='org.lfedge.eve.config.EdgeDevConfig.contentInfo', index=14,
      number=20, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='volumes', full_name='org.lfedge.eve.config.EdgeDevConfig.volumes', index=15,
      number=21, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='controllercert_confighash', full_name='org.lfedge.eve.config.EdgeDevConfig.controllercert_confighash', index=16,
      number=22, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='maintenance_mode', full_name='org.lfedge.eve.config.EdgeDevConfig.maintenance_mode', index=17,
      number=24, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='controller_epoch', full_name='org.lfedge.eve.config.EdgeDevConfig.controller_epoch', index=18,
      number=25, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='baseos', full_name='org.lfedge.eve.config.EdgeDevConfig.baseos', index=19,
      number=26, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='global_profile', full_name='org.lfedge.eve.config.EdgeDevConfig.global_profile', index=20,
      number=27, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='local_profile_server', full_name='org.lfedge.eve.config.EdgeDevConfig.local_profile_server', index=21,
      number=28, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='profile_server_token', full_name='org.lfedge.eve.config.EdgeDevConfig.profile_server_token', index=22,
      number=29, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='vlans', full_name='org.lfedge.eve.config.EdgeDevConfig.vlans', index=23,
      number=30, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='bonds', full_name='org.lfedge.eve.config.EdgeDevConfig.bonds', index=24,
      number=31, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='edgeview', full_name='org.lfedge.eve.config.EdgeDevConfig.edgeview', index=25,
      number=32, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='disks', full_name='org.lfedge.eve.config.EdgeDevConfig.disks', index=26,
      number=33, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='shutdown', full_name='org.lfedge.eve.config.EdgeDevConfig.shutdown', index=27,
      number=34, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='device_name', full_name='org.lfedge.eve.config.EdgeDevConfig.device_name', index=28,
      number=35, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_name', full_name='org.lfedge.eve.config.EdgeDevConfig.project_name', index=29,
      number=36, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='org.lfedge.eve.config.EdgeDevConfig.project_id', index=30,
      number=37, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='enterprise_name', full_name='org.lfedge.eve.config.EdgeDevConfig.enterprise_name', index=31,
      number=38, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='enterprise_id', full_name='org.lfedge.eve.config.EdgeDevConfig.enterprise_id', index=32,
      number=39, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='snapshots', full_name='org.lfedge.eve.config.EdgeDevConfig.snapshots', index=33,
      number=40, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=265,
  serialized_end=1772,
)


_CONFIGREQUEST = _descriptor.Descriptor(
  name='ConfigRequest',
  full_name='org.lfedge.eve.config.ConfigRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='configHash', full_name='org.lfedge.eve.config.ConfigRequest.configHash', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='integrity_token', full_name='org.lfedge.eve.config.ConfigRequest.integrity_token', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
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
  serialized_start=1774,
  serialized_end=1834,
)


_CONFIGRESPONSE = _descriptor.Descriptor(
  name='ConfigResponse',
  full_name='org.lfedge.eve.config.ConfigResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='config', full_name='org.lfedge.eve.config.ConfigResponse.config', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='configHash', full_name='org.lfedge.eve.config.ConfigResponse.configHash', index=1,
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
  serialized_start=1836,
  serialized_end=1926,
)

_EDGEDEVCONFIG.fields_by_name['id'].message_type = config_dot_devcommon__pb2._UUIDANDVERSION
_EDGEDEVCONFIG.fields_by_name['apps'].message_type = config_dot_appconfig__pb2._APPINSTANCECONFIG
_EDGEDEVCONFIG.fields_by_name['networks'].message_type = config_dot_netconfig__pb2._NETWORKCONFIG
_EDGEDEVCONFIG.fields_by_name['datastores'].message_type = config_dot_storage__pb2._DATASTORECONFIG
_EDGEDEVCONFIG.fields_by_name['base'].message_type = config_dot_baseosconfig__pb2._BASEOSCONFIG
_EDGEDEVCONFIG.fields_by_name['reboot'].message_type = config_dot_devcommon__pb2._DEVICEOPSCMD
_EDGEDEVCONFIG.fields_by_name['backup'].message_type = config_dot_devcommon__pb2._DEVICEOPSCMD
_EDGEDEVCONFIG.fields_by_name['configItems'].message_type = config_dot_devcommon__pb2._CONFIGITEM
_EDGEDEVCONFIG.fields_by_name['systemAdapterList'].message_type = config_dot_devmodel__pb2._SYSTEMADAPTER
_EDGEDEVCONFIG.fields_by_name['deviceIoList'].message_type = config_dot_devmodel__pb2._PHYSICALIO
_EDGEDEVCONFIG.fields_by_name['networkInstances'].message_type = config_dot_netinst__pb2._NETWORKINSTANCECONFIG
_EDGEDEVCONFIG.fields_by_name['cipherContexts'].message_type = config_dot_acipherinfo__pb2._CIPHERCONTEXT
_EDGEDEVCONFIG.fields_by_name['contentInfo'].message_type = config_dot_storage__pb2._CONTENTTREE
_EDGEDEVCONFIG.fields_by_name['volumes'].message_type = config_dot_storage__pb2._VOLUME
_EDGEDEVCONFIG.fields_by_name['baseos'].message_type = config_dot_baseosconfig__pb2._BASEOS
_EDGEDEVCONFIG.fields_by_name['vlans'].message_type = config_dot_devmodel__pb2._VLANADAPTER
_EDGEDEVCONFIG.fields_by_name['bonds'].message_type = config_dot_devmodel__pb2._BONDADAPTER
_EDGEDEVCONFIG.fields_by_name['edgeview'].message_type = config_dot_edgeview__pb2._EDGEVIEWCONFIG
_EDGEDEVCONFIG.fields_by_name['disks'].message_type = config_dot_storage__pb2._DISKSCONFIG
_EDGEDEVCONFIG.fields_by_name['shutdown'].message_type = config_dot_devcommon__pb2._DEVICEOPSCMD
_EDGEDEVCONFIG.fields_by_name['snapshots'].message_type = config_dot_storage__pb2._SNAPSHOTCONFIG
_CONFIGRESPONSE.fields_by_name['config'].message_type = _EDGEDEVCONFIG
DESCRIPTOR.message_types_by_name['EdgeDevConfig'] = _EDGEDEVCONFIG
DESCRIPTOR.message_types_by_name['ConfigRequest'] = _CONFIGREQUEST
DESCRIPTOR.message_types_by_name['ConfigResponse'] = _CONFIGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

EdgeDevConfig = _reflection.GeneratedProtocolMessageType('EdgeDevConfig', (_message.Message,), {
  'DESCRIPTOR' : _EDGEDEVCONFIG,
  '__module__' : 'config.devconfig_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.EdgeDevConfig)
  })
_sym_db.RegisterMessage(EdgeDevConfig)

ConfigRequest = _reflection.GeneratedProtocolMessageType('ConfigRequest', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGREQUEST,
  '__module__' : 'config.devconfig_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.ConfigRequest)
  })
_sym_db.RegisterMessage(ConfigRequest)

ConfigResponse = _reflection.GeneratedProtocolMessageType('ConfigResponse', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGRESPONSE,
  '__module__' : 'config.devconfig_pb2'
  # @@protoc_insertion_point(class_scope:org.lfedge.eve.config.ConfigResponse)
  })
_sym_db.RegisterMessage(ConfigResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
