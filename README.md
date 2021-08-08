# **Online Store** #

We are members of the engineering team of an online store. When we look at ratings for our online store application, we received the following facts:

1. Customers were able to put items in their cart, check out, and then pay. After several days, many of our customers received calls from our Customer Service department stating that their orders have been canceled due to stock unavailability.
2. These bad reviews generally come within a week after our 12.12 event, in which we held a large flash sale and set up other major discounts to promote our store.

After checking in with our Customer Service and Order Processing departments, we received the following additional facts:

1. Our inventory quantities are often misreported, and some items even go as far as having a negative inventory quantity.
2. The misreported items are those that performed very well on our 12.12 event.
3. Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of orders, and thus requested help from our Customer Service department to call our customers and notify them that we have had to cancel their orders.

And Now

 **1. Describe what you think happened that caused those bad reviews during our 12.12 event and why it happened**

 I think this happened because there was no stock validation at the time the order was made. And at the same time, there is a *Race Condition*, where the final result is sometimes not as desired.

 **2. Describe what you think happened that caused those bad reviews during our 12.12 event and why it happened**

 The solution from my analysis is to add stock validation at the time the order is made. Then add *Mutex* to handle *Race Condition*

 **3. Based on your proposed solution, build a Proof of Concept that demonstrates technically how your solution will work.**

 For the Proof of Concept that I made, can be seen in the following lines
<br>
<br>

# **How To Use PoC**

## **Installation**
First
```
git clone https://github.com/fauzanmh/online-store
```
Then
```
go mod vendor && go get -u github.com/swaggo/swag/cmd/swag
```
<br>

## **Setting Environment**
First, copy and rename main.example.json to main.json then change the Env of the Postgres database to what you have, example:
```json
{
    "pg": {
        "host": "localhost",
        "port": "5432",
        "dbname": "online_store",
        "user": "postgres",
        "password": "",
        "sslmode": "disable",
        "max_open_connection": 10,
        "max_idle_connection": 5,
        "max_connection_lifetime": "5m"
    }
}
```
<br>

## **Migration**
First, install [Golang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) and then setting connection database same as before
```
DB_USER=postgres
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=online_store
DB_SSL=disable
```
Then

```
migrate -source file:./script/migration/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL}
```
Or
```makefile
make migration-up
```
<br>

## **Run Service**
```
swag init && go run main.go
```
Or
```
make run 
```
<br>

## **Documentation**

### **1. Endpoint(s)**
Using Swagger
```
http://localhost:8099/api/swagger/index.html
```

### **2. Database Design**
In folder
```
script/migration
```
<br>

## **Run Testing**
```
go test -v -cover -coverprofile=cover.out ./unit_test
```
Or
```
make test
```