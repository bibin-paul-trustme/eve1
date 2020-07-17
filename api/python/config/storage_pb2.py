# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/storage.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from config import devcommon_pb2 as config_dot_devcommon__pb2
from config import acipherinfo_pb2 as config_dot_acipherinfo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='config/storage.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config'),
  serialized_pb=_b('\n\x14\x63onfig/storage.proto\x1a\x16\x63onfig/devcommon.proto\x1a\x18\x63onfig/acipherinfo.proto\"P\n\rSignatureInfo\x12\x15\n\rintercertsurl\x18\x01 \x01(\t\x12\x15\n\rsignercerturl\x18\x02 \x01(\t\x12\x11\n\tsignature\x18\x03 \x01(\x0c\"\xa6\x01\n\x0f\x44\x61tastoreConfig\x12\n\n\x02id\x18\x64 \x01(\t\x12\x16\n\x05\x64Type\x18\x01 \x01(\x0e\x32\x07.DsType\x12\x0c\n\x04\x66qdn\x18\x02 \x01(\t\x12\x0e\n\x06\x61piKey\x18\x03 \x01(\t\x12\x10\n\x08password\x18\x04 \x01(\t\x12\r\n\x05\x64path\x18\x05 \x01(\t\x12\x0e\n\x06region\x18\x06 \x01(\t\x12 \n\ncipherData\x18\x07 \x01(\x0b\x32\x0c.CipherBlock\"\xaa\x01\n\x05Image\x12\'\n\x0euuidandversion\x18\x01 \x01(\x0b\x32\x0f.UUIDandVersion\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06sha256\x18\x03 \x01(\t\x12\x18\n\x07iformat\x18\x04 \x01(\x0e\x32\x07.Format\x12\x1f\n\x07siginfo\x18\x05 \x01(\x0b\x32\x0e.SignatureInfo\x12\x0c\n\x04\x64sId\x18\x06 \x01(\t\x12\x11\n\tsizeBytes\x18\x08 \x01(\x03\"\x8e\x01\n\x05\x44rive\x12\x15\n\x05image\x18\x01 \x01(\x0b\x32\x06.Image\x12\x10\n\x08readonly\x18\x05 \x01(\x08\x12\x10\n\x08preserve\x18\x06 \x01(\x08\x12\x1b\n\x07\x64rvtype\x18\x08 \x01(\x0e\x32\n.DriveType\x12\x17\n\x06target\x18\t \x01(\x0e\x32\x07.Target\x12\x14\n\x0cmaxsizebytes\x18\n \x01(\x03\"\xc6\x01\n\x0b\x43ontentTree\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0c\n\x04\x64sId\x18\x02 \x01(\t\x12\x0b\n\x03URL\x18\x03 \x01(\t\x12\x18\n\x07iformat\x18\x04 \x01(\x0e\x32\x07.Format\x12\x0e\n\x06sha256\x18\x05 \x01(\t\x12\x14\n\x0cmaxSizeBytes\x18\x06 \x01(\x04\x12\x1f\n\x07siginfo\x18\x07 \x01(\x0b\x32\x0e.SignatureInfo\x12\x13\n\x0b\x64isplayName\x18\x08 \x01(\t\x12\x18\n\x10generation_count\x18\t \x01(\x03\"\\\n\x13VolumeContentOrigin\x12&\n\x04type\x18\x01 \x01(\x0e\x32\x18.VolumeContentOriginType\x12\x1d\n\x15\x64ownloadContentTreeID\x18\x02 \x01(\t\"\xe4\x01\n\x06Volume\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12$\n\x06origin\x18\x02 \x01(\x0b\x32\x14.VolumeContentOrigin\x12)\n\tprotocols\x18\x03 \x03(\x0e\x32\x16.VolumeAccessProtocols\x12\x17\n\x0fgenerationCount\x18\x04 \x01(\x03\x12\x14\n\x0cmaxsizebytes\x18\x05 \x01(\x03\x12\x10\n\x08readonly\x18\x06 \x01(\x08\x12\x13\n\x0b\x64isplayName\x18\x07 \x01(\t\x12\x12\n\nclear_text\x18\x08 \x01(\x08\x12\x11\n\tmount_dir\x18\t \x01(\t*p\n\x06\x44sType\x12\r\n\tDsUnknown\x10\x00\x12\n\n\x06\x44sHttp\x10\x01\x12\x0b\n\x07\x44sHttps\x10\x02\x12\x08\n\x04\x44sS3\x10\x03\x12\n\n\x06\x44sSFTP\x10\x04\x12\x17\n\x13\x44sContainerRegistry\x10\x05\x12\x0f\n\x0b\x44sAzureBlob\x10\x06*k\n\x06\x46ormat\x12\x0e\n\nFmtUnknown\x10\x00\x12\x07\n\x03RAW\x10\x01\x12\x08\n\x04QCOW\x10\x02\x12\t\n\x05QCOW2\x10\x03\x12\x07\n\x03VHD\x10\x04\x12\x08\n\x04VMDK\x10\x05\x12\x07\n\x03OVA\x10\x06\x12\x08\n\x04VHDX\x10\x07\x12\r\n\tCONTAINER\x10\x08*G\n\x06Target\x12\x0e\n\nTgtUnknown\x10\x00\x12\x08\n\x04\x44isk\x10\x01\x12\n\n\x06Kernel\x10\x02\x12\n\n\x06Initrd\x10\x03\x12\x0b\n\x07RamDisk\x10\x04*I\n\tDriveType\x12\x10\n\x0cUnclassified\x10\x00\x12\t\n\x05\x43\x44ROM\x10\x01\x12\x07\n\x03HDD\x10\x02\x12\x07\n\x03NET\x10\x03\x12\r\n\tHDD_EMPTY\x10\x04*1\n\x15VolumeAccessProtocols\x12\x0c\n\x08VAP_NONE\x10\x00\x12\n\n\x06VAP_9P\x10\x01*N\n\x17VolumeContentOriginType\x12\x10\n\x0cVCOT_UNKNOWN\x10\x00\x12\x0e\n\nVCOT_BLANK\x10\x01\x12\x11\n\rVCOT_DOWNLOAD\x10\x02\x42=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3')
  ,
  dependencies=[config_dot_devcommon__pb2.DESCRIPTOR,config_dot_acipherinfo__pb2.DESCRIPTOR,])

_DSTYPE = _descriptor.EnumDescriptor(
  name='DsType',
  full_name='DsType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DsUnknown', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DsHttp', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DsHttps', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DsS3', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DsSFTP', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DsContainerRegistry', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DsAzureBlob', index=6, number=6,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1169,
  serialized_end=1281,
)
_sym_db.RegisterEnumDescriptor(_DSTYPE)

DsType = enum_type_wrapper.EnumTypeWrapper(_DSTYPE)
_FORMAT = _descriptor.EnumDescriptor(
  name='Format',
  full_name='Format',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='FmtUnknown', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RAW', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='QCOW', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='QCOW2', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VHD', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VMDK', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OVA', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VHDX', index=7, number=7,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CONTAINER', index=8, number=8,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1283,
  serialized_end=1390,
)
_sym_db.RegisterEnumDescriptor(_FORMAT)

Format = enum_type_wrapper.EnumTypeWrapper(_FORMAT)
_TARGET = _descriptor.EnumDescriptor(
  name='Target',
  full_name='Target',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TgtUnknown', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Disk', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Kernel', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Initrd', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RamDisk', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1392,
  serialized_end=1463,
)
_sym_db.RegisterEnumDescriptor(_TARGET)

Target = enum_type_wrapper.EnumTypeWrapper(_TARGET)
_DRIVETYPE = _descriptor.EnumDescriptor(
  name='DriveType',
  full_name='DriveType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Unclassified', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CDROM', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HDD', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NET', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HDD_EMPTY', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1465,
  serialized_end=1538,
)
_sym_db.RegisterEnumDescriptor(_DRIVETYPE)

DriveType = enum_type_wrapper.EnumTypeWrapper(_DRIVETYPE)
_VOLUMEACCESSPROTOCOLS = _descriptor.EnumDescriptor(
  name='VolumeAccessProtocols',
  full_name='VolumeAccessProtocols',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='VAP_NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VAP_9P', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1540,
  serialized_end=1589,
)
_sym_db.RegisterEnumDescriptor(_VOLUMEACCESSPROTOCOLS)

VolumeAccessProtocols = enum_type_wrapper.EnumTypeWrapper(_VOLUMEACCESSPROTOCOLS)
_VOLUMECONTENTORIGINTYPE = _descriptor.EnumDescriptor(
  name='VolumeContentOriginType',
  full_name='VolumeContentOriginType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='VCOT_UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VCOT_BLANK', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VCOT_DOWNLOAD', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1591,
  serialized_end=1669,
)
_sym_db.RegisterEnumDescriptor(_VOLUMECONTENTORIGINTYPE)

VolumeContentOriginType = enum_type_wrapper.EnumTypeWrapper(_VOLUMECONTENTORIGINTYPE)
DsUnknown = 0
DsHttp = 1
DsHttps = 2
DsS3 = 3
DsSFTP = 4
DsContainerRegistry = 5
DsAzureBlob = 6
FmtUnknown = 0
RAW = 1
QCOW = 2
QCOW2 = 3
VHD = 4
VMDK = 5
OVA = 6
VHDX = 7
CONTAINER = 8
TgtUnknown = 0
Disk = 1
Kernel = 2
Initrd = 3
RamDisk = 4
Unclassified = 0
CDROM = 1
HDD = 2
NET = 3
HDD_EMPTY = 4
VAP_NONE = 0
VAP_9P = 1
VCOT_UNKNOWN = 0
VCOT_BLANK = 1
VCOT_DOWNLOAD = 2



_SIGNATUREINFO = _descriptor.Descriptor(
  name='SignatureInfo',
  full_name='SignatureInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='intercertsurl', full_name='SignatureInfo.intercertsurl', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='signercerturl', full_name='SignatureInfo.signercerturl', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='signature', full_name='SignatureInfo.signature', index=2,
      number=3, type=12, cpp_type=9, label=1,
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
  serialized_start=74,
  serialized_end=154,
)


_DATASTORECONFIG = _descriptor.Descriptor(
  name='DatastoreConfig',
  full_name='DatastoreConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='DatastoreConfig.id', index=0,
      number=100, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dType', full_name='DatastoreConfig.dType', index=1,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fqdn', full_name='DatastoreConfig.fqdn', index=2,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='apiKey', full_name='DatastoreConfig.apiKey', index=3,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='password', full_name='DatastoreConfig.password', index=4,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dpath', full_name='DatastoreConfig.dpath', index=5,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='region', full_name='DatastoreConfig.region', index=6,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cipherData', full_name='DatastoreConfig.cipherData', index=7,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=157,
  serialized_end=323,
)


_IMAGE = _descriptor.Descriptor(
  name='Image',
  full_name='Image',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuidandversion', full_name='Image.uuidandversion', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='Image.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sha256', full_name='Image.sha256', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='iformat', full_name='Image.iformat', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='siginfo', full_name='Image.siginfo', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dsId', full_name='Image.dsId', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sizeBytes', full_name='Image.sizeBytes', index=6,
      number=8, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=326,
  serialized_end=496,
)


_DRIVE = _descriptor.Descriptor(
  name='Drive',
  full_name='Drive',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='Drive.image', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='readonly', full_name='Drive.readonly', index=1,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='preserve', full_name='Drive.preserve', index=2,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='drvtype', full_name='Drive.drvtype', index=3,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='target', full_name='Drive.target', index=4,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='maxsizebytes', full_name='Drive.maxsizebytes', index=5,
      number=10, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=499,
  serialized_end=641,
)


_CONTENTTREE = _descriptor.Descriptor(
  name='ContentTree',
  full_name='ContentTree',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='ContentTree.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dsId', full_name='ContentTree.dsId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='URL', full_name='ContentTree.URL', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='iformat', full_name='ContentTree.iformat', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sha256', full_name='ContentTree.sha256', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='maxSizeBytes', full_name='ContentTree.maxSizeBytes', index=5,
      number=6, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='siginfo', full_name='ContentTree.siginfo', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='displayName', full_name='ContentTree.displayName', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='generation_count', full_name='ContentTree.generation_count', index=8,
      number=9, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=644,
  serialized_end=842,
)


_VOLUMECONTENTORIGIN = _descriptor.Descriptor(
  name='VolumeContentOrigin',
  full_name='VolumeContentOrigin',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='VolumeContentOrigin.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='downloadContentTreeID', full_name='VolumeContentOrigin.downloadContentTreeID', index=1,
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
  serialized_start=844,
  serialized_end=936,
)


_VOLUME = _descriptor.Descriptor(
  name='Volume',
  full_name='Volume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='Volume.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='origin', full_name='Volume.origin', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='protocols', full_name='Volume.protocols', index=2,
      number=3, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='generationCount', full_name='Volume.generationCount', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='maxsizebytes', full_name='Volume.maxsizebytes', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='readonly', full_name='Volume.readonly', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='displayName', full_name='Volume.displayName', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='clear_text', full_name='Volume.clear_text', index=7,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mount_dir', full_name='Volume.mount_dir', index=8,
      number=9, type=9, cpp_type=9, label=1,
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
  serialized_start=939,
  serialized_end=1167,
)

_DATASTORECONFIG.fields_by_name['dType'].enum_type = _DSTYPE
_DATASTORECONFIG.fields_by_name['cipherData'].message_type = config_dot_acipherinfo__pb2._CIPHERBLOCK
_IMAGE.fields_by_name['uuidandversion'].message_type = config_dot_devcommon__pb2._UUIDANDVERSION
_IMAGE.fields_by_name['iformat'].enum_type = _FORMAT
_IMAGE.fields_by_name['siginfo'].message_type = _SIGNATUREINFO
_DRIVE.fields_by_name['image'].message_type = _IMAGE
_DRIVE.fields_by_name['drvtype'].enum_type = _DRIVETYPE
_DRIVE.fields_by_name['target'].enum_type = _TARGET
_CONTENTTREE.fields_by_name['iformat'].enum_type = _FORMAT
_CONTENTTREE.fields_by_name['siginfo'].message_type = _SIGNATUREINFO
_VOLUMECONTENTORIGIN.fields_by_name['type'].enum_type = _VOLUMECONTENTORIGINTYPE
_VOLUME.fields_by_name['origin'].message_type = _VOLUMECONTENTORIGIN
_VOLUME.fields_by_name['protocols'].enum_type = _VOLUMEACCESSPROTOCOLS
DESCRIPTOR.message_types_by_name['SignatureInfo'] = _SIGNATUREINFO
DESCRIPTOR.message_types_by_name['DatastoreConfig'] = _DATASTORECONFIG
DESCRIPTOR.message_types_by_name['Image'] = _IMAGE
DESCRIPTOR.message_types_by_name['Drive'] = _DRIVE
DESCRIPTOR.message_types_by_name['ContentTree'] = _CONTENTTREE
DESCRIPTOR.message_types_by_name['VolumeContentOrigin'] = _VOLUMECONTENTORIGIN
DESCRIPTOR.message_types_by_name['Volume'] = _VOLUME
DESCRIPTOR.enum_types_by_name['DsType'] = _DSTYPE
DESCRIPTOR.enum_types_by_name['Format'] = _FORMAT
DESCRIPTOR.enum_types_by_name['Target'] = _TARGET
DESCRIPTOR.enum_types_by_name['DriveType'] = _DRIVETYPE
DESCRIPTOR.enum_types_by_name['VolumeAccessProtocols'] = _VOLUMEACCESSPROTOCOLS
DESCRIPTOR.enum_types_by_name['VolumeContentOriginType'] = _VOLUMECONTENTORIGINTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SignatureInfo = _reflection.GeneratedProtocolMessageType('SignatureInfo', (_message.Message,), dict(
  DESCRIPTOR = _SIGNATUREINFO,
  __module__ = 'config.storage_pb2'
  # @@protoc_insertion_point(class_scope:SignatureInfo)
  ))
_sym_db.RegisterMessage(SignatureInfo)

DatastoreConfig = _reflection.GeneratedProtocolMessageType('DatastoreConfig', (_message.Message,), dict(
  DESCRIPTOR = _DATASTORECONFIG,
  __module__ = 'config.storage_pb2'
  # @@protoc_insertion_point(class_scope:DatastoreConfig)
  ))
_sym_db.RegisterMessage(DatastoreConfig)

Image = _reflection.GeneratedProtocolMessageType('Image', (_message.Message,), dict(
  DESCRIPTOR = _IMAGE,
  __module__ = 'config.storage_pb2'
  # @@protoc_insertion_point(class_scope:Image)
  ))
_sym_db.RegisterMessage(Image)

Drive = _reflection.GeneratedProtocolMessageType('Drive', (_message.Message,), dict(
  DESCRIPTOR = _DRIVE,
  __module__ = 'config.storage_pb2'
  # @@protoc_insertion_point(class_scope:Drive)
  ))
_sym_db.RegisterMessage(Drive)

ContentTree = _reflection.GeneratedProtocolMessageType('ContentTree', (_message.Message,), dict(
  DESCRIPTOR = _CONTENTTREE,
  __module__ = 'config.storage_pb2'
  # @@protoc_insertion_point(class_scope:ContentTree)
  ))
_sym_db.RegisterMessage(ContentTree)

VolumeContentOrigin = _reflection.GeneratedProtocolMessageType('VolumeContentOrigin', (_message.Message,), dict(
  DESCRIPTOR = _VOLUMECONTENTORIGIN,
  __module__ = 'config.storage_pb2'
  # @@protoc_insertion_point(class_scope:VolumeContentOrigin)
  ))
_sym_db.RegisterMessage(VolumeContentOrigin)

Volume = _reflection.GeneratedProtocolMessageType('Volume', (_message.Message,), dict(
  DESCRIPTOR = _VOLUME,
  __module__ = 'config.storage_pb2'
  # @@protoc_insertion_point(class_scope:Volume)
  ))
_sym_db.RegisterMessage(Volume)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
