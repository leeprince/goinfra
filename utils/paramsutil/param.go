package paramsutil

import "regexp"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/7/3 20:42
 * @Desc:
 */

// ValidateEmail 检查给定的电子邮件地址是否有效
func ValidateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	return matched
}

// ValidatePhone 检查给定的电话号码是否为中国有效的手机号码
func ValidatePhone(phone string) bool {
	phoneRegex := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(phoneRegex, phone)
	return matched
}
