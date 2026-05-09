# Warteg System Backend (Go)

A secure backend for user authentication built with Go, Gin, and GORM (MySQL).

## Security Features
- **Argon2id/Bcrypt Hashing**: Securely stores passwords.
- **JWT Authentication**: Stateless session management.
- **Rate Limiting**: Protects against brute-force attacks.
- **Security Headers**: X-Content-Type-Options, X-Frame-Options, CSP, HSTS.
- **CORS**: Configurable cross-origin resource sharing.
- **Input Validation**: Strict validation using Gin binding.
- **Generic Error Messages**: Prevents information leakage (username enumeration).

## Prerequisites
- Go 1.20+
- MySQL Server

## Setup Instructions

1. **Clone the repository** (if not already in the directory).
2. **Configure Database**:
   - Create a database in MySQL (e.g., `warteg_db`).
   - Update `.env` with your MySQL credentials.
3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```
4. **Run the Server**:
   ```bash
   go run main.go
   ```

## API Endpoints

### Authentication
- `POST /api/auth/register`: Register a new user.
  ```json
  {
    "username": "johndoe",
    "email": "john@example.com",
    "password": "strongpassword123"
  }
  ```
- `POST /api/auth/login`: Login and receive a JWT.
  ```json
  {
    "username": "johndoe",
    "password": "strongpassword123"
  }
  ```

### User (Protected)
- `GET /api/user/profile`: Get the logged-in user's profile.
  - Requires header: `Authorization: Bearer <your_jwt_token>`

## Development Note
The database table `users` will be automatically created on the first run via GORM's AutoMigrate feature.
