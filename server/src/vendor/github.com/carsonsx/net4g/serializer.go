package net4g

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g/util"
	"github.com/golang/protobuf/proto"
	"reflect"
)

func Serialize(serializer Serializer, v interface{}, h ...interface{}) (data []byte, err error) {
	var header interface{}
	if len(h) > 0 {
		header = h[0]
	}
	data, err = serializer.Serialize(v, header)
	return
}

func NewRawPackById(id int, data ...[]byte) *RawPack {
	rp := new(RawPack)
	rp.Id = id
	if len(data) > 0 {
		rp.Data = data[0]
	}
	return rp
}

func NewRawPackByKey(key string, data ...[]byte) *RawPack {
	rp := new(RawPack)
	rp.Key = key
	if len(data) > 0 {
		rp.Data = data[0]
	}
	return rp
}

func NewRawPackByType(t reflect.Type, data ...[]byte) *RawPack {
	rp := new(RawPack)
	rp.Type = t
	if len(data) > 0 {
		rp.Data = data[0]
	}
	return rp
}

type RawPack struct {
	Id   int
	Key  string
	Type reflect.Type
	Data []byte
}

type Serializer interface {
	SetIdStartingValue(id int)
	SerializeId(v interface{}, id_at_most_one ...int) (id int, err error)
	SerializeKey(v interface{}, key_at_most_one ...string) (key string, err error)
	DeserializeId(v interface{}, id_at_most_one ...int) (id int, err error)
	DeserializeKey(v interface{}, key_at_most_one ...string) (key string, err error)
	Serialize(v, h interface{}) (data []byte, err error)
	Deserialize(raw []byte) (v, h interface{}, rp *RawPack, err error)
	RangeId(f func(id int, t reflect.Type))
	RangeKey(f func(key string, t reflect.Type))
}

func NewEmptySerializer() *EmptySerializer {
	s := new(EmptySerializer)
	s.SerializerIdsOfType = make(map[reflect.Type]int)
	s.SerializerTypesOfId = make(map[int]reflect.Type)
	s.SerializerKeysOfType = make(map[reflect.Type]string)
	s.SerializerTypesOfKey = make(map[string]reflect.Type)
	s.DeserializerIdsOfType = make(map[reflect.Type]int)
	s.DeserializerTypesOfId = make(map[int]reflect.Type)
	s.DeserializerKeysOfType = make(map[reflect.Type]string)
	s.DeserializerTypesOfKey = make(map[string]reflect.Type)
	s.id = 1
	return s
}

type EmptySerializer struct {
	SerializerIdsOfType    map[reflect.Type]int
	SerializerTypesOfId    map[int]reflect.Type
	SerializerKeysOfType   map[reflect.Type]string
	SerializerTypesOfKey   map[string]reflect.Type
	DeserializerIdsOfType  map[reflect.Type]int
	DeserializerTypesOfId  map[int]reflect.Type
	DeserializerKeysOfType map[reflect.Type]string
	DeserializerTypesOfKey map[string]reflect.Type
	ids                    []int
	keys                   []string
	id                     int
	byId                   bool
}

func (s *EmptySerializer) SetIdStartingValue(id int) {
	s.id = id
}

func (s *EmptySerializer) SerializeId(v interface{}, id_at_most_one ...int) (id int, err error) {
	return s.registerId(s.SerializerIdsOfType, s.SerializerTypesOfId, v, id_at_most_one...)
}

func (s *EmptySerializer) DeserializeId(v interface{}, id_at_most_one ...int) (id int, err error) {
	return s.registerId(s.DeserializerIdsOfType, s.DeserializerTypesOfId, v, id_at_most_one...)
}

func (s *EmptySerializer) registerId(ids map[reflect.Type]int, types map[int]reflect.Type, v interface{}, id_at_most_one ...int) (id int, err error) {

	t := reflect.TypeOf(v)
	if t == nil || t.Kind() != reflect.Ptr {
		panic("interface type must be a pointer")
	}

	if len(s.SerializerKeysOfType) > 0 {
		panic("can not registered id and key in one serializer")
	}

	if len(id_at_most_one) > 1 {
		panic("only mapping one type with one id")
	}

	if _id, ok := ids[t]; ok {
		text := fmt.Sprintf("%s has been registered by %d", t.String(), _id)
		log4g.Error(text)
		return 0, errors.New(text)
	}

	if len(id_at_most_one) == 1 {
		id = id_at_most_one[0]
	} else {
		id = s.id
	}

	ids[t] = id
	types[id] = t
	s.ids = append(s.ids, id)

	s.byId = true

	s.id++

	return
}

func (s *EmptySerializer) SerializeKey(v interface{}, key_at_most_one ...string) (key string, err error) {
	return s.registerKey(s.SerializerKeysOfType, s.SerializerTypesOfKey, v, key_at_most_one...)
}

func (s *EmptySerializer) DeserializeKey(v interface{}, key_at_most_one ...string) (key string, err error) {
	return s.registerKey(s.DeserializerKeysOfType, s.DeserializerTypesOfKey, v, key_at_most_one...)
}

func (s *EmptySerializer) registerKey(keys map[reflect.Type]string, types map[string]reflect.Type, v interface{}, key_at_most_one ...string) (key string, err error) {

	t := reflect.TypeOf(v)
	if t == nil || t.Kind() != reflect.Ptr {
		panic("interface type must be a pointer")
	}

	if len(s.SerializerIdsOfType) > 0 {
		panic("can not registered key and id in one serializer")
	}

	if len(key_at_most_one) > 1 {
		panic("only mapping one type with one key")
	}

	if _key, ok := keys[t]; ok {
		text := fmt.Sprintf("%s has been registered by %s", t.Elem().Name(), _key)
		log4g.Error(text)
		err = errors.New(text)
		return
	}

	if len(key_at_most_one) == 1 {
		key = key_at_most_one[0]
	} else {
		key = t.String()
	}

	keys[t] = key
	types[key] = t
	s.keys = append(s.keys, key)

	s.byId = false
	log4g.Info("%v register by key '%s'\n", t, key)

	return
}

func (s *EmptySerializer) RangeId(f func(id int, t reflect.Type)) {
	for _, id := range s.ids {
		f(id, s.SerializerTypesOfId[id])
	}
}

func (s *EmptySerializer) RangeKey(f func(key string, t reflect.Type)) {
	for _, key := range s.keys {
		f(key, s.SerializerTypesOfKey[key])
	}
}

func (s *EmptySerializer) PreprocessRawPack(rp *RawPack) {
	if rp.Type == nil {
		if rp.Id > 0 {
			rp.Type = s.SerializerTypesOfId[rp.Id]
		} else if rp.Key != "" {
			rp.Type = s.SerializerTypesOfKey[rp.Key]
		}
	}
	if rp.Id == 0 && rp.Key == "" && rp.Type != nil {
		if len(s.SerializerIdsOfType) > 0 {
			rp.Id = s.SerializerIdsOfType[rp.Type]
		}
		if len(s.SerializerKeysOfType) > 0 {
			rp.Key = s.SerializerKeysOfType[rp.Type]
		}
	}
}

func NewByteSerializer() Serializer {
	s := new(ByteSerializer)
	s.EmptySerializer = NewEmptySerializer()
	return s
}

type ByteSerializer struct {
	*EmptySerializer
}

func (s *ByteSerializer) Serialize(v, h interface{}) (data []byte, err error) {
	if rp, ok := v.(*RawPack); ok {
		s.PreprocessRawPack(rp)
		log4g.Trace("serialized - %v", rp)
		data = util.AddIntHeader(rp.Data, NetConfig.IdSize, uint64(rp.Id), NetConfig.LittleEndian)
	} else {
		data = v.([]byte)
	}
	if NetConfig.HeaderSize > NetConfig.IdSize {
		if h == nil {
			log4g.Panic("header cannot be nil")
		}
		header := h.([]byte)
		if NetConfig.HeaderSize != len(header)+NetConfig.IdSize {
			log4g.Panic("invalid header length: excepted %d, actual %d", NetConfig.HeaderSize-NetConfig.IdSize, len(header))
		}
		_data := data
		data = make([]byte, len(header)+len(_data))
		copy(data, header)
		copy(data[len(header):], _data)
	}
	return
}

func (s *ByteSerializer) Deserialize(raw []byte) (v, h interface{}, rp *RawPack, err error) {
	rp = new(RawPack)
	rp.Id = int(util.GetIntHeader(raw, NetConfig.IdSize, NetConfig.LittleEndian))
	rp.Data = raw[NetConfig.IdSize:]
	v = rp.Data
	log4g.Trace("deserialize - %v", *rp)
	return
}

func NewStringSerializer() Serializer {
	s := new(StringSerializer)
	s.EmptySerializer = NewEmptySerializer()
	return s
}

type StringSerializer struct {
	*EmptySerializer
}

func (s *StringSerializer) Serialize(v, h interface{}) (raw []byte, err error) {
	return []byte(v.(string)), nil
}

func (s *StringSerializer) Deserialize(raw []byte) (v, h interface{}, rp *RawPack, err error) {
	rp = new(RawPack)
	return string(raw), nil, rp, nil
}

func NewJsonSerializer() Serializer {
	s := new(JsonSerializer)
	s.EmptySerializer = NewEmptySerializer()
	return s
}

type JsonSerializer struct {
	*EmptySerializer
}

func (s *JsonSerializer) Serialize(v, h interface{}) (data []byte, err error) {

	if rp, ok := v.(*RawPack); ok {
		s.PreprocessRawPack(rp)
		if s.byId {
			if id, ok := s.SerializerIdsOfType[rp.Type]; ok {
				data = util.AddIntHeader(rp.Data, NetConfig.IdSize, uint64(id), NetConfig.LittleEndian)
			} else {
				err = errors.New(fmt.Sprintf("%v is not registed by any id", rp.Type))
				log4g.Error(err)
			}
		} else {
			if key, ok := s.SerializerKeysOfType[rp.Type]; ok {
				m := map[string]json.RawMessage{key: rp.Data}
				data, err = json.Marshal(m)
				if err != nil {
					log4g.Error(err)
					return
				}
				if log4g.IsTraceEnabled() {
					log4g.Trace("serialized %v - %s", rp.Type, string(data))
				}
			} else {
				log4g.Panic("%v is not registered by any key", rp.Type)
			}
		}
	} else {
		t := reflect.TypeOf(v)
		if t == nil || t.Kind() != reflect.Ptr {
			panic("value type must be a pointer")
		}
		if s.byId {
			if id, ok := s.SerializerIdsOfType[t]; ok {
				data, err = json.Marshal(v)
				if err != nil {
					log4g.Error(err)
					return
				}
				if log4g.IsTraceEnabled() {
					log4g.Trace("serializing %v - %v", t, v)
					log4g.Trace("serialized %v - %s", t, string(data))
				}
				data = util.AddIntHeader(data, NetConfig.IdSize, uint64(id), NetConfig.LittleEndian)
			} else {
				err = errors.New(fmt.Sprintf("%v is not registed by any id", t))
				log4g.Error(err)
			}
		} else {
			if key, ok := s.SerializerKeysOfType[t]; ok {
				m := map[string]interface{}{key: v}
				data, err = json.Marshal(m)
				if err != nil {
					log4g.Error(err)
					return
				}
				if log4g.IsTraceEnabled() {
					log4g.Trace("serialized %v - %s", t, string(data))
				}
			} else {
				log4g.Panic("%v is not registered by any key", t)
			}
		}
	}

	return
}

func (s *JsonSerializer) Deserialize(raw []byte) (v, h interface{}, rp *RawPack, err error) {

	rp = new(RawPack)

	if s.byId {
		if len(raw) < NetConfig.IdSize {
			text := fmt.Sprintf("message length [%d] is short than id size [%d]", len(raw), NetConfig.IdSize)
			err = errors.New(text)
			log4g.Error(err)
			return
		}

		rp.Id = int(util.GetIntHeader(raw, NetConfig.IdSize, NetConfig.LittleEndian))
		rp.Data = raw[NetConfig.IdSize:]
		var ok bool
		if rp.Type, ok = s.DeserializerTypesOfId[rp.Id]; ok {
			v = reflect.New(rp.Type.Elem()).Interface()
			if len(rp.Data) == 0 {
				return
			}
			err = json.Unmarshal(rp.Data, v)
			if err != nil {
				log4g.Error(err)
			} else {
				log4g.Trace("deserialize %v - %s", rp.Type, string(rp.Data))
			}
		} else {
			err = errors.New(fmt.Sprintf("id[%d] is not registered by any type", rp.Id))
			log4g.Error(err)
		}
	} else {
		var m_raw map[string]json.RawMessage
		err = json.Unmarshal(raw, &m_raw)
		if err != nil {
			log4g.Error(err)
			return
		}
		if len(m_raw) == 0 {
			text := fmt.Sprintf("invalid json: %v", string(raw))
			err = errors.New(text)
			log4g.Error(err)
			return
		}
		for rp.Key, rp.Data = range m_raw {
			var ok bool
			if rp.Type, ok = s.DeserializerTypesOfKey[rp.Key]; ok {
				v = reflect.New(rp.Type.Elem()).Interface()
				if len(rp.Data) == 0 {
					continue
				}
				err = json.Unmarshal(rp.Data, v)
				if err != nil {
					log4g.Error(err)
				} else {
					log4g.Trace("deserialize %v - %s", rp.Type, string(raw))
					break
				}
			} else {
				err = errors.New(fmt.Sprintf("key '%s' is not registered by any type", rp.Key))
				log4g.Error(err)
			}
		}
	}
	return
}

func NewProtobufSerializer() Serializer {
	s := new(ProtobufSerializer)
	s.EmptySerializer = NewEmptySerializer()
	return s
}

type ProtobufSerializer struct {
	*EmptySerializer
}

func (s *ProtobufSerializer) Serialize(v, h interface{}) (data []byte, err error) {

	if rp, ok := v.(*RawPack); ok {
		s.PreprocessRawPack(rp)
		data = util.AddIntHeader(rp.Data, NetConfig.IdSize, uint64(rp.Id), NetConfig.LittleEndian)
		if log4g.IsDebugEnabled() {
			bytes, _ := json.Marshal(v)
			log4g.Trace("serialize %d - %v", rp.Id, string(bytes))
		}
	} else {
		t := reflect.TypeOf(v)
		if t == nil || t.Kind() != reflect.Ptr {
			panic("value type must be a pointer")
		}
		if id, ok := s.SerializerIdsOfType[t]; ok {
			data, err = proto.Marshal(v.(proto.Message))
			if err != nil {
				log4g.Error(err)
				return
			}
			data = util.AddIntHeader(data, NetConfig.IdSize, uint64(id), NetConfig.LittleEndian)
			if log4g.IsDebugEnabled() {
				bytes, _ := json.Marshal(v)
				log4g.Trace("serialize %v - %v", t, string(bytes))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v is not registed by any id", t))
		}
	}

	return
}

func (s *ProtobufSerializer) Deserialize(raw []byte) (v, h interface{}, rp *RawPack, err error) {

	if len(raw) < NetConfig.IdSize {
		text := fmt.Sprintf("message length [%d] is short than id size [%d]", len(raw), NetConfig.IdSize)
		err = errors.New(text)
		log4g.Error(err)
		return
	}

	rp = new(RawPack)
	rp.Id = int(util.GetIntHeader(raw, NetConfig.IdSize, NetConfig.LittleEndian))
	rp.Data = raw[NetConfig.IdSize:]
	var ok bool
	if rp.Type, ok = s.DeserializerTypesOfId[rp.Id]; ok {
		v = reflect.New(rp.Type.Elem()).Interface()
		if len(rp.Data) == 0 {
			return
		}
		err = proto.UnmarshalMerge(rp.Data, v.(proto.Message))
		if err != nil {
			log4g.Error(err)
		} else {
			if log4g.IsDebugEnabled() {
				bytes, _ := json.Marshal(v)
				log4g.Trace("deserialize %v - %v", rp.Type, string(bytes))
			}
		}
	}
	return
}
