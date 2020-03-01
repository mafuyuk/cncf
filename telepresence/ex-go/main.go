package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "app-b")
}

// Endpoint which uses an external service "debugkit"
func debugkitHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://debugkit/info/hostname")
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Fprint(w, string(bytes))
}

// Endpoint which uses an external service "app-a"
func appAHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://" + cfg.AppASvc)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Fprint(w, string(bytes))
}

// Endpoint which returns a secret message
func secretHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, cfg.VerySecretMessage)
}

type Cfg struct {
	Port              string
	AppASvc           string
	VerySecretMessage string
}

var cfg Cfg

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/debugkit", debugkitHandler)
	http.HandleFunc("/app-a", appAHandler)
	http.HandleFunc("/secret", secretHandler)

	cfg = Cfg{}

	viper.SetConfigName("app-b")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("APPB")
	viper.BindEnv("PORT")
	viper.BindEnv("APPASVC")
	viper.BindEnv("VERYSECRETMESSAGE")
	viper.Unmarshal(&cfg)

	http.ListenAndServe(":"+cfg.Port, nil)
}
