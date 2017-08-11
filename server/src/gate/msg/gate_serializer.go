package msg

import (
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/log4g"
	"encoding/binary"
	"reflect"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/carsonsx/net4g/util"
	"container/list"
	"bytes"
)



func NewGateSerializer() net4g.Serializer {
	s := new(GateSerializer)
	s.EmptySerializer = net4g.NewEmptySerializer()
	return s
}

type GateSerializer struct {
	*net4g.EmptySerializer
}

var (
	WriteNo uint32
)

type SGHeader struct {
	Flag string
	Type int8
	Ver uint8
	GateId  string
	No uint32
	Zip bool
	Encrypt bool
	Length uint16
	Id uint8
}

func (h *SGHeader) toBytes() []byte {

	var header [32]byte

	//消息标示
	flagByte := []byte(h.Flag)
	header[0] = flagByte[0]
	header[1] = flagByte[1]
	//消息类型
	header[2] = byte(h.Type)
	//协议版本
	header[3] = byte(h.Ver)
	//闸机ID
	copy(header[4:16], []byte(h.GateId))
	//消息流水
	binary.BigEndian.PutUint32(header[16:20], h.No)
	//消息属性
	var zip byte
	if h.Zip {
		zip = 1
	} else {
		zip = 0
	}
	var encrypt byte
	if h.Encrypt {
		encrypt = 1 << 1
	} else {
		encrypt = 0
	}
	header[20] = zip | encrypt
	//消息总长度
	binary.BigEndian.PutUint16(header[21:23], h.Length)
	//消息id
	header[23] = byte(h.Id)

	return header[:]
}

func NewSGHeader(gateId string) *SGHeader {
	var h SGHeader
	//消息标示
	h.Flag = "SG"
	//消息类型
	h.Type = 1
	//协议版本
	h.Ver = 1
	//闸机ID
	h.GateId = gateId
	//消息属性
	h.Zip = false
	h.Encrypt = false
	return &h
}

func (h *SGHeader) fromBytes(header []byte) {

	//消息标示
	h.Flag = string(header[0:2])
	//消息类型
	h.Type = int8(header[2])
	//协议版本
	h.Ver = uint8(header[3])
	//闸机ID
	h.GateId = string(bytes.Trim(header[4:16], "\x00"))
	//消息流水
	h.No = binary.BigEndian.Uint32(header[16:20])
	//消息属性
	h.Zip = (header[20] & 1) == 1
	h.Encrypt = ((header[20] >> 1) & 1) == 1
	//消息总长度
	h.Length = binary.BigEndian.Uint16(header[21:23])
	//消息id
	h.Id = uint8(header[23])

}

type gateList struct {
	elements *list.List
}
func (gl *gateList) init() {
	if gl.elements == nil {
		gl.elements = list.New()
	}
}

func (gl *gateList) addByte(e byte) {
	gl.init()
	gl.elements.PushBack(e)
	if e == 0x10 {
		gl.elements.PushBack(e)
	}
}

func (gl *gateList) addBytes(els []byte) {
	for _, e := range els {
		gl.addByte(e)
	}
}

func (gl *gateList) addUint16(e uint16) {
	//low := (byte)(e & 0xff)
	//gl.addByte(low)
	//high := (byte)((e>>8) & 0xff)
	//gl.addByte(high)
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, e)
	gl.addBytes(b)
}

func (gl *gateList) toBytes() []byte {
	bytes := make([]byte, gl.elements.Len())
	idx := 0
	for e := gl.elements.Front(); e != nil; e = e.Next() {
		bytes[idx] = e.Value.(byte)
		idx++
	}
	return bytes
}

func (s *GateSerializer) Serialize(v, h interface{}) (data []byte, err error) {

	//if !s.Registered {
	//	panic("not registered any id or key")
	//}

	var id int

	if rp, ok := v.(*net4g.RawPack); ok {
		s.PreprocessRawPack(rp)
		id = rp.Id
		data = rp.Data
	} else {
		t := reflect.TypeOf(v)
		if t == nil || t.Kind() != reflect.Ptr {
			panic("value type must be a pointer")
		}

		var ok bool
		id, ok = s.SerializerIdsOfType[t]
		if !ok {
			err = errors.New(fmt.Sprintf("%v is not registed by any id", t))
			log4g.Error(err)
			return
		}

		data, err = json.Marshal(v)
		if err != nil {
			log4g.Error(err)
			return
		}

		if log4g.IsTraceEnabled() {
			log4g.Trace("serializing %v - %v", t, v)
			log4g.Trace("serialized %v - %s", t, string(data))
		}

	}

	sgHeader := h.(*SGHeader)
	sgHeader.Id = uint8(id)
	sgHeader.No = WriteNo
	WriteNo++
	sgHeader.Length = uint16(len(data) + 32)
	headBytes := sgHeader.toBytes()
	data = util.CombineBytes(headBytes, data)

	log4g.Debug("msg len: %d", len(data))

	//for testing
	//data = append(data, 0x10, 0x10)

	//封装成帧
	x := new(gateList)
	//长度
	x.addUint16(uint16(len(data) + 2)) //消息总长度加2个字节的校验
	log4g.Debug("frame len: %d", len(data) +2)
	//数据
	x.addBytes(data)
	//校验
	var sum uint16
	for i := range data {
		sum += uint16(data[i])
	}
	x.addUint16(sum)

	data = x.toBytes()

	return
}

func (s *GateSerializer) Deserialize(raw []byte) (v, h interface{}, rp *net4g.RawPack, err error) {

	//if !s.Registered {
	//	panic("not registered any id or key")
	//}

	//帧解析

	//去0x10
	met10 := false
	var realRaw []byte
	for _, b := range raw {
		if b == 0x10 {
			met10 = !met10
			if met10 {
				realRaw = append(realRaw, b)
			}
		} else {
			realRaw = append(realRaw, b)
		}
	}

	rawLen := len(realRaw)
	//长度
	dataLen := binary.BigEndian.Uint16(realRaw) - 2
	//数据
	data := realRaw[2:rawLen-2]
	//校验
	sum := binary.BigEndian.Uint16(realRaw[rawLen-2:])

	if dataLen != uint16(len(data)) {
		err = errors.New(fmt.Sprintf("wrong message length: except %d, actual %d", dataLen, len(data)))
		log4g.Error(err)
		return
	}

	var realSum uint16
	for i := range data {
		realSum += uint16(data[i])
	}

	if sum != realSum {
		err = errors.New("verify message data failed")
		log4g.Error(err)
		return
	}

	var sgHeader SGHeader
	sgHeader.fromBytes(data)
	h = &sgHeader

	rp = new(net4g.RawPack)
	rp.Id = int(sgHeader.Id)
	rp.Data = data[32:]

	var ok bool
	if rp.Type, ok = s.DeserializerTypesOfId[rp.Id]; ok {
		if log4g.IsTraceEnabled() {
			log4g.Trace("deserialize %v - %s", rp.Type, string(rp.Data))
		}
		v = reflect.New(rp.Type.Elem()).Interface()
		if len(rp.Data) == 0 {
			return
		}
		err = json.Unmarshal(rp.Data, v)
		if err != nil {
			log4g.Error(err)
			log4g.Error(string(rp.Data))
		}
	}

	return
}
