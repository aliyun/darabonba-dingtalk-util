package client

import (
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
)

func Test_GetTimestamp(t *testing.T) {
	utils.AssertEqual(t, 13, len(tea.StringValue(GetTimestamp())))
}

func Test_HasError(t *testing.T) {
	res := map[string]interface{}{
		"error_response": "ok",
	}
	out := tea.BoolValue(HasError(res))
	utils.AssertEqual(t, true, out)

	delete(res, "error_response")
	out = tea.BoolValue(HasError(res))
	utils.AssertEqual(t, false, out)

	res["errcode"] = 0
	out = tea.BoolValue(HasError(res))
	utils.AssertEqual(t, false, out)

	res["errcode"] = 10
	out = tea.BoolValue(HasError(res))
	utils.AssertEqual(t, true, out)
}

func Test_ComputeSignature(t *testing.T) {
	res := ComputeSignature(tea.String("accesskeysecret"), tea.String("canonicalString"))
	utils.AssertEqual(t, "NK+EVoCsJHmdwTSe5lzCj5GavWyQBO12rQaUfRScfz0=", tea.StringValue(res))
}

func Test_GetCanonicalStringForIsv(t *testing.T) {
	res := GetCanonicalStringForIsv(tea.String("timestamp"), tea.String("suiteTicket"))
	utils.AssertEqual(t, "timestamp\nsuiteTicket", tea.StringValue(res))

	res = GetCanonicalStringForIsv(tea.String("timestamp"), nil)
	utils.AssertEqual(t, "timestamp", tea.StringValue(res))
}
