# WebSync API

A Go-based API service for synchronizing product data between ERP and Magento 2 systems.

## Features

- 🔐 Authentication with JWT
- 🐇 Queue management with RabbitMQ
- 🔄 Product synchronization (ERP to Magento 2)
- 🖼️ Image synchronization (ERP to Magento 2)
- 📊 Daily sales synchronization (Magento to ERP)
- ⚡ Rate limiting
- 🔒 CORS support
- 🏥 Health check endpoints
- 📝 Swagger documentation

## Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- RabbitMQ 3.9 or higher
- Docker (optional)

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
# Application
APP_PORT=9545
GIN_MODE=debug # or release

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=password
DB_NAME=websync

# RabbitMQ
RABBITMQ_URL=amqp://guest:guest@localhost:5672/

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# AWS S3 (for image storage)
AWS_S3_REGION=ap-southeast-1
AWS_S3_ACCESS_KEY_ID=your-access-key
AWS_S3_SECRET_ACCESS_KEY=your-secret-key
AWS_S3_BUCKET=your-bucket

# Service URLs
STRAPI_URL=http://localhost:1337
SERVICE_URL=http://your-service-url
```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/captainkie/websync-api.git
cd websync-api
```

2. Install dependencies:

```bash
go mod download
```

3. Generate Swagger documentation:

```bash
swag init --pd -g ./cmd/main/main.go -o ./docs
```

4. Run the application:

```bash
go run cmd/main/main.go
```

## Docker Deployment

1. Build the Docker image:

```bash
docker build -t websync-api .
```

2. Run the container:

```bash
docker run -p 9545:9545 --env-file .env websync-api
```

## API Documentation

Swagger documentation is available at: `http://localhost:9545/docs/index.html`

### Health Check Endpoints

- `GET /health` - Basic health check
- `GET /health/live` - Liveness check
- `GET /health/ready` - Readiness check

### Authentication Endpoints

- `POST /api/authentication/login` - User login
- `POST /api/authentication/register` - User registration

### User Management

- `GET /api/users` - Get all users
- `GET /api/users/:id` - Get user by ID
- `POST /api/users` - Create new user
- `PATCH /api/users/:id` - Update user
- `DELETE /api/users/:id` - Delete user

### Queue Management

- `GET /api/queue/products` - Sync products
- `GET /api/queue/images` - Sync images
- `POST /api/queue/daily-sales` - Sync daily sales

## Error Handling

The API uses a standardized error response format:

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

## Rate Limiting

- Global rate limit: 1000 requests per minute
- IP-based rate limit: 100 requests per minute per IP
- Route-specific rate limits can be configured

## Development

### Project Structure

```
.
├── cmd/
│   └── main/
│       ├── main.go
│       ├── queue.go
│       └── cron.go
├── internal/
│   └── app/
│       ├── controllers/
│       ├── middleware/
│       ├── models/
│       ├── repository/
│       ├── routes/
│       └── service/
├── pkg/
│   ├── errors/
│   ├── helpers/
│   └── magento2/
├── docs/
├── config/
└── types/
```

### Running Tests

```bash
go test ./...
```

### Code Style

Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) for code style guidelines.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Captainkiez
