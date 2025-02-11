package main

import (
	"net/http"
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

		var data jsonInput
		readJSON(w, r, &data)

		// validate input
		if !inputValidator(jsonInput{Input1: data.Input1, Input2: data.Input2}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}

		payload := jsonResponse{
			Result: data.Input1 + data.Input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) subtractHandler(m *http.ServeMux) {

	m.HandleFunc("POST /api/subtract", func(w http.ResponseWriter, r *http.Request) {

		var data jsonInput
		readJSON(w, r, &data)

		// validate input
		if !inputValidator(jsonInput{Input1: data.Input1, Input2: data.Input2}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}

		payload := jsonResponse{
			Result: data.Input1 - data.Input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) multiplyHandler(m *http.ServeMux) {

	m.HandleFunc("POST /api/multiply", func(w http.ResponseWriter, r *http.Request) {

		var data jsonInput
		readJSON(w, r, &data)

		// validate input
		if !inputValidator(jsonInput{Input1: data.Input1, Input2: data.Input2}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}

		payload := jsonResponse{
			Result: data.Input1 * data.Input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) divideHandler(m *http.ServeMux) {

	m.HandleFunc("POST /api/divide", func(w http.ResponseWriter, r *http.Request) {

		var data jsonInput
		readJSON(w, r, &data)

		// validate input
		if !inputValidator(jsonInput{Input1: data.Input1, Input2: data.Input2}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}

		payload := jsonResponse{
			Result: data.Input1 / data.Input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) sumHandler(m *http.ServeMux) {
	m.HandleFunc("POST /api/sum", func(w http.ResponseWriter, r *http.Request) {

		var data jsonArrayInput
		readJSON(w, r, &data)

		// validate input
		if !inputValidatorArray(jsonArrayInput{Inputarray: data.Inputarray}) {
			writeJSON(w, http.StatusBadRequest, jsonResponse{Message: "Input Validation Error"})
			return
		}
		result := 0
		for _, n := range data.Inputarray {
			result += n
		}
		payload := jsonResponse{
			Result: result,
		}
		writeJSON(w, http.StatusOK, payload)

	})
}
