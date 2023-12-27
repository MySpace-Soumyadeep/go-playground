package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/model"
	"test/views"
)

func DeleteTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			name := r.URL.Path[8:]
			fmt.Println(name)
			if name == "" || len(name) == 0 {
				w.Write([]byte("No item deleted!!!"))
				return
			} else {
				err := model.DeleteTodo(name)
				if err != nil {
					w.Write([]byte("Some error"))
					return
				}
				w.WriteHeader(http.StatusOK)
				var data = views.DeleteResponse{Status: "Item Deleted"}
				json.NewEncoder(w).Encode(data)
			}
		}
	}
}
