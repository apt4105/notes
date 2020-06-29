package models

import "time"

type Note struct {
	ID      int32
	Name    string
	Creator User
	Created time.Time
	Updated time.Time
}

type Collaboration struct {
	NoteID              int32
	Principal           User
	Read, Write, Delete bool
}
