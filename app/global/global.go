package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"oj/app/internal/model/config/config"
)

var (
	Config       *config.Config
	Logger       *zap.SugaredLogger
	MysqlDB      *sqlx.DB
	DbLike       *redis.Client
	DbCollection *redis.Client
	DbVerify     *redis.Client
	DbSubmit     *redis.Client
	RedisContext = context.Background()
)
