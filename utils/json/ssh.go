package json

import (
	"HFish/utils/log"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
)

var sshJson []byte

func init() {
	file, err := ioutil.ReadFile("./libs/ssh/config.json")

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "读取文件失败", err)
	}

	sshJson = file
}

func GetSsh() (*simplejson.Json, error) {
	res, err := simplejson.NewJson(sshJson)
	return res, err
}
