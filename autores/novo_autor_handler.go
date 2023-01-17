package autores

import (
	"net/http"

	"jornada-dev-eficiente/casa-do-codigo/errs"
	"jornada-dev-eficiente/casa-do-codigo/serializator"
	"jornada-dev-eficiente/casa-do-codigo/validator"
)

func (h *autoresHandler) NovoAutor(w http.ResponseWriter, r *http.Request) {

	var request NovoAutorRequest
	err := serializator.ReadJSON(r, &request)
	if err != nil {
		errs.BadRequestError(w, err)
		return
	}

	v := validator.New()
	if request.Validate(v); !v.Valid() {
		errs.ValidationError(w, v.Errors)
		return
	}

	err = serializator.WriteJSON(w, http.StatusCreated, request)
	if err != nil {
		errs.ServerError(w, err)
	}
}
