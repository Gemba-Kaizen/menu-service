package models

type FoodItem struct {
	Id          int64  `json:"id gorm:"primaryKey"`
	MerchantId  int64  `json:"merchant_id"`
	Name        string `json:"food_name"`
	Description string `json:"description"`
	Price       float64  `json:"price"`
}
