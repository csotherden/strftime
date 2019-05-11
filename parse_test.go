package strftime

import "testing"

func Test_parseFormat(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "The abbreviated name of the day of the week",
			args: args{
				f: "%a",
			},
			want: "Mon",
		},
		{
			name: "The full name of the day of the week",
			args: args{
				f: "%A",
			},
			want: "Monday",
		},
		{
			name: "The abbreviated month name",
			args: args{
				f: "%b",
			},
			want: "Jan",
		},
		{
			name: "The full month name",
			args: args{
				f: "%B",
			},
			want: "January",
		},
		{
			name: "The preferred date and time representation",
			args: args{
				f: "%c",
			},
			want: "Mon Jan 02 15:04:05 2006",
		},
		{
			name: "The day of the month as a decimal number (range 01 to 31)",
			args: args{
				f: "%d",
			},
			want: "02",
		},
		{
			name: "Equivalent to %m/%d/%y",
			args: args{
				f: "%D",
			},
			want: "01/02/06",
		},
		{
			name: "Like %d, the day of the month as a decimal number, but a leading zero is replaced by a space",
			args: args{
				f: "%e",
			},
			want: "2",
		},
		{
			name: "Equivalent to %Y-%m-%d (the ISO 8601 date format)",
			args: args{
				f: "%F",
			},
			want: "2006-01-02",
		},
		{
			name: "Equivalent to %b",
			args: args{
				f: "%h",
			},
			want: "Jan",
		},
		{
			name: "The hour as a decimal number using a 24-hour clock (range 00 to 23)",
			args: args{
				f: "%H",
			},
			want: "15",
		},
		{
			name: "The hour as a decimal number using a 12-hour clock (range 01 to 12)",
			args: args{
				f: "%I",
			},
			want: "03",
		},
		{
			name: "The hour (12-hour clock) as a decimal number (range 1 to 12); single digits are preceded by a blank.",
			args: args{
				f: "%l",
			},
			want: "3",
		},
		{
			name: "The month as a decimal number (range 01 to 12)",
			args: args{
				f: "%m",
			},
			want: "01",
		},
		{
			name: "The minute as a decimal number (range 00 to 59)",
			args: args{
				f: "%M",
			},
			want: "04",
		},
		{
			name: "A newline character",
			args: args{
				f: "%n",
			},
			want: "\n",
		},
		{
			name: "Either 'AM' or 'PM' according to the given time value. Noon is treated as 'PM' and midnight as 'AM'",
			args: args{
				f: "%p",
			},
			want: "PM",
		},
		{
			name: "Like %p but in lowercase: 'am' or 'pm'",
			args: args{
				f: "%P",
			},
			want: "pm",
		},
		{
			name: "The time in a.m. or p.m. notation",
			args: args{
				f: "%r",
			},
			want: "03:04:05 PM",
		},
		{
			name: "The time in 24-hour notation (%H:%M)",
			args: args{
				f: "%R",
			},
			want: "15:04",
		},
		{
			name: "The second as a decimal number (range 00 to 60)",
			args: args{
				f: "%S",
			},
			want: "05",
		},
		{
			name: "A tab character",
			args: args{
				f: "%t",
			},
			want: "\t",
		},
		{
			name: "The time in 24-hour notation (%H:%M:%S)",
			args: args{
				f: "%T",
			},
			want: "15:04:05",
		},
		{
			name: "Equivalent to %D",
			args: args{
				f: "%x",
			},
			want: "01/02/06",
		},
		{
			name: "Equivalent to %T",
			args: args{
				f: "%X",
			},
			want: "15:04:05",
		},
		{
			name: "The year as a decimal number without a century (range 00 to 99)",
			args: args{
				f: "%y",
			},
			want: "06",
		},
		{
			name: "The year as a decimal number including the century",
			args: args{
				f: "%Y",
			},
			want: "2006",
		},
		{
			name: "The +hhmm or -hhmm numeric timezone (that is, the hour and minute offset from UTC)",
			args: args{
				f: "%z",
			},
			want: "-0700",
		},
		{
			name: "The timezone name or abbreviation",
			args: args{
				f: "%Z",
			},
			want: "MST",
		},
		{
			name: "The date and time in date(1) format",
			args: args{
				f: "%+",
			},
			want: "Mon Jan 02 15:04:05 MST 2006",
		},
		{
			name: "A literal '%' character",
			args: args{
				f: "%%",
			},
			want: "%",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFormat(tt.args.f); got != tt.want {
				t.Errorf("parseFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
