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
	plog.Debug("prince log Debug SetReportCaller 01")
	plog.SetReportCaller(true)
	plog.Debug("prince log Debug SetReportCaller 01")
	plog.WithField("WithField01", "WithFieldValue01").Debug("prince log Debug WithField 01")

	plog.SetReportCaller(false)
	plog.Debug("prince log Debug SetReportCaller 02")
	plog.Debug("prince log Debug SetReportCaller 02")
	plog.WithField("WithField01", "WithFieldValue02").Debug("prince log Debug WithField 02")
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
