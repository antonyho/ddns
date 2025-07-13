# Task: Phase 1 - Foundation Setup

## Objective
Set up the foundation for the dynamic DNS platform by implementing shared packages, basic project structure, DNS service, and database schema.

## Todo Items

### 1. Project Structure & Build Setup
- [ ] Create root Go module and Makefile
- [ ] Set up Docker Compose for local development (PostgreSQL + Redis)
- [ ] Create basic project documentation structure

### 2. Shared Packages Implementation
- [ ] Create `pkg/models` - Database models and shared structs
- [ ] Create `pkg/jwt` - JWT validation with RS256/ES256 support
- [ ] Create `pkg/redis` - Redis connection and utilities
- [ ] Create `pkg/database` - PostgreSQL connection and utilities
- [ ] Add unit tests for all shared packages

### 3. Database Foundation
- [ ] Create database schema migrations
- [ ] Implement migration runner
- [ ] Add seed data for testing

### 4. DNS Service Implementation
- [ ] Create DNS service with Viper configuration
- [ ] Implement basic DNS query handling (A, AAAA, CNAME records)
- [ ] Add Redis cache integration with PostgreSQL fallback
- [ ] Set 60-second TTL for all DNS responses
- [ ] Add comprehensive logging
- [ ] Create unit and integration tests

### 5. Cache Service Implementation
- [ ] Create cache service with Redis Streams subscription
- [ ] Implement cache invalidation logic
- [ ] Add database record synchronization
- [ ] Handle cache startup from PostgreSQL
- [ ] Add error handling and retry logic
- [ ] Create unit tests

### 6. Development Environment
- [ ] Create Docker containers for all services
- [ ] Set up health check endpoints
- [ ] Implement graceful shutdown handling
- [ ] Add development scripts and documentation

## Implementation Notes

### Key Requirements
- All services use Viper for configuration
- Redis DB separation: 0=DNS cache, 1=Streams, 2=Rate limiting, 3=Sessions
- JWT with 5-minute expiration, 30-second clock skew
- Comprehensive error handling and logging
- Unit tests for all business logic

### Technology Decisions
- Go 1.21+ for all backend services
- `github.com/miekg/dns` for DNS functionality
- PostgreSQL for persistent storage
- Redis for caching and message streams
- Viper for configuration management

### Security Considerations
- Input validation on all inputs
- SQL injection prevention with prepared statements
- Rate limiting preparation
- Structured logging without sensitive data

## Review
[To be filled upon completion]