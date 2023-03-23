/*******
* @Author:qingmeng
* @Description:
* @File:errCode
* @Date:2023/3/22
 */

package errorx


// 代码
const (
	OK 				uint32=0
	defaultErrCode 	uint32= 100001
	ParamErrCode	 uint32= 100002
	InternalErrCode uint32= 100003
	DBErrCode       uint32= 100004
	RPCErrCode      uint32= 100005
)

//	错误信息
const (
	SUCCESS="success"
	ERRUSERID = "用户ID错误"
	ERRUSERNAME="用户名错误"
	ERRUSERPASSWORD="用户密码错误"
	ERRTOKEN="用户token错误"
	ERRDB="数据库繁忙,请稍后再试"
	ERRINTERNAL="服务器开小差了，请稍后再试"
)
