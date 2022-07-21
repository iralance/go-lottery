package conf

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool
}

// 系统中所有mysql主库 root:mysql123456@tcp(127.0.0.1:3306)/go-lottery?charset=utf-8
var DbMasterList = []DbConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "root",
		Database:  "go-lottery",
		IsRunning: true,
	},
}

var DbMaster DbConfig = DbMasterList[0]
