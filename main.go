package main

import (
	"jornada-dev-eficiente/casa-do-codigo/autores"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	autoresHandler := autores.NewAutoresHandler()

	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/api/autores", autoresHandler.NovoAutor)

	log.Println("Server started at por 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
