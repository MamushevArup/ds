package service

import (
	"sync"
)

type P struct {
	// in memory storage where key is question and value is option

	poll      map[string]map[int]string
	createdBy map[string]string
	mu        sync.Mutex
}

func NewPoll() *P {
	return &P{
		poll:      make(map[string]map[int]string),
		createdBy: make(map[string]string),
	}
}

func (p *P) CreatePoll(id, question string, options map[int]string) error {
	// lock for safe access
	p.mu.Lock()
	// assign question to the creator
	p.createdBy[question] = id
	// assign question to the options
	p.poll[question] = options
	p.mu.Unlock()
	return nil
}
