package warehouse

import "errors"

var ErrInvalidCategory = errors.New("WRONG PRODUCT CATEGORY")
var ErrInvalidID = errors.New("WRONG PRODUCT ID")
var ErrInvalidName = errors.New("WRONG PRODUCT NAME")
var ErrInvalidPrice = errors.New("WRONG PRODUCT PRICE")
var ErrInvalidQuantity = errors.New("WRONG PRODUCT QUANTITY")
var ErrNilWarehouse = errors.New("NIL WAREHOUSE")
var ErrProductNotFound = errors.New("PRODUCT NOT FOUND")
var ErrInvalidAmount = errors.New("WRONG AMOUNT")
var ErrNotEnoughProducts = errors.New("NOT ENOUGH PRODUCTS IN WAREHOUSE")
var ErrProductAlreadyExist = errors.New("PRODUCT ALREADY EXIST")
