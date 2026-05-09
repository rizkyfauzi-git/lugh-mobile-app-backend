# API Documentation - Warteg System Backend

Base URL (Lokal): `http://localhost:8080`
Base URL (Vercel): `https://your-app-name.vercel.app`

## Authentication Endpoints

### 1. Register User
Mendaftarkan akun pengguna baru.

- **URL**: `/api/auth/register`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "rizkyfauzi",
    "email": "rizky@example.com",
    "password": "password123"
  }
  ```
- **Success Response**:
  - **Code**: `201 Created`
  - **Content**: `{"message": "User registered successfully"}`

---

### 2. Login
Melakukan login untuk mendapatkan token akses (JWT).

- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "rizkyfauzi",
    "password": "password123"
  }
  ```
- **Success Response**:
  - **Code**: `200 OK`
  - **Content**:
    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "user": {
        "id": 1,
        "username": "rizkyfauzi",
        "email": "rizky@example.com",
        "created_at": "2026-05-09T..."
      }
    }
    ```
- **Error Response**:
  - **Code**: `401 Unauthorized`
  - **Content**: `{"error": "Invalid username or password"}`

---

## Protected Endpoints
*Membutuhkan Header: `Authorization: Bearer <TOKEN_JWT>`*

### 3. Get Profile
Mengambil data profil pengguna yang sedang login.

- **URL**: `/api/user/profile`
- **Method**: `GET`
- **Headers**:
  - `Authorization`: `Bearer your_jwt_token_here`
- **Success Response**:
  - **Code**: `200 OK`
  - **Content**:
    ```json
    {
      "id": 1,
      "username": "rizkyfauzi",
      "email": "rizky@example.com",
      "created_at": "2026-05-09T...",
      "updated_at": "2026-05-09T..."
    }
    ```
- **Error Response**:
  - **Code**: `401 Unauthorized`
  - **Content**: `{"error": "Invalid or expired token"}`

---

## Security Notes
1. **Password Hashing**: Semua password disimpan dalam format Bcrypt (hash).
2. **JWT Expiration**: Token berlaku selama 24 jam.
3. **Rate Limiting**: Maksimal 1 request per detik dengan burst 5 untuk endpoint auth guna mencegah brute-force.
4. **CORS**: Sudah diaktifkan untuk semua origin (`*`).
