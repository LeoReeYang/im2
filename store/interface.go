package store

import "github.com/LeoReeYang/im2/models"

type UserHandeler interface {
	Run()
	registe(*User)
	unregiste(*User)
	Transfer(*models.Message)
	GetAllUsers() []string
	PutUserToRegisterChannel(*User)
	PutUserToLeaveChannel(*User)
}
