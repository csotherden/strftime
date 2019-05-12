package strftime

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	type args struct {
		format string
		t      time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "The abbreviated name of the day of the week",
			args: args{
				format: "%a",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "Sat",
		},
		{
			name: "The full name of the day of the week",
			args: args{
				format: "%A",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "Saturday",
		},
		{
			name: "The abbreviated month name",
			args: args{
				format: "%b",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "May",
		},
		{
			name: "The full month name",
			args: args{
				format: "%B",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "May",
		},
		{
			name: "The preferred date and time representation",
			args: args{
				format: "%c",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "Sat May 11 23:45:24 2019",
		},
		{
			name: "The day of the month as a decimal number (range 01 to 31)",
			args: args{
				format: "%d",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "11",
		},
		{
			name: "Equivalent to %m/%d/%y",
			args: args{
				format: "%D",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "05/11/19",
		},
		{
			name: "Like %d, the day of the month as a decimal number, but a leading zero is replaced by a space",
			args: args{
				format: "%e",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "11",
		},
		{
			name: "Equivalent to %Y-%m-%d (the ISO 8601 date format)",
			args: args{
				format: "%F",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "2019-05-11",
		},
		{
			name: "Equivalent to %b",
			args: args{
				format: "%h",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "May",
		},
		{
			name: "The hour as a decimal number using a 24-hour clock (range 00 to 23)",
			args: args{
				format: "%H",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "23",
		},
		{
			name: "The hour as a decimal number using a 12-hour clock (range 01 to 12)",
			args: args{
				format: "%I",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "11",
		},
		{
			name: "The hour (12-hour clock) as a decimal number (range 1 to 12); single digits are preceded by a blank.",
			args: args{
				format: "%l",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "11",
		},
		{
			name: "The month as a decimal number (range 01 to 12)",
			args: args{
				format: "%m",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "05",
		},
		{
			name: "The minute as a decimal number (range 00 to 59)",
			args: args{
				format: "%M",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "45",
		},
		{
			name: "A newline character",
			args: args{
				format: "%n",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "\n",
		},
		{
			name: "Either 'AM' or 'PM' according to the given time value. Noon is treated as 'PM' and midnight as 'AM'",
			args: args{
				format: "%p",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "PM",
		},
		{
			name: "Like %p but in lowercase: 'am' or 'pm'",
			args: args{
				format: "%P",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "pm",
		},
		{
			name: "The time in a.m. or p.m. notation",
			args: args{
				format: "%r",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "11:45:24 PM",
		},
		{
			name: "The time in 24-hour notation (%H:%M)",
			args: args{
				format: "%R",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "23:45",
		},
		{
			name: "The second as a decimal number (range 00 to 60)",
			args: args{
				format: "%S",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "24",
		},
		{
			name: "A tab character",
			args: args{
				format: "%t",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "\t",
		},
		{
			name: "The time in 24-hour notation (%H:%M:%S)",
			args: args{
				format: "%T",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "23:45:24",
		},
		{
			name: "Equivalent to %D",
			args: args{
				format: "%x",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "05/11/19",
		},
		{
			name: "Equivalent to %T",
			args: args{
				format: "%X",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "23:45:24",
		},
		{
			name: "The year as a decimal number without a century (range 00 to 99)",
			args: args{
				format: "%y",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "19",
		},
		{
			name: "The year as a decimal number including the century",
			args: args{
				format: "%Y",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "2019",
		},
		{
			name: "The +hhmm or -hhmm numeric timezone (that is, the hour and minute offset from UTC)",
			args: args{
				format: "%z",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "-0400",
		},
		{
			name: "The timezone name or abbreviation",
			args: args{
				format: "%Z",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "EDT",
		},
		{
			name: "The date and time in date(1) format",
			args: args{
				format: "%+",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "Sat May 11 23:45:24 EDT 2019",
		},
		{
			name: "A literal '%' character",
			args: args{
				format: "%%",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "%",
		},
		{
			name: "The century number (year/100) as a 2-digit integer",
			args: args{
				format: "%C",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "20",
		},
		{
			name: "The ISO 8601 week-based year with century as a decimal number",
			args: args{
				format: "%G",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "2019",
		},
		{
			name: "Like %G, but without century, that is, with a 2-digit year (00–99)",
			args: args{
				format: "%g",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "19",
		},
		{
			name: "The day of the year as a decimal number (range 001 to 366)",
			args: args{
				format: "%j",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "131",
		},
		{
			name: "The hour (24-hour clock) as a decimal number (range 0 to 23)",
			args: args{
				format: "%k",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "23",
		},
		{
			name: "The number of seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC)",
			args: args{
				format: "%s",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "1557632724",
		},
		{
			name: "The day of the week as a decimal, range 1 to 7, Monday being 1",
			args: args{
				format: "%u",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "6",
		},
		{
			name: "The week number of the current year as a decimal number, range 00 to 53, starting with the first Sunday as the first day of week 01",
			args: args{
				format: "%U",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "18",
		},
		{
			name: "The ISO 8601 week number",
			args: args{
				format: "%V",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "19",
		},
		{
			name: "The day of the week as a decimal, range 0 to 6, Sunday being 0",
			args: args{
				format: "%w",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "6",
		},
		{
			name: "The week number of the current year as a decimal number, range 00 to 53, starting with the first Monday as the first day of week 01",
			args: args{
				format: "%W",
				t:      time.Date(2019, time.May, 11, 23, 45, 24, 0, time.Local),
			},
			want: "18",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.format, tt.args.t); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
