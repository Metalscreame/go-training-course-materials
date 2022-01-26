package main

import (
	"context"
	"database/sql"
	"fmt"
)

/*
Straight to the point, Golang Context is a powerful
tool that is used to share scoped data, cancellation,
and timeout when you are working with concurrent
programming with Go.

A little introduction about Contexts
A Context carries a deadline, a cancellation
signal, and other values across API boundaries, as you can see
from the context interface.

type Context interface {
  Deadline() (deadline time.Time, ok bool)
  Done() <-chan struct{}
  Err() error
  Value(key interface{}) interface{}
}

Deadline() returns the time when this context will be cancelled, if any.
Done() returns a channel that is closed when the context is cancelled or times out.
Err() returns the reason why this contexts was cancelled, after Done() is closed.
Value works like a key-value and is used to share data.
There are pretty much two approaches when you're working with contexts.
Either you propagate the original context or create a child context by calling WithCancel,
WithDeadline, WithTimeout or WithValue. Whenever you create a child, you receive a
new context from the original one, that said, when the original is cancelled, all contexts
derived from it are also cancelled.
*/

type ctxKey string

const dbDataKey ctxKey = "db-key"

type DbData struct {
	dbName string
	db     *sql.DB
}

func main() {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, dbDataKey, DbData{db: &sql.DB{}, dbName: "PostgreSQL"})

	err := processing(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("finished")
}

func processing(ctx context.Context) error {
	dbData, ok := ctx.Value(dbDataKey).(DbData)
	if !ok {
		return fmt.Errorf("no data found in context")
	}

	db := dbData.db
	name := dbData.dbName
	fmt.Printf("db %v, %v db name\n", db, name)
	return nil
}
