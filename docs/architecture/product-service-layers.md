# Alur Request Product Service

```mermaid
sequenceDiagram
    participant C as Client
    participant H as Handler
    participant R as Repository
    C->>H: GET /api/products/1
    H->>R: FindByID(1)
    R-->>H: Product atau ErrProductNotFound
    H-->>C: 200 atau 404 (JSON)
```

## Penjelasan

| Lapisan | File | Tugas |
|---|---|---|
| Handler | `internal/handler/product_handler.go` | Menerima request HTTP dan menulis response JSON |
| Repository | `internal/repository/product_repository.go` | Mengambil data produk (saat ini dari memori) |
| Model | `internal/model/product.go` | Definisi struktur data produk |
| Response | `internal/response/response.go` | Format response standar |