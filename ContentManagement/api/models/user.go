package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID
	FirstName  string
	LastName   string
	Email      string
	Created_at string
	Fullanme   string
	//TODO Create Role Concept
	//TODO Define Auth Process
}

func CreateUser(firstname string, lastname string, email string) *User {
	return &User{
		Id:         uuid.New(),
		FirstName:  firstname,
		LastName:   lastname,
		Email:      email,
		Fullanme:   firstname + " " + lastname,
		Created_at: time.Now().UTC().String(),
	}
}
