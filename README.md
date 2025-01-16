# Challange

Project RESTful API sederhana dengan studi kasus E-Commerce.

## Fitur

- **Autentikasi JWT**: Menggunakan token JWT untuk autentikasi pengguna.
- **Manajemen Pengguna**: Fungsi untuk login, registrasi, dan mengelola data pengguna.
- **Manajemen Produk**: Menambahkan, memperbarui, dan menghapus produk.
- **Pemesanan**: Melakukan pemesanan barang dan mengurangi stock.

## Bonus

- [x] API Documentation
- [x] Dockerized
- [ ] Automated Testing
- [x] Diagram Arsitektur

## Prasyarat

Sebelum menjalankan proyek ini, pastikan sudah menginstal:

- [Go](https://golang.org/dl/)
- [Docker](https://docker.com/)
<!-- - [PostgreSQL](https://www.postgresql.org/download/) atau database lainnya (tergantung pada konfigurasi `.env`) -->

## Instalasi

1. **Clone repositori ini**:

```bash
git clone https://github.com/username/repository-name.git
cd repository-name
```

2. **Instal dependensi**:
```bash
go mod tidy
```

3. **Buat file** .env:
```env
JWT_SECRET_KEY=your_jwt_secret_key
SERVER_PORT=8080
```

4. **Menjalankan aplikasi**:
```bash
go run main.go
```

## Docker
```bash
docker-compose up --build
```

## Dokumentasi API

- Dapat diakses di [Postman Documentation](https://documenter.getpostman.com/view/1475503/2sAYQZJCwZ) atau https://documenter.getpostman.com/view/1475503/2sAYQZJCwZ

## ER Diagram
![Diagram](diagram.drawio.png)