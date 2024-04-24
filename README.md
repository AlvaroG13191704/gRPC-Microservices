# gRPC (Remote Procedure Call)

## ¿Qué es RPC?
Es un protocolo que permite a un programa corriendo en una computadora llamar a un procedimiento o función que se ejecuta en otra computadora a través de la red.

- **Transparencia**: En este caso la computadora cliente puede llamar una función al igual que llamaría una función de manera local, pero sin la necesidad de preocuparse de los detalles de la red.
- **Ejecución remota**: La ejecución de la función no ocurre en la computadora cliente, sino que en el servidor.
- **Protocolo de comunicación**:  RCP define un estándar para el cliente y el servidor de como comunicarse.


## ¿Ahora qué es gRPC?
Es un framework de código abierto desarrollado por Google que permite a los desarrolladores crear aplicaciones cliente-servidor de manera sencilla y eficiente. gRPC utiliza el protocolo HTTP/2 para la comunicación. 

gRCP usa Protocol Buffers (también conocido como protofub) que es una interfaz de lenguaje para definir la estructura de la data que se va a intercambiar entre el cliente y servidor.

### Puntos clave
- **Multiplataforma y agnóstico de lenguaje:** gRPC esta diseñado para que se trabaje a través de diferentes lenguajes de programación, por ejemplo: Java, C++, Python, Go, Ruby y muchos más. Esto permite la comunicación efectiva entre diferentes tecnologías. 
- **Eficiencia en la serialización de los datos:** al usar el protocolo de Buffers por defecto, el cual es mucho más eficiente y compacto que otros formatos como el JSON o XML. Esto permite que sea más rápida la transferencia al reducir el ancho de banda en la red.
- **Comunicación Bidireccional:** gRPC soporta diferentes tipos de métodos, como el unario, server-side streaming, client-side streaming y bidireccional. 
- **Http2/Transport:** gRPC esta construido con HTTP2/protocol, lo que le provee características como la multiplexación, compresión de headers y control de como fluye la data.

## ¿Por qué usar gRPC?
- **Rendimiento:** Por lo anteriormente descrito, con el uso de HTTP2, el protocolo de buffers hace que gRPC sea una opción bastante sólida para construir sistemas distribuidos, donde los volúmenes de transferencia de datos son más grandes y la latencia es sensible. 
- **Facilidad de uso:** gRPC permite a los desarrolladores definir los servicios y los mensajes en un archivo proto y luego generar el código fuente en diferentes lenguajes de programación.
- **Interoperabilidad:** gRPC soporta múltiples lenguajes de programación, lo que permite a los desarrolladores construir sistemas distribuidos con diferentes tecnologías.
- **Seguridad:** gRPC soporta la autenticación y la autorización, lo que permite a los desarrolladores asegurar sus servicios.
- **Escalabilidad:** gRPC soporta diferentes tipos de métodos, lo que permite a los desarrolladores construir sistemas distribuidos escalables.


# ¿Qué haremos en este repositorio?

Ente repositorio haremos un ejemplo simulando la vida real y unos benchmarks para comparar el rendimiento de gRPC con respecto a otros protocolos de comunicación como REST.


## Escenario de la vida real

El ejemplo que tendremos es apegado a un escenario de la vida real, en este caso tendremos una red de microservicios para una aplicación que administra clínicas médicas. La idea es tener un servicio de doctores y otro servicio de citas. 

En algún punto del desarrollo de la aplicación, necesitaremos que el microservicio de los doctores (cliente) se comunique con el microservicio de las citas médicas (server) para obtener información sobre las citas de un doctor en específico. Esto con el objetivo de que no tengamos código duplicado y que los microservicios sean independientes entre sí. 

Para esto, usaremos gRPC para la comunicación entre los servicios, crearemos el servicio de doctores (cliente) el cual se podrá comunicar con el servicio de citas médicas (server) para obtener la información necesaria usando gRPC.


## Instalaciones globales necesarias

Instalación de Python:
https://www.python.org/downloads/

Instalación de Go:
https://golang.org/doc/install

Documentación oficial de gRPC:
https://grpc.io/

Documentación oficial de Protocol Buffers:
https://protobuf.dev/

## Levantar imagen de mongoDB 
Para que este ejemplo sea más apegado a un entorno de producción, usaremos una base de datos no relacional como MongoDB. Para esto, usaremos Docker para levantar una instancia de MongoDB y poder conectarnos a ella desde nuestros servicios.
```bash
docker compose up -d
```

## Servicio de doctores (Cliente)
El servicio de doctores estará hecho en Golang y usará el framework Fiber. Este servicio tendrá dos endpoints, uno para gRPC y otro para REST. Además cuenta con un directorio proto el cual contiene el archivo appointment.proto que define la estructura de los mensajes que se intercambiarán entre el cliente y el servidor.

Estas son las depenencias necesarioas a instalar para el servicio de doctores:

```bash
# grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
# Fiber v2
go get github.com/gofiber/fiber/v2
# Activar el path de go
export PATH="$PATH:$(go env GOPATH)/bin" # linux
$env:PATH += ";$(go env GOPATH)/bin" # windows
```


Para poder generar el código compilado del archivo `.protto` se debe ejecutar el siguiente comando dentro de la carpeta `proto`. Los archivos generados son `appointment.pb.go` y `appointment_grpc.pb.go`. El primero contiene la estructura de los mensajes y el segundo contiene la definición de los servicios.
```bash
cd proto
protoc --go_out=. --go-grpc_out=. appointment.proto
```

## Servicio de citas médicas con python (Server)

El servicio de citas médicas es mucho más simple ya que solo tendrá la funcionalidad de escuchar las peticiones del cliente y devolver la información necesaria. Este servicio estará hecho en Python en un entorno virtual.

Para crear el entorno virtual y activarlo, se deben ejecutar los siguientes comandos:

```bash
python -m venv env
source env/bin/activate # linux
env/Scripts/activate # windows
```

Luego se deben instalar las dependencias necesarias para el servicio de citas médicas:

```bash
pip install grpcio
pip install grpcio-tools
pip install pymongo
```

Al igual que el cliente de doctores, se debe generar el código compilado del archivo `.proto` para poder usarlo en el servidor. Para esto, se debe ejecutar el siguiente comando en la raíz del proyecto de python. Los archivos generados son `appointment_pb2.py` y `appointment_pb2_grpc.py`. El primero contiene la estructura de los mensajes y el segundo contiene la definición de los servicios.

```bash
python -m grpc_tools.protoc -I./proto --python_out=./ --pyi_out=./ --grpc_python_out=./ ./proto/appointment.proto   
```

## Servicio de citas médicas con Golang (Server)

El servicio de citas médicas en Golang es similar al servicio de doctores, solo que este solo tendrá la funcionalidad de escuchar las peticiones del cliente y devolver la información necesaria. Este servicio estará hecho en Golang.

Para instalar las dependencias necesarias para el servicio de citas médicas en Golang, se deben ejecutar los siguientes comandos:

```bash
go get google.golang.org/grpc
go get go.mongodb.org/mongo-driver/mongo
```

### Levantar los servicios

Para levantar tanto el cliente, como el servidor hecho con gRPC y REST, se deben ejecutar los siguientes comandos en la raíz de cada proyecto:

**Es importante que el contenedor de mongo este corriendo.**

```bash
go run main.go
```
