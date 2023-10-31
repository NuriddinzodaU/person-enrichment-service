package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	response "person-service/internal/models"
	"strconv"
)

func (h Handler) GetAllPerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vCount := mux.Vars(r)["count"]
		vPage := mux.Vars(r)["page"]

		count, err := strconv.Atoi(vCount)
		if err != nil {
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}

		page, err := strconv.Atoi(vPage)
		if err != nil {
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}

		resp, err := h.svc.RepositoryInstance().GetAllPerson(count, page)
		if err != nil {
			log.Print(err)
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		response.ToJson(w, http.StatusOK, resp)
	}
}
