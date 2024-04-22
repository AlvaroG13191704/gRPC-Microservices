from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Appointment(_message.Message):
    __slots__ = ("id", "doctor_id", "patient_name", "description")
    ID_FIELD_NUMBER: _ClassVar[int]
    DOCTOR_ID_FIELD_NUMBER: _ClassVar[int]
    PATIENT_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    id: int
    doctor_id: str
    patient_name: str
    description: str
    def __init__(self, id: _Optional[int] = ..., doctor_id: _Optional[str] = ..., patient_name: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class GetAppointmentsRequest(_message.Message):
    __slots__ = ("doctor_id",)
    DOCTOR_ID_FIELD_NUMBER: _ClassVar[int]
    doctor_id: str
    def __init__(self, doctor_id: _Optional[str] = ...) -> None: ...

class GetAppointmentsResponse(_message.Message):
    __slots__ = ("appointments",)
    APPOINTMENTS_FIELD_NUMBER: _ClassVar[int]
    appointments: _containers.RepeatedCompositeFieldContainer[Appointment]
    def __init__(self, appointments: _Optional[_Iterable[_Union[Appointment, _Mapping]]] = ...) -> None: ...
