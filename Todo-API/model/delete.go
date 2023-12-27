package model

import "fmt"

func DeleteTodo(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
