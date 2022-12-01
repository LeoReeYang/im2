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

func (cm *ConnectionMeta) Get(name string) *connection.Connection {
	cm.rmmutex.RLock()
	defer cm.rmmutex.RUnlock()

	if user, ok := cm.conns[name]; ok {
		return user
	}

	return nil
}

func (cm *ConnectionMeta) Put(c *connection.Connection) {
	cm.rmmutex.Lock()
	defer cm.rmmutex.Unlock()

	cm.conns[c.GetName()] = c
}

func (cm *ConnectionMeta) Remove(name string) {
	cm.rmmutex.Lock()
	defer cm.rmmutex.Unlock()

	delete(cm.conns, name)
}

func (cm *ConnectionMeta) All() []string {
	cm.rmmutex.RLock()
	defer cm.rmmutex.RUnlock()

	conns := []string{}
	for name := range cm.conns {
		conns = append(conns, name)
	}

	return conns
}

func (cm *ConnectionMeta) CheckConnections() {
	ticker := time.NewTicker(CheckPeriod)
	done := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			cm.rmmutex.Lock()
			for name, conn := range cm.conns {
				if conn.GetAlive() == connection.Offline {
					delete(cm.conns, name)
					color.HiMagenta("[ %s ] connection removed. ", name)
				}
			}
			cm.rmmutex.Unlock()
		case <-done:
			return
		}
	}
}
