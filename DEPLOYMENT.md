# Deployment Guide

This guide provides detailed instructions for deploying the WebSync API in different environments.

## Prerequisites

- Docker and Docker Compose
- MySQL 8.0+
- RabbitMQ 3.9+
- AWS S3 bucket (for image storage)
- Domain name (optional)

## Environment Setup

1. Create a `.env` file with the following variables:

```env
# Application
APP_PORT=9545
GIN_MODE=release

# Database
DB_HOST=your-db-host
DB_PORT=3306
DB_USERNAME=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=websync

# RabbitMQ
RABBITMQ_URL=amqp://guest:guest@your-rabbitmq-host:5672/

# JWT
JWT_SECRET=your-secure-secret-key
JWT_EXPIRATION=24h

# AWS S3
AWS_S3_REGION=ap-southeast-1
AWS_S3_ACCESS_KEY_ID=your-access-key
AWS_S3_SECRET_ACCESS_KEY=your-secret-key
AWS_S3_BUCKET=your-bucket

# Service URLs
STRAPI_URL=http://your-strapi-url
SERVICE_URL=http://your-service-url
```

## Docker Deployment

### 1. Build the Image

```bash
docker build -t websync-api:latest .
```

### 2. Run with Docker

```bash
docker run -d \
  --name websync-api \
  -p 9545:9545 \
  --env-file .env \
  websync-api:latest
```

### 3. Run with Docker Compose

Create a `docker-compose.yml` file:

```yaml
version: "3.8"

services:
  api:
    build: .
    ports:
      - "9545:9545"
    env_file:
      - .env
    depends_on:
      - db
      - rabbitmq
    restart: unless-stopped

  db:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: your-root-password
      MYSQL_DATABASE: websync
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - "3306:3306"

  rabbitmq:
    image: rabbitmq:3.9-management
    ports:
      - "5672:5672"
      - "15672:15672"
    volumes:
      - rabbitmq_data:/var/lib/rabbitmq

volumes:
  mysql_data:
  rabbitmq_data:
```

Run with:

```bash
docker-compose up -d
```

## Production Deployment

### 1. Security Considerations

- Use HTTPS with valid SSL certificates
- Set up a reverse proxy (Nginx/Apache)
- Configure firewall rules
- Use strong passwords and secrets
- Enable rate limiting
- Set up monitoring and logging

### 2. Nginx Configuration

Example Nginx configuration:

```nginx
server {
    listen 80;
    server_name your-domain.com;
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl;
    server_name your-domain.com;

    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    location / {
        proxy_pass http://localhost:9545;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

### 3. Monitoring Setup

1. Set up Prometheus metrics:

```bash
# Add to your Docker run command
-e ENABLE_METRICS=true
```

2. Configure Grafana dashboards for:
   - API response times
   - Error rates
   - Queue processing metrics
   - Resource usage

### 4. Backup Strategy

1. Database backups:

```bash
# Daily backup
mysqldump -u root -p websync > backup_$(date +%Y%m%d).sql
```

2. Configuration backups:

```bash
# Backup .env and other configs
tar -czf config_backup_$(date +%Y%m%d).tar.gz .env docker-compose.yml
```

## Scaling

### Horizontal Scaling

1. Use a load balancer (e.g., Nginx, HAProxy)
2. Configure multiple API instances
3. Use external Redis for session management
4. Set up database replication

### Vertical Scaling

1. Increase container resources:

```yaml
services:
  api:
    deploy:
      resources:
        limits:
          cpus: "2"
          memory: 2G
```

2. Optimize database:
   - Add indexes
   - Configure connection pooling
   - Enable query caching

## Troubleshooting

### Common Issues

1. Database Connection Issues:

   - Check credentials
   - Verify network connectivity
   - Check database logs

2. RabbitMQ Connection Issues:

   - Verify RabbitMQ is running
   - Check credentials
   - Check network connectivity

3. API Performance Issues:
   - Check resource usage
   - Review logs
   - Monitor queue sizes

### Logs

View logs with:

```bash
# Docker
docker logs websync-api

# Docker Compose
docker-compose logs -f api
```

## Maintenance

### Regular Tasks

1. Update dependencies:

```bash
go get -u ./...
go mod tidy
```

2. Database maintenance:

```sql
OPTIMIZE TABLE your_table;
ANALYZE TABLE your_table;
```

3. Log rotation:

```bash
# Configure logrotate
/var/log/websync-api/*.log {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
}
```

### Health Checks

1. Monitor endpoints:

   - `/health`
   - `/health/live`
   - `/health/ready`

2. Set up alerts for:
   - High error rates
   - Slow response times
   - Queue backlogs
   - Resource usage
