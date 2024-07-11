package main

import (
	"backend/app"
	"backend/clients"
	"log"
)

func main() {
    // Inicializa la base de datos
    clients.InitDB()

    // Configura y corre el servidor
    router := app.SetupRouter()
    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
