package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Message string `json:"message,omitempty"`
	Result  int    `json:"result,omitempty"`
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

func (c *Config) indexHandler(m *http.ServeMux) {
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload := jsonResponse{
			Message: "Welcome to Calculator APP",
		}
		writeJSON(w, http.StatusOK, payload)
	})
}

func (c *Config) addHandler(m *http.ServeMux) {

	m.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		payload := jsonResponse{
			Result: input1 + input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) subtractHandler(m *http.ServeMux) {

	m.HandleFunc("POST /subtract", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		payload := jsonResponse{
			Result: input1 - input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) multiplyHandler(m *http.ServeMux) {

	m.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		payload := jsonResponse{
			Result: input1 * input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}

func (c *Config) divideHandler(m *http.ServeMux) {

	m.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		input1, _ := strconv.Atoi(r.URL.Query().Get("input1"))
		input2, _ := strconv.Atoi(r.URL.Query().Get("input2"))
		payload := jsonResponse{
			Result: input1 / input2,
		}
		writeJSON(w, http.StatusOK, payload)
	})

}
