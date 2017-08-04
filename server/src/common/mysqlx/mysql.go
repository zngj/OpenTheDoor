package mysqlx

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/ziutek/mymysql/godrv"
	"fmt"
	"github.com/carsonsx/log4g"
	"runtime/debug"
	"common/file"
)

var url string
var _db *sql.DB

var config struct {
	Addr     string `json:"addr"`
	DB       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init()  {

	file.LoadJsonConfig("mysql.json", &config)

	url = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Asia%%2FChongqing",
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

func GetDBConn() *sql.DB  {
	return _db
}

type MyTx struct {
	*sql.Tx
	count int
}

func (t *MyTx) Begin()  {
	t.count++
}

func Exec(t *MyTx, query string, args ...interface{}) error {
	if t == nil {
		tx, err := _db.Begin()
		if err != nil {
			return err
		}
		t = new (MyTx)
		t.Tx = tx
	}
	t.Begin()
	var err error
	defer func() { t.Commit(err) }()
	_, err = t.Exec(query, args...)
	if err != nil {
		log4g.Error(err)
	}
	return err
}

func Exists(query string, args ...interface{}) (bool, error) {
	rows, err := GetDBConn().Query(query, args...)
	if err != nil {
		log4g.Error(err)
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}


func (t *MyTx) Commit(err error)  {
	defer func() { t.rollback(err) }()
	if err != nil {
		return
	}
	if t.count == 1 {
		err = t.Tx.Commit()
		if err != nil {
			log4g.Error(err)
			return
		}
	}
	if t.count > 0 {
		t.count--
	}
	//do nothing if count is 0
}


func (t *MyTx) rollback(err error)  {
	rb := false
	if r := recover(); r != nil {
		log4g.Error("********************* Data Access Panic *********************")
		log4g.Error(r)
		log4g.Error(string(debug.Stack()))
		log4g.Error("********************* Data Access Panic *********************")
		rb = true
	} else if err != nil {
		rb = true
	}
	if rb && t.count > 0 {
		log4g.Error("rollback because of error")
		e := t.Tx.Rollback()
		if e != nil {
			log4g.Error(e)
		}
		if t.count > 0 {
			t.count = 0
		}
	}
}
