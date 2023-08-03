package test

import (
	"errors"
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/18 01:34
 * @Desc:
 */

func TestReturnErr(t *testing.T) {
	err := ReturnErr01()
	fmt.Println("TestReturnErr:", err)
}

func ReturnErr01() (err error) {
	s, err := ReturnErr02()
	if err != nil {
		fmt.Println("ReturnErr02 err:", err)
		return
	}
	fmt.Println(s)
	return
}
func ReturnErr02() (s string, err error) {
	return "", errors.New("ReturnErr02")
}

func TestReturnForScopeErr(t *testing.T) {
	err := ReturnForScopeErr01()
	fmt.Println("ReturnForScopeErr01:", err)

	// 会报错：因为在其他作用域的使用了返回参数的变量，进行了重新声明赋值，导致报错。
	// 解决：在其他作用域的使用了返回参数的变量时，不要重新声明，只需赋值。
	//	如：var s string; s, err = xxx();
	//	如：s, serr := xxx(); err = serr;
	//	如：_, err = xxx();
	err = ReturnForScopeErr0101()
	fmt.Println("ReturnForScopeErr0101:", err)
}

func ReturnForScopeErr01() (err error) {
	s, err := ReturnErr02()
	if err != nil {
		fmt.Println("ReturnErr02 err:", err)
		return
	}
	fmt.Println(s)
	return
}
func ReturnForScopeErr0101() (err error) {
	/*for i := 0; i < 3; i++ {
		// 报错：inner declaration of var err error
		s, err := ReturnErr02()
		if err != nil {
			// 报错：result parameter err not in scope at return
			fmt.Println("ReturnErr02 err:", err)
			return
		}
		// 报错：result parameter err not in scope at return
		fmt.Println(s)
		return
	}*/

	// 解决：方式一；var s string; s, err = xxx();
	for i := 0; i < 3; i++ {
		var s string
		s, err = ReturnErr02()
		if err != nil {
			fmt.Println("ReturnErr02 err:", err)
			return
		}
		fmt.Println(s)
		return
	}

	// 解决：方式二；s, serr := xxx(); err = serr;
	for i := 0; i < 3; i++ {
		s, serr := ReturnErr02()
		if serr != nil {
			err = serr
			fmt.Println("ReturnErr02 err:", err)
			return
		}
		fmt.Println(s)
		return
	}

	// 解决：方式三 如：_, err = xxx();
	for i := 0; i < 3; i++ {
		_, err = ReturnErr02()
		if err != nil {
			fmt.Println("ReturnErr02 err:", err)
			return
		}
		return
	}

	return
}

func TestReturnIfScopeErr(t *testing.T) {
	var err error

	err = ReturnIfScopeErr01()
	fmt.Println("ReturnForScopeErr01:", err)

	err = ReturnIfScopeErr02()
	fmt.Println("ReturnIfScopeErr02:", err)
}

func ReturnIfScopeErr01() (err error) {
	if true {
		/*
			报错：
			.\return_error_test.go:112:3: result parameter err not in scope at return
				.\return_error_test.go:109:6: inner declaration of var err error

			解决：在其他作用域的使用了返回参数的变量时，不要重新声明，只需赋值。
				如：var s string; s, err = xxx();
				如：s, serr := xxx(); err = serr;
				如：_, err = xxx();

		*/
		_, err = ReturnErr02()
		if err != nil {
			fmt.Println("ReturnIfScopeErr01 err:", err)
			return
		}
		return
	}

	return
}

func ReturnIfScopeErr02() (err error) {
	s, err := ReturnErr02()
	if err != nil {
		fmt.Println("ReturnIfScopeErr02 err:", err)
		return
	}
	fmt.Println("ReturnIfScopeErr02 s:", s)

	return
}
