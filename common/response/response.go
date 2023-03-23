/*******
* @Author:qingmeng
* @Description:
* @File:response.go
* @Date:2023/3/16
 */

package response

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"tik-tok-brief/common/errorx"
)

//type Body struct {
//	Code int         `json:"code"`
//	Msg  string      `json:"msg"`
//	Data interface{} `json:"data,omitempty"`
//}


//http 参数错误返回
func ParseErr(c context.Context,w http.ResponseWriter, err error){
	httpx.WriteJsonCtx(c,w,http.StatusBadRequest,errorx.NewParamErr(err.Error()).ResponseData())
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var errResponse errorx.CodeErrorResponse
	if err != nil {
		switch e := err.(type) {
		case errorx.CodeError: // API部分的自定义错误
			errResponse=*e.ResponseData()
		default:
			//给前端统一返回的内部错误
			errResponse.StatusCode= int32(errorx.InternalErrCode)
			errResponse.StatusMsg=errorx.ERRINTERNAL

			//处理rpc错误
			if grpcErr,ok:=status.FromError(err);ok{
				errResponse.StatusCode= int32(errorx.RPCErrCode)
				//rpc传过来的自定义错误
				if uint32(grpcErr.Code())==errorx.ParamErrCode{
					errResponse.StatusCode= int32(errorx.ParamErrCode)
					errResponse.StatusMsg=grpcErr.Message()
				}
			}
		}
		httpx.OkJson(w, errResponse)
		return
	}
	httpx.OkJson(w,resp)
}