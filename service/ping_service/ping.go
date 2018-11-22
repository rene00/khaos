package ping_service

import (
	"github.com/rene00/khaos/models"
)

type Ping struct {
	Username string
}

func (p *Ping) Add() error {
	ping := map[string]interface{}{
		"username": p.Username,
	}
	if err := models.AddPing(ping); err != nil {
		return err
	}
	return nil
}
