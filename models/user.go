package models

type User struct {
	Id      uint64 `json:"id"`
	Friends uint64
	Block   uint64
}
type Friends struct {
	Uid     uint64 `json:"uid"`
	Friends []uint64
}

type Blocks struct {
	Uid    uint64 `json:"uid"`
	Blocks []uint64
}
