package main

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/29 15:06
 * @Desc:
 */

func TestCreateJwt(t *testing.T) {
	CreateJwtNew()
}

func TestCreateJwtNewWithClaims(t *testing.T) {
	CreateJwtNewWithClaims()
}
