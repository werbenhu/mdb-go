//
//  @File : mdb.go
//	@Author : WerBen
//  @Email : 289594665@qq.com
//	@Time : 2021/2/5 15:48 
//	@Desc : TODO ...
//

package mdb

import (
	"context"
	"github.com/werbenhu/mdb-go/mongo"
	"gopkg.in/mgo.v2"
	"log"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex
var instance *Mdb

type Opt struct {
	Context context.Context
	PoolSize int
	Host string //	mongodb://username:password@1270.0.01:27017
}

func (opt *Opt) Build() []OptItem{
	return []OptItem {
		OptCtx(opt.Context),
		OptHost(opt.Host),
		OptPoolSize(opt.PoolSize),
	}
}

type Mdb struct {
	opt *Opt
	mdbArray []*mongo.Session
	mu sync.Mutex
}

type IOptItem interface {
	apply(*Opt)
}

type OptItem struct {
	inject func(opt *Opt)
}

func (item *OptItem) apply(opt *Opt) {
	item.inject(opt)
}

func NewOptItem(inject func(opt *Opt)) OptItem {
	return OptItem{
		inject: inject,
	}
}

func OptCtx(ctx context.Context) OptItem {
	return NewOptItem(func(opt *Opt) {
		opt.Context = ctx
	})
}

func OptHost(host string) OptItem {
	return NewOptItem(func(opt *Opt) {
		opt.Host = host
	})
}

func OptPoolSize(size int) OptItem {
	return NewOptItem(func(opt *Opt) {
		opt.PoolSize = size
	})
}

func Get() *mongo.Session {
	if nil == instance {
		log.Fatalf("Error the mdb is not initialized \n")
		return nil
	}

	return instance.Get()
}

func Collection(database string, name string) *mgo.Collection {
	if nil == instance {
		log.Fatalf("Error the mdb is not initialized \n")
		return nil
	}
	db := instance.Get()
	return db.DB(database).C(name)
}

func Init(opts...OptItem) func() {

	mu.Lock()
	defer mu.Unlock()

	instance = new(Mdb)
	// default options
	opt := &Opt{
		Context:context.Background(),
		PoolSize: 1,
		Host: "mongodb://root:123456@127.0.0.1:27017",
	}

	// set options by args
	for _, o := range opts {
		o.apply(opt)
	}
	instance.opt = opt

	if opt.PoolSize < 1 {
		log.Fatalf("Error pool size illegal\n")
	}

	instance.mdbArray = make([]*mongo.Session, opt.PoolSize)
	return instance.Destroy
}

func (m *Mdb) Destroy() {
	for k, v := range m.mdbArray {
		if nil != v {
			v.Close()
			m.mdbArray[k] = nil
		}
	}
}

func (m *Mdb) Collection(database string, name string) *mgo.Collection {
	db := m.Get()
	return db.DB(database).C(name)
}

func (m *Mdb) Get() *mongo.Session {
	length := len(m.mdbArray)
	index := 0
	if 1 < length {
		rand.Seed(time.Now().UnixNano())
		index = rand.Intn(length)
	}
	if nil == m.mdbArray[index]{
		mu.Lock()
		defer mu.Unlock()
		m.mdbArray[index] = mongo.New(m.opt.Host)
	}
	return m.mdbArray[index]
}