package models

import (
	"log"
	"time"
)

type Todo struct{
	ID int
	Content string
	UserId string
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string)(err error){
	cmd := `INSERT INTO todos (
	content,
	user_id,
	created_at
	) values ($1, $2, $3)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}