package server

import (
	"app/api/llog"
	"app/api/presentation/handler"
	"app/api/presentation/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server interface {
	Serve()
	Route(appHandler *handler.AppHandler)
}

type server struct {
	Handler *mux.Router
	Addr    string
}

func New(addr string) Server {
	r := mux.NewRouter()

	srv := &server{
		Addr:    addr,
		Handler: r,
	}
	return srv
}

func (s *server) Serve() {
	llog.Info(fmt.Sprintf("server running %s...", s.Addr))
	llog.Fatal(http.ListenAndServe(s.Addr, s.Handler))
}

func (s *server) Route(appHandler *handler.AppHandler) {

	s.Handler.Use(middleware.CommonMiddleware)
	authRouter := s.Handler.PathPrefix("/").Subrouter()
	authRouter.Use(middleware.AuthMiddleware)
	adminRouter := s.Handler.PathPrefix("/").Subrouter()
	// TODO: middleware

	s.Handler.HandleFunc("/ping", pingHandler).Methods("GET")

	s.Handler.HandleFunc("/login", appHandler.AuthHandler.Login).Methods("POST")
	s.Handler.HandleFunc("/account", appHandler.UserHandler.Create).Methods("POST")
	s.Handler.HandleFunc("/users", appHandler.UserHandler.GetAll).Methods("GET")
	s.Handler.HandleFunc("/users/{id}", appHandler.UserHandler.GetByID).Methods("GET")
	s.Handler.HandleFunc("/users/{id}/follows", appHandler.UserHandler.GetFollows).Methods("GET")
	s.Handler.HandleFunc("/users/{id}/followers", appHandler.UserHandler.GetFollowers).Methods("GET")
	s.Handler.HandleFunc("/categories", appHandler.CategoryHandler.GetAll).Methods("GET")
	s.Handler.HandleFunc("/tags", appHandler.TagHandler.GetAll).Methods("GET")
	s.Handler.HandleFunc("/tags/{id}", appHandler.TagHandler.GetByID).Methods("GET")
	s.Handler.HandleFunc("/threads", appHandler.ThreadHandler.GetAll).Methods("GET")
	s.Handler.HandleFunc("/threads/{id}", appHandler.ThreadHandler.GetByID).Methods("GET")
	s.Handler.HandleFunc("/threads/{id}/members", appHandler.ThreadHandler.GetMembersByThreadID).Methods("GET")

	{
		authRouter.HandleFunc("/logout", appHandler.AuthHandler.Logout).Methods("DELETE")
		authRouter.HandleFunc("/account", appHandler.UserHandler.GetMe).Methods("GET")
		authRouter.HandleFunc("/account/profile", appHandler.UserHandler.UpdateProfile).Methods("PUT")
		authRouter.HandleFunc("/account/user-id", appHandler.UserHandler.UpdateUserID).Methods("PUT")
		authRouter.HandleFunc("/account/password", appHandler.UserHandler.UpdatePassword).Methods("PUT")
		authRouter.HandleFunc("/account", appHandler.UserHandler.DeleteMe).Methods("DELETE")
		authRouter.HandleFunc("/users/{followedUUID}/follows", appHandler.UserHandler.Follow).Methods("POST")
		authRouter.HandleFunc("/users/{followedUUID}/follows", appHandler.UserHandler.Unfollow).Methods("DELETE")
		authRouter.HandleFunc("/tags", appHandler.TagHandler.Create).Methods("POST")
		authRouter.HandleFunc("/threads", appHandler.ThreadHandler.Create).Methods("POST")
		authRouter.HandleFunc("/threads/{id}", appHandler.ThreadHandler.Update).Methods("PUT")
		authRouter.HandleFunc("/threads/{id}", appHandler.ThreadHandler.Delete).Methods("DELETE")
		authRouter.HandleFunc("/threads/{id}/members", appHandler.ThreadHandler.Join).Methods("POST")
		authRouter.HandleFunc("/threads/{id}/members", appHandler.ThreadHandler.Leave).Methods("DELETE")
		authRouter.HandleFunc("/threads/{id}/members/{userID}", appHandler.ThreadHandler.ForceToLeave).Methods("DELETE")
		authRouter.HandleFunc("/threads/{threadID}/messages", appHandler.MessageHandler.GetByThreadID).Methods("GET")
		authRouter.HandleFunc("/threads/{threadID}/messages", appHandler.MessageHandler.Create).Methods("POST")
		authRouter.HandleFunc("/threads/{threadID}/messages/{messageID}", appHandler.MessageHandler.AddFavorite).Methods("POST")

		// TODO: impl
		authRouter.HandleFunc("/account/tags", appHandler.TagHandler.AddTagToUser).Methods("POST")
		authRouter.HandleFunc("/account/tags/{tagID}", appHandler.TagHandler.RemoveTagFromUser).Methods("DELETE")
		authRouter.HandleFunc("/threads/{threadID}/tags", appHandler.TagHandler.AddTagToThread).Methods("POST")
		authRouter.HandleFunc("/threads/{threadID}/tags/{tagID}", appHandler.TagHandler.RemoveTagFromThread).Methods("DELETE")

	}

	{
		adminRouter.HandleFunc("/categories", appHandler.CategoryHandler.Create).Methods("POST")
		adminRouter.HandleFunc("/categories/{id}", appHandler.CategoryHandler.Update).Methods("PUT")
		adminRouter.HandleFunc("/categories/{id}", appHandler.CategoryHandler.Delete).Methods("DELETE")

	}

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
