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

El ejemplo que tendremos es apegado a un escenario de la vida real, en este caso tendremos una red de microservicios para una aplicación que administra clínicas médicas. La idea es tener separado el microservicio de los doctores y el de las citas médicas. 

En algún punto del desarrollo de la aplicación, necesitaremos que el microservicio de los doctores se comunique con el microservicio de las citas médicas para obtener información sobre las citas de un doctor en específico. Esto con el objetivo de que no tengamos código duplicado y que los microservicios sean independientes entre sí. 

Para esto, usaremos gRPC para la comunicación entre los microservicios, crearemos un microservicio en Python usando FastAPI y otro en Go usando Fiber. Ambos endpoints implementarán gRPC y REST.


## Instalaciones globales necesarias

Instalación de Python:
https://www.python.org/downloads/

Instalación de Go:
https://golang.org/doc/install

Documentación oficial de gRPC:
https://grpc.io/

Documentación oficial de Protocol Buffers:
https://protobuf.dev/

## Levantar imagen de mongoDB en un docker-compose
```bash
docker compose up -d
```

## Servicio de doctores
El servicio de doctores estará hecho en Golang y usará el framework Fiber. Este servicio tendrá dos endpoints, uno para gRPC y otro para REST.

Para iniciar crearemos un directorio llamado `go-service` y dentro de este crearemos un archivo `main.go` con el siguiente contenido:

```bash
mkadir go-service
cd go-service
```

Creamos un archivo `go.mod` con el nombre del módulo y las dependencias necesarias:

```bash
go mod init <nombre-del-modulo>/go-service
```

Instalamos las dependencias globales necesarias para gRPC y Protocol Buffers, luego activamos el path de binarios de Go:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH="$PATH:$(go env GOPATH)/bin" # linux
$env:PATH += ";$(go env GOPATH)/bin" # windows
```

Instalar estas depencias en caso de que no estén instaladas:

```bash
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get google.golang.org/grpc
go get github.com/gofiber/fiber/v2
```

Creamos un archivo directorio llamado `proto` y dentro de este un archivo `doctor.proto` con el siguiente contenido:

```proto