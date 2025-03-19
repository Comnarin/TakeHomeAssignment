package services

import (
	"calculateDiscount/Models"
	"calculateDiscount/Requests"
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

func CalculatePercentageByCategory(price float64, productReq requests.ProductReq, discount models.Discount) float64 {
	if productReq.Product.Category == discount.ProductCategory {
		totalPrice := productReq.Product.Price * float64(productReq.Quantity)
		return CalculatePercentageDiscount(discount.Amount, totalPrice)
	}
	return price
}

func CalculateDiscountByPoint(price float64, point int) float64 {
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

	var coupon *models.Discount
	var categoryPercentage *models.Discount
	var point *models.Discount
	var seasonal *models.Discount

	for _, discount := range cart.Discounts {
		switch discount.DiscountCategory {
		case "Coupon":
			if coupon == nil || discount.Amount > coupon.Amount {
				coupon = &discount
			}
		case "OnTop":
			if discount.DiscountName == "PercentageByCategory" {
				categoryPercentage = &discount
			} else if discount.DiscountName == "Point" {
				point = &discount
			}

		case "Seasonal":
			seasonal = &discount
		}
	}

	if coupon != nil {
		if coupon.DiscountName == "FixedAmount" {
			total = CalculateFixedAmountDiscount(coupon.Amount, total)
		} else if coupon.DiscountName == "Percentage" {
			total = CalculatePercentageDiscount(coupon.Amount, total)
		}
	}

	if point != nil && categoryPercentage != nil {

		calPoint := CalculateDiscountByPoint(total, point.Point)
		categoryPercentageAmount := 0.0
		for _, product := range cart.Products {
			calCategory := CalculatePercentageByCategory(total, product, *categoryPercentage)
			categoryPercentageAmount += calCategory
		}

		if calPoint < categoryPercentageAmount {
			total = calPoint
		} else {
			total = categoryPercentageAmount
		}

	} else if point != nil {

		total = CalculateDiscountByPoint(total, point.Point)

	} else if categoryPercentage != nil {

		categoryPercentageAmount := 0.0
		for _, product := range cart.Products {
			calCategory := CalculatePercentageByCategory(total, product, *categoryPercentage)
			categoryPercentageAmount += calCategory
		}
	}

	if seasonal != nil {
		total = CalculateSpecialDiscount(total, seasonal.Condition, seasonal.Amount)
	}


	return total
}
