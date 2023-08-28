package test

import (
	"fmt"
	"github.com/leeprince/goinfra/http/httpcli"
	"github.com/leeprince/goinfra/utils/idutil"
	"log"
	"net/http"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/29 00:04
 * @Desc:
 */

func TestNewHttpClientDefautl(t *testing.T) {
	url := "http://127.0.0.1:18080/"
	client := httpcli.NewHttpClient()
	
	respBody, resp, err := client.
		WithURL(url).
		WithMethod(http.MethodPost).
		Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp:%+v \n", resp)
	fmt.Printf("respBody:%+v \n", string(respBody))
}

func TestNewHttpClientDefautlLogId(t *testing.T) {
	url := "http://127.0.0.1:18080/"
	client := httpcli.NewHttpClient()
	
	respBody, resp, err := client.
		WithURL(url).
		WithMethod(http.MethodPost).
		WithLogID(idutil.UniqIDV3()).
		Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp:%+v \n", resp)
	fmt.Printf("respBody:%+v \n", string(respBody))
}

func TestNewHttpClientPing(t *testing.T) {
	url := "http://127.0.0.1:18080/ping"
	client := httpcli.NewHttpClient()
	
	respBody, resp, err := client.
		WithURL(url).
		WithMethod(http.MethodPost).
		Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp:%+v \n", resp)
	fmt.Printf("respBody:%+v \n", string(respBody))
}

func TestNewHttpClientReportAdd(t *testing.T) {
	url := "http://127.0.0.1:18080/report/add"
	client := httpcli.NewHttpClient()
	
	req := &ReportOrderResultReq{
		OrderId:        "OrderId-01",
		TicketNumber:   "TicketNumber-02",
		ContactPhone:   "ContactPhone-03",
		IsTransfer:     1,
		IsOccupySeat:   1,
		CompleteStatus: 1,
		FailReason:     "FailReason-07",
		MachineName:    "MachineName-08",
	}
	respBody, resp, err := client.
		WithURL(url).
		WithMethod(http.MethodPost).
		WithBody(req).
		Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp:%+v \n", resp)
	fmt.Printf("respBody:%+v \n", string(respBody))
}
