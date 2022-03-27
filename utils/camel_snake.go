package utils

import (
    "github.com/leeprince/goinfra/consts"
    "regexp"
    "strings"
    "unicode"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午10:10
 * @Desc:   驼峰命名法 & 蛇形命名法
 */

// 驼峰命名法：小驼峰
func ToLowerCamel(s string) string {
    var underlineRune []rune
    for i, strRune := range s {
        if i == 0 {
            underlineRune = append(underlineRune, strRune)
            continue
        }
        if !unicode.IsLetter(strRune) {
            underlineRune = append(underlineRune, consts.UnderlineRune)
            continue
        }
        if unicode.IsUpper(strRune) {
            underlineRune = append(underlineRune, consts.UnderlineRune, strRune)
            continue
        }
        underlineRune = append(underlineRune, strRune)
    }
    
    underlineStrSlice := strings.Split(string(underlineRune), consts.UnderlineStr)
    var wordLowerCamel []string
    for i, word := range underlineStrSlice {
        if i == 0 {
            wordLowerCamel = append(wordLowerCamel, strings.ToLower(word))
            continue
        }
        wordLowerCamel = append(wordLowerCamel, strings.Title(strings.ToLower(word)))
    }
    return strings.Join(wordLowerCamel, "")
}

// 驼峰命名法：大驼峰
func ToUpperCamel(s string) string {
    lowerCamel := ToLowerCamel(s)
    lowerCamelRune := []rune(lowerCamel)
    lowerCamelRune[0] = unicode.ToUpper(lowerCamelRune[0])
    return string(lowerCamelRune)
}

// 蛇形命名法
//  下划线开头/结尾都保留
func ToSnake(s string) string {
    var underlineRune []rune
    for i, strRune := range s {
        if i == 0 {
            underlineRune = append(underlineRune, unicode.ToLower(strRune))
            continue
        }
        if !unicode.IsLetter(strRune) {
            underlineRune = append(underlineRune, consts.UnderlineRune)
            continue
        }
        if unicode.IsUpper(strRune) {
            underlineRune = append(underlineRune, consts.UnderlineRune, unicode.ToLower(strRune))
            continue
        }
        underlineRune = append(underlineRune, strRune)
    }
    underlineStr := string(underlineRune)
    
    re := regexp.MustCompile("(_+)")
    return re.ReplaceAllString(underlineStr, consts.UnderlineStr)
}
