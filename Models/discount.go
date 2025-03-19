package models

type Discount struct {
	Id               int
	DiscountName     string
	DiscountCategory string
	ProductCategory  string
	Amount           float64
	Point            int
	Condition        float64
}

// กำหนดให้ discount name มีค่าเป็น FixedAmount ,Percentage ,Point ,PercentageByCategory, และ Seasonal
// กำหนดให้ discount category มีค่าเป็น Coupon, OnTop, และ Seasonal