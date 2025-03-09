.env;

DB_HOST=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=library
DB_PORT=5432

PORT=8080

POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=library

Books:

GET /api/v1/books (with pagination)
GET /api/v1/books/:id (with author and reviews)
POST /api/v1/books
PUT /api/v1/books/:id
DELETE /api/v1/books/:id

Authors:

GET /api/v1/authors (with books)
GET /api/v1/authors/:id
POST /api/v1/authors
PUT /api/v1/authors/:id
DELETE /api/v1/authors/:id


Reviews:

GET /api/v1/books/:id/reviews
POST /api/v1/books/:id/reviews
PUT /api/v1/reviews/:id
DELETE /api/v1/reviews/:id

Containerization with Docker (Dockerfile and docker-compose.yaml)
Swagger Documentation

- Separation of Models and DTOs
- Swagger Documentation Created
- CORS middleware
- Health check endpoint created. The goal is that the app builds faster than the db and throws a connection request to the db and the application fails to start.
- Better error handling
- Consistent response formats


docker-compose up --build



