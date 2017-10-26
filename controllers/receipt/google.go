package receipt

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

var googleTokenPath = "https://accounts.google.com/o/oauth2/token"

func refreshToken(iapConfig IAPConfig) (googleTokenRes, error) {
	var tokenRes googleTokenRes
	config := map[string]string{
		"client_id":     iapConfig.GoogleClientID,
		"client_secret": iapConfig.GoogleClientSecret,
		"refresh_token": iapConfig.GoogleRefToken,
		"grant_type":    "refresh_token",
	}

	urlData := url.Values{}
	for k, v := range config {
		urlData.Set(k, v)
	}

	response, err := http.PostForm(googleTokenPath, urlData)
	if err != nil {
		return tokenRes, err
	}
	if response.StatusCode >= 400 {
		err = errors.New(response.Status)
		return tokenRes, err
	}
	defer response.Body.Close()
	resultBuf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tokenRes, err
	}

	err = json.Unmarshal(resultBuf, &tokenRes)
	if err != nil {
		return tokenRes, err
	}
	return tokenRes, nil
}

func validateGoogle(receipt string, iapConfig IAPConfig) (googleValidateRes, error) {
	var receiptReq googleReceiptReq
	var validateResult googleValidateRes
	err := json.Unmarshal([]byte(receipt), &receiptReq)
	if err != nil {
		return validateResult, err
	}

	// 组织checkURL
	var checkURL string
	packageNameURL := receiptReq.PackageName
	productIDURL := receiptReq.ProductID
	purchaseTokenURL := receiptReq.PurchaseToken
	accessToken, refErr := refreshToken(iapConfig)
	if refErr != nil {
		return validateResult, refErr
	}
	accessTokenURL := accessToken.AccessToken
	if receiptReq.AutoRenewing { // 订阅模式

		checkURL = "https://www.googleapis.com/androidpublisher/v2/applications/" +
			packageNameURL +
			"/purchases/subscriptions/" +
			productIDURL +
			"/tokens/" + purchaseTokenURL +
			"?access_token=" + accessTokenURL
	} else { // 普通内建购买
		checkURL = "https://www.googleapis.com/androidpublisher/v2/applications/" +
			packageNameURL +
			"/purchases/products/" +
			productIDURL +
			"/tokens/" + purchaseTokenURL +
			"?access_token=" + accessTokenURL
	}

	response, err := http.Get(checkURL)
	if err != nil {
		return validateResult, err
	}
	if response.StatusCode >= 400 {
		err = errors.New(response.Status)
		return validateResult, err
	}
	defer response.Body.Close()

	resultBuf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return validateResult, err
	}

	json.Unmarshal(resultBuf, &validateResult)
	return validateResult, err
}
