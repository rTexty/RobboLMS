package store

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Course struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"` // Bonus: Author field
	CreatedAt   time.Time `json:"created_at"`
}

type Store struct {
	db *sql.DB
}

func New(dbURL string) (*Store, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	s := &Store{db: db}
	if err := s.migrate(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS courses (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		author TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := s.db.Exec(query)
	return err
}

func (s *Store) CreateCourse(c *Course) error {
	query := `INSERT INTO courses (title, description, author) VALUES ($1, $2, $3) RETURNING id, created_at`
	return s.db.QueryRow(query, c.Title, c.Description, c.Author).Scan(&c.ID, &c.CreatedAt)
}

func (s *Store) GetCourses(authorFilter string) ([]Course, error) {
	query := "SELECT id, title, description, author, created_at FROM courses"
	args := []interface{}{}

	if authorFilter != "" {
		query += " WHERE author ILIKE $1"
		args = append(args, "%"+authorFilter+"%")
	}

	query += " ORDER BY created_at DESC"

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []Course{}
	for rows.Next() {
		var c Course
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.Author, &c.CreatedAt); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func (s *Store) DeleteCourse(id int) error {
	res, err := s.db.Exec("DELETE FROM courses WHERE id = $1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("course not found")
	}
	return nil
}

func (s *Store) UpdateCourse(c *Course) error {
	res, err := s.db.Exec("UPDATE courses SET title=$1, description=$2, author=$3 WHERE id=$4", c.Title, c.Description, c.Author, c.ID)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("course not found")
	}
	return nil
}
