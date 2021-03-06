//
//  @File : main.go
//	@Author : WerBen
//  @Email : 289594665@qq.com
//	@Time : 2021/2/5 16:03 
//	@Desc : TODO ...
//

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/werbenhu/mdb-go/mdb"
	"github.com/werbenhu/mdb-go/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	mdbHost = "mongodb://root:pwd@127.0.0.1:27017"
)

func main() {
	mdbDestroy := mdb.Init(
		mdb.OptHost(mdbHost),
		mdb.OptPoolSize(2))
	defer mdbDestroy()

	//opt := &mdb.Opt{
	//	Host: mdbHost,
	//	PoolSize: 2,
	//	Context: context.Background(),
	//}
	//mdbDestroy := mdb.Init(opt.Build()...)
	//defer mdbDestroy()

	r := gin.Default()

	r.GET("/mdb", func(c *gin.Context) {

		var ret []map[string]interface{}
		mc := mdb.Collection("bugDb", "list")
		err := mc.Find(bson.M{"openid": "Dy473ymJqxtqL"}).All(&ret)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": ret,
		})
	})

	r.GET("/mongo", func(c *gin.Context) {

		mongoDb := mongo.New(mdbHost)
		mongoDb.Collection("bugDb", "list")
		defer mongoDb.Close()

		var ret []map[string]interface{}
		mc := mdb.Collection("bugDb", "list")
		err := mc.Find(bson.M{"openid": "Dy473ymJqxtqL"}).All(&ret)
		if err != nil {
			c.JSON(200, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": ret,
		})
	})

	r.Run("0.0.0.0:9008")
}

