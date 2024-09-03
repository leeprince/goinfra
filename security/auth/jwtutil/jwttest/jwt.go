package main

import (
	"github.com/golang-jwt/jwt/v5"
	jsoniter "github.com/json-iterator/go"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/6/22 16:48
 * @Desc:
 */

// GenerateJWT 生成JWT Token
// 	JWT由三部分组成：头部（header）、载荷（payload）和签名（signature）。
// 		头部通常由两部分组成：令牌类型（默认：Auth）和所使用的加密算法。
//		载荷（payload）是包含令牌中存储的信息的主要部分。它包括一些声明（claim），例如用户ID、用户名、过期时间等等。
//		签名（signature）是将头部和载荷组合后使用密钥生成的哈希值，用于验证令牌的真实性。
//	在HTTP请求添加名为Authorization的header，形式如下
//		Authorization: Bearer <token>
// @secretKey: Auth 加解密密钥
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func GenerateJWT(payloadKey, secret string, seconds int64, payload interface{}) (string, error) {
	// 头部通常由两部分组成：令牌类型（默认：Auth）和所使用的加密算法
	// 创建一个声明:载荷（payload）是包含令牌中存储的信息的主要部分。它包括一些声明（claim），例如用户ID、用户名、过期时间等等。
	iat := time.Now().Unix()
	payloadString, err := jsoniter.MarshalToString(payload)
	if err != nil {
		return "", err
	}
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	
	// 载荷（payload）使用自定义键名
	// claims["payload"] = payloadString
	claims[payloadKey] = payloadString
	
	// 签名（signature）是将头部和载荷组合后使用密钥生成的哈希值，用于验证令牌的真实性。
	signingMethod := jwt.SigningMethodHS256
	
	// 创建一个新的JWT，设置签名（signature）算法、载荷（payload）
	token := jwt.NewWithClaims(signingMethod, claims)
	
	// 生成JWT
	return token.SignedString([]byte(secret))
}
