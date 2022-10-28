package store

import (
	"im/models"
	"sync"
)

type UserStore struct {
	users map[uint64]*models.User
	uLock sync.RWMutex
}

func (us *UserStore) NewUserStore() *UserStore {
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
