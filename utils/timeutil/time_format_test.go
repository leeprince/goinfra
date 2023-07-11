package timeutil

import (
	"fmt"
	"github.com/leeprince/goinfra/consts"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/20 11:50
 * @Desc:
 */

func TestToUnix1(t *testing.T) {
	var v string
	v = "2023-05-18 16:17:18"
	fmt.Println(">>>")
	fmt.Println(ToLocalUnix(v, consts.TimeLayoutV1))
	
	v = "2023-05-18 16:17"
	fmt.Println(">>>")
	fmt.Println(ToLocalUnix(v, consts.TimeLayoutV2))
	
	v = "2023-05-18 16"
	fmt.Println(">>>")
	fmt.Println(ToLocalUnix(v, consts.TimeLayoutV3))
	
	v = "2023-05-18"
	fmt.Println(">>>")
	fmt.Println(ToLocalUnix(v, consts.TimeLayoutV4))
}

func TestToLocalUnix(t *testing.T) {
	type args struct {
		timeStr    string
		timeLayout string
	}
	tests := []struct {
		name         string
		args         args
		wantTimeUnix int64
		wantErr      bool
	}{
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18 16:17:18",
				timeLayout: consts.TimeLayoutV1,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18 16:17",
				timeLayout: consts.TimeLayoutV2,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18 16",
				timeLayout: consts.TimeLayoutV3,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18",
				timeLayout: consts.TimeLayoutV4,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-2023",
				timeLayout: consts.TimeLayoutV5,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-23",
				timeLayout: consts.TimeLayoutV6,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023-05-18",
				timeLayout: consts.TimeLayoutV4,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "05-18-23",
				timeLayout: consts.TimeLayoutV6,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
		{
			name: "",
			args: args{
				timeStr:    "2023/05/18",
				timeLayout: consts.TimeLayoutV41,
			},
			wantTimeUnix: 0,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTimeUnix, err := ToLocalUnix(tt.args.timeStr, tt.args.timeLayout)
			fmt.Println(gotTimeUnix, err)
		})
	}
}

func TestFormat(t *testing.T) {
	var ts string
	ts = Year()
	fmt.Println("Year", ts)
	
	ts = Month()
	fmt.Println("Month", ts)
	
	ts = Data()
	fmt.Println("Data", ts)
	
	ts = DataTime()
	fmt.Println("DataTime", ts)
	
	ts = DataTimeF(time.Now())
	fmt.Println("DataTimeF", ts)
	
	ts = DataTimeNanosecond()
	fmt.Println("DataTimeNanosecond", ts)
	
	ts = DataTimeMicrosecond()
	fmt.Println("DataTimeMicrosecond", ts)
	
	ts = DataTimeMillisecond()
	fmt.Println("DataTimeMillisecond", ts)
	
}
