package common

import (
  "strings"
  "bytes"
  "fmt"
  "log"
  "net/http"
  "net/url"
  "io/ioutil"
  "crypto/tls"
  "encoding/json"
)

func ApiRequest(full_url string, header map[string]string, params url.Values, data interface{}, method string) (ret interface{}, str string) {
  ret = nil
  defer func() {
    if err := recover(); err != nil {
      log.Println(err)
    }
  }()

  var req *http.Request
  var err error

  switch method {
    case "GET", "DELETE" :
      if nil != params {
        full_url += fmt.Sprintf("?%v", params.Encode())
      }
      req, err = http.NewRequest(method, full_url, nil)
    case "POST" :
      if nil != data {
        jsonData, err := json.Marshal(data)
        if err != nil { return  }
        req, err = http.NewRequest(method, full_url, bytes.NewBuffer(jsonData))
      } else if nil != params {
        req, err = http.NewRequest(method, full_url, strings.NewReader(params.Encode()))
      } else {
        req, err = http.NewRequest(method, full_url, nil)
      }
      if err != nil {
        log.Println(err)
        return
      }
      req.Header.Set("Content-Type", "application/json")
    default:
      log.Println("invalid request method")
      return
  }

  if err != nil {
    log.Println(err)
    return
  }

  if nil != header  {
    for key, value := range header  {
      req.Header.Add(key, value)
    }
  }

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
  client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	str = string(respBody)
  log.Println(str)
  log.Println(resp.StatusCode)
	byt := []byte(str)

	var dat interface{}
	if err = json.Unmarshal(byt, &dat); err != nil {
		log.Println(str)
		log.Println(err)
		return
	}

  ret = dat
	return
}
