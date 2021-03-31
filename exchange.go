package polis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Opts struct {
	MaxThreadDepth        uint
	MaxConnCurrentStreams uint

	PersistThreads    bool
	AllowMessageEdits bool
}

type Exchange struct{}

func NewExchange() (*Exchange, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rdb.Subscribe(context.Background(), "mychan1")
	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	defer pubsub.Close()
	msg := <-pubsub.Channel()
}
