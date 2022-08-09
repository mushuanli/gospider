package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
)


type config struct {
	AuthUser string `json:"auth_user"`
	AuthPass string `json:"auth_pass"`
	AuthIOT string `json:"auth_iot"`
	CookieUName string `json:"cookie_uname"`
	CookieUPass string `json:"cookie_upwd"`

	ApiList []string `json:"apiList"`
	ApiUsrInfo string `json:"api_usrinfo"`
	ApiDevList string `json:"api_dev"`
	ApiUsrPwd string `json:"api_userpwd"`
	ApiDevData string `json:"api_devdata"`
	ApiSwitch string `json:"api_switch"`
}

var CfgInfo config


func initLog(filename string) error {
	dir := filepath.Dir(filename)
	if dir == "" {
		dir = "."
	}

	file := dir + "/message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	//gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	return nil
}


func Init(filename string) bool {
	err := initLog(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Printf("Successfully Opened config file: %s\n", filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()


	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	json.Unmarshal(byteValue, &CfgInfo)
	return true
}

func GetApi( field string) string {
	
    r := reflect.ValueOf(CfgInfo)
    f := reflect.Indirect(r).FieldByName(field)
    return f.String()
}