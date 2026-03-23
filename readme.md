# Nikium IDE

A monorepo for a web IDE with Go backend, React frontend, and code execution worker.

## Architecture

```
nikium_ide/
├── backend/                    # Go HTTP API Server (Gin)
│   ├── cmd/server/             # Main entry point
│   └── internal/auth/          # Authentication (login, signup, JWT)
│
├── worker/                    # Go Code Runner (Gin)
│   ├── cmd/                    # Main entry point
│   └── internal/runner/        # Docker code execution
│
└── frontend/                   # React Frontend (Vite + Tailwind)
    └── src/                    # React components
```

## Tech Stack

- **Backend**: Go, Gin (HTTP framework), JWT
- **Worker**: Go, Gin, Docker (code execution)
- **Frontend**: React, Vite, Tailwind CSS
- **Package Manager**: pnpm

## Getting Started

### Prerequisites

- Node.js (v20+)
- pnpm (`npm install -g pnpm`)
- Go (v1.22+)
- Docker (for code execution)

### Installation

```bash
pnpm install
```

## Development

### Run Frontend

```bash
pnpm run dev
```

### Run Backend (API Server)

```bash
pnpm run dev:backend
```

### Run Worker (Code Executor)

```bash
pnpm run dev:worker
```

---

## API Documentation

### Backend Server (:8080)

#### Health Check

```
GET /health
```

**Test:**

```bash
curl http://localhost:8080/health
```

**Response:**

```json
{
  "status": "ok"
}
```

---

#### User Signup

```
POST /api/auth/signup
```

**Test:**

```bash
curl -X POST http://localhost:8080/api/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpass123",
    "email": "test@example.com"
  }'
```

**Request Body:**

```json
{
  "username": "string (required)",
  "password": "string (required)",
  "email": "string (required)"
}
```

**Response (Success):**

```json
{
  "message": "User created succesfully"
}
```

**Response (Error):**

```json
{
  "error": "username and password required"
}
```

---

#### User Login

```
POST /api/auth/login
```

**Test:**

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpass123"
  }'
```

**Request Body:**

```json
{
  "username": "string (required)",
  "password": "string (required)"
}
```

**Response (Success):**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expiresAt": 1713993600
}
```

**Response (Error):**

```json
{
  "error": "invalid credentials"
}
```

---

#### User Logout

```
POST /api/auth/logout
```

**Test:**

```bash
curl -X POST http://localhost:8080/api/auth/logout
```

**Response:**

```json
{
  "message": "logged out successfully"
}
```

---

#### Submit Code

```
POST /api/code/submit
```

**Test:**

```bash
curl -X POST http://localhost:8080/api/code/submit \
  -H "Content-Type: application/json" \
  -d '{
    "code": "print(\"Hello, World!\")",
    "input": "",
    "language": "python"
  }'
```

**Request Body:**

```json
{
  "code": "string (required)",
  "input": "string (optional)",
  "language": "string (optional, e.g., python, javascript, go)"
}
```

**Response:**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "queued",
  "message": "Code submitted for execution"
}
```

---

#### Get Code Result

```
GET /api/code/:id
```

**Test:**

```bash
curl http://localhost:8080/api/code/550e8400-e29b-41d4-a716-446655440000
```

**Response:**

```json
{
  "status": "completed",
  "output": "Hello, World!",
  "exitCode": 0
}
```

---

### Worker Server (:8081)

#### Health Check

```
GET /health
```

**Test:**

```bash
curl http://localhost:8081/health
```

**Response:**

```json
{
  "status": "ok"
}
```

---

#### Run Code

```
POST /run
```

**Test:**

```bash
curl -X POST http://localhost:8081/run \
  -H "Content-Type: application/json" \
  -d '{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "userId": "123e4567-e89b-12d3-a456-426614174000",
    "code": "print(\"Hello from Docker!\")",
    "input": "",
  }'
```

**Request Body:**

```json
{
  "id": "uuid (required)",
  "userId": "uuid (required)",
  "code": "string (required)",
  "input": "string (optional)"
}
```

**Response (Success):**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "output": "Hello from Docker!\n",
  "error": "",
  "exitCode": 0,
  "duration": "245.791421ms"
}
```

**Response (Error - Runtime Error):**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "output": "",
  "error": "exec format error",
  "exitCode": 1,
  "duration": "123.456ms"
}
```

**Response (Error - Timeout):**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "userId": "123e4567-e89b-12d3-a456-426614174000",
  "output": "",
  "error": "execution timed out",
  "exitCode": -1,
  "duration": "10s"
}
```

---

## Environment Variables

Create `backend/.env`:

```env
JWT_SECRET=your-secret-key-change-in-production
PORT=8080
```
