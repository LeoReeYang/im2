package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Friends  uint64 `json:"friends"`
	Block    uint64 `json:"block"`
}
type Friends struct {
	gorm.Model
	Uid     uint64   `json:"uid"`
	Friends []uint64 `json:"friends"`
}

type Blocks struct {
	gorm.Model
	Uid    uint64   `json:"uid"`
	Blocks []uint64 `json:"block"`
}
