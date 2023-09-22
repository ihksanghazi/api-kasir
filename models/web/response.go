package web

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Pagination struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	CurrentPage int         `json:"current_page"`
	TotalPage   int64       `json:"total_page"`
	Data        interface{} `json:"data"`
}

type Report struct {
	TotalTransaction float64 `json:"total_transaction"`
	ProductSold      float64 `json:"product_sold"`
	TotalSales       float64 `json:"total_sales"`
	TotalProfit      float64 `json:"total_profit"`
}
