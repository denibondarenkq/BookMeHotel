package entity

type FoodPlan struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	PriceFactor float64 `json:"price_factor"`
}
