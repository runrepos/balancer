package storage

import (
	//r "superapp/internal/storage/redis"
	//m "superapp/internal/storage/mongodb"
	g "superapp/internal/storage/mysql_gorm"
	"superapp/config"
	//	"gorm.io/gorm"
	// "context"
	"log"
)

type Storage struct {
	//Redis *r.R //*redis.Client
	//Mongo *m.M
	Mysql *g.MySq
}

// var mongo_ctx = context.Background()



func InitStorage(c *config.Cfg) *Storage { //*redis.Client {
	//redis:=r.InitRedis(config.Config.RedisConnectionUrl);
	mysql, err := g.InitDriver(c.MYSQL_CONNECT_STRING) ;  if err != nil { log.Fatal(err) }
	return &Storage{
		Mysql:  mysql,
		//Redis: r.InitRedis(c.RedisConnectionUrl),
		//Mongo: m.InitMongo(&mongo_ctx, c.MongoConnectionUrl, c.MongoDb, c.MongoColl),
	}
}
