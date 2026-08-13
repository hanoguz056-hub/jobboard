# Job Board

Full-stack job board platform built with Go and React TypeScript.

## Tech Stack

**Backend:** Go, Fiber, PostgreSQL, JWT, Swagger  
**Frontend:** React, TypeScript, TailwindCSS, TanStack Query

## Features

- JWT authentication with roles (Employer, Candidate, Admin)
- Companies and job listings management
- PDF resume upload for applications
- Swagger UI documentation
- Docker + PostgreSQL

## Run locally

### Backend

```bash
cd backend
docker compose up -d
go run cmd/main.go
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```
