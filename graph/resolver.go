package graph
//go:generate go run github.com/99designs/gqlgen generate

import (	
	"gorm.io/gorm"
	"github.com/KingBean4903/graphql-vod-platform/internal/realtime"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *gorm.DB
	PubSub *realtime.RedisPubSub
}
