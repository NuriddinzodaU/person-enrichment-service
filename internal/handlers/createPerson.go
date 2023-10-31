package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"person-service/internal/models"
	response "person-service/internal/models"
)

func (h Handler) CreatePerson() http.HandlerFunc {
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
		person.Age, err = getAge(person.Name)
		if err != nil {
			log.Print(err)
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		person.Gender, err = getGender(person.Name)
		if err != nil {
			log.Print(err)
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		person.Nationality, err = getNationality(person.Name)
		if err != nil {
			log.Print(err)
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}

		err = h.svc.RepositoryInstance().CreatePerson(&person)
		if err != nil {
			log.Print(err)
			response.ToJson(w, http.StatusBadRequest, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func getGender(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Gender string `json:"gender"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	return data.Gender, nil
}

func getNationality(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Country []struct {
			CountryID   string  `json:"country_id"`
			Probability float64 `json:"probability"`
		} `json:"country"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	if len(data.Country) > 0 {
		return data.Country[0].CountryID, nil
	}

	return "", nil
}

func getAge(name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data struct {
		Age int `json:"age"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	return data.Age, nil
}
