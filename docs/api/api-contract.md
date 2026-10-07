# API Contract

## Format Response Standar

Sukses:
```json
{ "success": true, "message": "ok", "data": {} }
```

Gagal:
```json
{ "success": false, "message": "pesan error", "data": null }
```

## Product Service

### GET /api/products
Mengambil semua produk.

Response 200:
```json
{
  "success": true,
  "message": "ok",
  "data": [
    { "id": 1, "name": "Keripik Singkong", "price": 15000, "stock": 40, "category_id": 2 }
  ]
}
```

### GET /api/products/{id}
Mengambil satu produk. Response 404 jika tidak ditemukan.