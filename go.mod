module github.com/OkoliDaniel/waitlist-backend

go 1.21.1

replace internal/database v0.0.0 => ./internal/database

require internal/database v0.0.0

require (
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/dghubble/go-twitter v0.0.0-20221104224141-912508c3888b // indirect
	github.com/dghubble/gologin/v2 v2.4.0 // indirect
	github.com/dghubble/oauth1 v0.7.2 // indirect
	github.com/dghubble/sling v1.4.1 // indirect
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.1.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/sqlc-dev/sqlc/cmd/sqlc v1.24.0 // indirect
	github.com/pressly/goose/v3/cmd/goose v3.16.0 // indirect

)
