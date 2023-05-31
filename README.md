
# My Gram

This repository is a development of the final project at digitalent.

MyGram is a clone of Instagram, later the features here will be similar to the original application. For now it only has simple features such as login, register, upload photos and leave comments. Because this is still in the development stage, there will be additional features later.


## Built With


- Go
- Gin
- PostgreSQL
- GORM
- Wire
- Godotenv
- Golang JWT
- Gin Swagger


## Run Locally

Clone the project

```bash
  git clone https://github.com/ariwiraa/my-gram.git
```

Go to the project directory

```bash
  cd my-gram
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run main.go
```


## Configuration

Change .env

Change the database configuration according to your own in the .env file

If you want to run swagger, you must first install swagger. **Skip this if you have already done the installer**

```bash
  go install github.com/swaggo/swag/cmd/swag@latest
```

Initialized Swag

```bash
  swag init
```
## Documentation
Since this has not been uploaded to the server yet, while to run it on local first

[Documentation With GinSwagger](http://localhost:8080/swagger/index.html)


## Authors

- [@ariwiraa](https://www.github.com/ariwiraa)

