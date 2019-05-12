package strftime

import (
	"bytes"
	"time"
)

func parseFormat(f string, specs map[rune]string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(f)*2))

	for i := 0; i < len(f); i++ {
		if i+1 == len(f) {
			buf.WriteByte(f[i])
			break
		}

		if f[i] == '%' {
			if replace, found := specs[rune(f[i+1])]; found {
				buf.WriteString(replace)
				i++
			} else {
				buf.WriteByte(f[i])
			}
		} else {
			buf.WriteByte(f[i])
		}
	}
	return buf.String()
}

func secondaryFormat(t time.Time, formattedTime string) string {
	specs := getSecondarySpecs(t)
	return parseFormat(formattedTime, specs)
}
