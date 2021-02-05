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
	mdbUrl = "mongodb://root:bu12345@183.134.21.92:27017"
)

func main() {
	mdbDestroy := mdb.Init(
		mdb.OptUrl(mdbUrl),
		mdb.OptPoolSize(2))
	defer mdbDestroy()

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

		mongoDb := mongo.New(mdbUrl)
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
