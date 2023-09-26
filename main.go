package main

import (
	di "github.com/ariwiraa/my-gram/wire"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// port := os.Getenv("APP_PORT")
	godotenv.Load(".env")

	// router := routes.NewRouter(userHandler)
	router := di.InitializedServer()

	router.Run(":8080")
}
