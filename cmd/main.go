package main // отвечает за запуск програаммы

import (
	Pet_project_ToDoApp "Pet-project-ToDoApp"
	"Pet-project-ToDoApp/pkg/handler"
	"Pet-project-ToDoApp/pkg/repository"
	"Pet-project-ToDoApp/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"

	//"github.com/zhashkevych/todo-app"
	//"github.com/zhashkevych/todo-app/pkg/handler"
	"log"
)

func main() {
	//handlers := new(handler.Handler)

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variable: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Pet_project_ToDoApp.Server)                                          // инициализируем экземпляр сервера с помощью ключевого слова
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil { // запустим сервер //viper key испровить ошибку
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
