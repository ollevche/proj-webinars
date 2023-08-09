package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting server")

	s := &Storage{}

	h := &cityHandler{
		s: s,
	}

	r := mux.NewRouter()

	r.Handle("/cities", auth(http.HandlerFunc(h.getCities))).
		Methods(http.MethodGet)

	r.Handle("/cities", auth(http.HandlerFunc(h.postCities))).
		Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}

type User struct {
	Username string
	Password string
}

var adminUser = User{
	Username: "admin",
	Password: "admin",
}

// Написати вебсервер, який по GET отримує 2 числа і знак (*/+-), а відповідає результатом цієї операції

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != adminUser.Username || password != adminUser.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type cityHandler struct {
	s *Storage
}

type singleCityReqBody struct {
	City string
}

func (h *cityHandler) postCities(w http.ResponseWriter, r *http.Request) {
	var reqBody singleCityReqBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.s.CreateCity(reqBody.City)
}

func (h *cityHandler) getCities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling request")

	jsonMessage, err := json.Marshal(h.s.GetCities())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonMessage)
}
