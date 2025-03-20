package main

// import (
// 	requests "calculateDiscount/requests"
// 	services "calculateDiscount/Services"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	jsonFile, err := os.Open("cart.json")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer jsonFile.Close()

	
// 	var req requests.Cart

// 	byteValue, err := os.ReadFile("cart.json")
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}

// 	err = json.Unmarshal(byteValue, &req)
// 	if err != nil {
// 		fmt.Println("Error unmarshaling JSON:", err)
// 		return
// 	}
	
// 	total := services.ApplyDiscount(req)
// 	fmt.Println("Total Price = ",total)

// }