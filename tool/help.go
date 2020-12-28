package tool

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/spf13/cast"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
)

const (
	regular = `^1([38][0-9]|14[57]|5[^4])\d{8}$`
)

func ValidateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func FormatData(code interface{}, data interface{}, msg interface{}) (res map[string]interface{}) {
	if msg == nil {
		msg = "ok"
	}
	res = make(map[string]interface{})
	res["status"] = code
	res["data"] = data
	res["msg"] = msg

	return
}
func FormatDataVos(code interface{}, msg interface{}, outMap map[string]interface{}) (res map[string]interface{}) {
	if msg == nil {
		msg = "ok"
	}
	res = make(map[string]interface{})
	for k, v := range outMap {
		res[k] = v
	}

	res["status"] = code
	res["msg"] = msg

	return
}
func Unique() string {
	data := GetRandomString(10)
	return cast.ToString(Md5(cast.ToString(data)))
}
func GetRandomString(l int) string {
	str := "abcdefjhjkmnpqrstuvwxyz12345678"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)

}
func UniqueNew() string {
	data := GenerateRangeNum(1, 10)
	return cast.ToString(Md5(cast.ToString(data)))
}
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	//fmt.Printf("rand is %v\n", randNum)
	return randNum
}
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructToMapViaJson(obj interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &m)
	return
}
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Alihost() (host string) {
	host = "http://wmfrontdata.oss-cn-beijing.aliyuncs.com/"
	return
}

// is_numeric()
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.Trim(str, " \\t\\n\\r\\v\\f")
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9,Point,Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
