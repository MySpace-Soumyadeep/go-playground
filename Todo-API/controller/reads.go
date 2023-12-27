package controller

import (
	"encoding/json"
	"net/http"
	"test/model"
)

func ReadAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			//if name not entered---> read all entries
			if name == "" || len(name) == 0 {
				data, err := model.ReadAll()
				if err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			} else {
				//if name  entered---> read by name
				data, err := model.ReadByName(name)
				if err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			}
		}
	}
}

// func ReadByName() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			name := r.URL.Query().Get("name")
// 			if name == "" || len(name) == 0 {
// 				w.Write([]byte("No name entered!!!"))
// 				return
// 			} else {
// 				data, err := model.ReadByName(name)
// 				if err != nil {
// 					w.Write([]byte(err.Error()))
// 					return
// 				}
// 				w.WriteHeader(http.StatusOK)
// 				json.NewEncoder(w).Encode(data)
// 			}
// 		}
// 	}
// }
