package models

import (
	"time"
)

type Ping struct {
	Model
	AuthID      uint
	DateCreated time.Time
}

func (p *Ping) Add() error {
	p.DateCreated = time.Now()
	if err := db.Create(p).Error; err != nil {
		return err
	}
	return nil
}
