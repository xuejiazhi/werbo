package utils

import (

	//"fmt"
	"encoding/base64"
	"reflect"
	"regexp"
	"strings"
	"unicode/utf8"
	// "strconv"
)

//截取字符串函数
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
} //end func Substr

//替换字符串
func StrReplace(str, old, new string) string {
	//返回替换的字符串
	return strings.Replace(str, old, new, -1)
} //end func ReplaceAll

//转成小写
func StrToLower(str string) string {
	return strings.ToLower(str)
}

//转成大写
func StrToUpper(str string) string {
	return strings.ToUpper(str)
}

//是否TXT包含在str中
func StrContains(str, txt string) bool {
	return strings.Contains(str, txt)
}

//判断两个字符串是否相等
func StrEqual(str, txt string) bool {
	return strings.EqualFold(str, txt)
}

//将字符串分割
func StrSplit(str, sep string) []string {
	return strings.Split(str, sep)
}

//去掉两边空格
func StrTrim(str string) string {
	return strings.TrimSpace(str)
}

//parse str
func ParseStr(str string) map[string]string {
	data := make(map[string]string)
	tmpData := strings.Split(str, "&")
	for _, v := range tmpData {
		item := strings.Split(v, "=")
		if len(item) != 2 {
			continue
		}
		kItem := reflect.ValueOf(item[0])
		key := kItem.String()
		vItem := reflect.ValueOf(item[1])
		value := vItem.String()
		data[key] = value
	}
	return data
}

/**
将字符串反转
**/
func ReverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

/**
http build query
*/
func HttpBuildQuery(data map[string]string) string {
	var content string
	index := 0
	length := len(data)
	for k, v := range data {
		index++
		// 组装数据
		if index == length {
			content += k + "=" + v
		} else {
			content += k + "=" + v + "&"
		}
	}
	return content
}

//过滤小图标
func FilterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		}
	}
	return new_content
}

//base64 生成编码
func Base64Encode(content string) string {
	input := []byte(content)
	return base64.StdEncoding.EncodeToString(input)
}

//base64解码
func Base64Decode(content string) string {
	//解码
	decodeBytes, _ := base64.StdEncoding.DecodeString(content)
	//判断是否是base64
	reStr := string(decodeBytes)
	// lenStr := len([]byte(content))
	if content == Base64Encode(reStr) { //是转码成功
		return reStr
	} else {
		return content
	}
}

//过滤html标签
func FilterHtmlTag(content string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src := re.ReplaceAllStringFunc(content, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = StrTrim(re.ReplaceAllString(src, "\n"))
	//返回
	return src
}
