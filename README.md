
---

## **Persyaratan Sistem**
Pastikan Anda memiliki:

1. **Docker** (versi terbaru)
2. **Go** (versi 1.22.6 atau lebih baru)
3. Akses ke file konfigurasi yang relevan dan dataset yang sesuai

---

## **Langkah-Langkah Menjalankan API**

### 1. **Menjalankan API dengan Docker Compose**
Untuk memulai layanan, cukup jalankan perintah berikut:

```bash
docker compose up -d
```

Perintah ini akan menjalankan semua layanan yang dibutuhkan dalam mode latar belakang (*detached mode*). Pastikan semua kontainer berjalan dengan lancar.

### 2. **Masuk ke dalam Container**
Setelah semua kontainer berjalan, masuk ke kontainer aplikasi utama untuk menjalankan perintah lebih lanjut:

```bash
docker exec -it krs-management-go-app-1 bash
```

Anda sekarang berada di dalam lingkungan kontainer aplikasi.

---

## **Melakukan Unit Test dan Memasukkan Data**
API ini dilengkapi dengan unit test untuk memastikan integritas dan validitas data saat dimasukkan ke dalam sistem.

### 1. **TestDataInsertion1**
Proses ini bertanggung jawab untuk memasukkan data awal seperti mahasiswa, mata kuliah, ruang, sesi kuliah, dan tahun ajaran:

```bash
go test -v -run TestDataInsertion1 -timeout 30m
```

### 2. **TestDataInsertion2**
Proses ini menangani data tambahan seperti tagihan mahasiswa, IP semester, jadwal input KRS, dan data terkait lainnya:

```bash
go test -v -run TestDataInsertion2 -timeout 30m
```

### 3. **TestDataInsertion3**
Proses ini menangani data kompleks seperti validasi KRS, sesi kuliah bentrok, jadwal tawar, dan log KRS:

```bash
go test -v -run TestDataInsertion3 -timeout 30m
```

> **Catatan:** Jika dataset besar atau pengujian membutuhkan waktu lebih lama, Anda dapat menyesuaikan timeout menggunakan flag `-timeout`.

---

Selamat menggunakan API KRS Management! dengan membuka file documentation yang terdapat payload postman yang bisa anda import 🚀




## Overview
This is the API documentation for the Student Information System. All endpoints require authentication via a JWT token, except for the login endpoint. The token must be included in the `Authorization` header as `Bearer <token>`.

---

## Authentication
### Login
**POST** `/api/auth/login`

Request:
```json
{
  "nim_dinus": "176bd4688305a3ae546b35b52aed75c8",
  "pass_mhs": "your_password"
}
```

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "token": "<your_jwt_token>"
  }
}
```

---

## Endpoints

### 1. Get Offered KRS (All schedules for the academic year)
**GET** `/api/students/krs-offers?kode-ta=20232`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": 283583,
      "tahun_ajaran": 20232,
      "kelompok": "A24.8801",
      "matakuliah": "BAHASA INDONESIA",
      "sks": 2,
      "hari": "-",
      "jam_mulai": "",
      "jam_selesai": "",
      "ruang": "",
      "jns_jam": "3"
    },
    ...
  ]
}
```

---

### 2. Get KRS Schedule
**GET** `/api/students/krs-schedule`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "ta": 20241,
    "prodi": "B12",
    "tgl_mulai": "2024-08-14 15:00:00",
    "tgl_selesai": "2024-08-14 19:00:00"
  }
}
```

---

### 3. Update Class Type
**PUT** `/api/students/class`

**Headers:**
- Authorization: `Bearer <token>`

Request:
```json
{
  "kelas": 1
}
```

Response (Success):
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "message": "success update class"
  }
}
```

Response (Error):
```json
{
  "code": 400,
  "status": "BAD REQUEST",
  "data": {
    "error": "you have added a total of 1 Krs, you can't change the class type"
  }
}
```

---

### 4. Get Student Status
**GET** `/api/students/status`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "nim_dinus": "176bd4688305a3ae546b35b52aed75c8",
    "ta_masuk": "2022",
    "prodi": "B12",
    "akdm_stat": "aktif",
    "date_reg": "",
    "spp_status": "unpaid",
    "kelas": "pagi"
  }
}
```

---

### 5. Get Temporary KRS Schedule
**GET** `/api/students/krs`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id_schedule": 276907,
      "krs_record_id": 3976939,
      "tahun_ajaran": "20232",
      "kode_mata_kuliah": "B12.6406",
      "kelompok": "B12.4.2",
      "nama_mata_kuliah": "STATISTIK I",
      "jumlah_sks": 3,
      "hari": "SELASA",
      "jam_mulai": "07:00:00",
      "jam_selesai": "09:30:00",
      "ruang": "D.4.1",
      "jns_jam": "3"
    }
  ]
}
```

---

### 6. Permit KRS Entry Out of Schedule
**GET** `/api/students/permit`

**Headers:**
- Authorization: `Bearer <token>`

Response (Not Found):
```json
{
  "code": 404,
  "status": "NOT FOUND",
  "data": {
    "error": "error cant get permit status where nim"
  }
}
```

Response (Allowed):
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "message": "allowed insert krs"
  }
}
```

---

### 7. Get KRS Validation Status
**GET** `/api/students/krs-status?kode-ta=20232`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "validate": "Not Validated",
    "tahun_ajaran": "20231",
    "dipaketkan": "tidak di paketkan",
    "tahun_masuk": "2022",
    "sks": 23,
    "ips": "3.83"
  }
}
```

---

### 8. Get Student Scores
**GET** `/api/students/scores`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "kode_matkul": "B12.6401",
      "mata_kuliah": "AKUNTANSI KEUANGAN MENENGAH II",
      "sks": 3,
      "category": "T",
      "jenis_matkul": "wajib",
      "nilai": "A"
    },
    ...
  ]
}
```

---

### 9. Get Offered Prodi Schedules
**GET** `/api/students/schedule-prodi?kode-ta=20232`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": 276958,
      "TahunAjaran": "20232",
      "KodeMataKuliah": "B12.6608",
      "Kelompok": "B12.6.1",
      "NamaMataKuliah": "",
      "JumlahSKS": 0,
      "Hari": "-",
      "JamMulai": "",
      "JamSelesai": "",
      "Ruang": "",
      "jns_jam": "1",
      "StatusPemilihan": "Bisa",
      "StatusKrs": ""
    },
    ...
  ]
}
```

---

### 10. Get All Schedules (With Conflicts)
**GET** `/api/students/schedule-conflict?kode-ta=20232`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": 276154,
      "tahun_ajaran": "20232",
      "kelompok": "A14.7205",
      "nama_mata_kuliah": "LITERASI INFORMASI",
      "jumlah_sks": 2,
      "hari": "SENIN",
      "jam_mulai": "07:00:00",
      "jam_selesai": "08:40:00",
      "ruang": "Kulino",
      "jns_jam": "1",
      "status_bentrok": "",
      "keterangan_slot": "31/42"
    },
    ...
  ]
}
```

---

### 11. Add Schedule by ID
**GET** `/api/students/schedule/283093?kode-ta=20232`

**Headers:**
- Authorization: `Bearer <token>`

Response (Success):
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "message": "SUCCESS ADD NEW SCHEDULE"
  }
}
```

Response (Error):
```json
{
  "code": 400,
  "status": "BAD REQUEST",
  "data": {
    "error": "krs bentrok bang"
  }
}
```

---

### 12. Get Logs
**GET** `/api/students/log?kode-ta=20232`

**Headers:**
- Authorization: `Bearer <token>`

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id_rec": 3976939,
      "nim_dinus": "176bd4688305a3ae546b35b52aed75c8",
      "kode_mata_kuliah": "B12.6406",
      "aksi": 0,
      "id_jadwal": 276907,
      "ip_address": 0,
      "last_update": "2024-12-27T00:22:48+07:00"
    }
  ]
}
```

---

### 13. Delete KRS
**DELETE** `/api/students/krs/4004069?kode-ta=20232`

**Headers:**
- Authorization: `Bearer <token>`

Response (Success):
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "message": "success delete schedule where id krs = 3976939"
  }
}
```

Response (Error):
```json
{
  "code": 400,
  "status": "BAD REQUEST",
  "data": {
    "error": "wrong id_krs 3976939 not found"
  }
}
```

---

### 14. Update Validation Status
**PUT** `/api/students/validate`

**Headers:**
- Authorization: `Bearer <token>`

Request:
```json
{
  "job_host": "haha",
  "job_agent": "ghihih",
  "ta": 20232
}
```

Response:
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "message": "success update validate status"
  }
}
```



## License
This project is licensed under the MIT License.

