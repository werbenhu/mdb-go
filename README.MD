# Example

## global connection-pool-全局的连接池
```
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/werbenhu/mdb-go/mdb"
	"gopkg.in/mgo.v2/bson"
)

const (
	mdbUrl = "mongodb://root:pwd@127.0.0.1:27017"
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
	r.Run("0.0.0.0:9008")
}
```

## independent connection-单独使用一个连接
```
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/werbenhu/mdb-go/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	mdbUrl = "mongodb://root:pwd@127.0.0.1:27017"
)

func main() {
	r := gin.Default()

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

```