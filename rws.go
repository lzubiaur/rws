package rws

import (
  "net/http"
  "net/url"
  "os"
)

var endpoint string = "https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706"

func Get() (resp *http.Response, err error) {
  v := url.Values{}
  v.Add("applicationId",os.Getenv("RWS_APPLICATION_ID"))
  v.Add("keyword","test")
  resp,err = http.Get(endpoint + "?" + v.Encode())
  return
}
