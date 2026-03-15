@echo off
echo Starting backend and frontend...

:: Start frontend in a new cmd window
start "Frontend (Angular)" cmd /k "cd frontend && npm start"

:: Start backend in a new cmd window
start "Backend (Golang)" cmd /k "cd backend && go run ./cmd/server/main.go"

echo Development servers are starting in separate windows.
