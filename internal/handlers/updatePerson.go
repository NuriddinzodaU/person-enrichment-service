package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"person-service/internal/models"
	response "person-service/internal/models"
)

func (h Handler) UpdatePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			person models.Person
		)
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			log.Print(err)
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}

		err = h.svc.RepositoryInstance().UpdatePerson(&person)
		if err != nil {
			log.Print(err)
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
