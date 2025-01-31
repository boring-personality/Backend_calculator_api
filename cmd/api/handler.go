package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (c *Config) indexHandler(m *http.ServeMux) {

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload := jsonResponse{
			Message: "Welcome to Calculator APP",
		}
		writeJSON(w, http.StatusOK, payload)
	})
}

func (c *Config) addHandler(m *http.ServeMux) {

	m.HandleFunc("POST /api/add", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		fmt.Println(r.Body)
		// validate input
		if !inputValidator(jsonInput{Input1: input1, Input2: input2}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}

		payload := jsonResponse{
			Result: input1 + input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) subtractHandler(m *http.ServeMux) {

	m.HandleFunc("POST /api/subtract", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))

		// validate input
		if !inputValidator(jsonInput{Input1: input1, Input2: input2}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}

		payload := jsonResponse{
			Result: input1 - input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) multiplyHandler(m *http.ServeMux) {

	m.HandleFunc("POST /api/multiply", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))

		// validate input
		if !inputValidator(jsonInput{Input1: input1, Input2: input2}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}

		payload := jsonResponse{
			Result: input1 * input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) divideHandler(m *http.ServeMux) {

	m.HandleFunc("POST /api/divide", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))

		// validate input
		if !inputValidator(jsonInput{Input1: input1, Input2: input2}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}

		payload := jsonResponse{
			Result: input1 / input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}
