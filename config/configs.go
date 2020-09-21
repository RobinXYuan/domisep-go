package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-ini/ini"
	"github.com/go-playground/validator/v10"
)

var (
	//Cfg Config file
	Cfg *ini.File

	//RunMode System run mode debug/release
	RunMode string

	//HTTPPort port number
	HTTPPort int

	//JWTSecret JWT secret key
	JWTSecret string
)

//LoadConfigs load system configs
func LoadConfigs() {
	useJSONNameAsValidationTag()
	loadConfigs()
}

func useJSONNameAsValidationTag() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func loadConfigs() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		panic("Cannot load config file")
	}

	loadBaseConfigs()
	loadServerConfigs()
	loadAppConfigs()
}

func getRunMode() string {
	return Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func getSection(sectionName string) *ini.Section {
	sec, err := Cfg.GetSection(sectionName)
	if err != nil {
		panic(fmt.Sprintf("Load %s config failed!", sectionName))
	}

	return sec
}

func loadBaseConfigs() {
	RunMode = getRunMode()
}

func loadServerConfigs() {
	sec := getSection("server")

	RunMode = getRunMode()

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
}

func loadAppConfigs() {
	sec := getSection("app")

	JWTSecret = sec.Key("JWT_SECRET").MustString("11ca6be05d2043f2bfa680fbe4d3883e")
}
