package main

import (
	"LinkaAja/packages"
	"LinkaAja/src/controllers"
	"LinkaAja/src/repositories"
	"LinkaAja/src/routers"
	"LinkaAja/src/usecases"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	packages.LoadMongo()
}

func main() {
	db := packages.LoadDatabase()
	repository := repositories.NewRepo(db)
	usecase := usecases.NewUC(repository)
	ctrl := controllers.NewCtrl(usecase)
	route := routers.NewRouter(ctrl)
	group := route.RouterGroup()
	log.Printf("[SERVER] starting at port : %v", os.Getenv("SERVER_PORT"))
	log.Fatalln(
		http.ListenAndServe(
			fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")),
			packages.AllowCORS(group),
		),
	)
}
