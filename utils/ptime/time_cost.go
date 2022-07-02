package ptime

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/20 下午11:18
 * @Desc:
 */
import (
    "github.com/leeprince/goinfra/plog"
    "time"
)

type TimeCost struct {
    start  time.Time
    middle time.Time
    end    time.Time
    runing bool
}

func NewTimeCost(msg string) *TimeCost {
    now := time.Now()
    timeCost := &TimeCost{start: now, middle: now, runing: true}
    
    plog.Info(msg + " > now time:", now.UnixNano())
    return timeCost
}

// 统计从middle开始到当前时间
func (s *TimeCost) Duration(msg string) int64 {
    if !s.runing {
        plog.Info(msg +" > Duration err:!s.runing")
        return 0
    }
    
    tmp := time.Now()
    duration := tmp.UnixNano() - s.middle.UnixNano()
    s.middle = tmp
    
    plog.Info(msg + " > cost time(ns):", duration)
    return duration
}

// 统计从NewTimeCost开始到当前时间
func (s *TimeCost) Stop(msg string) int64 {
    if !s.runing {
        plog.Info(msg + " > Stop err:!s.runing")
        return 0
    }
    
    s.runing = false
    tmp := time.Now()
    
    duration := tmp.UnixNano() - s.start.UnixNano()
    s.end = tmp
    
    plog.Info(msg+" > end cost time(ns):", duration)
    return duration
}
