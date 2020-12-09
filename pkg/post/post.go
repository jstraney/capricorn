package post

import (
  "database/sql"
  "github.com/gin-gonic/gin"
)

func Index(db *sql.DB) gin.HandlerFunc {
  return func (c *gin.Context) {
    c.JSON(200, gin.H{
      "yis": "boi",
    })
  }
}

func Create(db *sql.DB) gin.HandlerFunc {
  return func (c *gin.Context) {
    c.JSON(200, gin.H{
      "yis": "boi",
    })
  }
}
