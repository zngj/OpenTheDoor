package mysqlx

import (
	"database/sql"
	"github.com/carsonsx/log4g"
	"runtime/debug"
	"reflect"
	"runtime"
	"errors"
)

var ErrNotFound = errors.New("not found")

func NewDao() *Dao {
	return new(Dao)
}

type Dao struct {
	Db *sql.DB
	Tx *sql.Tx
	db_counter int
	tx_counter int
}

func (dao *Dao) Connect() {
	if dao.Db == nil {
		log4g.Debug("connected database")
		dao.Db = GetDBConn()
	}
	//dao.db_counter++
	//log4g.Debug("db counter: %d", dao.db_counter)
}

//func (dao *Dao) Disconnect() {
//	if dao.db_counter == 1 {
//		if err := dao.Db.Close(); err != nil {
//			log4g.Error(err)
//		} else {
//			log4g.Debug("disconnected database")
//		}
//	}
//	if dao.db_counter > 0 {
//		dao.db_counter--
//	}
//	log4g.Debug("db counter: %d", dao.db_counter)
//}

func (dao *Dao) BeginTx() error {
	if dao.tx_counter == 0 {
		dao.Connect()
		var err error
		dao.Tx, err = dao.Db.Begin()
		if err != nil {
			log4g.Error(err)
			return err
		}
	}
	dao.tx_counter++
	return nil
}

func (dao *Dao) CommitTx(errs ...error) {
	defer func() { dao.rollback(errs...) }()
	if len(errs) == 0 || errs[0] != nil {
		return
	}
	if dao.tx_counter == 1 {
		err := dao.Tx.Commit()
		if err != nil {
			log4g.Error(errs)
		}
		//dao.Disconnect()
	}
	if dao.tx_counter > 0 {
		dao.tx_counter--
	}

	//do nothing if tx_counter is 0
}

func (dao *Dao) rollback(err ...error)  {
	rb := false
	if r := recover(); r != nil {
		log4g.Error("********************* Data Access Panic *********************")
		log4g.Error(r)
		log4g.Error(string(debug.Stack()))
		log4g.Error("********************* Data Access Panic *********************")
		rb = true
	} else if len(err) > 0 && err[0]!= nil {
		rb = true
	}
	if rb && dao.tx_counter > 0 {
		log4g.Error("rollback because of error")
		e := dao.Tx.Rollback()
		if e != nil {
			log4g.Error(e)
		}
		if dao.tx_counter > 0 {
			dao.tx_counter = 0
		}
	}
}

func (dao *Dao) Exec(query string, args ...interface{}) error {
	dao.BeginTx()
	var err error
	defer func() { dao.CommitTx(err) }()
	_, err = dao.Tx.Exec(query, args...)
	if err != nil {
		log4g.Error(err)
	}
	return err
}

type result struct {
	rows *sql.Rows
	err error
}

type InvalidResultError struct {
	Type reflect.Type
}

func (e *InvalidResultError) Error() string {
	if e.Type == nil {
		return "dao: Result(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "dao: Result(non-pointer " + e.Type.String() + ")"
	}
	return "dao: Result(nil " + e.Type.String() + ")"
}

func (r *result) Result(dest ...interface{}) (err error) {
	if r.err != nil {
		return r.err
	}
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	//rv := reflect.ValueOf(v)
	//if rv.Kind() != reflect.Ptr || rv.IsNil() {
	//	return &InvalidResultError{reflect.TypeOf(v)}
	//}
	//
	//if r.rows != nil {
	//	defer r.rows.Close()
	//}
	//
	//columns, err := r.rows.Columns()
	//if err != nil {
	//	log4g.Error(err)
	//	return err
	//}

	defer r.rows.Close()

	if r.rows.Next() {
		err = r.rows.Scan(dest...)
		if err != nil {
			log4g.Error(err)
			return
		}
	} else {
		err = ErrNotFound
	}

	return
}

func (r *result) Map(dest ...interface{}) (records []map[string]interface{}, err error) {
	columns, err := r.rows.Columns()
	if err != nil {
		log4g.Error(err)
		return nil, err
	}

	defer r.rows.Close()

	for r.rows.Next() {
		err = r.rows.Scan(dest...)
		if err != nil {
			log4g.Error(err)
			return
		}
		record := make(map[string]interface{})
		for i := range columns {
			record[columns[i]] = dest[i]
		}
		records = append(records, record)
	}
	return
}

//func newErrorResult(err error) *result {
//	r := new(result)
//	r.err = err
//	return r
//}

func (dao *Dao) Query(query string, args ...interface{}) *result {

	dao.Connect()
	//defer dao.Disconnect()

	r := new(result)

	log4g.Debug("query sql: %s", query)
	log4g.Debug("query arg: %v", args)

	var rows *sql.Rows
	rows, err := dao.Db.Query(query, args...)
	if err != nil {
		log4g.Error(err)
		r.err = err
		return r
	}

	r.rows = rows

	return r
}