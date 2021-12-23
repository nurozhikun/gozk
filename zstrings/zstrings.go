package zstrings

import "bytes"

func HumpToUnderline(s string) string {
	buf := bytes.Buffer{}
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			if buf.Len() > 0 {
				buf.WriteByte('_')
			}
			buf.WriteByte(byte(v) + 32)
		} else {
			buf.WriteRune(v)
		}
	}
	return buf.String()
}
