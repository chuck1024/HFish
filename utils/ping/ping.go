package ping

import (
	"HFish/utils/conf"
	"HFish/utils/try"
	"net/http"
)

func Ping() {
	try.Try(func() {
		rpcStatus := conf.Get("rpc", "status")

		s := "Server"

		if rpcStatus == "2" {
			s = "Client"
		}

		resp, _ := http.Get("http://ping.hfish.io/test?s=" + s)
		defer resp.Body.Close()
	}).Catch(func() {})
}
