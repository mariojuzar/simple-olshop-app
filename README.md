# Simple Database App

### What's in this repository

- DDL to create table in mysql
- struct table in golang
- read CSV files to populate the table (from given csv file) (in `database service`)
- answer for question  for given test
- endpoint to display order details


## Endpoint to display Order Details

endpoint: `{host}:{port}/api/v1/orders/details/:id`

example in this app: `localhost:7098/api/v1/orders/details/1`

sample response:

```
{
   "server_time": "2019-12-26T15:51:44.90374+07:00",
   "code": 200,
   "message": "OK",
   "data": {
       "order_id": 1,
       "customer_name": "JoBerry",
       "employee_name": "AdamBarr",
       "shipping_method": "Federal Express",
       "product_list": [
           {
               "product_id": 4,
               "product_name": "Black Mug",
               "quantity": 15,
               "unit_price": 5,
               "discount": 0,
               "sub_total_price": 75
           },
           {
               "product_id": 1,
               "product_name": "Blue Mug",
               "quantity": 9,
               "unit_price": 6,
               "discount": 0,
               "sub_total_price": 54
           }
       ],
       "total_payment": 129
   }
}
```

## Run application

### Using go build

before build and run apps, ensure to migrate DDL.sql to your destination mysql

build application with script `go build main.go`

run application `./main`

run app with auto build `go run main.go`