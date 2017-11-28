//  Copyright (c) 2017 The Go Authors. All rights reserved.
//  load configure of json
//	author JEE
//  contact :xuejiazhi@gamil.com

package configure

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"werbo/utils"
)

type JsonCfg struct{
	FileName, Key, Value string
}

//Factory method
func  JsFactory(FileName, Key, Value string) *JsonCfg{
	//返回
	return &JsonCfg{FileName,Key,Value}
}//end func Factory

//LoadJsonConfig
//Get JSON configuration
func (this *JsonCfg) LoadConfig() string {
	//Get the configuration content
	confInfo := readJsonFile(this.FileName + "." + JSON_EXT_NAME)
	//Turn into interface type
	var param = make(map[string]interface{}, 0)
	//Turn into json
	json.Unmarshal([]byte(confInfo), &param)
	// Transformation of subordinate configuration 
	sonParam, _ := utils.Inter2Map(param[this.Key])
	//return 
	ret,_ := utils.Inter2Str(sonParam["date"])
	return ret
} //end func LoadConfig

//readfile of json
func  readJsonFile(filename string) string {
	//readfile
	file, err := os.Open(CFG_PATH + "/" + filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no "+filename+" config file!", err)
		os.Exit(1)
	}
	
	//close file
	defer file.Close()
	inputReader := bufio.NewReader(file)
	reStr := ""

	//Circulation retrieval
	for {
		//read line
		inputStr, readerError := inputReader.ReadString('\n')
		//Remove notes
		inputStr = removeJsonNote(inputStr)
		//Is there a child configuration?
		sonCfgName := getJsonSonCfgName(inputStr)
		// If  have, pick up the configuration 
		if sonCfgName != "" {
			// Recursive call 
			inputStr = readJsonFile(sonCfgName) + "\n"
		}
		reStr = reStr + inputStr
		if readerError == io.EOF {
			break
		}
	}
	//return
	return reStr
} //read file

//comment remove
func removeJsonNote(str string) string {
	//Is it a comment
	pattern = `\/\/.*`
	reg = regexp.MustCompile(pattern)
	regStr := reg.FindAllString(str, 1)
	retStr := ""
	if len(regStr) > 0 {
		retStr = strings.Replace(str, regStr[0], "", -1)
	} else {
		retStr = str
	}
	//Is it a blank line
	if strings.TrimSpace(retStr) == "" {
		return ""
	}
	return retStr
}

//Is there a child configuration
func getJsonSonCfgName(str string) string {
	//Is there a child configuration
	pattern = `\$\$.*\$\$`
	reg = regexp.MustCompile(pattern)
	regStr := reg.FindAllString(str, -1)
	if len(regStr) > 0 {
		return strings.TrimSpace(strings.Replace(str, "$$", "", -1))
	} else {
		return ""
	}
} //end func getSonCfgName
