package vitamin

import (
	"VitaminApp/cors"
	"VitaminApp/graph/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func SetupRoutes() {
	vitaminsHandler := http.HandlerFunc(vitaminsHandler) //remember, HandlerFunc is just a "Handler" that has the capability to call a function
	vitaminHandler := http.HandlerFunc(vitaminHandler)
	http.Handle("/api/vitamins", cors.MiddlewareHandler(vitaminsHandler))
	http.Handle("/api/vitamins/", cors.MiddlewareHandler(vitaminHandler))

}

func vitaminsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vitamins, err := GetVitaminList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		vitaminsJSON, err := json.Marshal(vitamins)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(vitaminsJSON)
	case http.MethodPost:
		//declare a variable with the struct type
		var newVitamin model.NewVitamin
		//read in the json payload coming in from the outside world
		vitaminBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//unmarshal the data
		err = json.Unmarshal(vitaminBytes, &newVitamin)
		//check for errors
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//call addVitamin
		err = AddVitamin(newVitamin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	case http.MethodOptions: //preflight request to allow the webservice to return the CORS specific headers to determine whether the browser should allow traffic to the webserver or not
		return
	}
}

func vitaminHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "vitamins/")
	vitaminId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		vitamin, err := getVitaminById(vitaminId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		vitaminJSON, err := json.Marshal(vitamin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(vitaminJSON)
	case http.MethodPut:
		var updatedVitamin Vitamin
		updatedBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(updatedBytes, &updatedVitamin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		updatedVitamin.VitaminId = vitaminId
		err = updateVitamin(updatedVitamin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodDelete:
		err := deleteVitamin(vitaminId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodOptions:
		return
	}

}
