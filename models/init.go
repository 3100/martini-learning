package model

import (
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/mattn/go-sqlite3"
  "log"
  "time"
)

func initDb() *gorp.DbMap {
  db, err:= sql.Open("sqlite3", "/tmp/post_db.bin")
  checkErr(err, "sql.Open failed")

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
  dbmap.AddTableWithName(Post{}, "post").SetKeys(true, "Id")
  err = dbmap.CreateTablesIfNotExists()
  checkErr(err, "Create tables failed")
  return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
    log.Fatalln(msg, err)
  }
}
