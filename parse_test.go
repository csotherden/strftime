package strftime

import (
	"reflect"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	type args struct {
		format     string
		timeString string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "ANSIC",
			args: args{
				format:     "%a %b  %e %H:%M:%S %Y",
				timeString: "Wed Feb  4 21:00:57 2009",
			},
			want:    timeMustParse(time.ANSIC, "Wed Feb  4 21:00:57 2009"),
			wantErr: false,
		},
		{
			name: "UnixDate",
			args: args{
				format:     "%a %b  %e %H:%M:%S %Z %Y",
				timeString: "Thu Feb  4 21:00:57 PST 2010",
			},
			want:    timeMustParse(time.UnixDate, "Thu Feb  4 21:00:57 PST 2010"),
			wantErr: false,
		},
		{
			name: "RubyDate",
			args: args{
				format:     "%a %b %e %H:%M:%S %z %Y",
				timeString: "Thu Feb 04 21:00:57 -0800 2010",
			},
			want:    timeMustParse(time.RubyDate, "Thu Feb 04 21:00:57 -0800 2010"),
			wantErr: false,
		},
		{
			name: "RFC850",
			args: args{
				format:     "%A, %e-%b-%y %H:%M:%S %Z",
				timeString: "Thursday, 04-Feb-10 21:00:57 PST",
			},
			want:    timeMustParse(time.RFC850, "Thursday, 04-Feb-10 21:00:57 PST"),
			wantErr: false,
		},
		{
			name: "RFC1123",
			args: args{
				format:     "%a, %e %b %Y %H:%M:%S %Z",
				timeString: "Thu, 04 Feb 2010 21:00:57 PST",
			},
			want:    timeMustParse(time.RFC1123, "Thu, 04 Feb 2010 21:00:57 PST"),
			wantErr: false,
		},
		{
			name: "RFC1123Z",
			args: args{
				format:     "%a, %e %b %Y %H:%M:%S %z",
				timeString: "Thu, 04 Feb 2010 21:00:57 -0800",
			},
			want:    timeMustParse(time.RFC1123Z, "Thu, 04 Feb 2010 21:00:57 -0800"),
			wantErr: false,
		},
		/*{
			name:"RFC3339",
			args:args{
				format: "%Y-%m-%dT%H:%M:%S",
				timeString: "2010-02-04T21:00:57-08:00",
			},
			want:timeMustParse(time.ANSIC, "2010-02-04T21:00:57-08:00"),
			wantErr: false,
		},
		{
			name:"custom",
			args:args{
				format: "%a %b  %e %H:%M:%S %Y",
				timeString: "2006-01-02 15:04:05-07",
			},
			want:timeMustParse(time.ANSIC, "2006-01-02 15:04:05-07"),
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.format, tt.args.timeString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func timeMustParse(layout, timeString string) time.Time {
	t, err := time.Parse(layout, timeString)
	if err != nil {
		panic(err)
	}

	return t
}
