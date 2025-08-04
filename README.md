# Evermos Backend Mini Project

Mini project backend menggunakan **Golang + Fiber + GORM + MySQL**, dibuat sebagai bagian dari program Virtual Internship Rakamin x Evermos.

## 🚀 Tech Stack
- Go 
- Fiber (Web Framework)
- GORM (ORM)
- MySQL (DB)
- JWT (Auth)
- REST API
- EMSIFA API (Wilayah)

## 👉 API testing
[Download Postman Collection](mini-project-evermos/blob/main/Mini%20Project%20Evermos.postman_collection.json)


## 🛠 SQL Setup

Sebelum menjalankan aplikasi, lakukan setup awal database dan akun admin:

### 1. Buat Database
```sql
CREATE DATABASE toko_db;
```

### 2. Insert User Admin

Gunakan query berikut untuk menambahkan user dengan role admin (untuk testing endpoint kategori):

```sql
INSERT INTO users (nama, email, password, is_admin, created_at, updated_at)
VALUES (
  'Admin',
  'admin@example.test',
  '$2a$10$jRnt4t.1EHI6Kgizu08Mi.VHPLwdAP.bCRaacmYaK3ZLPC1HnfMM2', -- hash dari "rahasia123"
  true,
  NOW(), NOW()
);
```

> 💡 Password di atas adalah `rahasia123` (sudah dalam bentuk hash bcrypt)

### 3. Login sebagai Admin

Gunakan endpoint berikut untuk login sebagai admin:

```http
POST /auth/login

{
  "email": "admin@example.test",
  "kata_sandi": "rahasia123"
}
```

Setelah login, gunakan token untuk mengakses endpoint admin seperti:
- `POST /categories`
- `PUT /categories/:id`
- `DELETE /categories/:id`

