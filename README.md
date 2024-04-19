# gRCP (Remote Procedure Call)

## ¿Qué es RPC?
Es un protocolo que permite a un programa corriendo en una computadora llamar a un procedimiento o función que se ejecuta en otra computadora a través de la red.

- **Transparencia**: En este caso la computadora cliente puede llamar una función al igual que llamaría una función de manera local, pero sin la necesidad de preocuparse de los detalles de la red.
- **Ejecución remota**: La ejecución de la función no ocurre en la computadora cliente, sino que en el servidor.
- **Protocolo de comunicación**:  RCP define un estándar para el cliente y el servidor de como comunicarse.


## ¿Ahora qué es gRCP?
Es un framework de código abierto desarrollado por Google que permite a los desarrolladores crear aplicaciones cliente-servidor de manera sencilla y eficiente. gRPC utiliza el protocolo HTTP/2 para la comunicación. 