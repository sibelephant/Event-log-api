# Event-App API 🎉

A modern REST API for event management built with Go, Gin, and Prisma. This API allows users to create, manage, and attend events with a clean and efficient architecture.

## 🚀 Features

- **User Management**: Create, read, and manage user accounts
- **Event Management**: Create and list events with full details
- **Attendee Management**: Register users for events, manage attendance
- **Database Integration**: SQLite with Prisma ORM for type-safe database operations
- **Environment Configuration**: Secure environment variable management
- **Live Reload**: Air for development with hot reloading
- **Clean Architecture**: Modular structure with separated concerns

## 🛠️ Tech Stack

- **Language**: Go 1.24+
- **Web Framework**: [Gin](https://gin-gonic.com/)
- **Database**: SQLite (development) / PostgreSQL (production ready)
- **ORM**: [Prisma Client Go](https://github.com/steebchen/prisma-client-go)
- **Live Reload**: [Air](https://github.com/cosmtrek/air)
- **Environment**: [godotenv](https://github.com/joho/godotenv)

## 📁 Project Structure

```
Event-App/
├── cmd/api/                 # Application entry point
│   ├── main.go             # Main server setup
│   ├── routes.go           # Route definitions
│   ├── users.go            # User handlers
│   ├── events.go           # Event handlers
│   └── attendees.go        # Attendee handlers
├── internal/
│   ├── config/             # Configuration management
│   │   └── config.go       # Environment config loader
│   ├── database/           # Database connection
│   │   └── database.go     # Prisma client setup
│   └── models/             # Data models
│       ├── users.go        # User request/response models
│       ├── events.go       # Event request/response models
│       └── attendees.go    # Attendee request/response models
├── prisma/
│   └── schema.prisma       # Prisma schema definition
├── .env.example           # Environment variables template
├── .air.toml              # Air live reload configuration
└── README.md              # This file
```

## 🏗️ Database Schema

### Users

```prisma
model Users {
  id        Int         @id @default(autoincrement())
  email     String      @unique
  name      String
  password  String
  events    Events[]    @relation("EventOwner")
  attendees Attendees[]
}
```

### Events

```prisma
model Events {
  id          Int         @id @default(autoincrement())
  owner_id    Int
  name        String
  description String
  date        DateTime
  location    String
  owner       Users       @relation("EventOwner", fields: [owner_id], references: [id], onDelete: Cascade)
  attendees   Attendees[]
}
```

### Attendees

```prisma
model Attendees {
  id       Int    @id @default(autoincrement())
  user_id  Int
  event_id Int
  user     Users  @relation(fields: [user_id], references: [id], onDelete: Cascade)
  event    Events @relation(fields: [event_id], references: [id], onDelete: Cascade)
}
```

## 🚀 Quick Start

### Prerequisites

- Go 1.24 or higher
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/sibelephant/Event-log-api.git
cd Event-log-api
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Environment

```bash
# Copy the example environment file
cp .env.example .env

# Edit .env with your configuration
# DATABASE_URL, PORT, JWT_SECRET
```

### 4. Generate Prisma Client

```bash
go run github.com/steebchen/prisma-client-go generate
```

### 5. Set Up Database

```bash
# Create and migrate database
go run github.com/steebchen/prisma-client-go db push
```

### 6. Run the Application

```bash
# Development (with live reload)
air

# Or run directly
go run ./cmd/api
```

The API will be available at `http://localhost:8080`

## 📚 API Documentation

### Base URL

```
http://localhost:8080/api/v1
```

### 👥 User Endpoints

#### Create User

```http
POST /users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com"
}
```

#### Get All Users

```http
GET /users
```

#### Get User by ID

```http
GET /users/{id}
```

### 🎪 Event Endpoints

#### Create Event

```http
POST /events
Content-Type: application/json

{
  "name": "Tech Conference 2025",
  "description": "Annual technology conference",
  "date": "2025-12-25T15:04:05Z",
  "location": "San Francisco, CA",
  "owner_id": 1
}
```

#### Get All Events

```http
GET /events
```

### 👥 Attendee Endpoints

#### Register User for Event

```http
POST /attendees
Content-Type: application/json

{
  "user_id": 1,
  "event_id": 2
}
```

**Response:**

```json
{
  "message": "Successfully registered for event",
  "attendee": {
    "id": 1,
    "user_id": 1,
    "event_id": 2
  }
}
```

#### Get All Attendees

```http
GET /attendees
```

#### Get Event Attendees

```http
GET /events/{event_id}/attendees
```

**Response:**

```json
{
  "event": {
    "id": 2,
    "name": "Workshop",
    "description": "Programming workshop",
    "date": "2024-11-15T14:00:00Z",
    "location": "San Francisco"
  },
  "attendees_count": 2,
  "attendees": [
    {
      "id": 1,
      "user_id": 1,
      "event_id": 2,
      "user": {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com"
      }
    }
  ]
}
```

#### Get User's Attending Events

```http
GET /user/{user_id}/attending-events
```

#### Remove User from Event (by attendee ID)

```http
DELETE /attendees/{id}
```

#### Remove User from Event (by user and event ID)

```http
DELETE /user/{user_id}/events/{event_id}
```

## 🔧 Development

### Environment Variables

Create a `.env` file based on `.env.example`:

```env
DATABASE_URL="file:./dev.db"
PORT=8080
JWT_SECRET="your-super-secret-jwt-key-change-this-in-production"
```

### Live Reload with Air

The project includes Air configuration for live reloading during development:

```bash
# Install Air (if not already installed)
go install github.com/cosmtrek/air@latest

# Run with live reload
air
```

### Database Management

```bash
# Generate Prisma client after schema changes
go run github.com/steebchen/prisma-client-go generate

# Update database schema
go run github.com/steebchen/prisma-client-go db push

# Reset database (development only)
go run github.com/steebchen/prisma-client-go db reset
```

## 🧪 Testing

### Test the API with curl

```bash
# Create a user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }'

# Create an event
curl -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sample Event",
    "description": "This is a test event",
    "date": "2025-12-25T15:04:05Z",
    "location": "Test Location",
    "owner_id": 1
  }'

# Register user for event
curl -X POST http://localhost:8080/api/v1/attendees \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "event_id": 1
  }'

# Get all users
curl http://localhost:8080/api/v1/users

# Get all events
curl http://localhost:8080/api/v1/events

# Get event attendees
curl http://localhost:8080/api/v1/events/1/attendees

# Get user's attending events
curl http://localhost:8080/api/v1/user/1/attending-events
```

## 🚀 Deployment

### Production Environment

1. **Set up production database** (PostgreSQL recommended)
2. **Update environment variables**:
   ```env
   DATABASE_URL="postgresql://user:password@localhost:5432/eventapp?sslmode=require"
   PORT=8080
   JWT_SECRET="your-production-secret"
   ```
3. **Build the application**:
   ```bash
   go build -o eventapp ./cmd/api
   ```
4. **Run the binary**:
   ```bash
   ./eventapp
   ```

### Docker (Optional)

```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
CMD ["./main"]
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙋‍♂️ Support

If you have any questions or need help, please:

1. Check the [API Documentation](#-api-documentation)
2. Review existing [Issues](https://github.com/sibelephant/Event-log-api/issues)
3. Create a new issue if needed

## 🔮 Roadmap

- [ ] JWT Authentication middleware
- [x] Event attendance management ✅
- [ ] Event search and filtering
- [ ] File upload for event images
- [ ] Email notifications
- [ ] API rate limiting
- [ ] Comprehensive test suite
- [ ] API documentation with Swagger

---

**Built with ❤️ using Go and Prisma**
