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

type UserStatus string

const (
	Enabled  UserStatus = "enabled"
	Disabled UserStatus = "disabled"
)

func (s UserStatus) String() string {
	switch s {
	case Enabled:
		return "enabled"
	case Disabled:
		return "disabled"
	}
	return "unknown"
}
