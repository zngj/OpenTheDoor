package sqlx

import (
	"database/sql"
)


type NullString struct {
	sql.NullString
}

func (nx *NullString) String(v ...string) string {
	if len(v) > 0 {
		nx.NullString.String = v[0]
		nx.Valid = true
		return v[0]
	}
	return nx.NullString.String
}

type NullInt64 struct {
	sql.NullInt64
}

func (nx *NullInt64) Int8(v ...int8) int8 {
	if len(v) > 0 {
		nx.NullInt64.Int64 = int64(v[0])
		nx.Valid = true
		return v[0]
	}
	return int8(nx.NullInt64.Int64)
}

func (nx *NullInt64) Uint8() uint8 {
	return uint8(nx.NullInt64.Int64)
}

func (nx *NullInt64) Int16() int16 {
	return int16(nx.NullInt64.Int64)
}

func (nx *NullInt64) Uint16() uint16 {
	return uint16(nx.NullInt64.Int64)
}

func (nx *NullInt64) Int() int {
	return int(nx.NullInt64.Int64)
}

func (nx *NullInt64) Uint() uint {
	return uint(nx.NullInt64.Int64)
}

func (nx *NullInt64) Int32() int32 {
	return int32(nx.NullInt64.Int64)
}

func (nx *NullInt64) Uint32() uint32 {
	return uint32(nx.NullInt64.Int64)
}

func (nx *NullInt64) Int64() int64 {
	return nx.NullInt64.Int64
}

func (nx *NullInt64) Uint64() uint64 {
	return uint64(nx.NullInt64.Int64)
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (nx *NullFloat64) Float32(v ...float32) float32 {
	if len(v) > 0 {
		nx.NullFloat64.Float64 = float64(v[0])
		nx.Valid = true
		return v[0]
	}
	return float32(nx.NullFloat64.Float64)
}

func (nx *NullFloat64) Float64(v ...float64) float64 {
	if len(v) > 0 {
		nx.NullFloat64.Float64 = v[0]
		nx.Valid = true
		return v[0]
	}
	return nx.NullFloat64.Float64
}