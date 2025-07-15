# todo-app
A full-stack to-do list application with a Go (Gin) backend, MongoDB database, and a modern React + TypeScript frontend. Features include task creation, editing, completion, and deletion with real-time updates and filter options (All, Active, Completed).

## Features
- Add, edit, complete, and delete tasks
- Filter tasks (All, Active, Completed)
- Real-time updates
- RESTful API
- Modern UI

## Folder Structure
```
todo-app/
├── backend/         # Go (Gin) backend API
│   ├── main.go
│   ├── go.mod
│   ├── config/
│   ├── controllers/
│   ├── models/
│   ├── routes/
│   ├── services/
│   ├── middlewares/
│   └── utils/
├── frontend/        # React + TypeScript frontend
│   ├── src/
│   │   ├── components/
│   │   ├── context/
│   │   ├── hooks/
│   │   ├── services/
│   │   ├── styles/
│   │   └── ...
│   ├── public/
│   ├── package.json
│   └── ...
└── README.md        # Project overview (this file)
```

## Getting Started

### Backend (Go + Gin + MongoDB)
1. `cd backend`
2. Set environment variables (optional, defaults shown):
   - `MONGO_URI` (default: mongodb://localhost:27017)
   - `MONGO_DB` (default: todoapp)
   - `PORT` (default: 8080)
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```

### Frontend (React + TypeScript)
1. `cd frontend`
2. Install dependencies:
   ```sh
   npm install
   ```
3. Start the development server:
   ```sh
   npm run dev
   ```
4. Open [http://localhost:5173](http://localhost:5173) in your browser.

## License
MIT
