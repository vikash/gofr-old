package main

import (
	"errors"
	"fmt"
	"time"

	ezgoErr "github.com/zopsmart/ezgo/pkg/gofr/errors"

	"github.com/go-redis/redis/v8"

	"github.com/zopsmart/ezgo/pkg/gofr"
)

func main() {
	// Create a new application
	a := gofr.New()

	// Add all the routes
	a.GET("/hello", HelloHandler)
	a.POST("/hello", PostHandler)
	a.GET("/error", ErrorHandler)
	a.GET("/redis", RedisHandler)
	a.GET("/trace", TraceHandler)
	a.GET("/mysql", MysqlHandler)

	// Run the application
	a.Run()
}

func HelloHandler(c *gofr.Context) (interface{}, error) {
	name := c.Param("name")
	if name == "" {
		c.Log("Name came empty")
		name = "World"
	}

	return fmt.Sprintf("Hello %s!", name), nil
}

// Use case for 200 and 201 status codes in ezgo.
type entity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func PostHandler(c *gofr.Context) (interface{}, error) {
	var e entity
	err := c.Bind(&e)
	if err != nil {
		return nil, errors.New("error JSON format")
	}

	const query = `INSERT INTO names(id, name) values(?,?)`
	row, err := c.DB.Exec(query, e.ID, e.Name)
	if err != nil {
		// Duplicate entry
		return nil, nil
	}

	id, _ := row.LastInsertId()
	if id == -1 {
		return -1, errors.New("record not found")
	}

	return id, ezgoErr.NewEntity{}
}

func ErrorHandler(c *gofr.Context) (interface{}, error) {
	return nil, errors.New("some error occurred")
}

func RedisHandler(c *gofr.Context) (interface{}, error) {
	val, err := c.Redis.Get(c, "test").Result()
	if err != nil && err != redis.Nil { // If key is not found, we are not considering this an error and returning "".
		return nil, err
	}

	return val, nil
}

func TraceHandler(c *gofr.Context) (interface{}, error) {
	defer c.Trace("traceHandler").End()

	span2 := c.Trace("handler-work")
	<-time.After(time.Millisecond * 1) // Waiting for 1ms to simulate workload
	span2.End()

	return "Tracing Success", nil
}

func MysqlHandler(c *gofr.Context) (interface{}, error) {
	var value int
	err := c.DB.QueryRowContext(c, "select 2+2").Scan(&value)

	return value, err
}
