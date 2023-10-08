package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leeprince/goinfra/utils/timeutil"
	"log"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/29 14:52
 * @Desc:
 */

const (
	SignedStringKey = "hello"
)

const (
	RegisteredClaimsIssuer  = "prince"
	RegisteredClaimsSubject = "lee"
)

func CreateJwtNew() {
	var (
		key []byte
		t   *jwt.Token
		s   string
		err error
	)

	key = []byte(SignedStringKey)
	t = jwt.New(jwt.SigningMethodHS256)
	fmt.Printf("t: %+v \n", t)

	s, err = t.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("s: %+v \n", s)
}

func CreateJwtByNewWithClaims() (tokenString string, err error) {
	var (
		key []byte
		t   *jwt.Token
	)

	key = []byte(SignedStringKey)
	t = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    RegisteredClaimsIssuer,
			Subject:   RegisteredClaimsSubject,
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(timeutil.AfterSecond(60)),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	)
	fmt.Printf("t: %+v \n", t)

	tokenString, err = t.SignedString(key)
	if err != nil {
		fmt.Println("SignedString err:", err)
		return
	}
	fmt.Printf("tokenString: %+v \n", tokenString)

	return
}

func ParseWithClaims(tokenString string) (b bool, err error) {
	var (
		key []byte
	)

	key = []byte(SignedStringKey)

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("jwt.ParseWithClaims err:", err)
		return
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		fmt.Printf("token.Claims.(*jwt.RegisteredClaims) claims：%v", claims)
		b = true
	} else {
		fmt.Println("token.Claims.(*jwt.RegisteredClaims) err:", err)
	}

	return
}

/******************* 官网：example_test.go ***********************/
// Example creating a token using a custom claims type. The RegisteredClaims is embedded
// in the custom type to allow for easy encoding, parsing and validation of registered claims.
func ExampleNewWithClaims_customClaimsType() {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	fmt.Printf("foo: %v\n", claims.Foo)

	// Create claims while leaving out some of the optional fields
	claims = MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	//Output: foo: bar
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNTE2MjM5MDIyfQ.xVuY2FZ_MRXMIEgVQ7J-TFtaucVFRXUzHm9LmV41goM <nil>
}

// Example creating a token using a custom claims type.  The RegisteredClaims is embedded
// in the custom type to allow for easy encoding, parsing and validation of standard claims.
func ExampleParseWithClaims_customClaimsType() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIn0.QAWg1vGvnqRuCFTMcPkjZljXHh8U3L_qUjszOtQbeaA"

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.RegisteredClaims.Issuer)
	} else {
		fmt.Println(err)
	}

	// Output: bar test
}

/******************* 官网：example_test.go -end ***********************/
