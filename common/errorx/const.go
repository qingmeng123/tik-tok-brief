/*******
* @Author:qingmeng
* @Description:
* @File:errCode
* @Date:2023/3/22
 */

package errorx

// 代码
const (
	OK              uint32 = 0
	defaultErrCode  uint32 = 100001
	ParamErrCode    uint32 = 100002
	InternalErrCode uint32 = 100003
	DBErrCode       uint32 = 100004
	RPCErrCode      uint32 = 100005
)

// 错误信息
const (
	SUCCESS         = "success"
	ERRUSERID       = "用户ID错误"
	ERRUSERNAME     = "用户名错误"
	ERRUSERPASSWORD = "用户密码错误"
	ERRTOKEN        = "用户token错误"
	ERRFILEPARAM    = "文件格式不对"
	ERRFILEUPLOAD   = "上传文件出错啦"
	ERRFOLLOWUSER   = "已关注过该用户"
	ERRUNFOLLOWUSER = "未关注该用户"
	ERRLIKE         = "赞操作错误"
	ERRCOMMENT      = "评论操作错误"
	ERRTIMEPARAM    = "时间转化错误"

	ERRDB       = "数据库繁忙,请稍后再试"
	ERRINTERNAL = "服务器开小差了，请稍后再试"
)
