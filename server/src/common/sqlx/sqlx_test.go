package sqlx

import (
	"database/sql"
	"fmt"
	"github.com/carsonsx/log4g"
	"reflect"
	"testing"
	"time"
)

func TestNull(t *testing.T) {

}

func errorTest() (err error) {

	var aa []*User

	t := reflect.TypeOf(aa)
	fmt.Println(t.Kind())
	fmt.Println(t.Elem().Kind())
	fmt.Println(t.Elem().Elem().Kind())
	fmt.Println(reflect.TypeOf(aa).Elem().Elem())

	return
}

func TestStruct(t *testing.T) {
	user := new(User)
	NewDao().Query("select * from user").One(user)
	log4g.Debug(log4g.JsonFunc(user))
}

func TestAll(t *testing.T) {
	var users []*User
	NewDao().Query("select * from user").All(&users)
	log4g.Debug(log4g.JsonFunc(users))
}

type User struct {
	Id         int            `json:"id"`
	UserName   sql.NullString `db:"user_name"`
	Password   sql.NullString
	Sex        NullInt64
	InsertTime *time.Time `db:"insert_time"`
}
