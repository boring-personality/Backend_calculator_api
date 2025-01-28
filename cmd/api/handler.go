package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Result int `json:"result"`
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func add(m *http.ServeMux) {

	m.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		payload := jsonResponse{
			Result: input1 + input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func subtract(m *http.ServeMux) {

	m.HandleFunc("POST /subtract", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		payload := jsonResponse{
			Result: input1 - input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func multiply(m *http.ServeMux) {

	m.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		payload := jsonResponse{
			Result: input1 * input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func divide(m *http.ServeMux) {

	m.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		payload := jsonResponse{
			Result: input1 / input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}
