package services

import (
	"calculateDiscount/models"
	"calculateDiscount/requests"
	"math"
)

func CalculateFixedAmountDiscount(amount float64, price float64) float64 {
	if price < amount {
		return 0
	}
	return price - amount
}

func CalculatePercentageDiscount(amount float64, price float64) float64 {
	return price - (price * amount / 100)
}
func CalculatePercentageByCategory(cart requests.Cart, discount models.Discount, price float64) float64 {
	for _, product := range cart.Products {
		if product.Product.Category == *discount.ProductCategory {
			productTotalPrice := product.Product.Price * float64(product.Quantity)
			price -= productTotalPrice * (*discount.Amount) / 100
		}
	}
	return price
}
func CalculatePointDiscount(price float64, point int) float64 {
	maxPoint := int(0.2 * price)

	if point > maxPoint {
		point = maxPoint
	}

	return price - float64(point)
}

func CalculateSpecialDiscount(price float64, condition float64, discount float64) float64 {
	discountMultiplier := math.Floor(price / condition)
	return price - (discountMultiplier * discount)
}

func CalculateTotalPrice(cart requests.Cart) float64 {
	var total float64
	for _, product := range cart.Products {
		total += product.Product.Price * float64(product.Quantity)
	}
	return total
}

func ApplyDiscount(cart requests.Cart) float64 {
	total := CalculateTotalPrice(cart)

	var fixAmount *models.Discount
	var percentage *models.Discount
	var categoryPercentage *models.Discount
	var point *models.Discount
	var seasonal *models.Discount

	for _, discount := range cart.Discounts {
		switch discount.DiscountName {
		case "fixedAmount":
			fixAmount = &discount
		case "Percentage":
			percentage = &discount
		case "Point":
			point = &discount
		case "PercentageByCategory":
			categoryPercentage = &discount
		case "Seasonal":
			seasonal = &discount
		}
	}

	switch {
	case fixAmount != nil && percentage != nil:
		calFixAmount := CalculateFixedAmountDiscount(*fixAmount.Amount, total)
		calPercentage := CalculatePercentageDiscount(*percentage.Amount, total)
		if calFixAmount < calPercentage {
			total = calFixAmount
		} else {
			total = calPercentage
		}
	case fixAmount != nil:
		total = CalculateFixedAmountDiscount(*fixAmount.Amount, total)
	case percentage != nil:
		total = CalculatePercentageDiscount(*percentage.Amount, total)
	}

	switch {
	case point != nil && categoryPercentage != nil:
		calPoint := CalculatePointDiscount(total, *point.Point)
		calPercengateByCategory := CalculatePercentageByCategory(cart, *categoryPercentage, total)
		if calPoint < calPercengateByCategory {
			total = calPoint
		} else {
			total = calPercengateByCategory
		}
	case point != nil:
		total = CalculatePointDiscount(total, *point.Point)
	case categoryPercentage != nil:
		total = CalculatePercentageByCategory(cart, *categoryPercentage, total)
	}

	if seasonal != nil {
		total = CalculateSpecialDiscount(total, *seasonal.Condition, *seasonal.Amount)
	}

	if total < 0 {
		total = 0
	}

	return total
}
