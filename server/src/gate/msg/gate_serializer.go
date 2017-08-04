package msg

import (
	"reflect"
	"github.com/carsonsx/net4g"
)

func NewGateSerializer() net4g.Serializer {
	s := new(GateSerializer)
	s.EmptySerializer = net4g.NewEmptySerializer()
	return s
}

type GateSerializer struct {
	*net4g.EmptySerializer
}

func (s *GateSerializer) Serialize(v interface{}) (data []byte, err error) {

	if !s.IsRegistered() {
		panic("not registered any id or key")
	}

	t := reflect.TypeOf(v)
	if t == nil || t.Kind() != reflect.Ptr {
		panic("value type must be a pointer")
	}

	return
}

func (s *GateSerializer) Deserialize(raw []byte) (v interface{}, rp *net4g.RawPack, err error) {

	if !s.IsRegistered() {
		panic("not registered any id or key")
	}


	return
}
