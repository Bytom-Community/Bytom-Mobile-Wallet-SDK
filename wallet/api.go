package wallet

import (
	"github.com/bytom-community/mobile/wallet/pseudohsm"
	"github.com/bytom-community/mobile/wallet/errors"
)

const (
	// SUCCESS indicates the rpc calling is successful.
	SUCCESS = "success"
	// FAIL indicated the rpc calling is failed.
	FAIL = "fail"
)

// Response describes the response standard.
type Response struct {
	Status      string      `json:"status,omitempty"`
	Code        string      `json:"code,omitempty"`
	Msg         string      `json:"msg,omitempty"`
	ErrorDetail string      `json:"error_detail,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

//NewSuccessResponse success response
func NewSuccessResponse(data interface{}) Response {
	return Response{Status: SUCCESS, Data: data}
}

//NewErrorResponse error response
func NewErrorResponse(err error) Response {
	response := FormatErrResp(err)
	return response
}

//FormatErrResp format error response
func FormatErrResp(err error) (response Response) {
	response = Response{Status: FAIL}
	root := errors.Root(err)
	// Some types cannot be used as map keys, for example slices.
	// If an error's underlying type is one of these, don't panic.
	// Just treat it like any other missing entry.
	defer func() {
		if err := recover(); err != nil {
			response.ErrorDetail = ""
		}
	}()

	if info, ok := respErrFormatter[root]; ok {
		response.Code = info.ChainCode
		response.Msg = info.Message
		response.ErrorDetail = err.Error()
	} else {
		response.Code = respErrFormatter[ErrDefault].ChainCode
		response.Msg = respErrFormatter[ErrDefault].Message
		response.ErrorDetail = err.Error()
	}
	return response
}

// API is the scheduling center for server
type API struct {
	Hsm *pseudohsm.HSM
}
