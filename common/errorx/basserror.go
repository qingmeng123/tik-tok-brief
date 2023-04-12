/*******
* @Author:qingmeng
* @Description:
* @File:basserror
* @Date:2023/3/16
 */

package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CodeError 自定义错误
type CodeError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

// CodeErrorResponse 自定义的错误响应
type CodeErrorResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// Error CodeError实现error接口
func (e CodeError) Error() string {
	return e.Msg
}

// ResponseData 返回自定义类型的错误响应
func (e CodeError) ResponseData() *CodeErrorResponse {
	return &CodeErrorResponse{
		StatusCode: int32(e.Code),
		StatusMsg:  e.Msg,
	}
}

//-----------返回给前端的错误--------------

// NewParamErr 参数错误
func NewParamErr(msg string) CodeError {
	return CodeError{
		Code: ParamErrCode,
		Msg:  msg,
	}
}

// NewInternalErr 返回内部错误
func NewInternalErr() CodeError {
	return CodeError{
		Code: InternalErrCode,
		Msg:  ERRINTERNAL,
	}
}

//-------------rpc服务相互调用返回的错误--------------

// NewStatusParamErr rpc返回的参数错误
func NewStatusParamErr(msg string) error {
	return status.Error(codes.Code(ParamErrCode), msg)
}

// NewStatusDBErr rpc返回的数据库错误
func NewStatusDBErr() error {
	return status.Error(codes.Code(DBErrCode), ERRDB)
}

func NewStatusTxErr() error {
	return status.Error(codes.Aborted, ERRINTERNAL)
}

func NewStatusParamTxErr(msg string) error {
	return status.Error(codes.Aborted, msg)
}
