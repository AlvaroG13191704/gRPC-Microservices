import logging
import appointment_pb2_grpc
import appointment_pb2
import grpc
from concurrent import futures
# plots
import pandas as pd
import matplotlib.pyplot as plt

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


def generatePlot():
  # Create a DataFrame
  data = {
    'Http': {
      'Time': [ 120, 99, 89, 89, 87],
      'Weight': [5.1, 4.3, 3.8, 3.5,3.5]
    },
    'gRPC': {
      'Time': [89, 67, 60, 61, 61],
      'Weight': [0.9, 1.5,1.4, 1.5,1.5] 
    }
  }
  # Convert to DataFrame
  http_df = pd.DataFrame(data['Http'])
  http_df['Protocol'] = 'Http'
  grpc_df = pd.DataFrame(data['gRPC'])
  grpc_df['Protocol'] = 'gRPC'
  df = pd.concat([http_df, grpc_df])

  # Create subplots
  fig, axs = plt.subplots(2)

  # Plot Time
  for name, group in df.groupby('Protocol'):
    axs[0].plot(group['Time'], marker='o', linestyle='-', ms=5, label=name)
  axs[0].set_xlabel('Tiempo en ms')
  axs[0].set_ylabel('')
  axs[0].legend()

  # Plot Weight
  for name, group in df.groupby('Protocol'):
      axs[1].plot(group['Weight'], marker='o', linestyle='-', ms=5, label=name)
  axs[1].set_xlabel('Peso en MB')
  axs[1].set_ylabel('')
  axs[1].legend()

  plt.tight_layout()
  plt.show()



if __name__ == "__main__":
  logging.basicConfig()
  generatePlot()
  server()
