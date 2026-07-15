# FrameLog — Development Roadmap (Milestone Canvas)

---

# Before Touching Go or Vue

This is the only "study first" section.

**Don't spend weeks watching tutorials.**

Your goal is to know *just enough* to start building.

## Go (2–3 days max)

### Learn

* [ ] Project structure (`go mod`)
* [ ] Packages
* [ ] Functions
* [ ] Structs
* [ ] Methods
* [ ] Interfaces (basic understanding only)
* [ ] Error handling
* [ ] Pointers (basic)
* [ ] Slices & Maps
* [ ] JSON marshaling
* [ ] `context.Context`
* [ ] Goroutines (only know what they are)
* [ ] Standard library (`net/http`, `os`, `log`, `time`)

### Don't Learn Yet

❌ Channels

❌ Reflection

❌ Generics

❌ Concurrency patterns

❌ Clean Architecture

❌ DDD

❌ Dependency Injection frameworks

---

## Gin

Just learn

* [ ] Router
* [ ] Middleware
* [ ] Route Groups
* [ ] Request Binding
* [ ] JSON Response

---

## PostgreSQL

Must know

* [ ] SELECT
* [ ] INSERT
* [ ] UPDATE
* [ ] DELETE
* [ ] JOIN
* [ ] GROUP BY
* [ ] ORDER BY
* [ ] LIMIT
* [ ] Foreign Keys

---

## sqlc

Only

* [ ] Write SQL
* [ ] Generate Go code

Nothing more.

---

## Vue

Only refresh

* [ ] Composition API
* [ ] `<script setup>`
* [ ] Reactive state
* [ ] Components
* [ ] Props
* [ ] Emits
* [ ] Router
* [ ] Pinia
* [ ] Fetch API

Don't optimize yet.

---

# Development Philosophy

Every milestone must produce a working application.

Never break main.

Never leave half-finished features.

---

# Milestone 0 — Foundation

## Objective

Create a runnable development environment.

---

## Tasks

Backend

* [x] Initialize Go module
* [x] Install Gin
* [ ] Connect PostgreSQL
* [ ] Docker Compose
* [x] Health endpoint

Frontend

* [x] Create Vue project
* [ ] Install Tailwind
* [x] Install Pinia
* [x] Install Router

Project

* [ ] Git repository
* [ ] README
* [ ] Environment variables
* [ ] Folder structure

---

## Done When

* Backend starts
* Frontend starts
* PostgreSQL connected
* `/health` returns OK
* Docker Compose works

---

## Git Commit

```
chore: bootstrap project foundation
```

---

## Next

Authentication

---

# Milestone 1 — Authentication

## Objective

Users can create accounts and login.

---

## Tasks

Backend

* [ ] User table
* [ ] Password hashing
* [ ] Register endpoint
* [ ] Login endpoint
* [ ] JWT generation
* [ ] Auth middleware

Frontend

* [ ] Register page
* [ ] Login page
* [ ] Logout
* [ ] Protected routes
* [ ] User store

---

## Done When

A new user can

* Register
* Login
* Stay logged in
* Logout

---

## Git Commit

```
feat: implement authentication
```

---

## Next

Media database

---

# Milestone 2 — Media Search

## Objective

Users can search and import media.

---

## Tasks

Backend

* [ ] TMDB integration
* [ ] Search endpoint
* [ ] Save media
* [ ] Get media details

Frontend

* [ ] Search page
* [ ] Results page
* [ ] Media page
* [ ] Add button

---

## Done When

Search works.

Movies can be imported.

Media page opens.

---

## Git Commit

```
feat: add media search and import
```

---

## Next

Tracking

---

# Milestone 3 — Watch Tracking

## Objective

Track viewing status.

---

## Tasks

Statuses

* [ ] Want To Watch
* [ ] Watching
* [ ] Completed
* [ ] Dropped

Backend

* [ ] WatchEntry table
* [ ] CRUD endpoints

Frontend

* [ ] Status selector
* [ ] Dashboard widgets

---

## Done When

User can move movies through all statuses.

Dashboard updates.

---

## Git Commit

```
feat: implement watch tracking
```

---

## Next

Ratings

---

# Milestone 4 — Ratings & Reviews

## Objective

Allow personal opinions.

---

## Tasks

* [ ] Rating
* [ ] Review
* [ ] Notes
* [ ] Started date
* [ ] Finished date

Frontend

* [ ] Rating stars
* [ ] Review editor

---

## Done When

Movies can be reviewed.

Ratings appear.

---

## Git Commit

```
feat: add ratings and reviews
```

---

## Next

Dashboard

---

# Milestone 5 — Dashboard

## Objective

Personal home page.

---

## Tasks

Widgets

* [ ] Watching
* [ ] Watchlist
* [ ] Recently Finished
* [ ] Recent Reviews

Backend

* [ ] Dashboard endpoint

---

## Done When

Dashboard feels useful.

---

## Git Commit

```
feat: build dashboard
```

---

## Next

Statistics

---

# Milestone 6 — Statistics

## Objective

Visualize viewing habits.

---

## Tasks

Backend

* [ ] Monthly stats
* [ ] Yearly stats
* [ ] Genre counts
* [ ] Runtime

Frontend

* [ ] Charts
* [ ] Year selector

---

## Done When

User sees yearly recap.

---

## Git Commit

```
feat: add statistics dashboard
```

---

## Next

Lists

---

# Milestone 7 — Lists

## Objective

Organize media.

---

## Tasks

* [ ] Favorite list
* [ ] Custom lists
* [ ] Add/remove media
* [ ] List page

---

## Done When

Users create their own collections.

---

## Git Commit

```
feat: implement custom lists
```

---

## Next

Recommendations

---

# Milestone 8 — Recommendations

## Objective

Share movies.

---

## Tasks

Backend

* [ ] Recommendation table
* [ ] Send recommendation
* [ ] Recommendation inbox

Frontend

* [ ] Recommend dialog
* [ ] Inbox

---

## Done When

Friend receives recommendation.

---

## Git Commit

```
feat: add recommendation system
```

---

## Next

Friends

---

# Milestone 9 — Friends

## Objective

Basic social layer.

---

## Tasks

* [ ] Friend requests
* [ ] Accept request
* [ ] Profile page
* [ ] Activity feed

---

## Done When

Users follow each other.

---

## Git Commit

```
feat: implement friends and activity feed
```

---

## Next

Movie Night

---

# Milestone 10 — Movie Night

## Objective

Watch together.

---

## Tasks

* [ ] Create room
* [ ] Invite friends
* [ ] Vote
* [ ] Winner selection

---

## Done When

Users can vote on a movie.

---

## Git Commit

```
feat: add movie night feature
```

---

## Next

Polish

---

# Milestone 11 — Production Ready

## Objective

Make FrameLog deployable.

---

## Tasks

Backend

* [ ] Logging
* [ ] Error handling
* [ ] Validation
* [ ] Rate limiting

Frontend

* [ ] Loading states
* [ ] Error states
* [ ] Empty states
* [ ] Responsive UI

Infrastructure

* [ ] Docker
* [ ] CI
* [ ] VPS deployment
* [ ] HTTPS

---

## Done When

FrameLog is online.

---

## Git Commit

```
chore: prepare production release
```

---

# Every Feature Workflow

```
1. Design database

↓

2. Write migration

↓

3. Write SQL

↓

4. Generate sqlc

↓

5. Repository

↓

6. Service

↓

7. Handler

↓

8. API test

↓

9. Frontend API integration

↓

10. UI

↓

11. Loading state

↓

12. Error state

↓

13. Empty state

↓

14. Manual testing

↓

15. Commit

↓

16. Merge
```

---

# Project Rules (Our Team Agreement)

1. **No premature optimization.**
2. **No refactoring without a reason.**
3. **One milestone at a time.**
4. **Every coding session ends with a working app.**
5. **Don't introduce a new dependency unless it solves a current problem.**
6. **If a task takes more than one session, split it into smaller tasks.**
7. **`main` is always deployable.**
8. **Finish before polishing.**
9. **Document architectural decisions as they arise instead of redesigning from scratch.**
10. **Celebrate completed milestones before planning the next one.**

---

## My suggestion: one more document

I think this roadmap should be your **execution plan**, not your **specification**.

I'd create two living documents:

1. **`ROADMAP.md`** — the milestone plan above (what to build and in what order).
2. **`ARCHITECTURE.md`** — your project's source of truth for:

   * Database schema and ER diagrams
   * API endpoint contracts
   * Folder structure
   * Naming conventions
   * Coding standards
   * Authentication flow
   * Deployment architecture
   * Design decisions and rationale (ADR-style)

This separation helps a lot, especially if you like structure: the roadmap tells you **what's next**, while the architecture document tells you **how things fit together**. As the project grows, you'll always know where to look instead of trying to keep everything in your head.
