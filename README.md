# Saludos en GO

## Instalación
Ejecuta el siguiente comando
```bash
go get -u github.com/twinoner/greetings
```

## Uso
como utilizar

```go
package main

import (
    "fmt"
    "github.com/twinoner/greetings"
)

func main() {
    message, err := greetings.Hello("Minguel")

    if err != nil {
        fmt.Println("ocurrio un error: ", err)
        return
    }

    fmt.Println(message)
}

```