# Installation

Clone projek ini. lalu jalankan dengan 2 cara berikut:

## Docker

Buka file .env masing-masing projek, lalu ubah setingan dari localhost ke docker, kemudian jalankan perintah berikut

```
$ docker compose up -d
```

## Localhost

Buat database terlebih dahulu di `postgres` dengan nama `gabungin`

Buka file .env masing-masing projek, lalu ubah setingan dari docker ke localhost.

Buka projek dengan kata `service` masing-masing pada terminal lalu jalankan dengan perintah berikut

- Makefile

```
make server
```

- Go run

```
go run cmd/main.go
```

# Frontend

Untuk frontend dijalankan terakhir, dengan perintah berikut

```
npm install

# Development
npm run dev

# Production
npm run build
npm run start
```

buka http://localhost:3000 untuk melihat frontend

# API

Import collection `API Gateway Gabungin.postman_collection.json` ke postman