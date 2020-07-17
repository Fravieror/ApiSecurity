package infrastructure

import (
	"apiSecurity/movement/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

// Save movement individualy
func SaveMovements(w http.ResponseWriter, r *http.Request) {
	// resp := rest.Get(fmt.Sprint("http://", os.Getenv("APIMOVEMENTS"), "/ping"))
	var movement domain.Movement
	err := json.NewDecoder(r.Body).Decode(&movement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error cast struct to json"))
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(r.Body).
		Post(fmt.Sprint("http://", os.Getenv("APIMOVEMENTS"), "/movement"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't save movements"))
		return
	}
	w.WriteHeader(resp.StatusCode())
	w.Write(resp.Body())
}

// Get all movements for client
func Movements(w http.ResponseWriter, r *http.Request) {
	cli, ok := r.URL.Query()["client"]
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message": "param client not sended"}`))
		return
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{"client": cli[0]}).
		Get(fmt.Sprint("http://", os.Getenv("APIMOVEMENTS"), "/movements"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't get movements"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Body())
}

// Get only one movement for id
func Movement(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message": "param client not sended"}`))
		return
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{"id": id[0]}).
		Get(fmt.Sprint("http://", os.Getenv("APIMOVEMENTS"), "/movement"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't get movements"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Body())
}
