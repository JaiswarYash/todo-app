# To-Do App Backend (Go + Gin + MongoDB)

## Setup & Run

1. Install Go 1.20+
2. Install and run MongoDB (local or cloud, e.g. MongoDB Atlas)
3. Set environment variables (optional, defaults shown):
   - `MONGO_URI` (default: mongodb://localhost:27017)
   - `MONGO_DB` (default: todoapp)
   - `PORT` (default: 8080)
   
   Example (Windows PowerShell):
   ```powershell
   $env:MONGO_URI="mongodb://localhost:27017"
   $env:MONGO_DB="todoapp"
   $env:PORT="8080"
   ```

4. Install dependencies:
   ```sh
   go mod tidy
   ```
5. Run the server:
   ```sh
   go run main.go
   ```

## API Endpoints

### Create Task
```sh
curl -X POST http://localhost:8080/api/tasks/ \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy milk"}'
```

### Get All Tasks
```sh
curl http://localhost:8080/api/tasks/
```

### Update Task
```sh
curl -X PUT http://localhost:8080/api/tasks/<id> \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy bread", "completed": true}'
```

### Delete Task
```sh
curl -X DELETE http://localhost:8080/api/tasks/<id>
```

- Replace `<id>` with the MongoDB ObjectID string returned from create/list endpoints. 