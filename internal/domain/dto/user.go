package dto

import "time"

type User struct {
	ID           int       `json:"id"`
	DocumentType string    `json:"document_type" validate:"required"`
	DocumentID   string    `json:"document_id" validate:"required,number"`
	Names        string    `json:"names" validate:"required"`
	Surnames     string    `json:"surnames" validate:"required"`
	Birthdate    time.Time `json:"birthdate" validate:"required"`
	Age          int
	Gender       string `json:"gender" validate:"required"`
	Email        string `json:"email" validate:"required_if=Age 18,email"`
	Phone        string `json:"phone" validate:"required_if=Age 18,required,number"`
}
