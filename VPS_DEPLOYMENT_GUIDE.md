# VPS Setup & Deployment Guide

Panduan ini menjelaskan cara menyiapkan VPS untuk menjalankan backend Warteg menggunakan Docker.

## 1. Persiapan di VPS (Sekali Saja)

### Install Docker & Docker Compose
Jalankan perintah ini di terminal VPS Anda:
```bash
# Update sistem
sudo apt update && sudo apt upgrade -y

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install Docker Compose
sudo apt install -y docker-compose
```

### Login ke GitHub Container Registry (GHCR)
Agar VPS bisa mengambil image dari GitHub, Anda butuh **Personal Access Token (PAT)**:
1. Buka [GitHub Settings > Developer Settings > Personal Access Tokens (Tokens classic)](https://github.com/settings/tokens).
2. Generate token baru dengan izin `read:packages`.
3. Di terminal VPS, jalankan:
```bash
echo "TOKEN_ANDA" | docker login ghcr.io -u USERNAME_GITHUB --password-stdin
```

---

## 2. Setup Folder Aplikasi di VPS

Buat folder untuk project Anda dan siapkan file konfigurasi:
```bash
mkdir -p ~/warteg-app
cd ~/warteg-app

# Buat file .env (Copy-paste isi .env lokal Anda ke sini)
nano .env

# Ambil file docker-compose.yml dari GitHub
curl -O https://raw.githubusercontent.com/rizkyfauzi-git/lugh-mobile-app-backend/main/docker-compose.yml
```

---

## 3. Menjalankan Aplikasi

Setiap kali Anda ingin menjalankan atau mengupdate aplikasi ke versi terbaru:

```bash
cd ~/warteg-app

# Tarik image terbaru dari GitHub
docker-compose pull

# Jalankan container (-d artinya berjalan di background)
docker-compose up -d
```

### Cek apakah sudah jalan:
```bash
docker ps
# Atau cek log jika ada masalah:
docker logs -f warteg-backend
```

---

## Tips Tambahan: Update Otomatis
Jika Anda ingin VPS otomatis update saat Anda push ke GitHub, Anda bisa menambahkan perintah ini di bagian akhir GitHub Actions (`ci.yml`) menggunakan SSH, atau yang paling simpel adalah dengan menjalankan script `docker-compose pull && docker-compose up -d` secara manual di VPS setiap kali Anda selesai push.
