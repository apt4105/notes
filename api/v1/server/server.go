package server

import (
	"context"
	gosql "database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/julienschmidt/httprouter"

	"github.com/apt4105/notes/blob"
	"github.com/apt4105/notes/config"
	"github.com/apt4105/notes/data"
	"github.com/apt4105/notes/data/sql"
)

var ctxServerKey = &struct{}{}

type Server struct {
	db   data.Store
	blob blob.Store
}

type FactoryFunc func(context.Context) (
	srv *Server, closer func() error, err error)

// TODO
func NewServerFactory(conf config.Server) (FactoryFunc, error) {
	db, err := gosql.Open("sqlite", conf.Data.Conn)
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) (
		srv *Server,
		closer func() error,
		err error) {

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return nil, nil, err
		}

		srv = &Server{
			db: &sql.Store{
				Q: tx,
			},
		}

		closer = func() error {
			return tx.Commit()
		}

		return srv, closer, nil
	}, nil
}

func ServerHandler() http.Handler {
	mux := &apiMux{*httprouter.New()}

	mux.Handle(GET, "/users/:id", (*Server).GetUser)
	mux.Handle(GET, "/users/:id/notes", (*Server).GetUserNotes)
	mux.Handle(GET, "/notes/:id", (*Server).GetNote)

	return mux
}

func ServerMiddleware(conf config.Server) (
	func(http.Handler) http.Handler, error) {

	serverFactory, err := NewServerFactory(conf)
	if err != nil {
		return nil, err
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			srv, closer, err := serverFactory(r.Context())
			if err != nil {
				http.Error(w,
					"something really bad happened",
					http.StatusInternalServerError)
				return
			}

			defer closer()

			r = r.Clone(
				context.WithValue(r.Context(), ctxServerKey, srv))

			next.ServeHTTP(w, r)
		})
	}, nil
}

type apiMux struct {
	mux httprouter.Router
}

func (mux *apiMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux.mux.ServeHTTP(w, r)
}

func (mux *apiMux) Handle(method, path string,
	h func(s *Server, w http.ResponseWriter, r *http.Request)) {

	mux.mux.Handler(method, path,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s := r.Context().Value(ctxServerKey).(*Server)
			h(s, w, r)
		}))
}
