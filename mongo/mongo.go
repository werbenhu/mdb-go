//
//  @File : mongo.go
//	@Author : WerBen
//  @Email : 289594665@qq.com
//	@Time : 2021/2/5 16:27 
//	@Desc : TODO ...
//

package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type Session struct {
	mgo.Session
}

func (s *Session) Collection(database string, name string) *mgo.Collection {
	return s.DB(database).C(name)
}

func New(url string) *Session {
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("Error pool size illegal\n")
		return nil
	}
	return &Session{*session}
}
