package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	g "oj/app/global"
	"time"
)

func MysqlSetup() {
	config := g.Config.DataBase.Mysql
	db, err := sqlx.Connect("mysql", config.GetDsn())
	if err != nil {
		g.Logger.Fatalf("initialize mysql db failed, err: %v", err)
	}
	err = db.Ping()
	if err != nil {
		g.Logger.Fatalf("connect to mysql db failed, err: %v", err)
	}
	g.Logger.Infof("initialize mysql db successfully")
	g.MysqlDB = db
}

func RedisSetup() {
	config := g.Config.DataBase.Redis

	likeDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: "",
		Password: config.Password,
		DB:       config.DbLike,
		PoolSize: 10000,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := likeDb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("connect to like redis instance failed, err: %v", err)
	}
	g.DbLike = likeDb

	g.Logger.Info("initialize like redis client successfully")
	collectionDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: "",
		Password: config.Password,
		DB:       config.DbCollection,
		PoolSize: 10000,
	})

	_, err = likeDb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("connect to collection redis instance failed, err: %v", err)
	}
	g.DbCollection = collectionDb

	g.Logger.Info("initialize collection redis client successfully")

	verifyDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: "",
		Password: config.Password,
		DB:       config.DbVerify,
		PoolSize: 10000,
	})

	_, err = likeDb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("connect to verify redis instance failed, err: %v", err)
	}
	g.DbVerify = verifyDb

	g.Logger.Info("initialize verify redis client successfully")

	submitDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: "",
		Password: config.Password,
		DB:       config.DbLook,
		PoolSize: 10000,
	})
	_, err = submitDb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("connect to submit redis instance failed, err: %v", err)
	}
	g.DbSubmit = submitDb

	g.Logger.Info("initialize submit redis client successfully")
}
