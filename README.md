# final-project-denizedizcan-propertyfinder
 Property Finder Go Bootcamp - Final Project

# Basket service

This is a REST API developed in Go
As a data base using postgres

in this api you can:
create user, login and show user data
Insert Products to the db
list products
Insert stock data to db
List stock data
Update stock data
Insert price data
list price data
add products to users basket
update quantity of products in the basket
remove products from the basket
and create an order

also when we add items to the basket or update it or delete it
system calculate a discount value in order the buisness rules

for example:

. Every fourth order whose total is more than given amount may have discount
depending on products. Products whose VAT is %1 don’t have any discount
but products whose VAT is %8 and %18 have discount of %10 and %15
respectively.

. If there are more than 3 items of the same product, then fourth and
subsequent ones would have %8 off.

. If the customer made purchase which is more than given amount in a month
then all subsequent purchases should have %10 off.

. Only one discount can be applied at a time. Only the highest discount should
be applied.

to run the api 

go run api/cmd/main.go

don't forget the fix the db url to connect your db in api/db/db.go

and its ready to use it will create the tables it self and constraint the primary and foreign keys for you.