package auth

import (
  "log"
  "net/http"
  "database/sql"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
  "github.com/jstraney/capricorn/pkg/jwt"
)

type Authorization struct {
  Username string `json:"username" binding:"required"`
  Secret string `json:"secret" binding:"required"`
}

func SignIn (db *sql.DB) gin.HandlerFunc {

  return func (c *gin.Context) {

    var auth Authorization

    var hash string

    c.BindJSON(&auth)

    err := db.QueryRow("SELECT username, hash FROM user WHERE username = ?", auth.Username).Scan(&auth.Username, &hash)

    if err != nil {
      // give 500 error
      log.Fatal(err)
      c.JSON(http.StatusInternalServerError, nil)
    }

    err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(auth.Secret))

    if err != nil {
      // give 403 error
      c.JSON(http.StatusUnauthorized, nil)
    }

    jsonWebToken, err := jwt.GenerateJWT(auth.Username, true)

    if err != nil {
      c.JSON(http.StatusInternalServerError, nil)
    }

    // set jsonWebToken in header
    c.JSON(http.StatusOK, gin.H{
      "success" : true,
      "token" : jsonWebToken,
    })

  }

}

func SignOut (db *sql.DB) gin.HandlerFunc {
  return func (c *gin.Context) {
  }
}

func Refresh (db *sql.DB) gin.HandlerFunc {
  return func (c *gin.Context) {
  }
}

// middleware to require authorization
func AuthorizeJWT (claimedRoles []string) gin.HandlerFunc {
  return func (c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    bearerJWT := authHeader[len("Bearer"):]
    token, err := jwt.ValidateJWT(bearerJWT)
    if token.Valid {
      // check the claimed roles match at least one
    } else {
      log.Fatal(err)
      c.AbortWithStatus(http.StatusUnauthorized)
    }
  }
}
