package serializator

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func ReadJSON(r *http.Request, target any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(target)
	if err != nil {
		return err
	}

	return nil
}
