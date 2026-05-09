# API Documentation - Warteg System Backend

Base URL (Lokal): `http://localhost:8080`
Base URL (Vercel): `https://lugh-mobile-backend-v1.vercel.app/`

## Authentication Endpoints

### 1. Register User
Mendaftarkan akun pengguna baru.
- **URL**: `/api/auth/register`
- **Method**: `POST`
- **Body**: `{"username": "...", "email": "...", "password": "...", "full_name": "...", "phone": "..."}`
- **Note**: Setelah pendaftaran berhasil, sistem secara otomatis akan membuatkan dompet default (**Cash & QRIS**) serta kategori dasar Warteg untuk user tersebut.

### 2. Login
Melakukan login untuk mendapatkan token akses (JWT).
- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Body**: `{"username": "...", "password": "..."}`
- **Note**: Bagi user lama yang belum memiliki data dompet, sistem akan melakukan **Auto-Seed** data default pada saat login pertama kali setelah update ini.

---

## Finance Endpoints (Warteg Management)
*Semua endpoint di bawah ini membutuhkan Header: `Authorization: Bearer <TOKEN_JWT>`*

### 3. Wallets (Dompet/Penyimpanan)
Mengelola dompet Kas (Cash) atau Digital (QRIS).
- **List Wallets**: `GET /api/wallets`
- **Create Wallet**: `POST /api/wallets`
  - **Body**: `{"name": "Cash"}`

### 4. Categories (Kategori)
Mengelola kategori pemasukan dan pengeluaran.
- **List Categories**: `GET /api/categories`
- **Create Category**: `POST /api/categories`
  - **Body**: `{"name": "Belanja Sayur", "type": "expense", "icon": "leaf"}`
  - **Note**: `type` harus `income` atau `expense`.

### 5. Transactions (Transaksi Warteg)
Mencatat uang masuk (penjualan) dan uang keluar (belanja/operasional).
- **List Transactions**: `GET /api/transactions`
  - **Query (Optional)**: `?type=income` atau `?type=expense`
- **Record Transaction**: `POST /api/transactions`
  - **Body**:
    ```json
    {
      "wallet_id": 1,
      "category_id": 2,
      "amount": 50000,
      "type": "expense",
      "description": "Beli daging ayam 1kg",
      "date": "2026-05-09T22:30:00Z"
    }
    ```
- **Financial Summary**: `GET /api/transactions/summary`
  - **Response**: `{"total_income": 500000, "total_expense": 200000, "balance": 300000}`

---

### 6. User Profile
- **URL**: `/api/user/profile`
- **Method**: `GET`
