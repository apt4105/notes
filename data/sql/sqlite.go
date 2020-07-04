package sql

import (
	"context"
	"database/sql"

	"github.com/apt4105/notes/models"
)

// Queryer is an interface for making
// queries to a database this might be an
// sql.Tx, sql.DB, or sql.Conn
type Queryer interface {
	ExecContext(ctx context.Context,
		query string,
		args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context,
		query string,
		args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context,
		query string,
		args ...interface{}) *sql.Row
}

type Store struct {
	Q Queryer
}

func (s *Store) UserByID(userID int32) (*models.User, error) {
	return nil, nil
}

func (s *Store) NoteByID(noteID int32) (*models.Note, error) {
	return nil, nil
}

func (s *Store) NotesByUserID(userID int32) ([]models.Note, error) {
	return nil, nil
}

func (s *Store) CollaborationsByNoteID(noteID int32) ([]models.Collaboration, error) {
	return nil, nil
}
