# Synapsis Rest API

## Requirement

Pada saat aplikasi ini dibuat dengan menggunakan requirement

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
