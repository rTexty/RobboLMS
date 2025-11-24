# Course Management System (LMS)

A full-stack LMS application built with Go, PostgreSQL, and React.

## Features
- **Backend**: Go (Golang) REST API with PostgreSQL.
- **Frontend**: React (Vite) with CSS styling.
- **Infrastructure**: Docker & Docker Compose.
- **Bonus**:
    - [x] Docker Containerization
    - [x] Unit Tests (Backend)
    - [x] Author Field
    - [x] Edit Course Functionality

## Prerequisites
- Docker & Docker Compose
- OR Go 1.22+ and Node.js 18+ (for local run)

## Quick Start (Docker)

1. **Clone the repository** (if applicable)
2. **Run with Docker Compose**:
   ```bash
   docker-compose up --build
   ```
3. **Access the application**:
   - Frontend: [http://localhost:3000](http://localhost:3000)
   - Backend API: [http://localhost:8080/api/courses](http://localhost:8080/api/courses)

## Local Development

### Backend
```bash
cd backend
# Set DB URL (ensure Postgres is running)
export DATABASE_URL=postgres://user:pass@localhost:5432/lms?sslmode=disable
go run cmd/server/main.go
```

### Frontend
```bash
cd frontend
npm install
npm run dev
```

## Testing
To run backend unit tests:
```bash
cd backend
go test ./...
```
# RobboLMS
