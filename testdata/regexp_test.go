package testdata

import (
	"fmt"
	"github.com/leeprince/goinfra/perror"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/testdata/constants"
	"github.com/leeprince/goinfra/testdata/message"
	"github.com/spf13/cast"
	"regexp"
	"strings"
	"sync"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/19 18:20
 * @Desc:
 */

func TestMatchRequired(t *testing.T) {
	type args struct {
		logID    string
		str      string
		seatType message.SeatType
	}
	var tests = []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				logID:    "",
				str:      "连座1,1F,选座不满足,可出其他坐席",
				seatType: message.SeatTypeHard,
			},
		},
		{
			name: "",
			args: args{
				str:      "1B,选座不满足,可出其他坐席",
				seatType: message.SeatTypeHard,
			},
		},
		{
			name: "",
			args: args{
				str:      "不接受其它坐席，不接受无座",
				seatType: message.SeatTypeHard,
			},
		},
		{
			name: "",
			args: args{
				str:      "邻座，1张D座，1张F座",
				seatType: message.SeatTypeHard,
			},
		},
		{
			name: "",
			args: args{
				str:      "靠窗，条件不满足，直接失败",
				seatType: message.SeatTypeHard,
			},
		},
		{
			name: "",
			args: args{
				str:      "在线选座1F,条件不满足直接失败",
				seatType: message.SeatTypeHard,
			},
		},
		{
			name: "",
			args: args{
				str:      "2张下铺，条件不满足，直接失败",
				seatType: message.SeatTypeHardSleeper,
			},
		},
		{
			name: "",
			args: args{
				str:      "1张下铺，条件不满足，直接失败",
				seatType: message.SeatTypeHardSleeper,
			},
		},
		{
			name: "",
			args: args{
				str:      "1张中铺，条件不满足，直接失败",
				seatType: message.SeatTypeHardSleeper,
			},
		},
		{
			name: "",
			args: args{
				str:      "1张中铺，条件不满足，可出票",
				seatType: message.SeatTypeHard,
			},
		},
		{
			name: "",
			args: args{
				str:      "1张下铺，条件不满足，直接失败  >>> 三个人的硬卧票，只要求必须有一个下铺",
				seatType: message.SeatTypeHardSleeper,
			},
		},
		{
			name: "",
			args: args{
				str:      "中下铺1人,条件不满足直接失败;",
				seatType: message.SeatTypeHardSleeper,
			},
		},
		{
			name: "",
			args: args{
				str:      "必须：10车厢",
				seatType: message.SeatTypeSecond,
			},
		},
		{
			name: "",
			args: args{
				str:      "必须：01车厢",
				seatType: message.SeatTypeSecond,
			},
		},
		{
			name: "",
			args: args{
				str:      "1张A座，1张B座，1张C座",
				seatType: message.SeatTypeSecond,
			},
		},
		{
			name: "",
			args: args{
				str:      "1张A座，2张B座，3张C座",
				seatType: message.SeatTypeSecond,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			personNumber := int32(2)
			seatPositionType, matchPositionSuffix, positionList, err := MatchRequired(tt.args.logID, personNumber, tt.args.str, tt.args.seatType)
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			fmt.Printf("seatPositionType:%+v \n", seatPositionType)
			fmt.Printf("matchPositionSuffix:%+v \n", matchPositionSuffix)
			fmt.Printf("positionList:%+v \n", positionList)
		})
	}
}

var (
	onceMatchRequired sync.Once
	
	// prince@TODO: 暂不支持 2023/8/12 18:23
	/*
		06车厢1人,条件不满足直接失败;在线选座1D,条件不满足可出票;
		06车厢2人,条件不满足直接失败;在线选座1D1F,条件不满足直接失败;条件不满足，直接失败
	*/
	reCarriageSuffix *regexp.Regexp // 匹配 {xx数字}车厢{N}人
	
	/*
		连座1,1F,选座不满足,可出其他坐席
	*/
	reSeatPositionTypeSpecifySeat *regexp.Regexp // 匹配 {N人}{座位后缀英文字母}
	
	/*
		1张A座，1张B座，1张C座
	*/
	reSpecifySeatRequired *regexp.Regexp // // 匹配 {N}张{英文字母}座，{N}张{英文字母}座，
	
	/*
		1张A座，1张B座，条件不满足，可出票
	*/
	reSeatPositionTypeSpecifySeatSuffix *regexp.Regexp // 匹配 {N}张{英文字母}座
	
	/*
		2张下铺，条件不满足，直接失败
	*/
	reSeatPositionTypeSleeper *regexp.Regexp // 匹配 {N}张{上/中/下}铺
	
	/*
		中上铺1人,条件不满足直接失败;
		下铺1人,条件不满足直接失败;
		上铺1人,条件不满足直接失败;
	*/
	reSeatPositionTypeSleeperSpecify *regexp.Regexp // 匹配 {上/中/下}铺{N}人
	
	/*
		必须：10车厢
	*/
	reSeatPositionTypeCarriage *regexp.Regexp // 匹配 必须：{xx数字}车厢
	
	/*
		06车厢1人,条件不满足直接失败
		06车厢2人,条件不满足直接失败
	*/
	reSeatPositionTypeCarriagePerson *regexp.Regexp // 匹配 {xx数字}车厢{N}人

)

// MatchRequired 注意：匹配有先后顺序，规则是复杂要求在前面
func MatchRequired(logID string, personNumber int32, str string, seatType message.SeatType) (seatPositionType message.SeatPositionType, matchPositionSuffix []string, positionList []message.Position, err error) {
	onceMatchRequired.Do(func() {
		reCarriageSuffix = regexp.MustCompile(`(\d+)车厢(\d+)人`)
		reSeatPositionTypeSpecifySeat = regexp.MustCompile(`(\d+)([A-Z])`)
		reSpecifySeatRequired = regexp.MustCompile(`(\d+)张([^\s，]+)座`)
		reSeatPositionTypeSpecifySeatSuffix = regexp.MustCompile(`(\d+)张([A-Z])座`)
		reSeatPositionTypeSleeper = regexp.MustCompile(`(\d+)张([上中下])铺`)
		reSeatPositionTypeSleeperSpecify = regexp.MustCompile(`([上中下])铺(\d+)人`)
		reSeatPositionTypeCarriage = regexp.MustCompile(`必须：(\d+)车厢`)
	})
	var match []string
	var matchList [][]string
	plogEntry := plog.WithField("method", "MatchRequired").
		WithField("str", str).
		WithField("seatType", seatType)
	plogEntry.Info("request")
	
	match = reSeatPositionTypeCarriage.FindStringSubmatch(str)
	if len(match) > 0 {
		plogEntry.WithField("match", match).Info(logID, "reSeatPositionTypeCarriage-匹配 必须：{xx数字}车厢")
		
		seatPositionType = message.SeatPositionTypeCarriage
		for i := 0; i < int(personNumber); i++ {
			position := message.Position{
				Carriage:   match[1],
				SeatNumber: "",
				Sleeper:    "",
			}
			positionList = append(positionList, position)
		}
		
		return
	}
	
	match = reSeatPositionTypeSpecifySeat.FindStringSubmatch(str)
	if len(match) > 0 {
		plogEntry.WithField("match", match).Info(logID, "reSeatPositionTypeSpecifySeat-匹配 {数字}{英文字母}")
		
		seatPositionType = message.SeatPositionTypeSpecifySeat
		for i := 0; i < cast.ToInt(match[1]); i++ {
			matchPositionSuffix = append(matchPositionSuffix, match[2])
		}
		return
	}
	
	matchList = reSpecifySeatRequired.FindAllStringSubmatch(str, -1)
	if len(matchList) > 0 {
		plogEntry.WithField("match", match).Info(logID, "reSeatPositionTypeSpecifySeatSuffix-匹配 {N}张{英文字母}座，{N}张{英文字母}座，")
		
		seatPositionType = message.SeatPositionTypeSpecifySeatRequired
		for _, match := range matchList {
			count := cast.ToInt(match[1])
			seat := match[2]
			for i := 0; i < count; i++ {
				matchPositionSuffix = append(matchPositionSuffix, seat)
			}
		}
		return
	}
	
	match = reSeatPositionTypeSpecifySeatSuffix.FindStringSubmatch(str)
	if len(match) > 0 {
		plogEntry.WithField("match", match).Info(logID, "reSpecifySeatRequired-匹配 {N}张{英文字母}座")
		
		seatPositionType = message.SeatPositionTypeSpecifySeat
		for i := 0; i < cast.ToInt(match[1]); i++ {
			matchPositionSuffix = append(matchPositionSuffix, match[2])
		}
		return
	}
	
	match = reSeatPositionTypeSleeper.FindStringSubmatch(str)
	if len(match) > 0 {
		plogEntry.WithField("match", match).Info(logID, "reSeatPositionTypeSleeper-匹配 {N}张{上/中/下}铺")
		if !IsSleeperSeatType(seatType) {
			err = perror.NewBizErr(constants.WinErrSeatMatchSleeper.Key(), constants.WinErrSeatMatchSleeper.Value())
			plogEntry.WithError(err).Error(logID, "reSeatPositionTypeSleeper-WinErrSeatSleeperMatch err")
			return
		}
		
		seatPositionType = message.SeatPositionTypeSleeper
		if strings.Contains(str, "，同包厢") {
			seatPositionType = message.SeatPositionTypeSameCarriageSleeper
		}
		
		for i := 0; i < cast.ToInt(match[1]); i++ {
			sleeper, ok := message.SleeperMap[match[2]]
			if !ok {
				err = perror.NewBizErr(constants.WinErrSeatSleeperMatch.Key(), constants.WinErrSeatSleeperMatch.Value())
				plogEntry.WithError(err).Error(logID, "reSeatPositionTypeSleeper-WinErrSeatSleeperMatch err")
				return
			}
			position := message.Position{
				Carriage:   "",
				SeatNumber: "",
				Sleeper:    sleeper,
			}
			positionList = append(positionList, position)
		}
		return
	}
	
	match = reSeatPositionTypeSleeperSpecify.FindStringSubmatch(str)
	if len(match) > 0 {
		plogEntry.WithField("match", match).Info(logID, "reSeatPositionTypeSleeperSpecify-匹配 {上/中/下}铺{N}人")
		if !IsSleeperSeatType(seatType) {
			err = perror.NewBizErr(constants.WinErrSeatMatchSleeper.Key(), constants.WinErrSeatMatchSleeper.Value())
			plogEntry.WithError(err).Error(logID, "reSeatPositionTypeSleeper-WinErrSeatSleeperMatch err")
			return
		}
		
		seatPositionType = message.SeatPositionTypeSleeper
		for i := 0; i < cast.ToInt(match[2]); i++ {
			sleeper, ok := message.SleeperMap[match[1]]
			if !ok {
				err = perror.NewBizErr(constants.WinErrSeatSleeperMatch.Key(), constants.WinErrSeatSleeperMatch.Value())
				plogEntry.WithError(err).Error(logID, "reSeatPositionTypeSleeper-WinErrSeatSleeperMatch err")
				return
			}
			position := message.Position{
				Carriage:   "",
				SeatNumber: "",
				Sleeper:    sleeper,
			}
			positionList = append(positionList, position)
		}
		return
	}
	
	match = reSeatPositionTypeCarriagePerson.FindStringSubmatch(str)
	if len(match) > 0 {
		plogEntry.WithField("match", match).Info(logID, "reSeatPositionTypeCarriagePerson-匹配 {xx数字}车厢{N}人")
		if !IsSleeperSeatType(seatType) {
			err = perror.NewBizErr(constants.WinErrSeatMatchSleeper.Key(), constants.WinErrSeatMatchSleeper.Value())
			plogEntry.WithError(err).Error(logID, "reSeatPositionTypeSleeper-WinErrSeatSleeperMatch err")
			return
		}
		
		seatPositionType = message.SeatPositionTypeCarriage
		for i := 0; i < cast.ToInt(match[2]); i++ {
			position := message.Position{
				Carriage:   match[1],
				SeatNumber: "",
				Sleeper:    "",
			}
			positionList = append(positionList, position)
		}
		return
	}
	
	// 不符合上面所有的规则
	err = perror.NewBizErr(constants.WinErrSeatMatch.Key(), constants.WinErrSeatMatch.Value())
	plogEntry.WithError(err).Error("不符合上面所有的规则，请联系运营商！")
	return
}

func IsSleeperSeatType(seatType message.SeatType) bool {
	if strings.Contains(string(seatType), "Sleeper") {
		return true
	}
	return false
}
