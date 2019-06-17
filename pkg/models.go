package pkg

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"time"
)

var (
	engine *xorm.Engine
)

func InitDBEngine(conn string) error {
	var err error
	engine, err = xorm.NewEngine("mysql", conn)
	if err != nil {
		return err
	}
	engine.SetMaxIdleConns(0)
	engine.SetConnMaxLifetime(time.Second)
	engine.SetMapper(core.GonicMapper{})

	if has, err := engine.IsTableExist(&Book{}); err != nil {
		return err
	} else if !has {
		engine.CreateTables(&Book{})
	}

	return nil
}

func GetDBEngine() (*xorm.Engine, error) {
	if engine == nil {
		return nil, fmt.Errorf("db engine is not initialized")
	}
	return engine, nil
}

type Book struct {
	ID int `json:"id" xorm:"int not null unique 'id'"`
	Name string `json:"name" xorm:"varchar(25) not null 'name'"`
	Author string `json:"author" xorm:"varchar(25) 'author'"`
	Description string `json:"description,omitempty" xorm:"varchar(255) 'description'"`
	Price *int `json:"price,omitempty" xorm:"int 'price'"`
}

func (b *Book) Validate() error {
	if b.ID == 0 {
		return fmt.Errorf("book id must be non empty")
	}
	if b.Name == "" {
		return fmt.Errorf("book name must be non empty")
	}
	return nil
}

func Merge(old, nu Book) Book {
	if nu.Name != "" {
		old.Name = nu.Name
	}
	if nu.Author != "" {
		old.Author = nu.Author
	}
	if nu.Description != "" {
		old.Description = nu.Description
	}
	if nu.Price != nil {
		old.Price = nu.Price
	}
	return old
}

func (b Book) TableName() string {
	return "book"
}