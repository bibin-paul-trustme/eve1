# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: certs/certs.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from evecommon import evecommon_pb2 as evecommon_dot_evecommon__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='certs/certs.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\024org.lfedge.eve.certsZ#github.com/lf-edge/eve/api/go/certs'),
  serialized_pb=_b('\n\x11\x63\x65rts/certs.proto\x1a\x19\x65vecommon/evecommon.proto\"(\n\x0fZControllerCert\x12\x15\n\x05\x63\x65rts\x18\x01 \x03(\x0b\x32\x06.ZCert\"\x99\x01\n\x05ZCert\x12\x36\n\x08hashAlgo\x18\x01 \x01(\x0e\x32$.org.lfedge.eve.common.HashAlgorithm\x12\x10\n\x08\x63\x65rtHash\x18\x02 \x01(\x0c\x12\x18\n\x04type\x18\x03 \x01(\x0e\x32\n.ZCertType\x12\x0c\n\x04\x63\x65rt\x18\x04 \x01(\x0c\x12\x1e\n\nattributes\x18\x05 \x01(\x0b\x32\n.ZCertAttr\"/\n\tZCertAttr\x12\x12\n\nis_mutable\x18\x01 \x01(\x08\x12\x0e\n\x06is_tpm\x18\x02 \x01(\x08*\xaf\x02\n\tZCertType\x12\x1d\n\x19\x43\x45RT_TYPE_CONTROLLER_NONE\x10\x00\x12 \n\x1c\x43\x45RT_TYPE_CONTROLLER_SIGNING\x10\x01\x12%\n!CERT_TYPE_CONTROLLER_INTERMEDIATE\x10\x02\x12&\n\"CERT_TYPE_CONTROLLER_ECDH_EXCHANGE\x10\x03\x12\x1f\n\x1b\x43\x45RT_TYPE_DEVICE_ONBOARDING\x10\n\x12\'\n#CERT_TYPE_DEVICE_RESTRICTED_SIGNING\x10\x0b\x12$\n CERT_TYPE_DEVICE_ENDORSEMENT_RSA\x10\x0c\x12\"\n\x1e\x43\x45RT_TYPE_DEVICE_ECDH_EXCHANGE\x10\rB;\n\x14org.lfedge.eve.certsZ#github.com/lf-edge/eve/api/go/certsb\x06proto3')
  ,
  dependencies=[evecommon_dot_evecommon__pb2.DESCRIPTOR,])

_ZCERTTYPE = _descriptor.EnumDescriptor(
  name='ZCertType',
  full_name='ZCertType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_CONTROLLER_NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_CONTROLLER_SIGNING', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_CONTROLLER_INTERMEDIATE', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_CONTROLLER_ECDH_EXCHANGE', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ONBOARDING', index=4, number=10,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_RESTRICTED_SIGNING', index=5, number=11,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ENDORSEMENT_RSA', index=6, number=12,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ECDH_EXCHANGE', index=7, number=13,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=296,
  serialized_end=599,
)
_sym_db.RegisterEnumDescriptor(_ZCERTTYPE)

ZCertType = enum_type_wrapper.EnumTypeWrapper(_ZCERTTYPE)
CERT_TYPE_CONTROLLER_NONE = 0
CERT_TYPE_CONTROLLER_SIGNING = 1
CERT_TYPE_CONTROLLER_INTERMEDIATE = 2
CERT_TYPE_CONTROLLER_ECDH_EXCHANGE = 3
CERT_TYPE_DEVICE_ONBOARDING = 10
CERT_TYPE_DEVICE_RESTRICTED_SIGNING = 11
CERT_TYPE_DEVICE_ENDORSEMENT_RSA = 12
CERT_TYPE_DEVICE_ECDH_EXCHANGE = 13



_ZCONTROLLERCERT = _descriptor.Descriptor(
  name='ZControllerCert',
  full_name='ZControllerCert',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='certs', full_name='ZControllerCert.certs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=48,
  serialized_end=88,
)


_ZCERT = _descriptor.Descriptor(
  name='ZCert',
  full_name='ZCert',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hashAlgo', full_name='ZCert.hashAlgo', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='certHash', full_name='ZCert.certHash', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='ZCert.type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cert', full_name='ZCert.cert', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='attributes', full_name='ZCert.attributes', index=4,
      number=5, type=11, cpp_type=10, label=1,
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
  serialized_start=91,
  serialized_end=244,
)


_ZCERTATTR = _descriptor.Descriptor(
  name='ZCertAttr',
  full_name='ZCertAttr',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='is_mutable', full_name='ZCertAttr.is_mutable', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_tpm', full_name='ZCertAttr.is_tpm', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=246,
  serialized_end=293,
)

_ZCONTROLLERCERT.fields_by_name['certs'].message_type = _ZCERT
_ZCERT.fields_by_name['hashAlgo'].enum_type = evecommon_dot_evecommon__pb2._HASHALGORITHM
_ZCERT.fields_by_name['type'].enum_type = _ZCERTTYPE
_ZCERT.fields_by_name['attributes'].message_type = _ZCERTATTR
DESCRIPTOR.message_types_by_name['ZControllerCert'] = _ZCONTROLLERCERT
DESCRIPTOR.message_types_by_name['ZCert'] = _ZCERT
DESCRIPTOR.message_types_by_name['ZCertAttr'] = _ZCERTATTR
DESCRIPTOR.enum_types_by_name['ZCertType'] = _ZCERTTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ZControllerCert = _reflection.GeneratedProtocolMessageType('ZControllerCert', (_message.Message,), dict(
  DESCRIPTOR = _ZCONTROLLERCERT,
  __module__ = 'certs.certs_pb2'
  # @@protoc_insertion_point(class_scope:ZControllerCert)
  ))
_sym_db.RegisterMessage(ZControllerCert)

ZCert = _reflection.GeneratedProtocolMessageType('ZCert', (_message.Message,), dict(
  DESCRIPTOR = _ZCERT,
  __module__ = 'certs.certs_pb2'
  # @@protoc_insertion_point(class_scope:ZCert)
  ))
_sym_db.RegisterMessage(ZCert)

ZCertAttr = _reflection.GeneratedProtocolMessageType('ZCertAttr', (_message.Message,), dict(
  DESCRIPTOR = _ZCERTATTR,
  __module__ = 'certs.certs_pb2'
  # @@protoc_insertion_point(class_scope:ZCertAttr)
  ))
_sym_db.RegisterMessage(ZCertAttr)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
