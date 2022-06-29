package providers

import (
	"context"
	"log"
	"time"

	"gitee.com/zhenyangze/gin-framework/helpers"
	"github.com/hibiken/asynq"
)

var nqClient *asynq.Client
var nqServer *asynq.Server

func InitAsynq() {
	var config *helpers.Config
	r := asynq.RedisClientOpt{
		Addr: config.GetStringByDefault("asynq.addr", config.GetStringByDefault("redis.addr", "localhost:6379")),
		DB:   config.GetIntByDefault("asynq.db", config.GetIntByDefault("redis.db", 0)),
	}
	nqClient = asynq.NewClient(r)

	srv := asynq.NewServer(r, asynq.Config{
		Concurrency: config.GetIntByDefault("asynq.db", 10),
	})

	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}

func AddNq(event string, times time.Duration, playload string, opts ...asynq.Option) error {
	t := asynq.NewTask(event, []byte(playload), opts...)
	_, err := nqClient.Enqueue(t, asynq.ProcessIn(times))
	if err != nil {
		return err
	}
	return nil
}

func handler(ctx context.Context, t *asynq.Task) error {
	payload := t.Payload()
	event := t.Type()
	Event.Publish(event, string(payload))
	return nil
}
