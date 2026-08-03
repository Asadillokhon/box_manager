package api

import (
	"box-manager/cmd/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (h *HTTPHandlers) decodeAndValidateClub(r *http.Request) (model.Club, error) {
	defer r.Body.Close()
	var club model.Club
	if err := json.NewDecoder(r.Body).Decode(&club); err != nil {
		return model.Club{}, fmt.Errorf("invalid JSON: %v", err)
	}
	if strings.TrimSpace(club.Name) == "" {
		return model.Club{}, fmt.Errorf("club name is required")
	}
	return club, nil
}

func (h *HTTPHandlers) CreateHTTPClub(w http.ResponseWriter, r *http.Request) {
	club, err := h.decodeAndValidateClub(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	created, err := h.store.CreateClub(club)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.writeJSON(w, http.StatusCreated, created)
}

func (h *HTTPHandlers) GetHTTPClub(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "undefined id parameter")
		return
	}
	club, err := h.store.GetClub(id)
	if err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, club)
}

func (h *HTTPHandlers) UpdateHTTPClub(w http.ResponseWriter, r *http.Request) {
	club, err := h.decodeAndValidateClub(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "undefined id parameter")
		return
	}

	if err := h.store.UpdateClub(id, club); err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	club.ID = id
	h.writeJSON(w, http.StatusOK, club)
}

func (h *HTTPHandlers) DeleteHTTPClub(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "undefined id parameter")
		return
	}
	if err := h.store.DeleteClub(id); err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *HTTPHandlers) ListHttpClubs(w http.ResponseWriter, r *http.Request) {
	clubs := h.store.ListClubs()
	h.writeJSON(w, http.StatusOK, clubs)
}
