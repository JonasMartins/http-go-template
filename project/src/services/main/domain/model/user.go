package model

import (
	base "project/src/pkg/model"
)

type User struct {
	Base     base.Base `json:"base"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}
