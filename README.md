# Article API - Sharing Vision Backend Test 2023

REST API untuk manajemen artikel dengan fitur CRUD lengkap menggunakan Golang, Gin Framework, dan PostgreSQL.

## Tech Stack

- **Go** 1.24.7
- **Gin** Web Framework
- **PostgreSQL** Database
- **golang-migrate** untuk database migration
- **validator/v10** untuk validasi input

## Project Structure

```
article-api/
├── config/              # Konfigurasi dan database connection
├── internal/
│   ├── handlers/        # HTTP handlers (controllers)
│   ├── models/          # Data models dan request/response structs
│   ├── repository/      # Database operations
│   └── validator/       # Input validation
├── db/migrations/       # Database migration files
├── main.go             # Entry point aplikasi
└── .env                # Environment variables
```

## Requirements

- Go 1.24.7 atau lebih tinggi
- PostgreSQL 12 atau lebih tinggi
- golang-migrate CLI (untuk menjalankan migration)

## Installation

### 1. Clone Repository

```bash
git clone <repository-url>
cd particle-be-tst
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Install golang-migrate

**Linux:**
```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

**MacOS:**
```bash
brew install golang-migrate
```

**Windows:**
```bash
scoop install migrate
```

### 4. Setup Database

Buat database PostgreSQL:

```bash
createdb article_db
```

### 5. Configuration

Copy `.env.example` ke `.env` dan sesuaikan konfigurasi:

```bash
cp .env.example .env
```

Edit `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=article_db
DB_SSLMODE=disable
SERVER_PORT=8080
```

### 6. Run Migration

```bash
migrate -path db/migrations -database "postgresql://postgres:your_password@localhost:5432/article_db?sslmode=disable" up
```

### 7. Run Application

```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## API Endpoints

### 1. Create Article

**Endpoint:** `POST /article/`

**Request Body:**
```json
{
  "title": "Judul artikel minimal 20 karakter",
  "content": "Content artikel minimal 200 karakter. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
  "category": "Technology",
  "status": "publish"
}
```

**Validation Rules:**
- `title`: required, minimal 20 karakter
- `content`: required, minimal 200 karakter
- `category`: required, minimal 3 karakter
- `status`: required, harus salah satu dari: `publish`, `draft`, `thrash`

**Response Success (201):**
```json
{
  "status": "success",
  "message": "Article created successfully",
  "data": {
    "id": 1,
    "title": "Judul artikel minimal 20 karakter",
    "content": "Content artikel minimal 200 karakter...",
    "category": "Technology",
    "status": "publish",
    "created_date": "2023-12-16T10:00:00Z",
    "updated_date": "2023-12-16T10:00:00Z"
  }
}
```

### 2. Get All Articles (with Pagination)

**Endpoint:** `GET /article/:limit/:offset`

**Example:** `GET /article/10/0`

**Response Success (200):**
```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "title": "Judul artikel",
      "content": "Content artikel...",
      "category": "Technology",
      "status": "publish",
      "created_date": "2023-12-16T10:00:00Z",
      "updated_date": "2023-12-16T10:00:00Z"
    }
  ]
}
```

### 3. Get Article by ID

**Endpoint:** `GET /article/:id`

**Example:** `GET /article/1`

**Response Success (200):**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "title": "Judul artikel",
    "content": "Content artikel...",
    "category": "Technology",
    "status": "publish",
    "created_date": "2023-12-16T10:00:00Z",
    "updated_date": "2023-12-16T10:00:00Z"
  }
}
```

**Response Error (404):**
```json
{
  "status": "error",
  "message": "Article not found"
}
```

### 4. Update Article

**Endpoint:** `PUT /article/:id`

**Example:** `PUT /article/1`

**Request Body:**
```json
{
  "title": "Updated judul artikel minimal 20 karakter",
  "content": "Updated content artikel minimal 200 karakter. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
  "category": "Updated Category",
  "status": "draft"
}
```

**Response Success (200):**
```json
{
  "status": "success",
  "message": "Article updated successfully",
  "data": {
    "id": 1,
    "title": "Updated judul artikel minimal 20 karakter",
    "content": "Updated content artikel...",
    "category": "Updated Category",
    "status": "draft",
    "created_date": "2023-12-16T10:00:00Z",
    "updated_date": "2023-12-16T10:30:00Z"
  }
}
```

### 5. Delete Article

**Endpoint:** `DELETE /article/:id`

**Example:** `DELETE /article/1`

**Response Success (200):**
```json
{
  "status": "success",
  "message": "Article deleted successfully"
}
```

**Response Error (404):**
```json
{
  "status": "error",
  "message": "Article not found"
}
```

## Testing with cURL

### Create Article
```bash
curl -X POST http://localhost:8080/article/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Test Article with Minimum Twenty Characters",
    "content": "This is a test article content that must be at least 200 characters long. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
    "category": "Technology",
    "status": "publish"
  }'
```

### Get All Articles
```bash
curl http://localhost:8080/article/10/0
```

### Get Article by ID
```bash
curl http://localhost:8080/article/1
```

### Update Article
```bash
curl -X PUT http://localhost:8080/article/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Article Title with Minimum Twenty Characters",
    "content": "This is an updated article content that must be at least 200 characters long. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
    "category": "Updated Category",
    "status": "draft"
  }'
```

### Delete Article
```bash
curl -X DELETE http://localhost:8080/article/1
```

## Database Schema

**Table: posts**

| Column | Type | Constraints |
|--------|------|-------------|
| id | SERIAL | PRIMARY KEY |
| title | VARCHAR(200) | NOT NULL |
| content | TEXT | NOT NULL |
| category | VARCHAR(100) | NOT NULL |
| created_date | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |
| updated_date | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP |
| status | VARCHAR(100) | NOT NULL, CHECK (publish/draft/thrash) |

**Indexes:**
- idx_posts_status (status)
- idx_posts_category (category)
- idx_posts_created_date (created_date)

## Error Handling

API menggunakan HTTP status codes standar:

- `200 OK` - Request berhasil
- `201 Created` - Resource berhasil dibuat
- `400 Bad Request` - Validasi error atau parameter invalid
- `404 Not Found` - Resource tidak ditemukan
- `500 Internal Server Error` - Server error

Format error response:

```json
{
  "status": "error",
  "message": "Error message",
  "errors": "Detailed error information"
}
```

## Development

### Build Application

```bash
go build -o article-api
```

### Run Tests

```bash
go test ./...
```

## License

MIT License
