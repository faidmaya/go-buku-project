# go-buku-project

## Setup lokal
1. copy .env.example -> .env dan isi sesuai environment
2. buat database Postgres, contoh: bukudb
3. jalankan SQL migrations (psql):
   psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f migrations/001_create_users.sql
   psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f migrations/002_create_categories.sql
   psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f migrations/003_create_books.sql
4. (opsional) buat user admin:
   - generate bcrypt hash; contoh script singkat:
     go run tools/hash_password.go yourpassword
   - kemudian INSERT INTO users(username, password, created_by) VALUES('admin', '<hash>', 'setup');
5. go mod tidy
6. go run ./cmd

## Endpoints
- POST /api/users/login -> {username, password} -> {token}
- GET /api/categories
- POST /api/categories (protected)
- GET /api/categories/:id
- DELETE /api/categories/:id (protected)
- GET /api/categories/:id/books
- GET /api/books
- POST /api/books (protected)
- GET /api/books/:id
- DELETE /api/books/:id (protected)

## Notes
- release_year must be between 1980 and 2024
- thickness is computed: total_page > 100 => "tebal", else "tipis"
