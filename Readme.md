# Go Fiber Blog Backend API
A production-ready RESTful API built in Go (Fiber) that powers a modern blogging platform.
It handles user authentication, post management, comments, and likes, following clean architecture and best practices.

## Features
- User Authentication (Sign Up, Sign In)
- Posts CRUD (Create, Read, Update, Delete)
- Comments on posts
- Likes for posts (one like per user per post)
- JWT Authentication & Authorization Middleware
- PostgreSQL Database Integration
- Password Hashing with bcrypt
- Graceful error handling and validation

TODO
- Setup Monitoring using grafana, prometheus, loki, tempo
- Containerize using Docker
- Deploy To Kubernetes Cluster

## Tech Stack
| Layer             | Technology              |
| ----------------- | ----------------------- |
| Language          | Go                      |
| Web Framework     | Fiber                   |
| Database          | PostgreSQL              |
| Auth              | JWT (JSON Web Tokens)   |
| Password Security | bcrypt                  |
| ORM/DB Access     | `database/sql`          |
| Deployment        | Docker, Kubernetes      |

## Database Schema
`USERS`
| Column     | Type      | Description      |
| ---------- | --------- | ---------------- |
| user_id    | BIGINT    | Primary key      |
| name       | TEXT      | User’s full name |
| email      | TEXT      | Unique email     |
| password   | TEXT      | Hashed password  |
| created_at | TIMESTAMP | Creation time    |

`POSTS`
| Column      | Type      | Description               |
| ----------- | --------- | ------------------------- |
| post_id     | BIGINT    | Primary key               |
| title       | VARCHAR   | Post title                |
| description | TEXT      | Post content              |
| author_id   | BIGINT    | References users(user_id) |
| created_at  | TIMESTAMP | Created date              |
| updated_at  | TIMESTAMP | Last updated date         |

`COMMENTS`
| Column     | Type      | Description               |
| ---------- | --------- | ------------------------- |
| comment_id | BIGINT    | Primary key               |
| post_id    | BIGINT    | References posts(post_id) |
| author_id  | BIGINT    | References users(user_id) |
| text       | TEXT      | Comment text              |
| created_at | TIMESTAMP | Created date              |
| updated_at | TIMESTAMP | Updated date              |

`LIKES`

| Column                   | Type      | Description                       |
| ------------------------ | --------- | --------------------------------- |
| like_id                  | BIGINT    | Primary key                       |
| user_id                  | BIGINT    | References users(user_id)         |
| post_id                  | BIGINT    | References posts(post_id)         |
| created_at               | TIMESTAMP | When post was liked               |
| UNIQUE(user_id, post_id) |           | Ensures a user can only like once |

