package model

import (
	"time"

	"fbc-bookings/internal/domain/dto"
)

type User struct {
	ID           int
	DocumentType string
	DocumentID   string
	Names        string
	Surnames     string
	Birthdate    time.Time
	Age          int
	Gender       string
	Email        string
	Phone        string
}

func (*User) TableName() string {
	return "users"
}

func (u *User) ToDomainDTO() dto.User {
	return dto.User{
		ID:           u.ID,
		DocumentType: u.DocumentType,
		DocumentID:   u.DocumentID,
		Names:        u.Names,
		Surnames:     u.Surnames,
		Birthdate:    u.Birthdate,
		Age:          u.Age,
		Gender:       u.Gender,
		Email:        u.Email,
		Phone:        u.Phone,
	}
}
