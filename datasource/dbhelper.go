package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/iralance/go-lottery/conf"
	"log"
	"sync"
)

var dbLock sync.Mutex
var masterInstance *xorm.Engine
var slaveInstance *xorm.Engine

// 得到唯一的主库实例
func InstanceDbMaster() *xorm.Engine {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()

	if masterInstance != nil {
		return masterInstance
	}
	return NewDbMaster()
}

func NewDbMaster() *xorm.Engine {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database,
	)
	instance, err := xorm.NewEngine(conf.DriverName, sourceName)
	if err != nil {
		log.Fatal("dbhelper.NewDbMaster NewEngine error ", err)
		return nil
	}
	instance.ShowSQL(true)
	masterInstance = instance
	return instance
}
