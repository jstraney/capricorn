package db

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func Init(user string, pass string, host string, port string, database string) (*sql.DB, error){

  db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, database))

  defer db.Close()

  if err != nil {
    return nil, err
  }

  return db, nil

}
