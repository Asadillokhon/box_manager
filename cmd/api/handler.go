package api

import (
	"box-manager/cmd/model"
	"box-manager/cmd/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	store store.Store
}

func NewHTTPHandler(store store.Store) *HTTPHandlers {
	return &HTTPHandlers{
		store: store,
	}
}

// writeJSON — хелпер для отправки JSON-ответов
func (h *HTTPHandlers) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// writeError — хелпер для отправки ошибок
func (h *HTTPHandlers) writeError(w http.ResponseWriter, status int, msg string) {
	h.writeJSON(w, status, map[string]string{"error": msg})
}

func (h *HTTPHandlers) decodeAndValidateFighter(r *http.Request) (model.Fighter, error) {
	defer r.Body.Close()

	var dto FighterDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return model.Fighter{}, fmt.Errorf("invalid request body")
	}
	if dto.FirstName == "" || dto.LastName == "" {
		return model.Fighter{}, fmt.Errorf("first_name and last_name are required")
	}
	birthday, err := time.Parse("2006-01-02", dto.BirthDate)
	if err != nil {
		return model.Fighter{}, fmt.Errorf("invalid birth_date format, expected YYYY-MM-DD")
	}
	return model.Fighter{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthDate: birthday,
		Weight:    dto.Weight,
		Category:  dto.Category,
	}, nil
}
func (h *HTTPHandlers) extractID(r *http.Request) (int, error) {
	// idStr = mux.Vars(r)["id"] можно так на одну строку
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		return 0, fmt.Errorf("missing id parameter")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %s", idStr)
	}
	return id, nil
}

/*
pattern : /box
method : POST
info : JSON in request body
*/
func (h *HTTPHandlers) CreateHTTPFighter(w http.ResponseWriter, r *http.Request) {
	// чтение
	fighter, err := h.decodeAndValidateFighter(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	created, err := h.store.CreateFighter(fighter)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.writeJSON(w, http.StatusCreated, created)
}

/*
pattern : /box/{id}
method : get
parametr : ID
*/
func (h *HTTPHandlers) GetHTTPFighter(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "undefined id parameter")
		return
	}
	fighter, err := h.store.GetFighter(id)
	if err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	h.writeJSON(w, http.StatusOK, fighter)
}

/*
pattern : /box/{id}
method : put
parametr : ID
*/
func (h *HTTPHandlers) UpdateHTTPFighter(w http.ResponseWriter, r *http.Request) {

	fighter, err := h.decodeAndValidateFighter(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "undefined id parameter")
		return
	}
	err = h.store.UpdateFighter(id, fighter)
	if err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	fighter.ID = id
	h.writeJSON(w, http.StatusOK, fighter)
}
func (h *HTTPHandlers) DeleteHTTPFighter(w http.ResponseWriter, r *http.Request) {
	id, err := h.extractID(r)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	err = h.store.DeleteFighter(id)
	if err != nil {
		h.writeError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPHandlers) ListHTTPFighter(w http.ResponseWriter, r *http.Request) {
	fighters := h.store.ListFighters()
	h.writeJSON(w, http.StatusOK, fighters)
}
