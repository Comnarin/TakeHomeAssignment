package services

import (
	"calculateDiscount/models"
	"calculateDiscount/requests"
	"fmt"
	"math"
	"strings"
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

func UpdateCart(cart requests.Cart, discount models.Discount, total float64) {
	switch strings.ToLower(discount.DiscountName) {
	case "fixedamount":
		for _, product := range cart.Products {
			weight := product.Product.Price / total
			product.Product.Price -= *discount.Amount * weight
		}
	case "percentage":
		for _, product := range cart.Products {
			product.Product.Price -= product.Product.Price * (*discount.Amount / 100)
		}

	}
}


func ApplyDiscount(cart requests.Cart) (float64, error) {

	if len(cart.Products) == 0 {
		return 0, fmt.Errorf("cart must contain at least one item")
	}

	var total float64

	for _, item := range cart.Products {
		if item.Product.Price < 0 {
			return 0, fmt.Errorf("item %s has an invalid price", item.Product.Name)
		}
		if item.Quantity < 0 {
			return 0, fmt.Errorf("item %s has an invalid quantity", item.Product.Name)
		}
		total += item.Product.Price * float64(item.Quantity)
	}

	var fixAmount *models.Discount
	var percentage *models.Discount
	var categoryPercentage *models.Discount
	var point *models.Discount
	var seasonal *models.Discount

	for _, discount := range cart.Discounts {

		discount.DiscountName = strings.ToLower(discount.DiscountName)

		if discount.Amount != nil && *discount.Amount < 0 {
			return 0, fmt.Errorf("discount %s has an invalid amount", discount.DiscountName)
		}
		if discount.Point != nil && *discount.Point < 0 {
			return 0, fmt.Errorf("discount %s has an invalid point", discount.DiscountName)
		}
		if discount.Condition != nil && *discount.Condition < 0 {
			return 0, fmt.Errorf("discount %s has an invalid condition", discount.DiscountName)
		}
		if discount.DiscountName != "fixedamount" && discount.DiscountName != "percentage" && discount.DiscountName != "percentagebycategory" && discount.DiscountName != "point" && discount.DiscountName != "seasonal" {
			return 0, fmt.Errorf("discount %s has an invalid discount name", discount.DiscountName)
		}

		switch discount.DiscountName {
		case "fixedamount":
			fixAmount = &discount
		case "percentage":
			percentage = &discount
		case "point":
			point = &discount
		case "percentagebycategory":
			categoryPercentage = &discount
		case "seasonal":
			seasonal = &discount
		}
	}

	switch {
	case fixAmount != nil && percentage != nil:
		calFixAmount := CalculateFixedAmountDiscount(*fixAmount.Amount, total)
		calPercentage := CalculatePercentageDiscount(*percentage.Amount, total)
		if calFixAmount < calPercentage {
			UpdateCart(cart, *fixAmount, total)
			total = calFixAmount
		} else {
			UpdateCart(cart, *percentage, total)
			total = calPercentage
		}
	case fixAmount != nil:
		UpdateCart(cart, *fixAmount, total)
		total = CalculateFixedAmountDiscount(*fixAmount.Amount, total)
	case percentage != nil:
		UpdateCart(cart, *percentage, total)
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

	return total, nil
}
