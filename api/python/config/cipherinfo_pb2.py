# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cipherinfo.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='cipherinfo.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\037com.zededa.cloud.uservice.protoZ$github.com/lf-edge/eve/api/go/config'),
  serialized_pb=_b('\n\x10\x63ipherinfo.proto\"\xc1\x01\n\nCipherInfo\x12\n\n\x02id\x18\x64 \x01(\t\x12-\n\x11keyExchangeScheme\x18\x01 \x01(\x0e\x32\x12.KeyExchangeScheme\x12+\n\x10\x65ncryptionScheme\x18\x02 \x01(\x0e\x32\x11.EncryptionScheme\x12\x14\n\x0cinitialValue\x18\x03 \x01(\x0c\x12\x12\n\npublicCert\x18\x04 \x01(\x0c\x12\x0e\n\x06sha256\x18\x05 \x01(\t\x12\x11\n\tsignature\x18\x06 \x01(\x0c*/\n\x11KeyExchangeScheme\x12\x0c\n\x08KEA_NONE\x10\x00\x12\x0c\n\x08KEA_ECDH\x10\x01*3\n\x10\x45ncryptionScheme\x12\x0b\n\x07SA_NONE\x10\x00\x12\x12\n\x0eSA_AES_256_CFB\x10\x01\x42G\n\x1f\x63om.zededa.cloud.uservice.protoZ$github.com/lf-edge/eve/api/go/configb\x06proto3')
)

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
  serialized_start=216,
  serialized_end=263,
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
  serialized_start=265,
  serialized_end=316,
)
_sym_db.RegisterEnumDescriptor(_ENCRYPTIONSCHEME)

EncryptionScheme = enum_type_wrapper.EnumTypeWrapper(_ENCRYPTIONSCHEME)
KEA_NONE = 0
KEA_ECDH = 1
SA_NONE = 0
SA_AES_256_CFB = 1



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
  serialized_start=21,
  serialized_end=214,
)

_CIPHERINFO.fields_by_name['keyExchangeScheme'].enum_type = _KEYEXCHANGESCHEME
_CIPHERINFO.fields_by_name['encryptionScheme'].enum_type = _ENCRYPTIONSCHEME
DESCRIPTOR.message_types_by_name['CipherInfo'] = _CIPHERINFO
DESCRIPTOR.enum_types_by_name['KeyExchangeScheme'] = _KEYEXCHANGESCHEME
DESCRIPTOR.enum_types_by_name['EncryptionScheme'] = _ENCRYPTIONSCHEME
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CipherInfo = _reflection.GeneratedProtocolMessageType('CipherInfo', (_message.Message,), dict(
  DESCRIPTOR = _CIPHERINFO,
  __module__ = 'cipherinfo_pb2'
  # @@protoc_insertion_point(class_scope:CipherInfo)
  ))
_sym_db.RegisterMessage(CipherInfo)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
