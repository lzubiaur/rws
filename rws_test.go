package rws

import (
  "testing"
  "io/ioutil"
)

func TestGet(t *testing.T) {
  resp,err := Get()
  if err != nil {
    t.Error("Connection error: %s",err)
  }
  if resp.StatusCode != 200 {
    t.Error()
    content,err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    t.Error(string(content[:]))
    if err != nil {
      t.Error(err)
    }
  }
}
