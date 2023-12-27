package controller

import (
	"encoding/json"
	"net/http"
	"test/model"
	"test/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// take some data
			// save it
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			// fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some Error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
