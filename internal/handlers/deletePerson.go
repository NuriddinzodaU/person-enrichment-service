package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	response "person-service/internal/models"
	"strconv"
)

func (h Handler) DeletePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]

		vID, err := strconv.Atoi(id)
		if err != nil {
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		err = h.svc.RepositoryInstance().DeletePerson(vID)
		if err != nil {
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		response.ToJson(w, http.StatusOK, err)
	}
}
