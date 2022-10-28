package models

type User struct {
	Id      uint64 `json:"id"`
	Friends []uint64
	Block   []uint64
}
