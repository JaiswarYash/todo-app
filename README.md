# todo-app
A full-stack to-do list application with a Go (Gin) backend, MongoDB database, and a modern React + TypeScript frontend. Features include task creation, editing, completion, and deletion with real-time updates and filter options (All, Active, Completed).

## Features
- Add, edit, complete, and delete tasks
- Filter tasks (All, Active, Completed)
- Real-time updates
- RESTful API
- Easy local development and deployment

## Tech Stack
- **Backend:** Go (Gin), MongoDB
- **Frontend:** React, TypeScript, Vite

## Getting Started

### Prerequisites
- [Go](https://golang.org/) 1.20+
- [Node.js](https://nodejs.org/) 16+
- [MongoDB](https://www.mongodb.com/) (local or Atlas)

### Backend Setup
```sh
cd backend
# Set environment variables as needed (see backend/README.md)
go mod tidy
go run main.go
```

### Frontend Setup
```sh
cd frontend
npm install
npm run dev
```

### API Endpoints
See `backend/README.md` for full API usage and examples.

## License
MIT

---

[GitHub Repository](https://github.com/JaiswarYash/todo-app)
