# FrameLog — ARCHITECTURE.md

> **Version:** 0.1
> **Status:** Living Document
> **Project:** FrameLog
> **Purpose:** Defines the architecture, conventions, and technical decisions of the project. This document should evolve as the application grows.

---

# 1. Vision

FrameLog is a **personal-first media tracking platform** where users can:

* Track movies and TV shows
* Maintain watchlists
* Record ratings and reviews
* View yearly statistics
* Organize media into lists
* Recommend media to friends
* Eventually support social features like activity feeds and movie nights

The application should remain **simple, fast, maintainable, and deployable on a single VPS**.

---

# 2. Architecture Overview

```
                     Internet
                         │
                         ▼
                    Caddy / Nginx
                         │
          ┌──────────────┴──────────────┐
          │                             │
          ▼                             ▼
     Vue Frontend                  Go REST API
                                      │
                                      ▼
                               Service Layer
                                      │
                                      ▼
                             Repository Layer
                                      │
                                      ▼
                                PostgreSQL
```

Future additions:

```
Redis
Background Workers
Email Service
Object Storage
```

---

# 3. Technology Stack

## Frontend

* Vue 3
* Vite
* TypeScript
* Vue Router
* Pinia
* TailwindCSS
* shadcn-vue
* TanStack Query
* VueUse

---

## Backend

* Go 1.24+
* Gin
* PostgreSQL
* pgx
* sqlc
* Goose
* JWT Authentication

---

## Infrastructure

* Docker Compose
* Caddy
* GitHub Actions
* Ubuntu VPS

---

# 4. High-Level Design

The backend follows a straightforward layered architecture:

```
HTTP Request

↓

Handler

↓

Service

↓

Repository

↓

Database
```

Responsibilities:

## Handler

Responsible for:

* Reading request
* Validation
* Calling service
* Returning response

No business logic.

---

## Service

Responsible for:

* Business rules
* Authorization
* Calculations
* Orchestration

No SQL.

---

## Repository

Responsible for:

* Database interaction
* sqlc generated methods
* Query execution

No HTTP knowledge.

---

# 5. Folder Structure

```
backend/

cmd/
    server/

internal/

    config/

    handler/

    middleware/

    service/

    repository/

    auth/

    models/

    validator/

db/

    migrations/

    queries/

pkg/

frontend/

src/

components/

layouts/

pages/

router/

stores/

services/

composables/

types/

assets/
```

---

# 6. Domain Model

The central entity is **Media**.

Everything revolves around Media.

```
Media

├── Movie
├── Series
├── Anime
├── Documentary
└── Mini Series
```

Never create separate Movie and Series tables.

Use a single Media table.

---

# 7. Database Design

## User

```
id

username

email

password_hash

avatar

bio

created_at

updated_at
```

---

## Media

```
id

tmdb_id

media_type

title

original_title

overview

poster_path

backdrop_path

release_date

runtime

language

country

created_at
```

---

## Genre

```
id

name
```

---

## MediaGenre

```
media_id

genre_id
```

Many-to-many relationship.

---

## WatchEntry

This is the heart of the application.

```
id

user_id

media_id

status

rating

started_at

finished_at

rewatch_count

favorite

notes

created_at

updated_at
```

---

## Review

```
id

user_id

media_id

content

spoiler

created_at
```

---

## CustomList

```
id

user_id

name

description
```

---

## ListItem

```
list_id

media_id
```

---

## Recommendation

```
id

sender_id

receiver_id

media_id

message

status

created_at
```

---

## Friendship

```
id

user_one

user_two

status
```

---

# 8. Entity Relationship

```
User

│

├───────────────┐

│               │

▼               ▼

WatchEntry   Review

│

▼

Media

│

▼

Genre
```

Later:

```
Media

↓

Recommendation

↓

Friend
```

---

# 9. Watch Status

Only these statuses exist.

```
WantToWatch

Watching

Completed

Dropped

Paused

Rewatching
```

Avoid adding more unless absolutely necessary.

---

# 10. REST API

Authentication

```
POST   /auth/register

POST   /auth/login

POST   /auth/logout

GET    /auth/me
```

---

Media

```
GET    /media/search

GET    /media/:id

POST   /media/import
```

---

Watch

```
GET    /watch

POST   /watch

PATCH  /watch/:id

DELETE /watch/:id
```

---

Reviews

```
POST   /reviews

PATCH  /reviews/:id

DELETE /reviews/:id
```

---

Statistics

```
GET /stats/year

GET /stats/month

GET /stats/genres
```

---

Lists

```
GET

POST

PATCH

DELETE
```

---

Recommendations

```
POST /recommend

GET /recommendations
```

---

Friends

```
POST /friends/request

POST /friends/accept

GET /friends
```

---

# 11. Authentication Flow

```
Register

↓

Hash Password

↓

Save User

↓

Login

↓

Verify Password

↓

Generate JWT

↓

Return Token

↓

Authenticated Requests
```

Future migration to secure HTTP-only cookies should be possible without changing business logic.

---

# 12. Request Lifecycle

```
Request

↓

Middleware

↓

Authentication

↓

Validation

↓

Handler

↓

Service

↓

Repository

↓

PostgreSQL

↓

Response
```

---

# 13. Error Handling

Standard JSON response:

Success

```json
{
  "success": true,
  "data": {}
}
```

Error

```json
{
  "success": false,
  "message": "Movie not found"
}
```

Validation

```json
{
  "success": false,
  "errors": {
    "title": "Required"
  }
}
```

---

# 14. Configuration

Environment variables

```
APP_ENV

PORT

DATABASE_URL

JWT_SECRET

TMDB_API_KEY

CORS_ORIGIN
```

No secrets committed to Git.

---

# 15. Development Workflow

Every feature follows this order.

```
Migration

↓

SQL Query

↓

sqlc Generate

↓

Repository

↓

Service

↓

Handler

↓

API Testing

↓

Frontend Integration

↓

UI

↓

Manual Testing

↓

Commit
```

Never skip steps.

---

# 16. Git Strategy

```
main

develop

feature/*
```

Merge only after the feature is fully working.

---

# 17. Coding Conventions

## Backend

* Keep handlers thin.
* Business logic belongs in services.
* Repository only talks to the database.
* Prefer explicit code over abstraction.
* Keep functions short.
* Return early.
* Write SQL first.

---

## Frontend

* One component, one responsibility.
* Pages fetch data.
* Components display data.
* Pinia stores global state only.
* Prefer composables over duplicated logic.

---

# 18. Non-Goals (v1)

The following are intentionally out of scope:

* GraphQL
* Microservices
* Event sourcing
* CQRS
* Redis (until required)
* WebSockets
* AI recommendations
* Real-time synchronization
* Elasticsearch
* Multi-tenant support

---

# 19. Future Enhancements (v2+)

* Friend activity feed
* Movie night rooms
* Voting system
* Notifications
* Episode progress tracking
* Streaming provider integration
* Calendar view
* Import from Letterboxd/Trakt
* Public profiles
* Mobile app
* Offline support

---

# 20. Architecture Principles

These are the rules that should guide every technical decision:

1. **Keep it boring.** Prefer proven, simple solutions over clever ones.
2. **Single responsibility.** Every package, service, and component should have one clear job.
3. **Database-first thinking.** Design the data model before writing handlers.
4. **SQL is a feature, not a liability.** Favor clear SQL with `sqlc` over hiding queries behind heavy ORMs.
5. **Explicit over magical.** Code should be easy to trace from request to database.
6. **Vertical slices over horizontal layers.** Finish a feature end-to-end before starting another.
7. **Optimize only when needed.** Measure first; optimize second.
8. **Deployability matters.** If it can't comfortably run on a small VPS, question the complexity.
9. **The app serves the user, not the architecture.** Every abstraction should earn its place.
10. **If a future version needs more sophistication, evolve into it rather than designing for it upfront.**

I recommend keeping this file as your project's "constitution." Whenever you're unsure about a design decision, update this document first, then implement the change. It will keep the project coherent as it grows.
