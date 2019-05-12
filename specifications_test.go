package strftime

import (
	"reflect"
	"testing"
	"time"
)

func Test_getSecondarySpecs(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want map[rune]string
	}{
		{
			name: "February 28th 2019",
			args: args{
				t: time.Date(2019, time.February, 28, 0, 0, 0, 0, time.FixedZone("EDT", -4*60*60)),
			},
			want: map[int32]string{
				103: "19",
				106: "059",
				107: "0",
				115: "1551326400",
				117: "4",
				119: "4",
				67:  "20",
				71:  "2019",
				85:  "08",
				86:  "9",
				87:  "08",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSecondarySpecs(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSecondarySpecs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yearWeek(t *testing.T) {
	type args struct {
		t     time.Time
		start time.Weekday
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "February 28th 2019 - Sunday",
			args: args{
				t:     time.Date(2019, time.February, 28, 0, 0, 0, 0, time.Local),
				start: time.Sunday,
			},
			want: 8,
		},
		{
			name: "February 28th 2019 - Monday",
			args: args{
				t:     time.Date(2019, time.February, 28, 0, 0, 0, 0, time.Local),
				start: time.Monday,
			},
			want: 8,
		},
		{
			name: "February 28th 2019 - Wednesday",
			args: args{
				t:     time.Date(2019, time.February, 28, 0, 0, 0, 0, time.Local),
				start: time.Wednesday,
			},
			want: 9,
		},
		{
			name: "January 2nd 2019 - Sunday",
			args: args{
				t:     time.Date(2019, time.January, 2, 0, 0, 0, 0, time.Local),
				start: time.Sunday,
			},
			want: 0,
		},
		{
			name: "January 1st 2018 - Monday",
			args: args{
				t:     time.Date(2018, time.January, 1, 0, 0, 0, 0, time.Local),
				start: time.Monday,
			},
			want: 1,
		},
		{
			name: "March 2nd 2016 - Sunday",
			args: args{
				t:     time.Date(2016, time.March, 2, 0, 0, 0, 0, time.Local),
				start: time.Sunday,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yearWeek(tt.args.t, tt.args.start); got != tt.want {
				t.Errorf("yearWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
