package api

import (
	"encoding/json"
	"lms-backend/internal/store"
	"net/http"
	"strconv"
)

type CourseStore interface {
	GetCourses(authorFilter string) ([]store.Course, error)
	CreateCourse(c *store.Course) error
	DeleteCourse(id int) error
	UpdateCourse(c *store.Course) error
}

type Handler struct {
	store CourseStore
}

func NewHandler(s CourseStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/courses", h.handleGetCourses)
	mux.HandleFunc("POST /api/courses", h.handleCreateCourse)
	mux.HandleFunc("DELETE /api/courses/{id}", h.handleDeleteCourse)
	mux.HandleFunc("PUT /api/courses/{id}", h.handleUpdateCourse)
}

func (h *Handler) handleGetCourses(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	courses, err := h.store.GetCourses(author)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.respondJSON(w, http.StatusOK, courses)
}

func (h *Handler) handleCreateCourse(w http.ResponseWriter, r *http.Request) {
	var c store.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if len(c.Title) < 3 {
		h.respondError(w, http.StatusBadRequest, "Title must be at least 3 characters long")
		return
	}

	if err := h.store.CreateCourse(&c); err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.respondJSON(w, http.StatusCreated, c)
}

func (h *Handler) handleDeleteCourse(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.store.DeleteCourse(id); err != nil {
		h.respondError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) handleUpdateCourse(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var c store.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if len(c.Title) < 3 {
		h.respondError(w, http.StatusBadRequest, "Title must be at least 3 characters long")
		return
	}

	c.ID = id
	if err := h.store.UpdateCourse(&c); err != nil {
		h.respondError(w, http.StatusNotFound, err.Error())
		return
	}
	h.respondJSON(w, http.StatusOK, c)
}

func (h *Handler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func (h *Handler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"error": message})
}
