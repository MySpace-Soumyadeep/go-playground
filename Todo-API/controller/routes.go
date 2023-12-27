package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/read", ReadAll())
	mux.HandleFunc("/delete/", DeleteTodo())
	return mux
}
