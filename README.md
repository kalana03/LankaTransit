# LankaTransit

LankaTransit is a full-stack travel booking system with a Go backend and a Vue 3 + Vite frontend.

## Prerequisites

Make sure you have the following installed:

- Go (1.26+ recommended)
- Node.js and npm

## 1. Clone and enter the project

```bash
cd /path/to/LankaTransit
```

## 2. Start the backend

```bash
cd backend
go mod tidy
go run .
```

The backend will start on:

- http://localhost:8080

## 3. Start the frontend

Open a new terminal and run:

```bash
cd frontend
npm install
npm run dev
```

The frontend will start on:

- http://localhost:5173

## 4. Access the app

Open your browser and visit:

- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

## Useful notes

- The backend uses SQLite and creates the database file in the backend folder.
- If you make changes to the Go backend, restart the backend process.
- If you make changes to the Vue frontend, Vite will hot-reload automatically.
