package data

import "github.com/apt4105/notes/models"


type Store interface {
	// user methods
	UserByID(userID int32) (*models.User, error)

	// note methods
	NoteByID(noteID int32) (*models.Note, error)
	NotesByUserID(userID int32) ([]models.Note, error)
	CollaborationsByNoteID(noteID int32) ([]models.Collaboration, error)
}
