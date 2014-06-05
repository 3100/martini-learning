package model

import (
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/mattn/go-sqlite3"
  "time"
)

type Post struct {
  Id      int64 `db:"post_id"`
  Created int64
  Title   string
  Body    string
}

func newPost(title, body string) Post {
  return Post{
    Created: time.Now().UnixNano(),
    Title:   title,
    Body:    body,
  }
}
