package model

import (
	base "project/src/pkg/model"
)

type Ping struct {
	Base    base.Base `json:"base"`
	Message string    `json:"message"`
}
