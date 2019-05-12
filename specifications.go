package strftime

import (
	"fmt"
	"strconv"
	"time"
)

var convSpecs = map[rune]string{
	'a': "Mon",
	'A': "Monday",
	'b': "Jan",
	'B': "January",
	'c': "Mon Jan 02 15:04:05 2006",
	'd': "02",
	'D': "01/02/06",
	'e': "2",
	//'%E': "", modifier
	'F': "2006-01-02",
	'h': "Jan",
	'H': "15",
	'I': "03",
	'l': "3",
	'm': "01",
	'M': "04",
	'n': "\n",
	//'%O': "", modifier
	'p': "PM",
	'P': "pm",
	'r': "03:04:05 PM",
	'R': "15:04",
	'S': "05",
	't': "\t",
	'T': "15:04:05",
	'x': "01/02/06",
	'X': "15:04:05",
	'y': "06",
	'Y': "2006",
	'z': "-0700",
	'Z': "MST",
	'+': "Mon Jan 02 15:04:05 MST 2006",
	'%': "%",
}

func getSecondarySpecs(t time.Time) map[rune]string {
	specs := make(map[rune]string)

	isoYear, isoWeek := t.ISOWeek()

	specs['C'] = strconv.Itoa(t.Year() / 100)
	specs['G'] = strconv.Itoa(isoYear)
	specs['g'] = strconv.Itoa(isoYear)[2:4]
	specs['j'] = fmt.Sprintf("%03d", t.YearDay())
	specs['k'] = strconv.Itoa(t.Hour())
	specs['s'] = strconv.Itoa(int(t.Unix()))
	specs['u'] = mondayWeekday[t.Weekday()]
	specs['U'] = fmt.Sprintf("%02d", yearWeek(t, time.Sunday))
	specs['V'] = strconv.Itoa(isoWeek)
	specs['w'] = strconv.Itoa(int(t.Weekday()))
	specs['W'] = fmt.Sprintf("%02d", yearWeek(t, time.Monday))
	return specs
}

var mondayWeekday = map[time.Weekday]string{
	time.Monday:    "1",
	time.Tuesday:   "2",
	time.Wednesday: "3",
	time.Thursday:  "4",
	time.Friday:    "5",
	time.Saturday:  "6",
	time.Sunday:    "7",
}

func yearWeek(t time.Time, start time.Weekday) int {
	offset := 0
	jan1 := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
	if start >= jan1.Weekday() {
		offset = 1
	}
	return int(float64(t.YearDay()))/7 + offset
}
