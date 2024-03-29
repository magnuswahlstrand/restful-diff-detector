# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='hello.proto',
  package='hello',
  syntax='proto3',
  serialized_pb=_b('\n\x0bhello.proto\x12\x05hello\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1d\n\nHelloReply\x12\x0f\n\x07message\x18\x01 \x01(\t\"(\n\x10\x44iffNotification\x12\x14\n\x0cresponseData\x18\x01 \x01(\t\"C\n\rDiffSubscribe\x12\x0c\n\x04path\x18\x01 \x01(\t\x12\x0e\n\x06period\x18\x02 \x01(\x05\x12\x14\n\x0csubscriberId\x18\x03 \x01(\t2?\n\x07Greeter\x12\x34\n\x08SayHello\x12\x13.hello.HelloRequest\x1a\x11.hello.HelloReply\"\x00\x32\x94\x01\n\x0e\x44iffSubscriber\x12<\n\tSubscribe\x12\x14.hello.DiffSubscribe\x1a\x17.hello.DiffNotification\"\x00\x12\x44\n\x0fSubscribeStream\x12\x14.hello.DiffSubscribe\x1a\x17.hello.DiffNotification\"\x00\x30\x01\x42\x30\n\x1bio.grpc.examples.helloworldB\x0fHelloWorldProtoP\x01\x62\x06proto3')
)




_HELLOREQUEST = _descriptor.Descriptor(
  name='HelloRequest',
  full_name='hello.HelloRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='hello.HelloRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=22,
  serialized_end=50,
)


_HELLOREPLY = _descriptor.Descriptor(
  name='HelloReply',
  full_name='hello.HelloReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='hello.HelloReply.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=52,
  serialized_end=81,
)


_DIFFNOTIFICATION = _descriptor.Descriptor(
  name='DiffNotification',
  full_name='hello.DiffNotification',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='responseData', full_name='hello.DiffNotification.responseData', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=83,
  serialized_end=123,
)


_DIFFSUBSCRIBE = _descriptor.Descriptor(
  name='DiffSubscribe',
  full_name='hello.DiffSubscribe',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='path', full_name='hello.DiffSubscribe.path', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='period', full_name='hello.DiffSubscribe.period', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subscriberId', full_name='hello.DiffSubscribe.subscriberId', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=125,
  serialized_end=192,
)

DESCRIPTOR.message_types_by_name['HelloRequest'] = _HELLOREQUEST
DESCRIPTOR.message_types_by_name['HelloReply'] = _HELLOREPLY
DESCRIPTOR.message_types_by_name['DiffNotification'] = _DIFFNOTIFICATION
DESCRIPTOR.message_types_by_name['DiffSubscribe'] = _DIFFSUBSCRIBE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

HelloRequest = _reflection.GeneratedProtocolMessageType('HelloRequest', (_message.Message,), dict(
  DESCRIPTOR = _HELLOREQUEST,
  __module__ = 'hello_pb2'
  # @@protoc_insertion_point(class_scope:hello.HelloRequest)
  ))
_sym_db.RegisterMessage(HelloRequest)

HelloReply = _reflection.GeneratedProtocolMessageType('HelloReply', (_message.Message,), dict(
  DESCRIPTOR = _HELLOREPLY,
  __module__ = 'hello_pb2'
  # @@protoc_insertion_point(class_scope:hello.HelloReply)
  ))
_sym_db.RegisterMessage(HelloReply)

DiffNotification = _reflection.GeneratedProtocolMessageType('DiffNotification', (_message.Message,), dict(
  DESCRIPTOR = _DIFFNOTIFICATION,
  __module__ = 'hello_pb2'
  # @@protoc_insertion_point(class_scope:hello.DiffNotification)
  ))
_sym_db.RegisterMessage(DiffNotification)

DiffSubscribe = _reflection.GeneratedProtocolMessageType('DiffSubscribe', (_message.Message,), dict(
  DESCRIPTOR = _DIFFSUBSCRIBE,
  __module__ = 'hello_pb2'
  # @@protoc_insertion_point(class_scope:hello.DiffSubscribe)
  ))
_sym_db.RegisterMessage(DiffSubscribe)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\033io.grpc.examples.helloworldB\017HelloWorldProtoP\001'))

_GREETER = _descriptor.ServiceDescriptor(
  name='Greeter',
  full_name='hello.Greeter',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=194,
  serialized_end=257,
  methods=[
  _descriptor.MethodDescriptor(
    name='SayHello',
    full_name='hello.Greeter.SayHello',
    index=0,
    containing_service=None,
    input_type=_HELLOREQUEST,
    output_type=_HELLOREPLY,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_GREETER)

DESCRIPTOR.services_by_name['Greeter'] = _GREETER


_DIFFSUBSCRIBER = _descriptor.ServiceDescriptor(
  name='DiffSubscriber',
  full_name='hello.DiffSubscriber',
  file=DESCRIPTOR,
  index=1,
  options=None,
  serialized_start=260,
  serialized_end=408,
  methods=[
  _descriptor.MethodDescriptor(
    name='Subscribe',
    full_name='hello.DiffSubscriber.Subscribe',
    index=0,
    containing_service=None,
    input_type=_DIFFSUBSCRIBE,
    output_type=_DIFFNOTIFICATION,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SubscribeStream',
    full_name='hello.DiffSubscriber.SubscribeStream',
    index=1,
    containing_service=None,
    input_type=_DIFFSUBSCRIBE,
    output_type=_DIFFNOTIFICATION,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DIFFSUBSCRIBER)

DESCRIPTOR.services_by_name['DiffSubscriber'] = _DIFFSUBSCRIBER

# @@protoc_insertion_point(module_scope)
