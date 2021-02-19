package model

import (
	"context"
	"log"

	"github.com/qiniu/qmgo"
)

func ProvideDbClient(conf *Config, logger *log.Logger) (client *qmgo.QmgoClient, ctx context.Context) {
	context := context.Background()
	client, err := qmgo.Open(context, &qmgo.Config{Uri: conf.Database.Host, Database: conf.Database.DatabaseName, Coll: conf.Database.TableName})
	if err != nil {
		logger.Println("err")
	}
	return client, context
}

/*
var url = "mongodb://127.0.0.1:27017"
var database = "pub"
var collection = "weather"

// 获取mongo数据库连接
func GetDbClient() (client *qmgo.QmgoClient, ctx context.Context) {
	context := context.Background()
	client, err := qmgo.Open(context, &qmgo.Config{Uri: url, Database: database, Coll: collection})
	if err != nil {
		fmt.Println("err")
	}
	return client, context

}
*/
