# 📺 GraphQL VOD Platform

A minimalist **Video-on-Demand** GraphQL API built in Go using `gqlgen`, PostgreSQL, Redis Pub/Sub, and JWT-based authentication. Users can register, upload videos, and post comments in real time.

## ⚙️ Features

- JWT-based authentication for secure mutations (`register`, `login`, `uploadVideo`, `postComment`)
- Video upload and metadata storage via PostgreSQL (using GORM)
- Real‑time comments with GraphQL Subscriptions backed by Redis
- Type-safe API with `gqlgen` generated schema and resolvers

## 🚀 Quick Start

### Prerequisites

- Go 1.24+
- Docker & Docker Compose

### Local launch

```bash
git clone https://github.com/KingBean4903/graphql-vod-platform.git
cd graphql-vod-platform
docker-compose up --build

```

By default
- Server runs on http://localhost:8800/query
- Playground available at http://localhost:8800/

## Auth Workflow
1. Register with register mutation → returns user object
2. Login with login mutation → returns JWT in token
3. Set HTTP headers for requests

Protected operations like uploadVideo and postComment require a valid JWT.

## Example usage:
# Register and login
mutation {
  register(input: {username:"alice", email:"a@example.com", password:"pass"}) {
    user {
      id
    }
  }
}

mutation {
  login(email:"a@example.com", password:"pass") {
    token
    user {
      id
      username
    }
  }
}

# Upload a video (with JWT header)
mutation {
  uploadVideo(input: {title:"My Vid", description:"Demo", url:"http://..."}) {
    id
    title
  }
}

# Subscribe and post a comment
subscription {
  commentAdded(videoID:"1") {
    id
    text
    author {
      username
    }
  }
}

mutation {
  postComment(videoID:"1", text:"Nice!") {
    id
    text
  }
}

## Development Notes
- After modifying .graphqls schema, run:

```bash
go generate ./...

```
