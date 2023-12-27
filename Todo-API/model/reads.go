package model

import (
	"fmt"
	"test/views"
)

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	// defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?", name)
	// defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
