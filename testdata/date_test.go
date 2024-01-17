package testdata

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/12 15:12
 * @Desc:
 */
type Date struct {
	OrderTime time.Time `gorm:"column:order_time;type:datetime;not null;DEFAULT '0000-00-00 00:00:00'" json:"order_time"` // 订单创建时间 格式：YYYY-MM-DD hh:mm:ss
}

func TestDateTimeUnmarshal(t *testing.T) {
	//jsonData := `{"order_time": "2024-01-02 13:35:11"}`
	jsonData := `{"order_time": "2024-01-02T13:35:11"}`

	var data Date
	// 报错：Error: parsing time "2024-01-02T13:35:11" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "Z07:00"
	// 解决办法：TestDateTimeUnmarshal();TestDateTimeUnmarshalSuccess();
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Order Time:", data.OrderTime)
}

type Date1 struct {
	OrderTime       time.Time `gorm:"column:order_time;type:datetime;not null;DEFAULT '0000-00-00 00:00:00'" json:"-"` // 订单创建时间 格式：YYYY-MM-DD hh:mm:ss
	OrderTimeString string    `json:"order_time"`                                                                      // 订单创建时间 格式：YYYY-MM-DD hh:mm:ss
}

func TestDateTimeUnmarshal1(t *testing.T) {
	//jsonData := `{"order_time": "2024-01-02 13:35:11"}`
	jsonData := `{"order_time": "2024-01-02T13:35:11"}`

	var data Date1
	// 报错：Error: parsing time "2024-01-02T13:35:11" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "Z07:00"
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("data:", data)

	loc := time.Local
	data.OrderTime, err = time.ParseInLocation("2006-01-02 15:04:05", data.OrderTimeString, loc)
	if err != nil {
		fmt.Println("Error1:", err)
		return
	}
	fmt.Println("data:", data)

}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1] // 去除双引号
	layout := "2006-01-02T15:04:05"
	ct.Time, err = time.Parse(layout, s)
	return
}

type DateSuccess struct {
	OrderTime CustomTime `json:"order_time"`
}

func TestDateTimeUnmarshalSuccess(t *testing.T) {
	jsonData := `{"order_time": "2024-01-02T13:35:11"}`
	var data DateSuccess
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Order Time:", data.OrderTime.Time)
}
