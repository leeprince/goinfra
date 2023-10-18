/*
 * @Date: 2020-07-23 15:17:40
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2021-04-15 11:54:46
 */
package gostreaming_test

import (
	"fmt"
	"math/rand"

	"testing"
	"time"

	"gitlab.yewifi.com/risk-control/risk-common/pkg/gostreaming"
	"gitlab.yewifi.com/risk-control/risk-common/pkg/redishelper"
)

func TestRedisStatusStorage(t *testing.T) {
	redisCli := redishelper.MustNewClient(redishelper.Config{
		Host:     "localhost",
		Port:     9999,
		Password: "",
	})
	statusStorage := gostreaming.NewRedisStatusStorage(redisCli)
	batch := gostreaming.NewBatch()
	batch.Get([]string{"version=1", "date=2020-07-23"}, "k1", []string{"is-man=true", "height=177"})
	batch.Set([]string{"version=1", "date=2020-07-23"}, "k1", []string{"is-man=true", "height=177"}, "dzp", 0)
	batch.Get([]string{"version=1", "date=2020-07-23"}, "k1", []string{"is-man=true", "height=177"})

	fmt.Println(batch)
	batchRsps, batchErrs, err := statusStorage.ExecBatch(batch)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("errs:")
	for i, err := range batchErrs {
		fmt.Printf("\ti: %d, err: %+v\n", i, err)
	}

	fmt.Println("rsps:")
	for i, rsp := range batchRsps {
		fmt.Printf("\ti: %d, rsp: %+v\n", i, rsp)
	}

}

func TestRedisStatusStorageZAdd(t *testing.T) {
	rand.Seed(0)

	redisCli := redishelper.MustNewClient(redishelper.Config{
		Host:     "localhost",
		Port:     9999,
		Password: "",
	})

	statusStorage := gostreaming.NewRedisStatusStorage(redisCli)
	batch := gostreaming.NewBatch()

	now := time.Now()
	for i := 0; i < 10; i++ {
		randN := rand.Int() % 1000
		batch.ZAdd([]string{"pk1", "pk2"}, "target1", []string{"desc1", "desc2"}, -float64(now.Add(time.Duration(randN)*time.Second).Unix()), randN)
	}

	_, batchErrs, err := statusStorage.ExecBatch(batch)
	if err != nil {
		t.Errorf("expect nil but got %+v", err)
		return
	}

	hasError := false
	for _, err := range batchErrs {
		if err != nil {
			t.Errorf("expect nil but got %+v", err)
			hasError = true
			continue
		}
	}
	if hasError {
		return
	}

	batch = gostreaming.NewBatch()
	randN := rand.Int() % 1000
	batch.ZRangeByScoreWithScores([]string{"pk1", "pk2"}, "target1", []string{"desc1", "desc2"},
		fmt.Sprintf("%f", -float64(now.Add(time.Duration(randN)*time.Second).Unix())),
		"+inf",
		2, 5,
	)

	results, batchErrs, err := statusStorage.ExecBatch(batch)
	if err != nil {
		t.Errorf("expect nil but got %+v", err)
		return
	}

	hasError = false
	for _, err := range batchErrs {
		if err != nil {
			t.Errorf("expect nil but got %+v", err)
			hasError = true
			continue
		}
	}
	if hasError {
		return
	}

	for _, result := range results {
		fmt.Println(result)
	}

}
