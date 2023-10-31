package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	response "person-service/internal/models"
	"strconv"
)

func (h Handler) GetPerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]

		vID, err := strconv.Atoi(id)
		if err != nil {
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		resp, err := h.svc.RepositoryInstance().GetPerson(int64(vID))
		if err != nil {
			log.Print(err)
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		response.ToJson(w, http.StatusOK, resp)
	}
}
