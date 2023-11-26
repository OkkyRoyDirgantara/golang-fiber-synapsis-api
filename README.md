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

## Installation from docker hub

pull the docker

```
docker pull okkyroydirgantara/golang-fiber-synapsis-api:latest
```

Create Network

```
docker network create db
```

Run MySQL

```
docker run -d -p 3306:3306 --network=db --name=mysql -e MYSQL_ROOT_PASSWORD=rahasia -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin -e MYSQL_HOST=db -e MYSQL_PORT=3306 -e MYSQL_DATABASE=synapsis mysql:latest
```

Run The Apps

```
docker run -d -p 3000:3000 --name=synapsis-rest --network=db -e DB_USERNAME=admin -e DB_PASSWORD=admin   -e DB_HOST=mysql -e DB_PORT=3306 -e DB_NAME=synapsis okkyroydirgantara/golang-fiber-synapsis-api:latest
```

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
