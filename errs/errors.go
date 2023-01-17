package errs

import (
	"jornada-dev-eficiente/casa-do-codigo/serializator"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, status int, message any) {
	output := map[string]any{"Error": message}

	err := serializator.WriteJSON(w, status, output)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ValidationError(w http.ResponseWriter, errors map[string]string) {
	errorResponse(w, http.StatusUnprocessableEntity, errors)
}

func ServerError(w http.ResponseWriter, err error) {
	log.Println(err)

	message := "Ops, algo deu errado! Tente novamente mais tarde..."
	errorResponse(w, http.StatusInternalServerError, message)
}

func BadRequestError(w http.ResponseWriter, err error) {
	log.Println(err)

	message := "Não foi possível processar esta requisição"
	errorResponse(w, http.StatusBadRequest, message)
}
