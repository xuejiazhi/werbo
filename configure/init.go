package configure
import(
	"regexp"
)

var (
	reg     *regexp.Regexp
	pattern string
	source  string
)

//const configure 
const(
	BASE_CFG = "json"   //Standard configuration
	CFG_PATH = "conf"   //configure Dir path
	JSON_EXT_NAME = "wb" //JSON config file ext
	JSON_SON_EXT_NAME = "swb" //JSON ext config file ext
)

//Definition Interface
type Jsons interface{
	JsFactory(string,string,string) *JsonCfg
}

type Confs interface{

}

type Cfg interface{
	LoadCfg() 
	Jsons
	Confs
}

//Main Get Configure
func LoadCfg(file,key,value string) (string,bool){
	switch BASE_CFG{
	case "json":
		//return Json configure
		return  JsFactory(file,key,value).LoadConfig(),true
	case "conf":
		return  JsFactory(file,key,value).LoadConfig(),true		
	default:
		//return Json configure
		return JsFactory(file,key,value).LoadConfig(),true
	}
	//return wrong
	return "Get Configure Fail!",false
}//end func LoadCfg