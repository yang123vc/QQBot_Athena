package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Trans struct {
	Sentences []struct {
		Trans   string `json:"trans"`
		Orig    string `json:"orig"`
		Backend int    `json:"backend"`
	} `json:"sentences"`
	Src      string `json:"src"`
	LdResult struct {
		Srclangs            []string  `json:"srclangs"`
		SrclangsConfidences []float64 `json:"srclangs_confidences"`
		ExtendedSrclangs    []string  `json:"extended_srclangs"`
	} `json:"ld_result"`
}

const (
	CNtoEN int = 0
	ANtoCN int = 1
)

func translate(origStr string, mode int) (sendStr string) {
	var trans Trans
	var resp *http.Response
	var err error
	switch mode {
	case CNtoEN:
		resp, err = http.Get("http://translate.google.cn/translate_a/single?client=gtx&dt=t&dj=1&ie=UTF-8&sl=auto&tl=en&q=" + origStr)
	case ANtoCN:
		resp, err = http.Get("http://translate.google.cn/translate_a/single?client=gtx&dt=t&dj=1&ie=UTF-8&sl=auto&tl=zh_CN&q=" + origStr)
	}
	//resp, err := http.Get("http://translate.google.cn/translate_a/single?client=gtx&dt=t&dj=1&ie=UTF-8&sl=auto&tl=en&q=" + origStr)
	if err != nil {
		return "网络不通"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "解析错误"
	}
	json.Unmarshal([]byte(string(body)), &trans)

	sendStr = trans.Sentences[0].Trans

	return
}
