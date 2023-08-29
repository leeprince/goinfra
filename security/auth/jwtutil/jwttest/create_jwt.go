package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leeprince/goinfra/utils/timeutil"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/29 14:52
 * @Desc:
 */

func CreateJwtNew() {
	var (
		key []byte
		t   *jwt.Token
		s   string
		err error
	)

	key = []byte("hello")
	t = jwt.New(jwt.SigningMethodHS256)
	fmt.Printf("t: %+v \n", t)

	s, err = t.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("s: %+v \n", s)
}

func CreateJwtNewWithClaims() {
	var (
		key []byte
		t   *jwt.Token
		s   string
		err error
	)

	key = []byte("hello")
	t = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "prince",
			Subject:   "p",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(timeutil.AfterSecond(60)),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	)
	fmt.Printf("t: %+v \n", t)

	s, err = t.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("s: %+v \n", s)
}
