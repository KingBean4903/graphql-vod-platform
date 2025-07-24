package realtime

import (
	"context"
	"fmt"
	"encoding/json"

	"github.com/KingBean4903/graphql-vod-platform/internal/redis"
)

type RedisPubSub struct {}

func NewRedisPubSub() *RedisPubSub {
	return &RedisPubSub
}

func (r *RedisPubSub) Publish(topic string, msg string) error {
	
		data, err := json.Marshal(msg)
		if err != nil {
				return err
		}

		return redis.Client.Publish(context.Background(), topic, data).Err()
}

func (r *RedisPubSub) Subscribe(topic string) (<-chan any, error) {
	
	out := make(chan any, 1)
	sub := redis.Client.Subscribe(context.Background(), topic)
	ch := sub.Channel()

	go func() { 
			for msg := range ch {
					var comment map[string]any
					if err := json.Unmarshal([]byte(msg.Payload), &comment); err == nil {
								out <- comment
					}
			}
	}()

	return out, nil

}
