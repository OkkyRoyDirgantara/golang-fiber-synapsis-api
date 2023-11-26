# Synapsis Rest API

## Requirement

Create with requirement

- Golang go1.19.3 windows/amd64
- MySQL v8.0.30

## Documentation API

Link Postman : [API Spec](https://www.postman.com/restless-satellite-772023/workspace/synapsis-api)

## Installation

setup installation if you want to run without docker
and configure file .env setup with your configuration

```
cp .env.example .env
```

```
go install
```

```
go get .
```

Run program

```
go run main.go
```

## Docker

Build a Docker image

```
docker-compose build
```

Running Docker image

```
docker-compose up -d
```

attention dont change file .env.example without configur docker configuration

## Flow the apps

1. Register customers
2. Login customers account
3. Get JWT token
4. Access endpoint with JWT token if middleware is available

Feature :

1. Register Customer
2. Login customer
3. Create Product
4. Add to cart
5. Simple Transaction (this is by soft delete product_cart)
