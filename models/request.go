package models

type Request struct {
	Status int
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateNameRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdatePasswordRequest struct {
}

type AddFriendRequest struct {
	FriendID uint `json:"friend_id" binding:"required"`
}
