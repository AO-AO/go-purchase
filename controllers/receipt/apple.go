package receipt

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 沙盒模式校验地址
var appleSandBoxHost = "sandbox.itunes.apple.com"
var prodHost = "buy.itunes.apple.com"
var applePath = "/verifyReceipt"

func validateApple(receipt string, sandboxMode bool, iapConfig IAPConfig) (appleValidateRes, error) {
	var checkURL string
	if sandboxMode {
		checkURL = appleSandBoxHost + applePath
	} else {
		checkURL = prodHost + applePath
	}
	checkURL = "https://" + checkURL
	applePassword := iapConfig.ApplePassword
	config := appleConfig{
		ReceiptData: receipt,
		Password:    applePassword,
	}
	result := appleValidateRes{}

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return result, err
	}

	ioBody := bytes.NewReader([]byte(jsonConfig))
	response, err := http.Post(checkURL, "application/json", ioBody)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	resultBuf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	json.Unmarshal(resultBuf, &result)
	return result, err
}
