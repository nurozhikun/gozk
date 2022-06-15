package zutils

import "strconv"

func StringFromInterface(i interface{}) string {
	s, _ := InterfaceToString(i)
	return s
}

func InterfaceToString(v interface{}) (string, bool) {
	switch v.(type) {
	case string:
		return v.(string), true
	default:
		if i, ok := InterfaceToInt(v); ok {
			return strconv.FormatInt(i, 10), true
		}
		if u, ok := InterfaceToUint(v); ok {
			return strconv.FormatUint(u, 10), true
		}
		return "", false
	}
}

func InterfaceToInt(v interface{}) (int64, bool) {
	switch v.(type) {
	case int8:
		return int64(v.(int8)), true
	case int16:
		return int64(v.(int16)), true
	case int32:
		return int64(v.(int32)), true
	case int64:
		return v.(int64), true
	case int:
		return int64(v.(int)), true
	case string:
		if i, err := strconv.ParseInt(v.(string), 10, 64); err != nil {
			return i, false
		} else {
			return i, true
		}
	default:
		return 0, false
	}
}

func InterfaceToUint(v interface{}) (uint64, bool) {
	switch v.(type) {
	case uint8:
		return uint64(v.(int8)), true
	case uint16:
		return uint64(v.(int16)), true
	case uint32:
		return uint64(v.(int32)), true
	case uint64:
		return v.(uint64), true
	case uint:
		return uint64(v.(int)), true
	case string:
		if i, err := strconv.ParseUint(v.(string), 10, 64); err != nil {
			return i, false
		} else {
			return i, true
		}
	default:
		return 0, false
	}
}
