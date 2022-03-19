package sentry_util

import (
    "errors"
    "fmt"
    "github.com/getsentry/sentry-go"
    "log"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/17 上午2:19
 * @Desc:
 */

const (
    dsn = "http://58be04091efa42feb1aa18390230bf2f@127.0.0.1:9100/1"
)

func TestInitSentryOffice(t *testing.T) {
    err := sentry.Init(sentry.ClientOptions{
        Dsn: dsn,
    })
    if err != nil {
        log.Fatalf("sentry.Init: %s", err)
    }
    // Flush buffered events before the program terminates.
    defer sentry.Flush(2 * time.Second)
    
    sentry.CaptureMessage("prince It works! 001")
    sentry.CaptureMessage("prince It works! 002")
    sentry.CaptureMessage("prince It works! 003")
}

func TestInitSentryUtil(t *testing.T) {
    err := Init(dsn)
    if err != nil {
        fmt.Println("Init err:", err)
        return
    }
    var eventID *sentry.EventID
    eventID = CaptureMessage("prince sentry CaptureMessage a010203", time.Second * 1)
    fmt.Println(eventID)
    eventID = CaptureMessage("prince sentry CaptureMessage b010203", time.Second * 1)
    fmt.Println(eventID)
}

func TestInitCaptureException(t *testing.T) {
    err := Init(dsn)
    if err != nil {
        fmt.Println("Init err:", err)
        return
    }
    var eventID *sentry.EventID
    eventID = CaptureException(errors.New("prince sentry CaptureException aa010203040506"), time.Second * 10)
    fmt.Println(*eventID)
    eventID = CaptureException(errors.New("prince sentry CaptureException ab010203040506"), time.Millisecond * 1)
    fmt.Println(*eventID)
}
