# Gorm Replicated

This library is a complementary library for [Gorm (v2)](https://github.com/go-gorm/gorm)
which resolves the first available pool passed to it.

## Installation

```bash
go get github.com/partizaans/gormreplicated
```

## Example

Before reading the example it is better to check the gorm [official docs](https://gorm.io/docs/dbresolver.html).

```golang
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/partizaans/gormreplicated"
)

func main() {
	db, err := gorm.Open(mysql.Open("db1_dsn"), &gorm.Config{})

	db.Use(dbresolver.Register(dbresolver.Config{
		// Use `db2`, `db1` as replica options
		// It can be as many as replicas you want
		// Note that the ordering is important
		Replicas: []gorm.Dialector{mysql.Open("db2_dsn"), mysql.Open("db1_dsn")},
		Policy:   gormreplicated.ReplicaPolicy{},
	}, "replicas"))
}
```