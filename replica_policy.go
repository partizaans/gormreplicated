package gormreplicated

import (
	"gorm.io/gorm"
)

// ReplicaPolicy chooses the first available connection pool from given pools
type ReplicaPolicy struct {
}

type pinger interface {
	Ping() error
}

func (ReplicaPolicy) Resolve(connPools []gorm.ConnPool) gorm.ConnPool {
	for _, conn := range connPools {
		thePinger, ok := conn.(pinger)
		if !ok {
			continue
		}
		if err := thePinger.Ping(); err != nil {
			continue
		}
		return conn
	}
	return nil
}
