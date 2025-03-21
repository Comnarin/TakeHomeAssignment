
# Take Home Assignment : Narin Sirinapuk

This module calculates the final order price by applying multiple discount campaigns to cart.


## Run Project

Run my project with go run main.go to use go fiber

```bash
go run main.go
```

or

Read Json file 
```bash
cd readJson
go run readJson.go
```  

## Discount Json References

#### 1 . Fix Amount Coupon

```
{
    "Id": 1,
    "DiscountName": "FixedAmount",
    "DiscountCategory": "Coupon",
    "Amount": 50.0
}
```

| Name | DiscountCategory    | Description                |
| :-------- | :------- | :------------------------- |
| `FixedAmount` | `Coupon` | Discount 50 ฿ |



#### 2. Percentage Coupon

```
{
    "Id": 2,
    "DiscountName": "Percentage",
    "DiscountCategory": "Coupon",
    "Amount": 10.0
}
```

| Name | DiscountCategory    | Description                |
| :-------- | :------- | :------------------------- |
| `FixedAmount` | `Coupon` | Discount 10% |



#### 3. Percentage By  Category OnTop

```
{
    "Id": 3,
    "DiscountName": "PercentageByCategory",
    "DiscountCategory": "Ontop",
    "ProductCategory" : "Clothings"
    "Amount": 15.0
}
```

| Name | DiscountCategory    | Description                |
| :-------- | :------- | :------------------------- |
| `PercentageByCategory` | `OnTop` | 15% Discount on clothings |

#### 4. Point Discount OnTop

```
{
    "Id": 4,
    "DiscountName": "Point",
    "DiscountCategory": "Ontop",
    "Point" : 68
}
```

| Name | DiscountCategory    | Description                |
| :-------- | :------- | :------------------------- |
| `Point` | `Ontop` | Discount 68 point = 68 ฿ |

#### 5. Seasonal Discount 

```
{
    "Id": 5,
    "DiscountName": "Seasonal",
    "DiscountCategory": "Seasonal",
    "Amount" : 30.0
    "Condition":400
}
```

| Name | DiscountCategory    | Description                |
| :-------- | :------- | :------------------------- |
| `Seasonal` | `Seasonal` | Discount 30฿ for Every 400฿  |



## Product Json Example

```
{
    "Product": {
            "Id": 1,
            "Name": "T-Shirt",
            "Price": 100.0,
            "Category": "Clothing"
            },
    "Quantity": 1
}
```
| Name | ProductCategory    | Price               |Quantity|
| :-------- | :------- | :------------------------- | :-----|
| `T-Shirt` | `Clothings` | 100 |1|

## Input Example

```
{
   {
    "Products": [
        {
            "Product": {
                "Id": 1,
                "Name": "T-Shirt",
                "Price": 100.0,
                "Category": "Clothing"
            },
            "Quantity": 1
        },
        {
            "Product": {
                "Id": 2,
                "Name": "Hat",
                "Price": 250.0,
                "Category": "Accessories"
            },
            "Quantity": 1
        },
        {
            "Product": {
                "Id": 3,
                "Name": "Hoodie",
                "Price": 700.0,
                "Category": "Accessories"
            },
            "Quantity": 1
        },
        {
            "Product": {
                "Id": 4,
                "Name": "Watch",
                "Price": 850.0,
                "Category": "Electronics"
            },
            "Quantity": 1
        },
        {
            "Product": {
                "Id": 5,
                "Name": "Bag",
                "Price": 640.0,
                "Category": "Accessories"
            },
            "Quantity": 1
        }
    ],
    "Discounts": [
        {
            "Id": 1,
            "DiscountName": "FixedAmount",
            "DiscountCategory": "Coupon",
            "Amount": 50.0
        },
        {
            "Id": 2,
            "DiscountName": "Percentage",
            "DiscountCategory": "Coupon",
            "Amount": 10.0
        },
        {
            "Id": 3,
            "DiscountName": "PercentageByCategory",
            "DiscountCategory": "OnTop",
            "ProductCategory": "Clothing",
            "Amount": 15.0
        },
        {
            "Id": 4,
            "DiscountName": "Point",
            "DiscountCategory": "OnTop",
            "Point": 68
        },
        {
            "Id": 5,
            "DiscountName": "Seasonal",
            "DiscountCategory": "Seasonal",
            "Amount": 30.0,
            "Condition": 400.0
        }
    ]
}
}

```

### Product
| Product | ProductCategory    | Price               |Quantity|
| :-------- | :------- | :------------------------- | :-----|
| T-Shirt | Clothings | 100.0 ฿ |1|
| Hat| Accessories| 250.0 ฿ |1|
| Hoodie| Accessories| 700.0 ฿ |1|
| Watch| Electronics| 850.0 ฿|1|
| Bag| Accessories| 640.0 ฿|1|

### Discount
| Campaign| Discount Category   | Amount               |Point|Condition|
| :-------- | :------- | :------------------------- | :-----|:---|
|FixedAmount | Coupon | 50.0 ฿ |-|-|
|Percentage| Coupon| 10% |-|-|
|PercentageByCategory| OnTop| 15% on Clothings |-|-|
| Point| OnTop| - |68|-|
| Seasonal| Seasonal| 40.0 ฿ |-|300.0 ฿|