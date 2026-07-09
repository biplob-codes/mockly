# Mockly

A lightweight fake REST API server for frontend and integration testing — built with Go's standard `net/http` (no framework, no external router).

Mockly serves realistic, pre-seeded JSON data over standard REST endpoints so you can build and test a UI before a real backend exists. It requires **no database, no setup, and no auth** — just run the binary and start hitting endpoints.

All seed data is Naruto-themed — users are shinobi (Sasuke, Kakashi, Itachi, Minato, and more), and their todos, posts, comments, and products follow the same world, so the data feels connected instead of generic.

> **Note:** Mockly does not persist any changes. `POST`, `PUT`, and `DELETE` requests return a realistic response, but the underlying data is never actually created, modified, or removed. Every `GET` always returns the same seeded dataset. This is intentional — it's a mock server, not a database.

## Features

- Zero setup — in-memory data, no database required
- Standard REST conventions (collection + single-item routes)
- Full CRUD-shaped endpoints for 5 resources: `users`, `posts`, `todos`, `products`, `comments`
- Naruto-themed seed data with realistic nested fields (addresses, company info, tags, SKUs, ratings)
- Self-describing root endpoint (`/`) listing all available resources
- Dedicated lightweight health check (`/health`)
- Simple, dependency-free codebase — easy to read and extend
- Configurable port via `PORT` environment variable

## Tech Stack

- Go (standard library only — `net/http`, `encoding/json`)
- No external router, no ORM, no database

## Getting Started

### Prerequisites

- Go 1.26 or later (uses `http.ServeMux` method + path pattern matching, e.g. `GET /users/{id}`)

### Installation

```bash
git clone https://github.com/biplob-codes/mockly.git
cd mockly
go run main.go
```

The server starts on port `8080` by default. Set a custom port with the `PORT` environment variable:

```bash
PORT=3000 go run main.go
```

### Root Endpoint

Hitting the root URL returns a self-describing index of the API — what it is, its version, and every available resource with its path and record count:

```bash
curl http://localhost:8080/
```

### Health Check

A separate, minimal endpoint for uptime monitors and load balancers:

```bash
curl http://localhost:8080/health
```

## API Reference

All responses are JSON. All resources support the same five operations.

| Method | Endpoint           | Description                     |
| ------ | ------------------ | ------------------------------- |
| GET    | `/{resource}`      | Get all records                 |
| GET    | `/{resource}/{id}` | Get a single record by ID       |
| POST   | `/{resource}`      | Create a record (not persisted) |
| PUT    | `/{resource}/{id}` | Update a record (not persisted) |
| DELETE | `/{resource}/{id}` | Delete a record (not persisted) |

Where `{resource}` is one of: `users`, `posts`, `todos`, `products`, `comments`.

### Example: Users

```bash
# Get all users
curl http://localhost:8080/users

# Get a single user
curl http://localhost:8080/users/1

# Create a user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Konohamaru Sarutobi",
    "username": "konohamarusarutobi",
    "email": "konohamaru@sarutobi.konoha",
    "phone": "555-0130",
    "website": "konohamarusarutobi.konoha",
    "avatar": "https://example.com/avatar.webp",
    "address": {
      "street": "Sarutobi Residence",
      "city": "Konohagakure",
      "state": "Land of Fire",
      "zipCode": "00015",
      "country": "Naruto World"
    },
    "company": {
      "name": "Konoha Hokage Office",
      "catchPhrase": "The Seventh Hokage in training",
      "industry": "Shinobi"
    }
  }'

# Update a user
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Name", "email": "updated@example.com"}'

# Delete a user
curl -X DELETE http://localhost:8080/users/1
```

### Resource Fields

**User**

```
id, name, username, email, phone, website, avatar,
address { street, city, state, zipCode, country },
company { name, catchPhrase, industry },
createdAt
```

**Post**

```
id, userId, title, body, tags[], views, createdAt
```

**Todo**

```
id, userId, title, completed, priority, dueDate
```

**Product**

```
id, name, description, price, category, sku, rating, inStock, createdAt
```

**Comment**

```
id, postId, name, email, body, createdAt
```

## Seed Data

All seed data lives in `internal/store/` and is Naruto-themed:

- **Users** — 14 shinobi (e.g. Sasuke Uchiha, Kakashi Hatake, Itachi Uchiha, Minato Namikaze, Tsunade Senju, Gaara Kazekage), each with a clan/village-themed address and company
- **Todos** — 15 in-world tasks tied to each user (e.g. training jutsu, sealing scrolls, restocking hospital supplies)
- **Posts** — 15 in-character blog posts reflecting each shinobi's story or philosophy
- **Comments** — 15 in-world reactions from other characters (Naruto, Sakura, Rock Lee, and more)
- **Products** — 15 ninja-world gear items (kunai, shuriken, scrolls, chakra weights, Akatsuki cloaks) across themed categories

## Project Structure

```
mockly/
├── cmd/
│   └── main.go              # entrypoint, route registration
├── internal/
│   ├── model/               # struct definitions (User, Post, Todo, Product, Comment)
│   ├── store/                # in-memory Naruto-themed seed data
│   └── handlers/             # HTTP handlers per resource, plus home and health
└── README.md
```

## Use Cases

- Build frontend UIs before the real backend API exists
- Practice API integration and HTTP client code
- Test error handling and edge cases with predictable data
- Learn/teach REST conventions with a minimal, readable Go codebase

## Contributing

Issues and pull requests are welcome. If you add a new resource, follow the existing pattern:

1. Add a struct + `New*` request struct in `internal/model/`
2. Add seed data in `internal/store/`
3. Add a handler file in `internal/handlers/` with `GetAll`, `GetById`, `Create`, `Update`, `Delete`
4. Register the five routes in `cmd/main.go`

## License

MIT
