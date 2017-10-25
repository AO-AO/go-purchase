package receipt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

var googleTokenPath = "https://accounts.google.com/o/oauth2/token"

func refreshToken(iapConfig IAPConfig) (googleTokenRes, error) {
	var tokenRes googleTokenRes
	config := map[string]string{
		"GoogleClientID":     iapConfig.GoogleClientID,
		"GoogleClientSecret": iapConfig.GoogleClientSecret,
		"GoogleRefToken":     iapConfig.GoogleRefToken,
	}

	urlData := url.Values{}
	urlData.Set("grant_type", "refresh_token")
	for k, v := range config {
		urlData.Set(k, v)
	}

	response, err := http.PostForm(googleTokenPath, urlData)
	if err != nil {
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
	packageNameURL, err := url.ParseRequestURI(receiptReq.PackageName)
	productIDURL, err := url.ParseRequestURI(receiptReq.ProductID)
	purchaseTokenURL, err := url.ParseRequestURI(receiptReq.PurchaseToken)
	accessTokenURL, err := url.ParseRequestURI(iapConfig.GooglePublicKeyStrLive)
	if err != nil {
		return validateResult, err
	}
	if receiptReq.AutoRenewing { // 订阅模式

		checkURL = "https://www.googleapis.com/androidpublisher/v2/applications/" +
			packageNameURL.String() +
			"/purchases/subscriptions/" +
			productIDURL.String() +
			"/tokens/" + purchaseTokenURL.String() +
			"?access_token=" + accessTokenURL.String()
	} else { // 普通内建购买
		checkURL = "https://www.googleapis.com/androidpublisher/v2/applications/" +
			packageNameURL.String() +
			"/purchases/products/" +
			productIDURL.String() +
			"/tokens/" + purchaseTokenURL.String() +
			"?access_token=" + accessTokenURL.String()
	}

	response, err := http.Get(checkURL)
	if err != nil {
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
