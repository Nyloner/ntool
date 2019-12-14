package ocr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Nyloner/ntool/common/utils"
	"github.com/Nyloner/ntool/logs"
)

var (
	apiKey      = "GtR0BKqQ5HWXbGeeeyuYdLnt"
	secretKey   = "pRqDZEfX4chYRzsuWNgRvBgfG4t86oGV"
	accessToken = ""
	apiUrl      = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
)

type tokenItem struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type respBody struct {
	ErrCode        int64 `json:"error_code"`
	WordsResultNum int64 `json:"words_result_num"`
	WordsResult    []struct {
		Words string `json:"words"`
	} `json:"words_result"`
}

func InitBaiDuOCR() error {
	if accessToken != "" {
		return nil
	}
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", apiKey, secretKey)
	resp, err := utils.GET(url)
	if err != nil {
		logs.Error("InitBaiDuOCR fail [url=%s err=%v]", url, err)
		return err
	}
	item := tokenItem{}
	if err := json.Unmarshal(resp.Content(), &item); err != nil {
		logs.Error("InitBaiDuOCR parse resp body fail [url=%s err=%v content=%#v]", url, err, string(resp.Content()))
		return err
	}
	accessToken = item.AccessToken
	logs.Info("InitBaiDuOCR success.")
	return nil
}

func OCRImage(imageUrl string) (text string, err error) {
	params := url.Values{
		"url": {imageUrl},
	}
	resp, err := http.PostForm(fmt.Sprintf("%s?access_token=%s", apiUrl, accessToken), params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	respBody := respBody{}
	if err := json.Unmarshal(content, &respBody); err != nil {
		logs.Error("OCRImage fail.[content]=%#v", content)
		return "", err
	}
	if respBody.ErrCode != 0 {
		logs.Error("OCRImage fail.[err_msg]=%#v", string(content))
		return "", fmt.Errorf("OCRImage fail.[err_msg]=%#v", string(content))
	}
	text = " "
	for _, words := range respBody.WordsResult {
		text += words.Words + " "
	}
	return text, nil
}
