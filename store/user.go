package store

import (
	"sync"

	"github.com/LeoReeYang/im2/models"
)

type UserStore struct {
	users map[uint64]*models.User
	uLock sync.RWMutex
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[uint64]*models.User),
	}
}

func (us *UserStore) CheckUserExist(uid uint64) bool {
	us.uLock.RLock()
	defer us.uLock.RUnlock()

	_, exist := us.users[uid]
	return exist
}

func (us *UserStore) CheckBlock(uid, fuid uint64) {

}
