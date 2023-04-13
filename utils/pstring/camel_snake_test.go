package pstring

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午10:10
 * @Desc:
 */

func TestRune(t *testing.T) {
	// msg := "leeprince"
	msg := "我爱您，中国"
	messageRunes := []rune(msg)
	fmt.Println(messageRunes)
	fmt.Println(string(messageRunes))

	messagebytes := []byte(msg)
	fmt.Println(messagebytes)
	fmt.Println(string(messagebytes))
}

func TestNewReplacer(t *testing.T) {
	// Copied from golint
	var commonInitialisms = []string{"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "UUID", "URI", "host", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS"}

	var commonInitialismsForReplacer []string
	var uncommonInitialismsForReplacer []string
	for i := len(commonInitialisms) - 1; i >= 0; i-- {
		initialism := commonInitialisms[i]
		toLower := strings.ToLower(initialism)
		stringsTitle := strings.Title(toLower)
		fmt.Println(initialism, toLower, stringsTitle)
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, stringsTitle)
		uncommonInitialismsForReplacer = append(uncommonInitialismsForReplacer, stringsTitle, initialism)
	}
}
func TestStrings(t *testing.T) {
	str := "myName"
	str1 := "my_Name"
	str2 := "MyName"
	str3 := "my_name"
	str4 := "myname"
	str5 := "MYNAME"

	ToLower := strings.ToLower(str)
	fmt.Println("ToLower:", ToLower)
	ToLower1 := strings.ToLower(str1)
	fmt.Println("ToLower1:", ToLower1)

	ToUpper := strings.ToUpper(str1)
	fmt.Println("ToUpper:", ToUpper)

	ToTitle := strings.ToTitle(str)
	fmt.Println("ToTitle:", ToTitle, strings.Title(str))
	ToTitle1 := strings.ToTitle(str1)
	fmt.Println("ToTitle1:", ToTitle1, strings.Title(str1))
	ToTitle2 := strings.ToTitle(str2)
	fmt.Println("ToTitle2:", ToTitle2, strings.Title(str2))
	ToTitle3 := strings.ToTitle(str3)
	fmt.Println("ToTitle3:", ToTitle3, strings.Title(str3))
	ToTitle4 := strings.ToTitle(str4)
	fmt.Println("ToTitle4:", ToTitle4, strings.Title(str4))
	ToTitle5 := strings.ToTitle(str5)
	fmt.Println("ToTitle5:", ToTitle5, strings.Title(strings.ToLower(str5)))

}

func TestChar(t *testing.T) {
	str := "P.Lee@xxx.com℃ᾭ"
	for _, i2 := range str {
		fmt.Println(">>>>>:", string(i2))
		fmt.Println(unicode.IsLower(i2))
		fmt.Println(unicode.IsLetter(i2))
		fmt.Println(unicode.IsUpper(i2))
		fmt.Println(unicode.IsTitle(i2))
	}
}

func TestToLowerCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "my_name"},
			want: "myName",
		},
		{
			args: args{s: "my_namE"},
			want: "myNamE",
		},
		{
			args: args{s: "myname"},
			want: "myname",
		},
		{
			args: args{s: "MyName"},
			want: "myName",
		},
		{
			args: args{s: "MyNamePrince"},
			want: "myNamePrince",
		},
		{
			args: args{s: "MyName,prince"},
			want: "myNamePrince",
		},
		{
			args: args{s: "MyName,Prince"},
			want: "myNamePrince",
		},
		{
			args: args{s: "MyName, , - prince"},
			want: "myNamePrince",
		},
		{
			args: args{s: "MyName prince"},
			want: "myNamePrince",
		},
		{
			args: args{s: "_MyName prince"},
			want: "MyNamePrince",
		},
		{
			args: args{s: "MyName prince_"},
			want: "myNamePrince",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToLowerCamel(tt.args.s)
			if got != tt.want {
				t.Errorf("ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUpperCamel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "id"},
			want: "Id",
		},
		{
			args: args{s: "my_name"},
			want: "MyName",
		},
		{
			args: args{s: "my_namE"},
			want: "MyNamE",
		},
		{
			args: args{s: "myname"},
			want: "Myname",
		},
		{
			args: args{s: "MyName"},
			want: "MyName",
		},
		{
			args: args{s: "MyNamePrince"},
			want: "MyNamePrince",
		},
		{
			args: args{s: "MyName,prince"},
			want: "MyNamePrince",
		},
		{
			args: args{s: "MyName,Prince"},
			want: "MyNamePrince",
		},
		{
			args: args{s: "MyName, , - prince"},
			want: "MyNamePrince",
		},
		{
			args: args{s: "MyName prince"},
			want: "MyNamePrince",
		},
		{
			args: args{s: "_MyName prince"},
			want: "MyNamePrince",
		},
		{
			args: args{s: "MyName prince_"},
			want: "MyNamePrince",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpperCamel(tt.args.s); got != tt.want {
				t.Errorf("ToUpperCamel() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.s, ">>>>>>>", tt.want)
		})
	}
}

func TestToSnake(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "my_name"},
			want: "my_name",
		},
		{
			args: args{s: "my_namE"},
			want: "my_nam_e",
		},
		{
			args: args{s: "myname"},
			want: "myname",
		},
		{
			args: args{s: "MyName"},
			want: "my_name",
		},
		{
			args: args{s: "MyNamePrince"},
			want: "my_name_prince",
		},
		{
			args: args{s: "MyName,prince"},
			want: "my_name_prince",
		},
		{
			args: args{s: "MyName,Prince"},
			want: "my_name_prince",
		},
		{
			args: args{s: "MyName, , - prince"},
			want: "my_name_prince",
		},
		{
			args: args{s: "MyName prince"},
			want: "my_name_prince",
		},
		{
			args: args{s: "_MyName prince"},
			want: "_my_name_prince",
		},
		{
			args: args{s: "MyName prince_"},
			want: "my_name_prince_",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSnake(tt.args.s); got != tt.want {
				t.Errorf("ToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}
