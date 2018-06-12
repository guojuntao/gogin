package db

import (
	"git.finogeeks.club/finochat/go-gin/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "github.com/globalsign/mgo"
	// "github.com/globalsign/mgo/bson"
	"net"
)

var (
	dbses Session
	cfg   config.Config
)

func init() {
	cfg = config.GetConfig()
	dbses = InitDB(cfg.MongoDomain, cfg.MongoPort, cfg.MongoAuth)
	itemCol := dbses.GetDBCol(cfg.DbName, cfg.ItemCollection)

	if err := EnsureIndex(itemCol, "id", true); err != nil {
		panic(err)
	}
}

type (
	Session struct {
		*mgo.Session
	}

	Collection struct {
		*mgo.Collection
	}
)

// 后台存储模型
type (
	Item struct {
		ID   string `bson:"id" json:"id"`
		Name string `bson:"name" json:"name"`
		Age  int    `bson:"age" json:"age"`
	}
)

func InitDB(domain string, port string, auth string) (ses Session) {
	addrs, err := net.LookupHost(domain)
	if err != nil {
		panic(err)
	}
	url := auth
	for i, addr := range addrs {
		if i == 0 {
			url = url + addr + ":" + port
		} else {
			url = url + "," + addr + ":" + port
		}
	}

	ses.Session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	ses.SetPoolLimit(50)
	return ses
}

func EnsureIndex(col Collection, primaryKey string, unique bool) (err error) {
	index := mgo.Index{
		Key:        []string{primaryKey},
		Unique:     unique,
		DropDups:   unique,
		Background: true,
		Sparse:     true,
	}
	err = col.EnsureIndex(index)
	return err
}

func (ses Session) GetDBCol(dbName string, colName string) (col Collection) {
	col.Collection = ses.DB(dbName).C(colName)
	return
}

func (ses Session) UninitDB() {
	ses.Close()
}

func FindItem(ID string) (obj Item, err error) {
	ses := Session{dbses.Session.Copy()}
	defer ses.Session.Close()
	col := ses.GetDBCol(cfg.DbName, cfg.ItemCollection)
	err = col.Find(bson.M{"id": ID}).One(&obj)
	return
}

func InsertItem(obj Item) (err error) {
	ses := Session{dbses.Session.Copy()}
	defer ses.Session.Close()
	col := ses.GetDBCol(cfg.DbName, cfg.ItemCollection)
	err = col.Insert(&obj)
	return
}

func UpdateItem(ID string, obj Item) (err error) {
	ses := Session{dbses.Session.Copy()}
	defer ses.Session.Close()
	col := ses.GetDBCol(cfg.DbName, cfg.ItemCollection)
	_, err = col.Upsert(bson.M{"id": ID}, obj)
	return
}

func DeleteItem(ID string) (err error) {
	ses := Session{dbses.Session.Copy()}
	defer ses.Session.Close()
	col := ses.GetDBCol(cfg.DbName, cfg.ItemCollection)
	err = col.Remove(bson.M{"id": ID})
	return
}
