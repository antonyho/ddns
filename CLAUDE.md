# Dynamic DNS Platform Implementation Guide

## Project Overview

Build a microservices-based dynamic DNS platform that provides free subdomains and custom domain DNS management with JWT-based API authentication using asymmetric key pairs.

## General Working Guidelines
1. Ignore the file `GEMINI.md`.
2. First think through the problem, read the codebase for relevant files, and write a plan to `tasks/todo.md`.
3. The plan should have a list of todo items that you can check off as you complete them.
4. Before you begin working, check in with me and I will verify the plan.
5. Then, begin working on the todo items, marking them as complete as you go.
6. On every step of progress, give me a high level explanation of what changes you have made.
7. Make every task and code change you do as simple as possible. We want to avoid making any massive or complex changes. Every change should impact as little code as possible. Everything is about simplicity.
8. Finally, add a review section to the todo.md file with a summary of the changes you made and any other relevant information.

## Architecture Summary

### Microservices:
1. **DNS Service** - Handles DNS queries
2. **Cache Service** - Listen to update event, caches records from Redis, stores records in PostgreSQL, loads records from PostgreSQL on start and updates cache records
3. **API Service** - RESTful API for DNS record management with JWT auth, publish DNS update events
4. **User Service** - RESTful API for user registration, email verification, public key management
5. **Worker Service** - Background jobs (emails, SSL certificates)
6. **Web Frontend** - Next.js dashboard for User and DNS managements

### Technology Stack:
- **Backend**: Go (all services except frontend)
- **Frontend**: Next.js with TypeScript
- **Database**: PostgreSQL (shared)
- **Cache/Queue**: Redis (caching + message streams)
- **DNS Library**: `github.com/miekg/dns`
- **Email**: AWS SES
- **Container**: Docker + Kubernetes(production service) + Docker Compose(local development)

## Implementation Order

### Phase 1: Foundation
1. Project structure setup
2. Shared packages (JWT, Redis, models)
3. DNS Service implementation
4. Database schema and migrations

### Phase 2: Core Features
1. API Service with JWT authentication
2. User Service with email verification
3. Worker Service for background tasks
4. Basic Web Frontend

### Phase 3: Production Features
1. Let's Encrypt integration
2. Rate limiting
3. Monitoring and logging
4. Performance optimization

## Project Structure

```
. (project root)
├── services/
│   ├── dns/
│   ├── cache/
│   ├── api/
│   ├── user/
│   ├── worker/
│   └── web/
├── pkg/
│   ├── models/
│   ├── jwt/
│   └── redis/
├── database/
│   └── migrations/
├── tasks/
│   └── todo.md
├── Makefile
└── README.md
```

## Code Quality Requirements (Priority Order)

1. **Good Practices**: Follow best practices
2. **Test Coverage**: implemented features must have unit tests
3. **Error Handling**: Always check errors, wrap with context
4. **Logging**: Structured logging with appropriate levels
5. **Documentation**: GoDoc comments for exported functions
6. **Security**: Input validation, SQL injection prevention, Rate limit against DDoS
7. **Performance**: Connection pooling, efficient queries

## Implementation Guidelines

### 1. Start with Shared Packages

Create common models.

### 2. DNS Service Implementation

Focus on:
- Viper-based configuration
- Clean separation of concerns
- Retrieve DNS records from Redis cache with PostgreSQL fallback
- Minimal 60-second TTL for all records
- Comprehensive logging

### 3. Cache Service Implementation

Focus on:
- Viper-based configuration
- Redis Streams subscriptions for cache invalidation and database record update

### 4. Database Schema

Create database schema.

### 5. Redis Configuration

Use Redis databases for separation:
- DB 0: DNS record cache
- DB 1: Message streams
- DB 2: Rate limiting
- DB 3: Web sessions

### 6. JWT Implementation

Create a robust JWT validation system:
- Support RS256 and ES256 algorithms
- Validate expiration with 30-second clock skew
- Check JTI blacklist
- Verify scope permissions
- Cache public keys in memory with TTL

### 7. API Conventions

- RESTful endpoints: `/api/v1/dns-records`, `/api/v1/domains`
- Consistent error responses: `{"error": "message", "code": "ERROR_CODE"}`
- Request/response logging middleware
- Proper HTTP status codes
- Publish DNS update events

### 8. Docker Setup

Each service should have:
- Multi-stage Dockerfile for minimal images
- Health check endpoints
- Graceful shutdown handling
- Environment-based configuration

### 9. Testing Strategy

- Unit tests for business logic
- Integration tests with test containers
- DNS query testing with mock data
- JWT validation edge cases

## Development Workflow

1. Run `make docker-up` to start PostgreSQL and Redis
2. Run `make migrate-up` to apply database schema
3. Develop individual services with hot reloading
4. Test with `make test`
5. Build containers with `make docker-build-all`

## Important Considerations

- DNS service must never have downtime
- JWT tokens expire in 5 minutes (configurable later based on user feedback)
- All timestamps in UTC
- Rate limiting from the start
- Prepare for horizontal scaling
- Cache invalidation must be reliable

## Task Planning Template

When creating `tasks/todo.md`, use this template:

```markdown
# Task: [Task Name]

## Objective
Brief description of what needs to be accomplished.

## Todo Items
- [ ] Task 1: Description
- [ ] Task 2: Description
- [ ] Task 3: Description

## Implementation Notes
Any relevant notes or considerations.

## Review
Summary of changes made and any relevant information.
```

## First Steps

1. Create the project structure
2. Implement shared packages with tests
3. Build DNS service with hardcoded responses
4. Add Redis caching with tests
5. Implement database integration
6. Add Redis Streams for cache invalidation
7. Create comprehensive tests for each component
8. Document API endpoints

Remember: 
- Start simple, make it work, then optimize
- Every change should be minimal and focused
- Write tests for all features
- The DNS service is the most critical component - it must be rock-solid before moving to other services

## Planned Features

- User's custom domain name (Not supported yet)
    - User's domain name ownership verification must be required in order to support this feature