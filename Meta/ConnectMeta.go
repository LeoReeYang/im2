package Meta

import (
	"github.com/fatih/color"

	"sync"
	"time"

	"github.com/LeoReeYang/im2/connection"
)

const CheckPeriod = time.Second * 5

type ConnectionMeta struct {
	conns   map[string]*connection.Connection
	rmmutex sync.RWMutex
}

func NewConnectionMeta() *ConnectionMeta {
	return &ConnectionMeta{
		conns: make(map[string]*connection.Connection),
	}
}

func (um *ConnectionMeta) Get(name string) *connection.Connection {
	um.rmmutex.RLock()
	defer um.rmmutex.RUnlock()

	if user, ok := um.conns[name]; ok {
		return user
	}

	return nil
}

func (um *ConnectionMeta) Put(c *connection.Connection) {
	um.rmmutex.Lock()
	defer um.rmmutex.Unlock()

	um.conns[c.GetName()] = c
}

func (um *ConnectionMeta) Remove(name string) {
	um.rmmutex.Lock()
	defer um.rmmutex.Unlock()

	delete(um.conns, name)
}

func (um *ConnectionMeta) GetAll() []string {
	um.rmmutex.RLock()
	defer um.rmmutex.RUnlock()

	conns := []string{}
	for name := range um.conns {
		conns = append(conns, name)
	}

	return conns
}

func (um *ConnectionMeta) CheckConnections() {
	ticker := time.NewTicker(CheckPeriod)
	done := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			um.rmmutex.Lock()
			for name, conn := range um.conns {
				if conn.GetAlive() == connection.Offline {
					delete(um.conns, name)
					color.HiMagenta("[ %s ] connection removed. ", name)
				}
			}
			um.rmmutex.Unlock()
		case <-done:
			return
		}
	}
}
