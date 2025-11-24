package api

import (
	"bytes"
	"encoding/json"
	"lms-backend/internal/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStore struct {
	courses []store.Course
}

func (m *MockStore) GetCourses(authorFilter string) ([]store.Course, error) {
	return m.courses, nil
}

func (m *MockStore) CreateCourse(c *store.Course) error {
	c.ID = len(m.courses) + 1
	m.courses = append(m.courses, *c)
	return nil
}

func (m *MockStore) DeleteCourse(id int) error {
	return nil
}

func (m *MockStore) UpdateCourse(c *store.Course) error {
	return nil
}

func TestGetCourses(t *testing.T) {
	mockStore := &MockStore{
		courses: []store.Course{
			{ID: 1, Title: "Test Course", Description: "Desc"},
		},
	}
	handler := NewHandler(mockStore)
	req, _ := http.NewRequest("GET", "/api/courses", nil)
	rr := httptest.NewRecorder()

	handler.handleGetCourses(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"title":"Test Course","description":"Desc","author":"","created_at":"0001-01-01T00:00:00Z"}]`
	if rr.Body.String() != expected+"\n" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateCourse(t *testing.T) {
	mockStore := &MockStore{}
	handler := NewHandler(mockStore)

	course := store.Course{Title: "New Course", Description: "New Desc"}
	body, _ := json.Marshal(course)
	req, _ := http.NewRequest("POST", "/api/courses", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.handleCreateCourse(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestCreateCourseInvalid(t *testing.T) {
	mockStore := &MockStore{}
	handler := NewHandler(mockStore)

	course := store.Course{Title: "No", Description: "Short title"}
	body, _ := json.Marshal(course)
	req, _ := http.NewRequest("POST", "/api/courses", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.handleCreateCourse(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
