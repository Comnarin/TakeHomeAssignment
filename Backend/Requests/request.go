package requests

import (
	"calculateDiscount/Models"
)

type ProductReq struct {
	Product models.Product
	Quantity int
}

type Cart struct {
	Products []ProductReq
	Discounts []models.Discount
}
