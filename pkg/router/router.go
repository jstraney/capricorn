package router

import (
  "os"
  "log"
  "github.com/jstraney/capricorn/pkg/post"
  "github.com/jstraney/capricorn/pkg/person"
  "github.com/jstraney/capricorn/pkg/auth"
  "github.com/jstraney/capricorn/pkg/db"
  "github.com/joho/godotenv"
  "github.com/gin-gonic/gin"
)

func Route() {

  err := godotenv.Load()

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  CapricornPort := os.Getenv("CAPRICORN_PORT")
  CapricornIface := os.Getenv("CAPRICORN_IFACE")

  CapricornMySQLUser:= os.Getenv("CAPRICORN_MYSQL_USER")
  CapricornMySQLPass := os.Getenv("CAPRICORN_MYSQL_PASS")
  CapricornMySQLHost := os.Getenv("CAPRICORN_MYSQL_HOST")
  CapricornMySQLPort := os.Getenv("CAPRICORN_MYSQL_PORT")
  CapricornMySQLDatabase := os.Getenv("CAPRICORN_MYSQL_DATABASE")

  r := gin.New()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  dbInstance, err := db.Init(CapricornMySQLUser, CapricornMySQLPass, CapricornMySQLHost, CapricornMySQLPort, CapricornMySQLDatabase)

  if err != nil {
    log.Fatalf("Error connecting to database")
  }

  api := r.Group("/api/v1")
  {
    api.GET("/person", person.Get(dbInstance))
    api.GET("/posts", post.Index(dbInstance))
    api.POST("/post", post.Create(dbInstance))
    api.POST("/auth/", auth.SignIn(dbInstance))
    api.POST("/auth/revoke", auth.SignOut(dbInstance))
  }

  r.Run(CapricornIface + ":" + CapricornPort)

}
