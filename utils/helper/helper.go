package helper

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func GetCurrentPage(page, pageSize int) (pageNew, pageSizeNew int) {
	if page == 0 {
		pageNew = 1
	}
	if pageSize == 0 {
		pageSizeNew = 10
	}
	return
}

func StringToInt64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 64)
	return
}
func StringToUint(s string) (i uint) {
	t, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	i = uint(t)
	return
}
func Int64ToString(i int64) (s string) {
	s = strconv.FormatInt(i, 10)

	return
}

func InterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func InterfaceToInt64(inVal interface{}) int64 {
	if inVal == nil {
		return 0
	}

	refValue := reflect.ValueOf(inVal)
	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
	}
	refType := reflect.TypeOf(inVal)
	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
	}

	switch refType.Kind() {
	case reflect.Bool:
		if refValue.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return int64(refValue.Uint())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return refValue.Int()
	case reflect.Float32, reflect.Float64:
		return int64(refValue.Float())
	case reflect.Complex64, reflect.Complex128:
		retValue, _ := strconv.ParseFloat(strconv.FormatComplex(refValue.Complex(), 'f', -1, 128), 64)
		return int64(retValue)
	}

	// 转换为字符串，在其中找数字
	re := regexp.MustCompile("-?[0-9]+")
	valueString := InterfaceToString(inVal)
	numberList := re.FindAllString(valueString, -1)
	if len(numberList) > 0 {
		rVal, err := strconv.ParseInt(numberList[0], 10, 64)
		if err == nil {
			return rVal
		}
	}

	return 0
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func GetRequestPath(path, prefix string) (uri string, id int64) {
	uri = strings.TrimPrefix(path, prefix)
	re := regexp.MustCompile(`^(.*)/(\d+)$`)
	matches := re.FindStringSubmatch(uri)
	if len(matches) == 3 {
		uri = matches[1]
		id = StringToInt64(matches[2])
	}

	return
}

// 获取操作系统
func GetPlatform(userAgent string) string {
	ua := strings.ToLower(userAgent)

	// 移动端
	if strings.Contains(ua, "android") {
		return "Android"
	} else if strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") || strings.Contains(ua, "ipod") {
		return "iOS"
	}

	// 桌面端
	if strings.Contains(ua, "windows") {
		return "Windows"
	} else if strings.Contains(ua, "macintosh") || strings.Contains(ua, "mac os") {
		return "MacOS"
	} else if strings.Contains(ua, "linux") {
		return "Linux"
	}

	return "Unknown"
}

// 获取浏览器类型
func GetBrowser(userAgent string) string {
	ua := strings.ToLower(userAgent)

	if strings.Contains(ua, "chrome") && !strings.Contains(ua, "edg") {
		return "Google Chrome"
	} else if strings.Contains(ua, "edg") {
		return "Microsoft Edge"
	} else if strings.Contains(ua, "firefox") {
		return "Mozilla Firefox"
	} else if strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome") {
		return "Apple Safari"
	} else if strings.Contains(ua, "opr") || strings.Contains(ua, "opera") {
		return "Opera"
	} else if strings.Contains(ua, "msie") || strings.Contains(ua, "trident") {
		return "Internet Explorer"
	}

	return "Unknown"
}

func ConvertToRestfulURL(url string) string {
	re := regexp.MustCompile(`(^.+?/[^/]+)/\d+$`)
	return re.ReplaceAllString(url, `$1/:id`)
}
