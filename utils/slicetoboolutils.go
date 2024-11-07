package utils

import (
	"database/sql"
	"time"
)

func BoolToInt8Slice(b bool) []int8 {
	if b {
		return []int8{1}
	}
	return []int8{0}
}
func BoolToUint8Slice(b bool) []uint8 {
	if b {
		return []uint8{1}
	}
	return []uint8{0}
}
func Uint8slicetoboolslices(int8Slice []uint8) []bool {
	boolSlice := make([]bool, len(int8Slice))
	for i, v := range int8Slice {
		boolSlice[i] = v != 0
	}
	return boolSlice
}

func Uint8slicetoboolslice(int8Slice []uint8) bool {
	boolSlice := make([]bool, len(int8Slice))
	for i, v := range int8Slice {
		boolSlice[i] = v != 0
		return v != 0
	}
	return false
}
func InterfaceToBool(any interface{}) bool {
	ret, ok := any.(bool)
	if !ok {
		return false
	}
	return ret
}
func NullTimeToPointTime(nulltime sql.NullTime) *time.Time {
	if nulltime.Valid {
		return &nulltime.Time
	} else {
		return nil
	}
}
func NullStringToPointString(nullstr sql.NullString) *string {
	if nullstr.Valid {
		return &nullstr.String
	} else {
		return nil
	}
}

func PointStringToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{String: "", Valid: false}
	} else {
		return sql.NullString{String: *s, Valid: true}
	}
}
