package strftime

import (
	"testing"
	"time"
)

func Test_parseFormat(t *testing.T) {
	type args struct {
		f string
		s map[rune]string
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
				s: convSpecs,
			},
			want: "Mon",
		},
		{
			name: "The full name of the day of the week",
			args: args{
				f: "%A",
				s: convSpecs,
			},
			want: "Monday",
		},
		{
			name: "The abbreviated month name",
			args: args{
				f: "%b",
				s: convSpecs,
			},
			want: "Jan",
		},
		{
			name: "The full month name",
			args: args{
				f: "%B",
				s: convSpecs,
			},
			want: "January",
		},
		{
			name: "The preferred date and time representation",
			args: args{
				f: "%c",
				s: convSpecs,
			},
			want: "Mon Jan 02 15:04:05 2006",
		},
		{
			name: "The day of the month as a decimal number (range 01 to 31)",
			args: args{
				f: "%d",
				s: convSpecs,
			},
			want: "02",
		},
		{
			name: "Equivalent to %m/%d/%y",
			args: args{
				f: "%D",
				s: convSpecs,
			},
			want: "01/02/06",
		},
		{
			name: "Like %d, the day of the month as a decimal number, but a leading zero is replaced by a space",
			args: args{
				f: "%e",
				s: convSpecs,
			},
			want: "2",
		},
		{
			name: "Equivalent to %Y-%m-%d (the ISO 8601 date format)",
			args: args{
				f: "%F",
				s: convSpecs,
			},
			want: "2006-01-02",
		},
		{
			name: "Equivalent to %b",
			args: args{
				f: "%h",
				s: convSpecs,
			},
			want: "Jan",
		},
		{
			name: "The hour as a decimal number using a 24-hour clock (range 00 to 23)",
			args: args{
				f: "%H",
				s: convSpecs,
			},
			want: "15",
		},
		{
			name: "The hour as a decimal number using a 12-hour clock (range 01 to 12)",
			args: args{
				f: "%I",
				s: convSpecs,
			},
			want: "03",
		},
		{
			name: "The hour (12-hour clock) as a decimal number (range 1 to 12); single digits are preceded by a blank.",
			args: args{
				f: "%l",
				s: convSpecs,
			},
			want: "3",
		},
		{
			name: "The month as a decimal number (range 01 to 12)",
			args: args{
				f: "%m",
				s: convSpecs,
			},
			want: "01",
		},
		{
			name: "The minute as a decimal number (range 00 to 59)",
			args: args{
				f: "%M",
				s: convSpecs,
			},
			want: "04",
		},
		{
			name: "A newline character",
			args: args{
				f: "%n",
				s: convSpecs,
			},
			want: "\n",
		},
		{
			name: "Either 'AM' or 'PM' according to the given time value. Noon is treated as 'PM' and midnight as 'AM'",
			args: args{
				f: "%p",
				s: convSpecs,
			},
			want: "PM",
		},
		{
			name: "Like %p but in lowercase: 'am' or 'pm'",
			args: args{
				f: "%P",
				s: convSpecs,
			},
			want: "pm",
		},
		{
			name: "The time in a.m. or p.m. notation",
			args: args{
				f: "%r",
				s: convSpecs,
			},
			want: "03:04:05 PM",
		},
		{
			name: "The time in 24-hour notation (%H:%M)",
			args: args{
				f: "%R",
				s: convSpecs,
			},
			want: "15:04",
		},
		{
			name: "The second as a decimal number (range 00 to 60)",
			args: args{
				f: "%S",
				s: convSpecs,
			},
			want: "05",
		},
		{
			name: "A tab character",
			args: args{
				f: "%t",
				s: convSpecs,
			},
			want: "\t",
		},
		{
			name: "The time in 24-hour notation (%H:%M:%S)",
			args: args{
				f: "%T",
				s: convSpecs,
			},
			want: "15:04:05",
		},
		{
			name: "Equivalent to %D",
			args: args{
				f: "%x",
				s: convSpecs,
			},
			want: "01/02/06",
		},
		{
			name: "Equivalent to %T",
			args: args{
				f: "%X",
				s: convSpecs,
			},
			want: "15:04:05",
		},
		{
			name: "The year as a decimal number without a century (range 00 to 99)",
			args: args{
				f: "%y",
				s: convSpecs,
			},
			want: "06",
		},
		{
			name: "The year as a decimal number including the century",
			args: args{
				f: "%Y",
				s: convSpecs,
			},
			want: "2006",
		},
		{
			name: "The +hhmm or -hhmm numeric timezone (that is, the hour and minute offset from UTC)",
			args: args{
				f: "%z",
				s: convSpecs,
			},
			want: "-0700",
		},
		{
			name: "The timezone name or abbreviation",
			args: args{
				f: "%Z",
				s: convSpecs,
			},
			want: "MST",
		},
		{
			name: "The date and time in date(1) format",
			args: args{
				f: "%+",
				s: convSpecs,
			},
			want: "Mon Jan 02 15:04:05 MST 2006",
		},
		{
			name: "A literal '%' character",
			args: args{
				f: "%%",
				s: convSpecs,
			},
			want: "%",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFormat(tt.args.f, tt.args.s); got != tt.want {
				t.Errorf("parseFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secondaryFormat(t *testing.T) {
	type args struct {
		t             time.Time
		formattedTime string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondaryFormat(tt.args.t, tt.args.formattedTime); got != tt.want {
				t.Errorf("secondaryFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
