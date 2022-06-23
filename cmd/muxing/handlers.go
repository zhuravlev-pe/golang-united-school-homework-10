package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func nameGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello, %s!", vars["PARAM"])
	if err != nil {
		log.Fatalln(err)
	}
}

func badGet(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func dataPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(w, "I got message:\n")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = w.Write(body)
}

func getNumericHeader(r *http.Request, key string) (int, error) {
	value := r.Header.Get(key)
	if value == "" {
		return 0, fmt.Errorf("header %q is missing", key)
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func headerGet(w http.ResponseWriter, r *http.Request) {
	a, err := getNumericHeader(r, "a")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, err := getNumericHeader(r, "b")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("a+b", fmt.Sprint(a+b))
	w.WriteHeader(http.StatusOK)
}
