package main

import (
	"go-api-template/internal/core/infra/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Cargar la conexión a la base de datos
	err := config.DbConection()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Configurar el enrutador principal
	router := mux.NewRouter()
	router.HandleFunc("/", HandleInit).Methods("GET")

	// Configurar las rutas de usuario

	// Iniciar el servidor
	log.Println("Servidor levantado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HandleInit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
