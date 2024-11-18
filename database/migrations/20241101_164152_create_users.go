package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUsers_20241101_164152 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsers_20241101_164152{}
	m.Created = "20241101_164152"

	migration.Register("CreateUsers_20241101_164152", m)
}

// Run the migrations
func (m *CreateUsers_20241101_164152) Up() {
	m.SQL("CREATE TABLE users(id SERIAL PRIMARY KEY, name VARCHAR(255), email VARCHAR(255) UNIQUE, password VARCHAR(255), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
}

// Reverse the migrations
func (m *CreateUsers_20241101_164152) Down() {
	m.SQL("DROP TABLE users")
}
