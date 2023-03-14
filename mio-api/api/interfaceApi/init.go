package interfaceApi

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"mio-api/database"
)

type InterfaceApi struct {
}

var db *gorm.DB
var client *redis.Client
var ctx context.Context

func init() {
	db = database.DB
	client = database.Client
	ctx = database.Ctx
}
