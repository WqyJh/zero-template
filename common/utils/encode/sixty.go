package encode

import (
	"encoding/binary"
	"strings"

	"github.com/google/uuid"
)

const (
	chars62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	count   = uint64(len(chars62))
)

func Reverse(s string) string {
	bytes := []byte(s)

	for from, to := 0, len(bytes)-1; from < to; from, to = from+1, to-1 {
		bytes[from], bytes[to] = bytes[to], bytes[from]
	}

	return string(bytes)
}

// 10进制转62进制
func Encode10To62(num uint64) string {
	var (
		remainder uint64
		sb        strings.Builder
	)
	for num > 0 {
		remainder = num % count
		num = num / count
		sb.WriteByte(chars62[remainder])
	}
	return Reverse(sb.String())
}

// 62进制转10进制
func Decode62To10(str string) uint64 {
	var result uint64
	var mul uint64 = 1
	for i := len(str) - 1; i >= 0; i-- {
		index := strings.IndexByte(chars62, str[i])
		result += mul * uint64(index)
		mul *= count
	}
	return result
}

func Shortuuid() string {
	uu := uuid.New()
	l1 := binary.BigEndian.Uint64(uu[:8])
	l2 := binary.BigEndian.Uint64(uu[8:])
	s1 := Encode10To62(l1)
	s2 := Encode10To62(l2)
	return s1 + s2
}
