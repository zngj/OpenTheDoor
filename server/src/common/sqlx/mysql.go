package sqlx

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/ziutek/mymysql/godrv"
	"common/file"
	"fmt"
	"github.com/carsonsx/log4g"
)

var url string
var _db *sql.DB

var config struct {
	Addr     string `json:"addr"`
	DB       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {

	defer func() {
		if err := recover(); err != nil {
			log4g.Error(err)
		}
	}()

	file.LoadJsonConfig("mysql.json", &config)

	if config.Username == "" {
		config.Username = "root"
	}

	if config.Password == "" {
		config.Password = "root"
	}

	url = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Asia%%2FChongqing&parseTime=true",
		config.Username,
		config.Password,
		config.Addr,
		config.DB)

	//url = fmt.Sprintf("tcp:%s*%s/%s/%s",
	//	mySQLConfig.Host,
	//	mySQLConfig.Database,
	//	mySQLConfig.User,
	//	mySQLConfig.Password)

	log4g.Info("db url: %s", url)

	var err error
	_db, err = sql.Open("mysql", url)
	//_db, err = sql.Open("mymysql", url)
	if err != nil {
		panic(err)
	}
	_db.SetMaxOpenConns(100)
	_db.SetMaxIdleConns(50)
	if err = _db.Ping(); err != nil {
		panic(err)
	}
}

func GetDBConn() *sql.DB {
	return _db
}
