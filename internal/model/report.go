package model

type AllReport struct {
	TotalRevenue     int        `json:"total_revenue"`
	TotalTransaction int        `json:"total_transaction"`
	BestSeller       BestSeller `json:"best_seller"`
}

type Report struct {
	TotalRevenue     int `json:"total_revenue"`
	TotalTransaction int `json:"total_transaction"`
}

type BestSeller struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Sold      int    `json:"sold"`
}
