import logging
import benchmark_pb2
import benchmark_pb2_grpc
import grpc
from concurrent import futures
# plots
import pandas as pd
import matplotlib.pyplot as plt



# Instancia del servicio
class BenchmarkService(benchmark_pb2_grpc.BenchmarkServicer):
  def Benchmark(self, request, context):
    print(f"Http Time received: {request.http.Time}")
    print(f"Http Weight received: {request.http.Weight}")
    print(f"gRPC Time received: {request.grpc.Time}")
    print(f"gRPC Weight received: {request.grpc.Weight}")

    # Convert repeated fields to lists
    http_time = list(request.http.Time)
    http_weight = list(request.http.Weight)
    grpc_time = list(request.grpc.Time)
    grpc_weight = list(request.grpc.Weight)

    # plot
    data = {
      'Http': {
        'Time': http_time,
        'Weight': http_weight
      },
      'gRPC': {
        'Time': grpc_time,
        'Weight': grpc_weight
      }
    }

    generatePlot(data)

    return benchmark_pb2.benchmarkResponse(
      message="Data processed successfully!"
    )
  

def server():
  port = "3003"
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  benchmark_pb2_grpc.add_BenchmarkServicer_to_server(BenchmarkService(), server)
  server.add_insecure_port(f'[::]:{port}')

  server.start()
  print(f"Server running on port {port}")
  server.wait_for_termination()


def generatePlot(data ):
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
    for i, time in enumerate(group['Time']):
      axs[0].text(i, time, str(time), ha='center')
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
  server()
