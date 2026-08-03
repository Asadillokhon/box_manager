package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandler *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandler: httpHandlers,
	}
}
func (s HTTPServer) StartServer() error {
	router := mux.NewRouter()
	// fighters
	router.Path("/fighters").Methods("POST").HandlerFunc(s.httpHandler.CreateHTTPFighter)
	router.Path("/fighters/{id}").Methods("GET").HandlerFunc(s.httpHandler.GetHTTPFighter)
	router.Path("/fighters/{id}").Methods("PUT").HandlerFunc(s.httpHandler.UpdateHTTPFighter)
	router.Path("/fighters/{id}").Methods("DELETE").HandlerFunc(s.httpHandler.DeleteHTTPFighter)
	router.Path("/fighters").Methods("GET").HandlerFunc(s.httpHandler.ListHTTPFighter)
	//club
	router.Path("/club").Methods("POST").HandlerFunc(s.httpHandler.CreateHTTPClub)
	router.Path("/club/{id}").Methods("GET").HandlerFunc(s.httpHandler.GetHTTPClub)
	router.Path("/club/{id}").Methods("PUT").HandlerFunc(s.httpHandler.UpdateHTTPClub)
	router.Path("/club/{id}").Methods("DELETE").HandlerFunc(s.httpHandler.DeleteHTTPClub)
	router.Path("/club").Methods("GET").HandlerFunc(s.httpHandler.ListHttpClubs)
	// tournament
	router.Path("/tournament").Methods("POST").HandlerFunc(s.httpHandler.CreateHTTPTournament)
	router.Path("/tournament/{id}").Methods("GET").HandlerFunc(s.httpHandler.GetHTTPTournament)
	router.Path("/tournament/{id}").Methods("PUT").HandlerFunc(s.httpHandler.UpdateHTTPTournament)
	router.Path("/tournament/{id}").Methods("DELETE").HandlerFunc(s.httpHandler.DeleteHTTPTournament)
	router.Path("/tournament").Methods("GET").HandlerFunc(s.httpHandler.ListHTTPTournament)
	router.Path("/tournament/{id}/participants").Methods("POST").HandlerFunc(s.httpHandler.AddParticipantHTTPTournament)
	router.Path("/tournament/{id}/participants/{fighterId}").Methods("DELETE").HandlerFunc(s.httpHandler.RemoveParticipantHTTPTournament)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
