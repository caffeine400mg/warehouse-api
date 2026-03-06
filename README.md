
--CREATE PRODUCT--
path: /products/create
method: POST
JSON:
{
    "product_name": "",
    "product_price": O,
    "product_category": "" //food, tools, clothes
}

--ADD PRODUCT--
path: /products/add/{id}
method: PATCH
JSON:
{
    "amount": 0
}

--REMOVE PRODUCT--
path: /products/remove/{id}
method: PATCH
JSON:
{
    "amount": 0
}

--SELL PRODUCT--
path: /products/sell/{id}
method: PATCH
JSON:
{
    "amount": 0
}

--SHOW ALL PRODUCTS--
path: /products/all
method: GET

--SHOW PROFIT--
path: /warehouse/profit
method: GET