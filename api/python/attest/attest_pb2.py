# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: attest.proto

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
  name='attest.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\037com.zededa.cloud.uservice.protoZ#github.com/lf-edge/eve/api/go/certs'),
  serialized_pb=_b('\n\x0c\x61ttest.proto\"f\n\nZAttestReq\x12 \n\x07reqType\x18\x01 \x01(\x0e\x32\x0f.ZAttestReqType\x12\x1c\n\x05quote\x18\x02 \x01(\x0b\x32\r.ZAttestQuote\x12\x18\n\x05\x63\x65rts\x18\x03 \x03(\x0b\x32\t.ZEveCert\"}\n\x0fZAttestResponse\x12\"\n\x08respType\x18\x01 \x01(\x0e\x32\x10.ZAttestRespType\x12 \n\x05nonce\x18\x02 \x01(\x0b\x32\x11.ZAttestNonceResp\x12$\n\tquoteResp\x18\x03 \x01(\x0b\x32\x11.ZAttestQuoteResp\"!\n\x10ZAttestNonceResp\x12\r\n\x05nonce\x18\x01 \x01(\x0c\"5\n\x0cZAttestQuote\x12\x12\n\nattestData\x18\x01 \x01(\x0c\x12\x11\n\tsignature\x18\x02 \x01(\x0c\":\n\x10ZAttestQuoteResp\x12&\n\x08response\x18\x01 \x01(\x0e\x32\x14.ZAttestResponseCode\"\x8f\x01\n\x08ZEveCert\x12#\n\x08hashAlgo\x18\x01 \x01(\x0e\x32\x11.ZEveCertHashType\x12\x10\n\x08\x63\x65rtHash\x18\x02 \x01(\x0c\x12\x1b\n\x04type\x18\x03 \x01(\x0e\x32\r.ZEveCertType\x12\x0c\n\x04\x63\x65rt\x18\x04 \x01(\x0c\x12!\n\nattributes\x18\x05 \x01(\x0b\x32\r.ZEveCertAttr\"!\n\x0cZEveCertAttr\x12\x11\n\tisMutable\x18\x01 \x01(\x08*Q\n\x0eZAttestReqType\x12\x13\n\x0f\x41TTEST_REQ_CERT\x10\x00\x12\x14\n\x10\x41TTEST_REQ_NONCE\x10\x01\x12\x14\n\x10\x41TTEST_REQ_QUOTE\x10\x02*Z\n\x0fZAttestRespType\x12\x14\n\x10\x41TTEST_RESP_CERT\x10\x00\x12\x15\n\x11\x41TTEST_RESP_NONCE\x10\x01\x12\x1a\n\x16\x41TTEST_RESP_QUOTE_RESP\x10\x02*O\n\x13ZAttestResponseCode\x12\x1b\n\x17\x41TTEST_RESPONSE_SUCCESS\x10\x00\x12\x1b\n\x17\x41TTEST_RESPONSE_FAILURE\x10\x01*:\n\x10ZEveCertHashType\x12\r\n\tHASH_NONE\x10\x00\x12\x17\n\x13HASH_SHA256_16bytes\x10\x01*\xa2\x01\n\x0cZEveCertType\x12\x1f\n\x1b\x43\x45RT_TYPE_DEVICE_ONBOARDING\x10\x00\x12\'\n#CERT_TYPE_DEVICE_RESTRICTED_SIGNING\x10\x01\x12$\n CERT_TYPE_DEVICE_ENDORSEMENT_RSA\x10\x02\x12\"\n\x1e\x43\x45RT_TYPE_DEVICE_ECDH_EXCHANGE\x10\x03\x42\x46\n\x1f\x63om.zededa.cloud.uservice.protoZ#github.com/lf-edge/eve/api/go/certsb\x06proto3')
)

_ZATTESTREQTYPE = _descriptor.EnumDescriptor(
  name='ZAttestReqType',
  full_name='ZAttestReqType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ATTEST_REQ_CERT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ATTEST_REQ_NONCE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ATTEST_REQ_QUOTE', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=578,
  serialized_end=659,
)
_sym_db.RegisterEnumDescriptor(_ZATTESTREQTYPE)

ZAttestReqType = enum_type_wrapper.EnumTypeWrapper(_ZATTESTREQTYPE)
_ZATTESTRESPTYPE = _descriptor.EnumDescriptor(
  name='ZAttestRespType',
  full_name='ZAttestRespType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ATTEST_RESP_CERT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ATTEST_RESP_NONCE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ATTEST_RESP_QUOTE_RESP', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=661,
  serialized_end=751,
)
_sym_db.RegisterEnumDescriptor(_ZATTESTRESPTYPE)

ZAttestRespType = enum_type_wrapper.EnumTypeWrapper(_ZATTESTRESPTYPE)
_ZATTESTRESPONSECODE = _descriptor.EnumDescriptor(
  name='ZAttestResponseCode',
  full_name='ZAttestResponseCode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ATTEST_RESPONSE_SUCCESS', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ATTEST_RESPONSE_FAILURE', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=753,
  serialized_end=832,
)
_sym_db.RegisterEnumDescriptor(_ZATTESTRESPONSECODE)

ZAttestResponseCode = enum_type_wrapper.EnumTypeWrapper(_ZATTESTRESPONSECODE)
_ZEVECERTHASHTYPE = _descriptor.EnumDescriptor(
  name='ZEveCertHashType',
  full_name='ZEveCertHashType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HASH_NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HASH_SHA256_16bytes', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=834,
  serialized_end=892,
)
_sym_db.RegisterEnumDescriptor(_ZEVECERTHASHTYPE)

ZEveCertHashType = enum_type_wrapper.EnumTypeWrapper(_ZEVECERTHASHTYPE)
_ZEVECERTTYPE = _descriptor.EnumDescriptor(
  name='ZEveCertType',
  full_name='ZEveCertType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ONBOARDING', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_RESTRICTED_SIGNING', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ENDORSEMENT_RSA', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CERT_TYPE_DEVICE_ECDH_EXCHANGE', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=895,
  serialized_end=1057,
)
_sym_db.RegisterEnumDescriptor(_ZEVECERTTYPE)

ZEveCertType = enum_type_wrapper.EnumTypeWrapper(_ZEVECERTTYPE)
ATTEST_REQ_CERT = 0
ATTEST_REQ_NONCE = 1
ATTEST_REQ_QUOTE = 2
ATTEST_RESP_CERT = 0
ATTEST_RESP_NONCE = 1
ATTEST_RESP_QUOTE_RESP = 2
ATTEST_RESPONSE_SUCCESS = 0
ATTEST_RESPONSE_FAILURE = 1
HASH_NONE = 0
HASH_SHA256_16bytes = 1
CERT_TYPE_DEVICE_ONBOARDING = 0
CERT_TYPE_DEVICE_RESTRICTED_SIGNING = 1
CERT_TYPE_DEVICE_ENDORSEMENT_RSA = 2
CERT_TYPE_DEVICE_ECDH_EXCHANGE = 3



_ZATTESTREQ = _descriptor.Descriptor(
  name='ZAttestReq',
  full_name='ZAttestReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='reqType', full_name='ZAttestReq.reqType', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='quote', full_name='ZAttestReq.quote', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='certs', full_name='ZAttestReq.certs', index=2,
      number=3, type=11, cpp_type=10, label=3,
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
  serialized_start=16,
  serialized_end=118,
)


_ZATTESTRESPONSE = _descriptor.Descriptor(
  name='ZAttestResponse',
  full_name='ZAttestResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='respType', full_name='ZAttestResponse.respType', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='nonce', full_name='ZAttestResponse.nonce', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='quoteResp', full_name='ZAttestResponse.quoteResp', index=2,
      number=3, type=11, cpp_type=10, label=1,
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
  serialized_start=120,
  serialized_end=245,
)


_ZATTESTNONCERESP = _descriptor.Descriptor(
  name='ZAttestNonceResp',
  full_name='ZAttestNonceResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='nonce', full_name='ZAttestNonceResp.nonce', index=0,
      number=1, type=12, cpp_type=9, label=1,
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
  serialized_start=247,
  serialized_end=280,
)


_ZATTESTQUOTE = _descriptor.Descriptor(
  name='ZAttestQuote',
  full_name='ZAttestQuote',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='attestData', full_name='ZAttestQuote.attestData', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='signature', full_name='ZAttestQuote.signature', index=1,
      number=2, type=12, cpp_type=9, label=1,
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
  serialized_start=282,
  serialized_end=335,
)


_ZATTESTQUOTERESP = _descriptor.Descriptor(
  name='ZAttestQuoteResp',
  full_name='ZAttestQuoteResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='response', full_name='ZAttestQuoteResp.response', index=0,
      number=1, type=14, cpp_type=8, label=1,
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
  serialized_start=337,
  serialized_end=395,
)


_ZEVECERT = _descriptor.Descriptor(
  name='ZEveCert',
  full_name='ZEveCert',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hashAlgo', full_name='ZEveCert.hashAlgo', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='certHash', full_name='ZEveCert.certHash', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='ZEveCert.type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cert', full_name='ZEveCert.cert', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='attributes', full_name='ZEveCert.attributes', index=4,
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
  serialized_start=398,
  serialized_end=541,
)


_ZEVECERTATTR = _descriptor.Descriptor(
  name='ZEveCertAttr',
  full_name='ZEveCertAttr',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='isMutable', full_name='ZEveCertAttr.isMutable', index=0,
      number=1, type=8, cpp_type=7, label=1,
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
  serialized_start=543,
  serialized_end=576,
)

_ZATTESTREQ.fields_by_name['reqType'].enum_type = _ZATTESTREQTYPE
_ZATTESTREQ.fields_by_name['quote'].message_type = _ZATTESTQUOTE
_ZATTESTREQ.fields_by_name['certs'].message_type = _ZEVECERT
_ZATTESTRESPONSE.fields_by_name['respType'].enum_type = _ZATTESTRESPTYPE
_ZATTESTRESPONSE.fields_by_name['nonce'].message_type = _ZATTESTNONCERESP
_ZATTESTRESPONSE.fields_by_name['quoteResp'].message_type = _ZATTESTQUOTERESP
_ZATTESTQUOTERESP.fields_by_name['response'].enum_type = _ZATTESTRESPONSECODE
_ZEVECERT.fields_by_name['hashAlgo'].enum_type = _ZEVECERTHASHTYPE
_ZEVECERT.fields_by_name['type'].enum_type = _ZEVECERTTYPE
_ZEVECERT.fields_by_name['attributes'].message_type = _ZEVECERTATTR
DESCRIPTOR.message_types_by_name['ZAttestReq'] = _ZATTESTREQ
DESCRIPTOR.message_types_by_name['ZAttestResponse'] = _ZATTESTRESPONSE
DESCRIPTOR.message_types_by_name['ZAttestNonceResp'] = _ZATTESTNONCERESP
DESCRIPTOR.message_types_by_name['ZAttestQuote'] = _ZATTESTQUOTE
DESCRIPTOR.message_types_by_name['ZAttestQuoteResp'] = _ZATTESTQUOTERESP
DESCRIPTOR.message_types_by_name['ZEveCert'] = _ZEVECERT
DESCRIPTOR.message_types_by_name['ZEveCertAttr'] = _ZEVECERTATTR
DESCRIPTOR.enum_types_by_name['ZAttestReqType'] = _ZATTESTREQTYPE
DESCRIPTOR.enum_types_by_name['ZAttestRespType'] = _ZATTESTRESPTYPE
DESCRIPTOR.enum_types_by_name['ZAttestResponseCode'] = _ZATTESTRESPONSECODE
DESCRIPTOR.enum_types_by_name['ZEveCertHashType'] = _ZEVECERTHASHTYPE
DESCRIPTOR.enum_types_by_name['ZEveCertType'] = _ZEVECERTTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ZAttestReq = _reflection.GeneratedProtocolMessageType('ZAttestReq', (_message.Message,), dict(
  DESCRIPTOR = _ZATTESTREQ,
  __module__ = 'attest_pb2'
  # @@protoc_insertion_point(class_scope:ZAttestReq)
  ))
_sym_db.RegisterMessage(ZAttestReq)

ZAttestResponse = _reflection.GeneratedProtocolMessageType('ZAttestResponse', (_message.Message,), dict(
  DESCRIPTOR = _ZATTESTRESPONSE,
  __module__ = 'attest_pb2'
  # @@protoc_insertion_point(class_scope:ZAttestResponse)
  ))
_sym_db.RegisterMessage(ZAttestResponse)

ZAttestNonceResp = _reflection.GeneratedProtocolMessageType('ZAttestNonceResp', (_message.Message,), dict(
  DESCRIPTOR = _ZATTESTNONCERESP,
  __module__ = 'attest_pb2'
  # @@protoc_insertion_point(class_scope:ZAttestNonceResp)
  ))
_sym_db.RegisterMessage(ZAttestNonceResp)

ZAttestQuote = _reflection.GeneratedProtocolMessageType('ZAttestQuote', (_message.Message,), dict(
  DESCRIPTOR = _ZATTESTQUOTE,
  __module__ = 'attest_pb2'
  # @@protoc_insertion_point(class_scope:ZAttestQuote)
  ))
_sym_db.RegisterMessage(ZAttestQuote)

ZAttestQuoteResp = _reflection.GeneratedProtocolMessageType('ZAttestQuoteResp', (_message.Message,), dict(
  DESCRIPTOR = _ZATTESTQUOTERESP,
  __module__ = 'attest_pb2'
  # @@protoc_insertion_point(class_scope:ZAttestQuoteResp)
  ))
_sym_db.RegisterMessage(ZAttestQuoteResp)

ZEveCert = _reflection.GeneratedProtocolMessageType('ZEveCert', (_message.Message,), dict(
  DESCRIPTOR = _ZEVECERT,
  __module__ = 'attest_pb2'
  # @@protoc_insertion_point(class_scope:ZEveCert)
  ))
_sym_db.RegisterMessage(ZEveCert)

ZEveCertAttr = _reflection.GeneratedProtocolMessageType('ZEveCertAttr', (_message.Message,), dict(
  DESCRIPTOR = _ZEVECERTATTR,
  __module__ = 'attest_pb2'
  # @@protoc_insertion_point(class_scope:ZEveCertAttr)
  ))
_sym_db.RegisterMessage(ZEveCertAttr)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
