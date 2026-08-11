package api

import (
	"box-manager/cmd/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (h *HTTPHandlers) decodeAndValidateTournament(r *http.Request) (model.Tournament, error) {
	defer r.Body.Close()
	var dto TournamentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return model.Tournament{}, fmt.Errorf("invalid request body")
	}
	date, err := time.Parse("2006-01-02", dto.Date)
	if err != nil {
		return model.Tournament{}, fmt.Errorf("invalid date format, expected YYYY-MM-DD")
	}

	if strings.TrimSpace(dto.Name) == "" {
		return model.Tournament{}, fmt.Errorf("tournament name is required")
	}
	tournament := model.Tournament{
		Name:      dto.Name,
		Date:      date,
		Location:  dto.Location,
		PrizeFund: dto.PrizeFund,
	}
	return tournament, nil
}

func (h *HTTPHandlers) CreateHTTPTournament(w http.ResponseWriter, r *http.Request) {
	tour, err := h.decodeAndValidateTournament(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	tournament, err := h.store.CreateTournament(r.Context(), tour)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusCreated, tournament)
}

func (h *HTTPHandlers) GetHTTPTournament(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	tournament, err := h.store.GetTournament(r.Context(), id)
	if err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, tournament)
}
func (h *HTTPHandlers) UpdateHTTPTournament(w http.ResponseWriter, r *http.Request) {
	tour, err := h.decodeAndValidateTournament(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	if err := h.store.UpdateTournament(r.Context(), id, tour); err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	tour.ID = id
	h.writeJSON(w, http.StatusOK, tour)
}
func (h *HTTPHandlers) DeleteHTTPTournament(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	if err := h.store.DeleteTournament(r.Context(), id); err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPHandlers) ListHTTPTournament(w http.ResponseWriter, r *http.Request) {
	tour := h.store.ListTournaments(r.Context())
	h.writeJSON(w, http.StatusOK, tour)
}

func (h *HTTPHandlers) AddParticipantHTTPTournament(w http.ResponseWriter, r *http.Request) {
	var dto AddParticipantDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	if err := h.store.AddParticipantToTournament(r.Context(), id, dto.FighterID, dto.Place); err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusCreated, dto)
}

func (h *HTTPHandlers) RemoveParticipantHTTPTournament(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tournamentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid tournament id")
		return
	}

	fighterID, err := strconv.Atoi(vars["fighterId"])
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid fighter id")
		return
	}

	if err := h.store.RemoveParticipantFromTournament(r.Context(), tournamentID, fighterID); err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
