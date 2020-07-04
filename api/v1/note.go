package models

import "time"

type Note struct {
	ID             int32           `json:"id"`
	Name           string          `json:"name"`
	Creator        User            `json:"creator"`
	Created        time.Time       `json:"created"`
	Updated        time.Time       `json:"updated"`
	Collaborations []Collaboration `json:"collaborations"`
}

type Collaboration struct {
	User      User `json:"user"`
	Read      bool `json:"read"`
	Write     bool `json:"write"`
	Delete    bool `json:"delete"`
}

