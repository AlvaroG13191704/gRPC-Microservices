from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Http(_message.Message):
    __slots__ = ("Time", "Weight")
    TIME_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_FIELD_NUMBER: _ClassVar[int]
    Time: _containers.RepeatedScalarFieldContainer[float]
    Weight: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, Time: _Optional[_Iterable[float]] = ..., Weight: _Optional[_Iterable[float]] = ...) -> None: ...

class Grpc(_message.Message):
    __slots__ = ("Time", "Weight")
    TIME_FIELD_NUMBER: _ClassVar[int]
    WEIGHT_FIELD_NUMBER: _ClassVar[int]
    Time: _containers.RepeatedScalarFieldContainer[float]
    Weight: _containers.RepeatedScalarFieldContainer[float]
    def __init__(self, Time: _Optional[_Iterable[float]] = ..., Weight: _Optional[_Iterable[float]] = ...) -> None: ...

class benchmarkRequest(_message.Message):
    __slots__ = ("http", "grpc")
    HTTP_FIELD_NUMBER: _ClassVar[int]
    GRPC_FIELD_NUMBER: _ClassVar[int]
    http: Http
    grpc: Grpc
    def __init__(self, http: _Optional[_Union[Http, _Mapping]] = ..., grpc: _Optional[_Union[Grpc, _Mapping]] = ...) -> None: ...

class benchmarkResponse(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...
