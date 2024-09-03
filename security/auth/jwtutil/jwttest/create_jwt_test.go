package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/29 15:06
 * @Desc:
 */

func TestCreateJwt(t *testing.T) {
	CreateJwtNew()
}

func TestCreateJwtByNewWithClaims(t *testing.T) {
	tokenString, err := CreateJwtByNewWithClaims()
	if err != nil {
		log.Fatal("tokenString err:", err)
	}
	fmt.Println("tokenString:", tokenString)
}

func TestParseWithClaims(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwcmluY2UiLCJzdWIiOiJsZWUiLCJleHAiOjE3MTg0NzI0NDJ9.DojzDlV2kBF5cmYp1ZbFzlrBcuPem5zA7DEkfQfMnzg"
	b, err := ParseWithClaims(tokenString)
	if err != nil {
		log.Fatal("ParseWithClaims err:", err)
	}
	fmt.Println("b:", b)
}

func TestCreateJwtByNewWithClaimsAndParse(t *testing.T) {
	tokenString, err := CreateJwtByNewWithClaims()
	if err != nil {
		log.Fatal("CreateJwtByNewWithClaims tokenString err:", err)
	}
	fmt.Println("CreateJwtByNewWithClaims tokenString:", tokenString)
	
	b, err := ParseWithClaims(tokenString)
	if err != nil {
		log.Fatal("ParseWithClaims err:", err)
	}
	fmt.Println("ParseWithClaims b:", b)
	
	time.Sleep(time.Second * 1)
}

func TestExampleNewWithClaims_customClaimsType(t *testing.T) {
	ExampleNewWithClaims_customClaimsType()
}
func TestExampleParseWithClaims_customClaimsType(t *testing.T) {
	ExampleParseWithClaims_customClaimsType()
}
