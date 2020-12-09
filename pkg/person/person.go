package person

import (
  "net/http"
  "database/sql"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/bcrypt"
)

func Get(db *sql.DB) gin.HandlerFunc {
  return func (c *gin.Context) {
  }
}

type CreatePersonPayload struct {
  Username string `json:"username" binding:"required"`
  Mobile string `json:"mobile" binding:"required"`
  Password string `json:"hash" binding:"required"`
}

func Create(db *sql.DB) gin.HandlerFunc {
  return func (c *gin.Context) {

    var json CreatePersonPayload

    hash, err := bcrypt.GenerateFromPassword([]byte(json.Password), 14)

    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "success" : false,
      })
    }

    result, err := db.Exec("INSERT INTO cap_person (username) VALUES $1, $2, $3", json.Username, json.Mobile, hash)

    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "success" : false,
      })
    }

    rowsAffected, err := result.RowsAffected()

    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "success" : false,
      })
    }

    c.JSON(http.StatusOK, gin.H{
      "success": true,
      "rowsAffected": rowsAffected,
    })

  }
}
