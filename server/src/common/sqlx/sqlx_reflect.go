package sqlx

import (
	"errors"
	"github.com/carsonsx/log4g"
	"reflect"
	"strings"
)

const DB_TAG = "db"

var ErrFieldNotFound = errors.New("not found field")

func Indirect(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// if this is a pointer and it's nil, allocate a new value and set it
func allocateValue(rv *reflect.Value) {
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		alloc := reflect.New(Indirect(rv.Type()))
		rv.Set(alloc)
	}
	if rv.Kind() == reflect.Map && rv.IsNil() {
		rv.Set(reflect.MakeMap(rv.Type()))
	}
}

func setTagValue(dest interface{}, tag string, v interface{}) error {
	rv := reflect.ValueOf(dest)
	rv = reflect.Indirect(rv)
	f, err := getFieldByTag(rv, tag)
	if err != nil {
		log4g.Error(ErrFieldNotFound)
		return ErrFieldNotFound
	}
	f.Set(reflect.ValueOf(v))
	return nil
}

func getValuesByTags(rv reflect.Value, tags []string, values []interface{}, ptr bool) {
	rv = reflect.Indirect(rv)
	for i, tag := range tags {
		f, err := getFieldByTag(rv, tag)
		if err != nil {
			continue
		}
		if ptr {
			values[i] = f.Addr().Interface()
		} else {
			values[i] = f.Interface()
		}
	}
}

func getFieldsByTags(rv reflect.Value, tags []string) {
	var values []interface{}
	for _, tag := range tags {
		f, err := getFieldByTag(rv, tag)
		if err != nil {
			continue
		}
		if f.IsValid() {
			values = append(values, f)
		}
	}
}

func getFieldByTag(rv reflect.Value, tag string) (f reflect.Value, err error) {
	rv = reflect.Indirect(rv)
	//getStructField(i, tag)
	cacheStructFields(rv.Type())
	key := getFieldKey(rv.Type(), tag)
	f = rv.Field(indexesByTag[key])
	return
}

var cachedTypes = make(map[reflect.Type]bool)
var fieldsByTag = make(map[string]*reflect.StructField)
var namesByTag = make(map[string]string)
var indexesByTag = make(map[string]int)

func getStructField(i interface{}, tag string) *reflect.StructField {
	t := reflect.TypeOf(i)
	k := getFieldKey(t, tag)
	sf := fieldsByTag[k]
	if sf != nil {
		return sf
	}
	cacheStructFields(t)
	sf = fieldsByTag[getFieldKey(t, tag)]
	if sf == nil {
		fieldsByTag[k] = nil
	}
	return sf
}

func getFieldKey(t reflect.Type, tag string) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String() + "_" + tag
}

func cacheStructFields(t reflect.Type) {
	if cachedTypes[t] {
		return
	}
	cachedTypes[t] = true
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	num := t.NumField()
	for i := 0; i < num; i++ {
		field := t.Field(i)
		tag := field.Tag.Get(DB_TAG)
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}
		key := getFieldKey(t, tag)
		fieldsByTag[key] = &field
		namesByTag[key] = field.Name
		indexesByTag[key] = i
	}
	//log4g.Debug(fieldsByTag)
	//log4g.Debug(namesByTag)
	//log4g.Debug(indexesByTag)
}
