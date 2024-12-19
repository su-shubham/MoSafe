# MoSafe Microservices Architecture

## Quick Start
```bash
# Clone and set up environment
git clone https://github.com/su-shubham/mosafe.git
cd mosafe
cp .env.example .env

# Start core infrastructure
docker compose -f docker-compose.infra.yml up -d

# Start services
docker compose up -d
```

## Core Services Overview

### API Gateway (Traefik)
- Entry point: `localhost:80`
- Dashboard: `localhost:8080`
- Configuration: `/traefik/config/`
- SSL termination and automatic cert management

### Authentication Service
- Port: 3001
- Database: PostgreSQL
- Features:
  - JWT token management
  - Role-based access control
  - Rate limiting
  - Session management

### Account Service
- Port: 3002
- Database: MongoDB
- Features:
  - Account CRUD operations
  - Balance management
  - Transaction history
  - Account validation

### Transaction Service
- Port: 3003
- Database: MongoDB
- Features:
  - Real-time transaction processing
  - Fraud detection
  - Transaction history
  - Payment processing

## Environment Configuration

Each service requires specific environment variables. Example configuration:

```env
# Auth Service
AUTH_JWT_SECRET=your_jwt_secret
AUTH_DB_URL=postgresql://user:pass@auth-db:5432/auth

# Account Service
ACCOUNT_DB_URL=mongodb://account-db:27017/accounts
REDIS_URL=redis://cache:6379

# Transaction Service
TRANSACTION_DB_URL=mongodb://transaction-db:27017/transactions
KAFKA_BROKERS=kafka:9092
```

## API Routes

### Authentication
```
POST   /api/auth/login
POST   /api/auth/register
POST   /api/auth/refresh
DELETE /api/auth/logout
GET    /api/auth/me
```

### Accounts
```
GET    /api/accounts
POST   /api/accounts
GET    /api/accounts/:id
PUT    /api/accounts/:id
DELETE /api/accounts/:id
GET    /api/accounts/:id/balance
```

### Transactions
```
GET    /api/transactions
POST   /api/transactions
GET    /api/transactions/:id
GET    /api/transactions/account/:accountId
POST   /api/transactions/transfer
```

## Error Handling

Standard error response format:
```json
{
  "error": {
    "code": "INVALID_ACCOUNT",
    "message": "Account not found",
    "details": {
      "accountId": "123",
      "reason": "Account has been closed"
    }
  },
  "requestId": "req_123abc",
  "timestamp": "2023-11-22T10:15:30Z"
}
```

## Database Schemas

### Account Collection
```javascript
{
  _id: ObjectId,
  userId: String,
  accountNumber: String,
  balance: Decimal128,
  currency: String,
  status: String,
  type: String,
  createdAt: Date,
  updatedAt: Date
}
```

### Transaction Collection
```javascript
{
  _id: ObjectId,
  sourceAccountId: ObjectId,
  destinationAccountId: ObjectId,
  amount: Decimal128,
  currency: String,
  type: String,
  status: String,
  metadata: Object,
  createdAt: Date
}
```

## Common Issues & Solutions

### Service Discovery Fails
```bash
# Check if Traefik is running
docker ps | grep traefik

# Verify network connectivity
docker network inspect mosafe_network

# Check service logs
docker compose logs -f auth-service
```

### Database Connection Issues
```bash
# Verify MongoDB connection
docker exec -it mongodb mongosh --eval "db.adminCommand('ping')"

# Check PostgreSQL
docker exec -it postgres pg_isready

# Reset database containers
docker compose down -v && docker compose up -d
```

## Security Considerations

- All inter-service communication uses mTLS
- Secrets managed via Docker secrets
- Regular security scanning with Trivy
- Rate limiting configured in Traefik
- CORS policies strictly enforced

## Performance Optimization

- Redis caching layer
- Connection pooling for databases
- Optimized Docker images (~100MB per service)
- CDN integration for static assets

## To Be Implemented

### Infrastructure
- [ ] Implement Kubernetes deployment manifests
- [ ] Set up GitLab CI/CD pipelines
- [ ] Configure automated backups for all databases
- [ ] Implement service mesh with Istio
- [ ] Set up ELK stack for centralized logging

### Security
- [ ] Implement OAuth2 authorization server
- [ ] Add API key rotation mechanism
- [ ] Set up WAF rules in Traefik
- [ ] Implement circuit breakers for all services
- [ ] Add request signature validation

### Monitoring
- [ ] Configure Grafana dashboards
- [ ] Set up alerting with PagerDuty
- [ ] Implement distributed tracing with Jaeger
- [ ] Add custom business metrics
- [ ] Set up SLO monitoring

### Performance
- [ ] Implement database sharding
- [ ] Add message queue for async operations
- [ ] Configure CDN for static assets
- [ ] Implement GraphQL federation
- [ ] Add server-side caching layer

### Testing
- [ ] Set up integration test environment
- [ ] Add contract testing with Pact
- [ ] Implement chaos engineering tests
- [ ] Add performance benchmarks
- [ ] Set up end-to-end testing pipeline

### Documentation
- [ ] Create API documentation with OpenAPI
- [ ] Add service dependency diagrams
- [ ] Document disaster recovery procedures
- [ ] Create runbooks for common issues
- [ ] Document security protocols
