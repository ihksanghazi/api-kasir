# api-kasir

## Description

"api-kasir" adalah sebuah sistem manajemen toko atau kasir yang berfokus pada pengelolaan produk dan transaksi. Sistem ini menyediakan berbagai fitur yang memungkinkan pengguna untuk melakukan berbagai operasi seperti menambahkan produk baru, mencari produk berdasarkan nama, melihat informasi produk, memperbarui data produk, menghapus produk, serta melakukan transaksi pembelian produk.

Selain itu, proyek ini juga dilengkapi dengan fitur pelaporan yang memungkinkan pengguna untuk melihat laporan transaksi, termasuk jumlah total transaksi, jumlah produk yang terjual, total penjualan, dan total keuntungan dalam rentang waktu tertentu.

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

### Get Product

Digunakan untuk mendapatkan informasi lengkap tentang produk berdasarkan **ID produk**.

#### Endpoint

```http
GET http://localhost:5000/api/product/<ProductId>
```

#### Parameter

- **ProductId** (string, required): ID unik produk yang ingin dilihat informasinya.

### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": "<ProductId>",
    "product_name": "Laptop",
    "purchase_price": 3000000,
    "selling_price": 4000000,
    "stock": 45,
    "created_at": "2023-09-22T01:19:12.341548+07:00",
    "updated_at": "2023-09-22T01:35:39.197295+07:00",
    "Transactions": [
      {
        "id": "<TransactionId>",
        "total": 37000000,
        "created_at": "2023-09-22T01:20:49.828702+07:00",
        "updated_at": "2023-09-22T01:20:49.866002+07:00"
      },
      .....
    ]
  }
}
```

> [!NOTE]
> Pastikan untuk menyediakan **ID produk** yang valid dalam URL untuk mengidentifikasi produk yang ingin dilihat informasinya.
> Data produk meliputi **nama produk**, **harga beli**, **harga jual**, **stok**, **tanggal pembuatan**, **tanggal pembaruan**, dan **riwayat transaksi** terkait produk.

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

### Create Transaction

Digunakan untuk membuat transaksi baru dalam sistem dengan daftar produk yang dibeli.

#### Endpoint

```http
POST http://localhost:5000/api/transaction
```

#### Request Body

<ul>
	<li> <bold>products</bold> (array, required): Daftar produk yang dibeli, setiap produk memiliki ID dan jumlah yang dibeli.</li>
	<ul>
		<li> <bold>product_id</bold> (string, required): ID unik produk yang dibeli.</li>
		<li> <bold>amount</bold> (integer, required): Jumlah produk yang dibeli.</li>
	</ul>
</ul>

Contoh Permintaan Body:

```json
{
	"products": [
		{
			"product_id": "8046372d-af76-461e-a33b-7badc7d1b48e",
			"amount": 1
		},
		{
			"product_id": "ad588905-9060-4ab6-a414-758c96e72b64",
			"amount": 1
		},
		{
			"product_id": "de0c7c2c-5108-46ed-a3bf-a22f66cff829",
			"amount": 1
		}
	]
}
```

#### Response

- **HTTP Status**: 201 Created
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 201,
	"status": "Created",
	"data": {
		"id": "<TransactionId>",
		"total": 11500000,
		"created_at": "2023-09-22T08:13:45.931252+07:00",
		"updated_at": "2023-09-22T08:13:45.992187+07:00",
		"transaction_details": [
			{
				"amount": 1,
				"products": {
					"product_name": "Monitor",
					"selling_price": 2500000
				}
			},
			.....
		]
	}
}
```

> [!NOTE]
> Permintaan POST ini digunakan untuk menambahkan pembelian produk ke dalam sistem, dengan memberikan daftar produk yang dibeli dan jumlahnya.
> Pastikan memeriksa status respons dan data transaksi untuk mendapatkan informasi tentang transaksi yang berhasil dilakukan, termasuk total harga dan rincian produk yang dibeli.

##

### Find Transaction

Digunakan untuk mendapatkan daftar transaksi dalam rentang tanggal tertentu.

#### Endpoint

```http
GET http://localhost:5000/api/transaction?startDate=2023-09-21&endDate=2023-09-30&page=1&limit=10
```

#### Parameter

- **startDate** (string, required): Tanggal awal rentang pencarian (format: yyyy-MM-dd).
- **endDate** (string, required): Tanggal akhir rentang pencarian (format: yyyy-MM-dd).
- **page** (integer, optional): Nomor halaman yang ingin diakses. Default: 1.
- **limit** (integer, optional): Jumlah transaksi yang ingin ditampilkan per halaman. Default: 5.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
  "code": 200,
  "status": "OK",
  "current_page": 1,
  "total_page": 1,
  "data": [
    {
      "id": "<TransactionId>",
      "total": 37000000,
      "created_at": "2023-09-22T01:20:49.828702+07:00",
      "updated_at": "2023-09-22T01:20:49.866002+07:00"
    },
    .....
  ]
}
```

> [!NOTE]
> Anda dapat menggunakan parameter **page** dan **limit** untuk mengontrol jumlah transaksi yang ditampilkan per halaman.
> Pastikan memeriksa status respons dan data transaksi untuk mendapatkan informasi tentang transaksi dalam rentang tanggal yang sesuai dengan pencarian Anda.

##

### Get Transaction

Digunakan untuk mendapatkan detail lengkap tentang sebuah transaksi berdasarkan ID transaksi.

#### Endpoint

```http
GET http://localhost:5000/api/transaction/<TransactionId>
```

Parameter

- **TransactionId** (string, required): ID unik transaksi yang ingin dilihat detailnya.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": "<TransactionId>",
    "total": 37000000,
    "created_at": "2023-09-22T01:20:49.828702+07:00",
    "updated_at": "2023-09-22T01:20:49.866002+07:00",
    "transaction_details": [
      {
        "amount": 2,
        "products": {
          "product_name": "Monitor",
          "selling_price": 2500000
        }
      },
      .....
    ]
  }
}
```

> [!NOTE]
> Pastikan untuk menyediakan ID transaksi yang valid dalam URL untuk mengidentifikasi transaksi yang ingin dilihat detailnya.
> Data transaksi mencakup total harga, tanggal pembuatan, tanggal pembaruan, dan rincian produk yang dibeli beserta jumlah dan harga jualnya.

##

### Get Report Transaction

Digunakan untuk mendapatkan laporan transaksi dalam rentang tanggal tertentu.

#### Endpoint

```http
GET http://localhost:5000/api/transaction/report?startDate=2023-09-01&endDate=2023-09-30
```

#### Parameter

- **startDate** (string, required): Tanggal awal rentang laporan (format: yyyy-MM-dd).
- **endDate** (string, required): Tanggal akhir rentang laporan (format: yyyy-MM-dd).

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"total_transaction": 3,
		"product_sold": 16,
		"total_sales": 66500000,
		"total_profit": 18000000
	}
}
```

> [!NOTE]
> Informasi dalam laporan meliputi **jumlah total transaksi**, **jumlah produk yang terjual**, **total penjualan**, dan **total keuntungan** dalam periode waktu yang dipilih.
> Pastikan untuk memeriksa status respons dan data laporan untuk mendapatkan informasi yang sesuai dengan rentang tanggal yang Anda minta.

##
