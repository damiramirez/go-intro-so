package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BodyRequest struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("GET /helloworld", HelloWorld)
	http.HandleFunc("GET /helloworld/{name}", HelloWorldQueryPath)
	http.HandleFunc("POST /helloworld", PostWorld)
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name := queryParams.Get("name")

	respuesta, err := json.Marshal(fmt.Sprintf("Hola %s! Como andas?", name))
	if err != nil {
		http.Error(w, "Error al codificar los datos como JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}


func HelloWorldQueryPath(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	respuesta, err := json.Marshal(fmt.Sprintf("Hola %s! Como andas?", name))
	if err != nil {
		http.Error(w, "Error al codificar los datos como JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

func PostWorld(w http.ResponseWriter, r *http.Request) {
	var request BodyRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respuesta, err := json.Marshal(fmt.Sprintf("Hola %s! Como andas?", request.Name))
	if err != nil {
		http.Error(w, "Error al codificar los datos como JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}

