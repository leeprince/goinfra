package test

import (
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils/idutil"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/5 下午10:29
 * @Desc:
 */

func TestPlog(t *testing.T) {
	plog.Debug("prince log Debug SetReportCaller")
	plog.Debug("prince log Debug SetReportCaller")
	plog.WithField("WithField01", "WithFieldValue").Debug("prince log Debug WithField")
}

func TestPlogSetOutputFile(t *testing.T) {
	plog.SetOutputFile("./", "application.log", true)
	plog.Debug("prince log Debug SetReportCaller")
	plog.Debug("prince log Debug SetReportCaller")
	plog.WithField("WithField01", "WithFieldValue").Debug("prince log Debug WithField")
}

func TestPlogSetOutputFileV1(t *testing.T) {
	plog.SetOutputFile("./", "application.log", false)
	plog.Debug("prince log Debug SetReportCaller")
	plog.Debug("prince log Debug SetReportCaller")
	plog.WithField("WithField01", "WithFieldValue").Debug("prince log Debug WithField")
}

func TestPlogSetOutputFileV2(t *testing.T) {
	plog.SetOutputFile("./logs/", "application.log", true)
	plog.Debug("prince log Debug SetReportCaller")
	plog.Debug("prince log Debug SetReportCaller")
	plog.WithField("WithField01", "WithFieldValue").Debug("prince log Debug WithField")
}

func TestPlogSetReportCaller(t *testing.T) {
	// plog.Debug("prince log Debug SetReportCaller 01")
	plog.SetReportCaller(true)
	plog.Info("prince log Info SetReportCaller 01")
	plog.Debug("prince log Debug SetReportCaller 01")
	plog.Error("prince log Error SetReportCaller 01")
	plog.Warn("prince log Warn SetReportCaller 01")
	plog.WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField 01")
	
	plog.SetReportCaller(false)
	plog.Info("prince log Info SetReportCaller 02")
	plog.Debug("prince log Debug SetReportCaller 02")
	plog.Error("prince log Error SetReportCaller 02")
	plog.Warn("prince log Warn SetReportCaller 02")
	plog.WithField("WithField01", "WithFieldValue02").Debug("prince log Debug WithField 02")
}

func TestPlogSetReportCallerLevel(t *testing.T) {
	plog.Debug("prince log Debug SetReportCaller 01")
	plog.SetReportCaller(true, plog.ErrorLevel)
	plog.Info("prince log Info SetReportCaller 01")
	plog.Debug("prince log Debug SetReportCaller 01")
	plog.Error("prince log Error SetReportCaller 01")
	plog.Warn("prince log Warn SetReportCaller 01")
	plog.WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField 01")
	
	plog.SetReportCaller(false)
	plog.Info("prince log Info SetReportCaller 02")
	plog.Debug("prince log Debug SetReportCaller 02")
	plog.Error("prince log Error SetReportCaller 02")
	plog.Warn("prince log Warn SetReportCaller 02")
	plog.WithField("WithField01", "WithFieldValue02").Debug("prince log Debug WithField 02")
}

func TestPlogSetReportCallerMore(t *testing.T) {
	plog.Debug("prince log Debug SetReportCaller 01")
	plog.SetReportCaller(true)
	plog.Info("prince log Info SetReportCaller 01")
	plog.Debug("prince log Debug SetReportCaller 01")
	plog.Error("prince log Error SetReportCaller 01")
	plog.Warn("prince log Warn SetReportCaller 01")
	plog.WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField 01")
	
	plog.SetReportCaller(false)
	plog.Info("prince log Info SetReportCaller 02")
	plog.Debug("prince log Debug SetReportCaller 02")
	plog.Error("prince log Error SetReportCaller 02")
	plog.Warn("prince log Warn SetReportCaller 02")
	plog.WithField("WithField01", "WithFieldValue02").Debug("prince log Debug WithField 02")
	
	// SetReportCaller 一次会增加一个钩子。所以在上面已经设置所有日志等级的情况下，重新设置指定等级是无效的，因为上面设置的钩子中已经包含所有日志等级
	plog.SetReportCaller(true, plog.ErrorLevel)
	plog.Info("prince log Info SetReportCaller 03")
	plog.Debug("prince log Debug SetReportCaller 03")
	plog.Error("prince log Error SetReportCaller 03")
	plog.Warn("prince log Warn SetReportCaller 03")
	plog.WithField("WithField01", "WithFieldValue03").Debug("prince log Debug WithField 03")
}

func TestPlogEntry(t *testing.T) {
	plogEntry := plog.LogID(idutil.UniqIDV3()).WithField("method", "TestPlogEntry")
	
	plogEntry.Info("request")
	
	plogEntry.Info("handler...")
	
	plogEntry.Info("response")
	
}

func TestPanicLog(t *testing.T) {
	plogEntry := plog.LogID(idutil.UniqIDV3()).WithField("method", "TestPlogEntry")
	
	plogEntry.Panic("request")
	
	plogEntry.Panic("handler...")
	
	plogEntry.Panic("response")
	
}
