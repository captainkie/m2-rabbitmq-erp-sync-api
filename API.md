# API Documentation

## Base URL

```
http://localhost:9545
```

## Authentication

All authenticated endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <your-token>
```

## Error Responses

All error responses follow this format:

```json
{
  "type": "ERROR_TYPE",
  "message": "Error message",
  "code": 400,
  "timestamp": "2024-01-20T10:30:00Z",
  "file": "file.go",
  "line": 42,
  "function": "FunctionName",
  "details": {}
}
```

Common error types:

- `VALIDATION_ERROR`: Input validation failed
- `NOT_FOUND`: Resource not found
- `UNAUTHORIZED`: Authentication required
- `FORBIDDEN`: Insufficient permissions
- `INTERNAL_ERROR`: Server error
- `BAD_REQUEST`: Invalid request
- `CONFLICT`: Resource conflict
- `RATE_LIMIT`: Rate limit exceeded

## Endpoints

### Health Check

#### GET /health

Basic health check endpoint.

**Response**

```json
{
  "status": "ok",
  "timestamp": "2024-01-20T10:30:00Z"
}
```

#### GET /health/live

Liveness check endpoint.

**Response**

```json
{
  "status": "ok",
  "timestamp": "2024-01-20T10:30:00Z"
}
```

#### GET /health/ready

Readiness check endpoint.

**Response**

```json
{
  "status": "ok",
  "timestamp": "2024-01-20T10:30:00Z"
}
```

### Authentication

#### POST /api/authentication/login

Login endpoint.

**Request Body**

```json
{
  "username": "string",
  "password": "string"
}
```

**Response**

```json
{
  "token": "string",
  "user": {
    "id": "string",
    "username": "string",
    "email": "string",
    "role": "string"
  }
}
```

#### POST /api/authentication/register

Registration endpoint.

**Request Body**

```json
{
  "username": "string",
  "email": "string",
  "password": "string"
}
```

**Response**

```json
{
  "id": "string",
  "username": "string",
  "email": "string",
  "role": "string"
}
```

### Users

#### GET /api/users

Get all users. Requires authentication.

**Query Parameters**

- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)
- `sort` (optional): Sort field (default: "created_at")
- `order` (optional): Sort order ("asc" or "desc", default: "desc")

**Response**

```json
{
  "data": [
    {
      "id": "string",
      "username": "string",
      "email": "string",
      "role": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  ],
  "pagination": {
    "total": 0,
    "page": 1,
    "limit": 10,
    "pages": 1
  }
}
```

#### GET /api/users/:id

Get user by ID. Requires authentication.

**Response**

```json
{
  "id": "string",
  "username": "string",
  "email": "string",
  "role": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

#### POST /api/users

Create new user. Requires authentication.

**Request Body**

```json
{
  "username": "string",
  "email": "string",
  "password": "string",
  "role": "string"
}
```

**Response**

```json
{
  "id": "string",
  "username": "string",
  "email": "string",
  "role": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

#### PATCH /api/users/:id

Update user. Requires authentication.

**Request Body**

```json
{
  "username": "string",
  "email": "string",
  "role": "string"
}
```

**Response**

```json
{
  "id": "string",
  "username": "string",
  "email": "string",
  "role": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

#### DELETE /api/users/:id

Delete user. Requires authentication.

**Response**

```json
{
  "message": "User deleted successfully"
}
```

### Queue Management

#### GET /api/queue/products

Sync products from ERP to Magento 2. Requires authentication.

**Query Parameters**

- `force` (optional): Force sync (default: false)

**Response**

```json
{
  "message": "Product sync started",
  "job_id": "string"
}
```

#### GET /api/queue/images

Sync images from ERP to Magento 2. Requires authentication.

**Query Parameters**

- `force` (optional): Force sync (default: false)

**Response**

```json
{
  "message": "Image sync started",
  "job_id": "string"
}
```

#### POST /api/queue/daily-sales

Sync daily sales from Magento 2 to ERP. Requires authentication.

**Request Body**

```json
{
  "date": "2024-01-20",
  "force": false
}
```

**Response**

```json
{
  "message": "Daily sales sync started",
  "job_id": "string"
}
```

## Rate Limiting

The API implements rate limiting to prevent abuse:

- Global rate limit: 1000 requests per minute
- IP-based rate limit: 100 requests per minute per IP
- Route-specific rate limits may apply

When rate limit is exceeded, the API returns:

```json
{
  "type": "RATE_LIMIT",
  "message": "Rate limit exceeded. Please try again later.",
  "code": 429,
  "timestamp": "2024-01-20T10:30:00Z",
  "details": {
    "retry_after": 300
  }
}
```

## WebSocket Events

The API provides real-time updates via WebSocket:

```
ws://localhost:9545/ws
```

### Events

#### Queue Status Updates

```json
{
  "type": "queue_status",
  "data": {
    "job_id": "string",
    "status": "processing|completed|failed",
    "progress": 0,
    "message": "string"
  }
}
```

#### Error Updates

```json
{
  "type": "error",
  "data": {
    "job_id": "string",
    "error": "string",
    "timestamp": "string"
  }
}
```

## Pagination

All list endpoints support pagination with the following query parameters:

- `page`: Page number (default: 1)
- `limit`: Items per page (default: 10)
- `sort`: Sort field (default: "created_at")
- `order`: Sort order ("asc" or "desc", default: "desc")

Response includes pagination metadata:

```json
{
  "data": [],
  "pagination": {
    "total": 0,
    "page": 1,
    "limit": 10,
    "pages": 1
  }
}
```

## Filtering

List endpoints support filtering with query parameters:

```
GET /api/users?role=admin&created_after=2024-01-01
```

## Sorting

List endpoints support sorting with query parameters:

```
GET /api/users?sort=username&order=asc
```

## Response Headers

Common response headers:

- `X-Request-ID`: Unique request identifier
- `X-RateLimit-Limit`: Rate limit for the current window
- `X-RateLimit-Remaining`: Remaining requests in the current window
- `X-RateLimit-Reset`: Time when the rate limit resets
