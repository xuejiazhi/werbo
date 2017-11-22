/**
*@author JEE
*
*读取配置
 */
package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	reg     *regexp.Regexp
	pattern string
	source  string
)

func LoadConfig(filename, item string) {
	//Get the configuration content
	confInfo := readFile(filename + ".wb")
	//转成interface
	var param = make(map[string]interface{}, 0)
	json.Unmarshal([]byte(confInfo), &param)
	//将下级配置转化
	sonParam, _ := InterfaceToMapInterface(param["base"])
	fmt.Println(sonParam)
	//返回
	// return sonParam[item]
}

//readfile
func readFile(filename string) string {
	//读取文件
	file, err := os.Open("conf/" + filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no "+filename+" config file!", err)
		os.Exit(1)
	}
	//关闭文件
	defer file.Close()
	inputReader := bufio.NewReader(file)
	reStr := ""
	//循环调取
	for {
		inputStr, readerError := inputReader.ReadString('\n')
		//移除注释
		inputStr = removeNote(inputStr)
		//是否有子配置
		sonCfgName := getSonCfgName(inputStr)
		//如果有,调取子配置
		if sonCfgName != "" {
			inputStr = readFile(sonCfgName) + "\n"
		}
		reStr = reStr + inputStr
		if readerError == io.EOF {
			break
		}
	}
	//返回
	return reStr
} //读取文件

//移除注释
func removeNote(str string) string {
	//是否是注释
	pattern = `\/\/.*`
	reg = regexp.MustCompile(pattern)
	regStr := reg.FindAllString(str, 1)
	retStr := ""
	if len(regStr) > 0 {
		retStr = strings.Replace(str, regStr[0], "", -1)
	} else {
		retStr = str
	}
	//是否是空行
	if strings.TrimSpace(retStr) == "" {
		return ""
	}
	return retStr
}

//是否有子配置
func getSonCfgName(str string) string {
	//是否是子配置
	pattern = `\$\$.*\$\$`
	reg = regexp.MustCompile(pattern)
	regStr := reg.FindAllString(str, -1)
	if len(regStr) > 0 {
		return strings.TrimSpace(strings.Replace(str, "$$", "", -1))
	} else {
		return ""
	}
} //end func getSonCfgName
