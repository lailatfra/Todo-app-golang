# Rest API Backend Todo-List

REST API untuk aplikasi Todo List menggunakan Golang.

## Teknologi

- Golang
- net/http
- JSON encoding
- CORS enabled

## Struktur Data

```json
{
  "id": 1,
  "title": "Belajar React - 09:00",
  "description": "Belajar React hooks",
  "status": "pending",
  "created_at": "2026-01-08T10:00:00Z"
}
```

## Endpoint API

Base URL: `http://localhost:8080`

### 1. Get All Todos
```
GET /todos
```
Response: Array of todos (200)

### 2. Get Todo by ID
```
GET /todos/{id}
```
Response: Single todo (200) atau 404

### 3. Create Todo
```
POST /todos
Body: {"title": "...", "description": "..."}
```
Response: Todo baru (201)
- Status default: `pending`
- ID dan CreatedAt auto generate

### 4. Update Todo

**Toggle Status (tanpa body):**
```
PUT /todos/{id}
```
Toggle antara pending dan done (200)

**Update Data (dengan body):**
```
PUT /todos/{id}
Body: {"title": "...", "description": "..."}
```
Update title dan description (200)

### 5. Delete Todo
```
DELETE /todos/{id}
```
Response: Success message (200)

## Cara Menjalankan

```bash
go run main.go
```

Server berjalan di `http://localhost:8080`

## Alur Backend

1. **GET /todos** - Return semua todo dari memory
2. **GET /todos/{id}** - Cari todo by ID, return 404 jika tidak ada
3. **POST /todos** - Parse body, generate ID, set status `pending`, tambah ke slice
4. **PUT /todos/{id}** - Jika tidak ada body: toggle status. Jika ada body: update data
5. **DELETE /todos/{id}** - Hapus todo dari slice by ID

## HTTP Status Codes

- 200: Success
- 201: Created
- 400: Bad Request
- 404: Not Found
