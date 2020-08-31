// This file is auto-generated, don't edit it. Thanks.
/**
 * This is for DingUtil
 */
package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"hash"
	"io"
	"strconv"
	"time"

	"github.com/alibabacloud-go/tea/tea"
)

func GetTimestamp() (_result *string) {
	return tea.String(tea.ToString(time.Now().UnixNano() / 1000000))
}

func HasError(res map[string]interface{}) (_result *bool) {
	if res["error_response"] != nil {
		return tea.Bool(true)
	}

	if res["errcode"] == nil {
		return tea.Bool(false)
	}

	byt, _ := json.Marshal(res["errcode"])
	num, err := strconv.Atoi(string(byt))
	if err == nil && num == 0 {
		return tea.Bool(false)
	}
	return tea.Bool(true)
}

func ComputeSignature(accessSecret *string, canonicalString *string) (_result *string) {
	h := hmac.New(func() hash.Hash { return sha256.New() }, []byte(tea.StringValue(accessSecret)))
	io.WriteString(h, tea.StringValue(canonicalString))
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return tea.String(signedStr)
}

func GetCanonicalStringForIsv(timestamp *string, suiteTicket *string) (_result *string) {
	res := tea.StringValue(timestamp)
	if suiteTicket != nil {
		res += "\n" + tea.StringValue(suiteTicket)
	}

	return tea.String(res)
}
