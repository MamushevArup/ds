package service

import (
	"errors"
	"sync"
)

type P struct {
	// in memory storage where key is question and value is option

	// poll key -> question, value : options
	poll map[string][]string
	// createdBy key -> question value : author_id
	createdBy map[string]string
	// question : { option : [ user_id ] }
	voted map[string]map[string][]string
	// for concurrent safe access to the map
	mu sync.Mutex
}

func NewPoll() *P {
	return &P{
		poll:      make(map[string][]string),
		createdBy: make(map[string]string),
		voted:     make(map[string]map[string][]string),
	}
}

func (p *P) CreatePoll(id, question string, options []string) error {
	// lock for safe access
	p.mu.Lock()
	// assign question to the creator
	p.createdBy[question] = id
	// assign question to the options
	p.poll[question] = options
	p.mu.Unlock()
	return nil
}
func (p *P) Vote(id, question, option string) (int, error) {
	p.mu.Lock()
	if p.createdBy[question] == "" {
		return 0, errors.New("question doesn't exist")
	}
	// check for option exist
	var checker bool
	for _, i2 := range p.poll[question] {
		if i2 == option {
			checker = true
		}
	}
	if !checker {
		return 0, errors.New("option doesn't exist")
	}
	// check for user_id > 0 and init the map for escape panic
	if _, ok := p.voted[question][option]; !ok {
		// If not, initialize a map for the question
		p.voted[question] = make(map[string][]string)
	}
	p.voted[question][option] = append(p.voted[question][option], id)
	p.mu.Unlock()
	// how many people voted for this option in particular question
	return len(p.voted[question][option]), nil
}
