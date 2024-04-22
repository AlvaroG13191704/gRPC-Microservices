import logging
import appointment_pb2_grpc
import appointment_pb2
import grpc
from concurrent import futures

# Instanciamos la clase que se creo en el archivo appointment_pb2_grpc.py
class AppointmentService(appointment_pb2_grpc.AppointmentsServiceServicer):
  def GetAppointments(self, request, context):

    # imprimir el request
    print("request: ", request)

    return appointment_pb2.GetAppointmentsResponse(appointments=[
      appointment_pb2.Appointment(
        id=1,
        doctor_id="1",
        patient_name="Juan Perez",
        description="Dolor de cabeza"
      )
    ])



def server():
  port = "3001"
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  appointment_pb2_grpc.add_AppointmentsServiceServicer_to_server(AppointmentService(), server)
  server.add_insecure_port(f'[::]:{port}')

  server.start()
  print(f"Server running on port {port}")
  server.wait_for_termination()


if __name__ == "__main__":
  logging.basicConfig()
  server()