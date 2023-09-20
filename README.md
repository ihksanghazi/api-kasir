# api-kasir

## Description

## Feature

### Create Product

Digunakan untuk membuat produk baru ke dalam sistem.

#### Endpoint

```http
POST http://localhost:5000/api/product
```

#### Request Body

- **product_name** (string, required): Nama produk.
- **purchase_price** (number, required): Harga beli produk.
- **selling_price** (number, required): Harga jual produk.
- **stock** (integer, required): Jumlah stok produk.

Contoh Request Body:

```json
{
	"product_name": "Monitor",
	"purchase_price": 2000000,
	"selling_price": 2500000,
	"stock": 60
}
```

#### Response

- **HTTP Status**: 201 Created
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"product_name": "Monitor",
		"purchase_price": 2000000,
		"selling_price": 2500000,
		"stock": 60
	}
}
```

> [!NOTE]
> Bidang seperti **product_name**, **purchase_price**, **selling_price**, dan **stock** harus diisi sesuai dengan informasi produk yang ingin didaftarkan.
> Pastikan memeriksa status respons untuk memastikan produk telah berhasil ditambahkan ke sistem.

##

### Find Product

Digunakan untuk mencari produk berdasarkan nama produk.

#### Endpoint

```http
GET http://localhost:5000/api/product?page=1&limit=3&search=
```

#### Parameter

- **page** (integer, optional): Nomor halaman yang ingin diakses. Default: 1.
- **limit** (integer, optional): Jumlah produk yang ingin ditampilkan per halaman. Default: 5.
- **search** (string, optional): Kata kunci pencarian untuk mencari produk berdasarkan nama produk.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"current_page": 1,
	"total_page": 2,
	"data": [
		{
			"id": "<ProductId>",
			"product_name": "Komputer",
			"purchase_price": 3000000,
			"selling_price": 4000000,
			"stock": 20
		},
		.....
```

> [!NOTE]
> Anda dapat menggunakan parameter **page** dan **limit** untuk mengontrol jumlah produk yang ditampilkan per halaman.
> Parameter **search** digunakan untuk mencari produk berdasarkan nama produk.
> Pastikan memeriksa status respons dan data hasil pencarian dengan cermat untuk mengakses informasi produk yang sesuai.

##

### Update Product

Digunakan untuk memperbarui informasi produk yang sudah ada dalam sistem.

#### Endpoint

```http
PUT http://localhost:5000/api/product/<ProductId>
```

#### Parameter

- **ProductId** (string, required): ID unik produk yang ingin diperbarui.

#### Request Body

- **product_name** (string, optional): Nama produk.
- **purchase_price** (number, optional): Harga beli produk.
- **selling_price** (number, optional): Harga jual produk.
- **stock** (integer, optional): Jumlah stok produk.

Contoh Request Body:

```json
{
	"product_name": "Komputer",
	"purchase_price": 0,
	"selling_price": 5000000,
	"stock": 0
}
```

### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"product_name": "Komputer",
		"purchase_price": 3000000,
		"selling_price": 5000000,
		"stock": 20
	}
}
```

> [!NOTE]
> Pastikan untuk mengganti nilai-nilai dalam request body sesuai dengan informasi yang ingin diperbarui.
> Anda perlu menyediakan ID produk yang valid dalam URL untuk mengidentifikasi produk yang akan diperbarui.
> Periksa status respons untuk memastikan bahwa perubahan informasi produk telah berhasil diterapkan.

##

### Delete Product

Digunakan untuk menghapus produk dari sistem berdasarkan ID produk.

#### Endpoint

```http
DELETE http://localhost:5000/api/product/<ProductId>
```

#### Parameter

- **ProdcutId** (string, required): ID unik produk yang ingin dihapus.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"message": "Success Delete Product With Id '<ProductId>'",
	"status": "OK"
}
```

> [!NOTE]
> Pastikan untuk menyediakan ID produk yang valid dalam URL untuk mengidentifikasi produk yang akan dihapus.
> Periksa status respons untuk memastikan bahwa penghapusan produk telah berhasil dilakukan. Pesan "Success Delete Product With Id 'ID Produk'" akan memberikan konfirmasi tentang produk yang dihapus.

##
