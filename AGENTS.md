# AI Agent Instructions - Prediction Market Platform

## Project Overview
Build a prediction market platform (like Kalshi) where users trade binary outcome contracts on real-world events. Frontend: Angular, Backend: Go, Database: PostgreSQL, Cache: Redis.

---

## Agent Roles & Responsibilities

### Agent 1: Backend Infrastructure Engineer
**Primary Focus**: Go backend, database, API architecture

**Tasks**:
1. Initialize Go project with clean architecture (handler → service → repository pattern)
2. Set up PostgreSQL with migrations for all tables (users, accounts, markets, orders, positions, transactions, resolutions, audit_logs)
3. Configure Redis for caching and pub/sub
4. Build REST API with Gin/Echo framework
5. Implement JWT authentication middleware
6. Create all API endpoints (auth, markets, orders, portfolio, admin)
7. Build the order matching engine (price-time priority, in-memory order book)
8. Implement WebSocket server for real-time updates (Gorilla WebSocket)
9. Add comprehensive error handling and logging (Zap/Logrus)
10. Write unit and integration tests

**Key Deliverables**:
- Complete Go backend with all API endpoints
- Order matching engine with sub-10ms latency
- WebSocket server for real-time data
- Database schema and migrations
- API documentation (OpenAPI/Swagger)

**Technical Requirements**:
- Use dependency injection
- Implement graceful shutdown
- Add rate limiting and security headers
- Use database transactions for atomic operations
- Implement connection pooling
- Add health check endpoints

---

### Agent 2: Frontend Engineer
**Primary Focus**: Angular application, UI/UX, real-time features

**Tasks**:
1. Initialize Angular project (v17+) with NgRx state management
2. Set up TailwindCSS/Angular Material for UI
3. Create authentication module (login, register, password reset, 2FA)
4. Build markets listing page with filters, search, and infinite scroll
5. Create market detail page with real-time price charts
6. Implement order book visualization with live updates
7. Build trading interface (buy/sell forms, order preview, execution)
8. Create portfolio dashboard with P&L tracking
9. Implement WebSocket service with auto-reconnection
10. Build admin panel for market creation and resolution
11. Add payment integration UI (Stripe Elements, Plaid Link)
12. Create mobile-responsive layouts
13. Write component tests and E2E tests

**Key Deliverables**:
- Complete Angular SPA with all user-facing features
- Real-time order book and price updates
- Portfolio management interface
- Admin dashboard
- Mobile-responsive design
- Comprehensive test coverage

**Technical Requirements**:
- Use Angular CLI commands (for example `ng g c`, `ng g m`, `ng g service`) for every scaffold so HTML, CSS, and TypeScript stay in their generated files.
- Keep templates, styles, and logic separated (no inline templates or styles) and favor small, modular components.
- Define shared interfaces and types in dedicated libraries or state folders instead of component files to keep reuse high.
- Enforce strict linting and unit or E2E tests in CI; do not merge frontend work without green `ng lint`, `ng test`, and Cypress runs.

- Use NgRx for state management
- Implement lazy loading for routes
- Add loading states and error handling
- Use RxJS for reactive programming
- Implement auth guards and interceptors
- Optimize for performance (OnPush change detection)

---

### Agent 3: DevOps & Infrastructure Engineer
**Primary Focus**: Deployment, CI/CD, monitoring, scaling

**Tasks**:
1. Create Dockerfiles for backend and frontend
2. Set up Kubernetes manifests (deployments, services, ingress)
3. Configure CI/CD pipeline (GitHub Actions)
4. Set up PostgreSQL with backups and replication
5. Configure Redis cluster for high availability
6. Implement monitoring (Prometheus + Grafana)
7. Set up logging aggregation (ELK stack or Loki)
8. Configure error tracking (Sentry)
9. Set up CDN for static assets
10. Implement SSL/TLS with cert management
11. Configure auto-scaling policies
12. Create disaster recovery procedures
13. Set up staging and production environments

**Key Deliverables**:
- Docker containers for all services
- Kubernetes cluster with auto-scaling
- CI/CD pipeline with automated testing
- Monitoring dashboards
- Logging and error tracking
- Production-ready infrastructure

**Technical Requirements**:
- Use Kubernetes secrets for sensitive data
- Implement blue-green deployments
- Add database connection pooling (PgBouncer)
- Configure load balancing
- Set up health checks and readiness probes
- Implement resource limits and requests

---

### Agent 4: Security & Compliance Engineer
**Primary Focus**: Security, KYC/AML, regulatory compliance

**Tasks**:
1. Implement KYC verification workflow (Persona/Onfido integration)
2. Build AML screening integration
3. Add identity document verification
4. Implement trading limits and enforcement
5. Create suspicious activity monitoring
6. Add IP geolocation blocking for restricted jurisdictions
7. Implement comprehensive audit logging
8. Add GDPR compliance features (data export/deletion)
9. Conduct security audit (OWASP vulnerabilities)
10. Set up bug bounty program
11. Implement DDoS protection
12. Add encryption for sensitive data
13. Create terms of service and privacy policy

**Key Deliverables**:
- Complete KYC/AML workflow
- Security audit report with fixes
- Compliance documentation
- Data privacy features
- Audit trail system
- Security monitoring

**Technical Requirements**:
- Use bcrypt for password hashing
- Implement TOTP for 2FA
- Add CSRF protection
- Use parameterized queries (SQL injection prevention)
- Implement rate limiting per user
- Add security headers (HSTS, CSP, etc.)

---

### Agent 5: Testing & Quality Assurance Engineer
**Primary Focus**: Testing, quality assurance, performance optimization

**Tasks**:
1. Write unit tests for Go backend (90%+ coverage)
2. Write component tests for Angular (Jest/Jasmine)
3. Create integration tests for API endpoints
4. Build E2E tests for critical user flows (Cypress/Playwright)
5. Conduct load testing (k6/Artillery)
6. Test order matching engine under stress
7. Perform WebSocket stress testing
8. Test database migration rollbacks
9. Conduct accessibility testing (WCAG 2.1)
10. Test cross-browser compatibility
11. Perform mobile responsiveness testing
12. Run security testing (penetration testing)
13. Conduct chaos engineering tests
14. Optimize performance (Lighthouse scores)

**Key Deliverables**:
- Comprehensive test suite with high coverage
- Load testing reports
- Performance optimization recommendations
- Accessibility compliance report
- Bug reports and fixes
- QA documentation

**Technical Requirements**:
- Use testify for Go testing
- Use Jest for Angular unit tests
- Use Cypress for E2E tests
- Achieve >80% code coverage
- Test error scenarios and edge cases
- Automate tests in CI/CD pipeline

---

## Cross-Agent Collaboration Points

### Integration Points
1. **Backend ↔ Frontend**: API contract (OpenAPI spec), WebSocket protocol
2. **Backend ↔ DevOps**: Environment configuration, secrets management, deployment strategy
3. **Backend ↔ Security**: Authentication flow, audit logging, encryption
4. **Frontend ↔ Security**: KYC UI, compliance notifications, secure forms
5. **All ↔ Testing**: Test data, test environments, automated testing

### Communication Protocol
- Use standardized API responses (JSON with consistent error format)
- Document all WebSocket message types and payloads
- Share environment variable requirements
- Coordinate database schema changes
- Align on coding standards and conventions

---

## Priority Order

### Phase 1 (Critical Path)
1. Backend: Database setup, authentication, basic API
2. Frontend: Authentication UI, basic layout
3. DevOps: Development environment, CI/CD basics

### Phase 2 (Core Features)
1. Backend: Market APIs, order matching engine
2. Frontend: Market listing, market details
3. Testing: Integration tests for critical paths

### Phase 3 (Trading)
1. Backend: Order placement, WebSocket updates
2. Frontend: Order book, trading interface
3. Testing: Load testing for trading engine

### Phase 4 (Finalization)
1. Backend: Payment integration, admin APIs
2. Frontend: Portfolio, admin panel
3. Security: KYC/AML implementation
4. DevOps: Production deployment
5. Testing: Full E2E and security testing

---

## Success Criteria

### Performance Targets
- API response time: <100ms (p99)
- Order matching latency: <10ms
- WebSocket update delivery: <50ms
- Database query time: <50ms (p95)
- Page load time: <2s (Lighthouse score >90)

### Quality Targets
- Code coverage: >80% backend, >70% frontend
- Zero critical security vulnerabilities
- 99.9% uptime SLA
- Zero data loss incidents
- WCAG 2.1 AA compliance

### Feature Completeness
- All 12 phases from checklist completed
- All critical user flows tested
- Documentation complete (API, deployment, user guides)
- Monitoring and alerting operational
- Disaster recovery tested

---

## Agent Best Practices

### General Guidelines
1. **Write Clean Code**: Follow SOLID principles, use meaningful names
2. **Test First**: Write tests before or alongside feature development
3. **Document Everything**: Code comments, API docs, README files
4. **Use Version Control**: Commit frequently with descriptive messages
5. **Handle Errors Gracefully**: Never expose internal errors to users
6. **Log Appropriately**: Use structured logging with proper levels
7. **Optimize Performance**: Profile before optimizing, measure results
8. **Security First**: Validate all inputs, sanitize all outputs
9. **Think Scale**: Design for horizontal scalability
10. **Collaborate**: Review other agents' work, provide feedback
11. **Guard Rails**: Treat lint, unit, integration, and performance tests as non-negotiable gates; do not merge code until every check is green and budgets are met.

### Code Review Checklist
- [ ] Code follows project conventions
- [ ] Tests are comprehensive and passing
- [ ] Documentation is updated
- [ ] Security best practices followed
- [ ] Performance considerations addressed
- [ ] Error handling is robust
- [ ] No hardcoded secrets or credentials
- [ ] Logs are meaningful and structured
- [ ] Database queries are optimized
- [ ] API changes are backward compatible (when possible)

---

## Technology Stack Reference

### Backend (Go)
```go
// Core libraries
gin-gonic/gin or labstack/echo  // Web framework
golang-jwt/jwt                   // JWT tokens
gorilla/websocket               // WebSocket
lib/pq or jackc/pgx             // PostgreSQL
go-redis/redis                  // Redis client
spf13/viper                     // Configuration
uber-go/zap                     // Logging
golang-migrate/migrate          // Migrations
stretchr/testify               // Testing
```

### Frontend (Angular)
```typescript
// Core dependencies
@angular/core, @angular/router
@ngrx/store, @ngrx/effects     // State management
@angular/material               // UI components
tailwindcss                     // Styling
chart.js or d3                  // Charts
rxjs                            // Reactive programming
@stripe/stripe-js               // Payments
socket.io-client or native WS   // WebSocket
```

### DevOps
```yaml
# Infrastructure
Docker
Kubernetes
PostgreSQL 15+
Redis 7+
Nginx/Traefik (Ingress)
Prometheus + Grafana
ELK Stack or Loki
Sentry
GitHub Actions or GitLab CI
```

---

## Emergency Procedures

### Critical Bugs in Production
1. Identify affected components
2. Roll back to last stable version
3. Create hotfix branch
4. Test fix thoroughly
5. Deploy with monitoring
6. Conduct post-mortem

### Performance Degradation
1. Check monitoring dashboards
2. Identify bottleneck (DB, API, network)
3. Scale affected services
4. Optimize queries/code
5. Add caching where appropriate

### Security Incident
1. Isolate affected systems
2. Assess scope and impact
3. Notify users if data compromised
4. Apply security patches
5. Conduct forensic analysis
6. Update security procedures

---

## Handoff Documentation Required

Each agent must provide:
1. **README**: Setup instructions, architecture overview
2. **API Documentation**: All endpoints with examples
3. **Deployment Guide**: Step-by-step deployment process
4. **Configuration Guide**: All environment variables explained
5. **Troubleshooting Guide**: Common issues and solutions
6. **Test Documentation**: How to run tests, coverage reports
7. **Performance Benchmarks**: Baseline metrics and targets
8. **Security Audit**: Findings and remediation status

---

## Final Checklist Before Launch

- [ ] All features from development checklist completed
- [ ] All tests passing (unit, integration, E2E)
- [ ] Security audit completed and vulnerabilities fixed
- [ ] Performance testing shows acceptable metrics
- [ ] Database backups configured and tested
- [ ] Monitoring and alerting operational
- [ ] Error tracking configured
- [ ] Documentation complete and reviewed
- [ ] Disaster recovery plan tested
- [ ] Compliance requirements met (KYC/AML)
- [ ] Terms of service and privacy policy finalized
- [ ] Customer support system ready
- [ ] Rollback procedures documented and tested
- [ ] Team trained on incident response
- [ ] Beta testing completed with feedback incorporated