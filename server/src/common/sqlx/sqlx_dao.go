package sqlx

import (
	"database/sql"
	"errors"
	"github.com/carsonsx/log4g"
	"reflect"
	"runtime"
	"runtime/debug"
)

var ErrNotFound = errors.New("not found")

func NewDao() *Dao {
	return new(Dao)
}

type Dao struct {
	Db         *sql.DB
	Tx         *sql.Tx
	db_counter int
	tx_counter int
}

func (dao *Dao) Connect() {
	if dao.Db == nil {
		//log4g.Trace("connected database")
		dao.Db = GetDBConn()
	}
}

func (dao *Dao) BeginTx() error {
	if dao.tx_counter == 0 {
		dao.Connect()
		var err error
		dao.Tx, err = dao.Db.Begin()
		if err != nil {
			log4g.Error(err)
			log4g.Error("\n" + string(debug.Stack()))
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
			log4g.Error(err)
			log4g.Error("\n" + string(debug.Stack()))
		}
		//dao.Disconnect()
	}
	if dao.tx_counter > 0 {
		dao.tx_counter--
	}

	//do nothing if tx_counter is 0
}

func (dao *Dao) rollback(err ...error) {
	rb := false
	if r := recover(); r != nil {
		log4g.Error("********************* Data Access Panic *********************")
		log4g.Error(r)
		log4g.Error(string(debug.Stack()))
		log4g.Error("********************* Data Access Panic *********************")
		rb = true
	} else if len(err) > 0 && err[0] != nil {
		rb = true
	}
	if rb && dao.tx_counter > 0 {
		log4g.Error("rollback because of error")
		e := dao.Tx.Rollback()
		if e != nil {
			log4g.Error(e)
			log4g.Error("\n" + string(debug.Stack()))
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

	log4g.Debug("exec  sql: %s", query)
	log4g.Debug("exec args: %v", args)

	_, err = dao.Tx.Exec(query, args...)
	if err != nil {
		log4g.Error(err)
		log4g.Error("\n" + string(debug.Stack()))
	}
	return err
}

func (dao *Dao) Query(query string, args ...interface{}) *result {

	dao.Connect()
	//defer dao.Disconnect()

	r := new(result)

	log4g.Debug("query  sql: %s", query)
	log4g.Debug("query args: %v", args)

	var rows *sql.Rows
	rows, err := dao.Db.Query(query, args...)
	if err != nil {
		log4g.Error(err)
		log4g.Error("\n" + string(debug.Stack()))
		r.err = err
		return r
	}

	r.rows = rows

	return r
}

type result struct {
	rows *sql.Rows
	err  error
}

type InvalidResultError struct {
	Type reflect.Type
}

func (e *InvalidResultError) Error() string {
	if e.Type == nil {
		return "sqlx: nil"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "sqlx: non-pointer " + e.Type.String()
	}
	return "sqlx: invalid " + e.Type.String() + ")"
}

func (r *result) Scan(dest ...interface{}) (err error) {
	if r.rows.Next() {
		err = r.rows.Scan(dest...)
		if err != nil {
			log4g.Error(err)
			log4g.Error("\n" + string(debug.Stack()))
			return
		}
	} else {
		err = ErrNotFound
		log4g.Warn("not found any record.")
	}
	return
}

func (r *result) One(v interface{}) (err error) {
	if r.err != nil {
		return r.err
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		err = &InvalidResultError{reflect.TypeOf(v)}
		log4g.Error(err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
			log4g.Error(err)
			log4g.Error("\n" + string(debug.Stack()))
		}
	}()
	defer r.rows.Close()
	cols, err := r.rows.Columns()
	if err != nil {
		log4g.Error(err)
		return
	}
	log4g.Trace("columns: %v", cols)
	dest := make([]interface{}, len(cols))
	getValuesByTags(rv, cols, dest, true)
	if r.rows.Next() {
		err = r.rows.Scan(dest...)
		if err != nil {
			log4g.Error("columns: %v", cols)
			log4g.Error(err)
			log4g.Error("\n" + string(debug.Stack()))
			return
		}
		log4g.Debug(log4g.JsonFunc(v))
	} else {
		err = ErrNotFound
		log4g.Warn("not found any record.")
	}
	return
}

func (r *result) All(v interface{}) (err error) {
	if r.err != nil {
		return r.err
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		err = &InvalidResultError{reflect.TypeOf(v)}
		log4g.Error(err)
		return
	}
	if Indirect(rv.Type()).Kind() != reflect.Slice {
		err = errors.New("required slice")
		log4g.Error(err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
			log4g.Error(err)
			log4g.Error("\n" + string(debug.Stack()))
		}
	}()
	defer r.rows.Close()
	cols, err := r.rows.Columns()
	if err != nil {
		log4g.Error(err)
		return
	}
	log4g.Trace("columns: %v", cols)
	isPtr := Indirect(rv.Type()).Elem().Kind() == reflect.Ptr
	elType := Indirect(Indirect(rv.Type()).Elem())
	dest := make([]interface{}, len(cols))
	rv = reflect.Indirect(rv)
	for r.rows.Next() {
		ep := reflect.New(elType)
		el := reflect.Indirect(ep)
		getValuesByTags(el, cols, dest, true)
		err = r.rows.Scan(dest...)
		if err != nil {
			log4g.Error(err)
			log4g.Error("\n" + string(debug.Stack()))
			return
		}
		log4g.Trace(ep)
		if isPtr {
			rv.Set(reflect.Append(rv, ep))
		} else {
			rv.Set(reflect.Append(rv, el))
		}
	}
	return
}

func (r *result) Exist() (exist bool, err error) {
	if r.err != nil {
		err = r.err
		return
	}
	exist = r.rows.Next()
	return
}

//func newErrorResult(err error) *result {
//	r := new(result)
//	r.err = err
//	return r
//}
