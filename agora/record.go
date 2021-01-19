// @Description  云端录制
// @Author  	 jiangyang  
// @Created  	 2021/1/5 4:34 下午
package agora

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	appID           = "d330164bbd86411d96d5684962bd0d14"
	GET_RESOURCE_ID = "https://api.agora.io/v1/apps/%s/cloud_recording/acquire"
)

type AcquireReq struct {
	Cname         string `json:"cname"`
	UID           string `json:"uid"`
	ClientRequest struct {
		ResourceExpiredHour int `json:"resourceExpiredHour"`
	} `json:"clientRequest"`
}

type AcquireResp struct {
	ResourceId string `json:"resourceId"`
}

func GenerateCredential() string {

	customerId := "c1e90ddcf1e148689a1b1f6dd908ffec"
	customerSecret := "a2622a8f03bb49779209026396bf2408"

	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Basic %s:%s", customerId, customerSecret)))
}

// 获取云端录制资源
func Acquire(acquireReq *AcquireReq) (*AcquireResp, error) {
	acquireResp:=&AcquireResp{}
	if err := Http("POST", fmt.Sprintf(GET_RESOURCE_ID, appID), acquireReq, acquireResp); err != nil {
		return nil, err
	}
	return acquireResp,nil
}

func Http(method string, url string, body interface{}, data interface{}) error {
	client := http.Client{}

	marshal, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/json;charset=utf-8")
	req.Header.Set("Authorization", GenerateCredential())
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(all, data)

}
