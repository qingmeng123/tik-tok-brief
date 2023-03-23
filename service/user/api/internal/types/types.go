// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id              int64  `json:"id"`
	Username        string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar,optional"`
	BackgroundImage string `json:"background_image,optional"`
	Signature       string `json:"Signature,optional"`
	TotalFavorited  int64  `json:"total_favorited,optional"`
	WorkCount       int64  `json:"work_count,optional"`
	FavoriteCount   int64  `json:"favorite_count,optional"`
}

type StatusResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,optional,"`
}

type RegisterRequest struct {
	Username string `form:"username" validate:"required,max=32,min=1"`
	Password string `form:"password" validate:"required,max=32,min=1"`
}

type RegisterResponse struct {
	StatusResponse
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type LoginReq struct {
	Username string `form:"username" validate:"required,max=32,min=1"`
	Password string `form:"password" validate:"required,max=32,min=1"`
}

type LoginResp struct {
	StatusResponse
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserInfoReq struct {
	UserId int64  `form:"user_id" validate:"required,gte=0"`
	Token  string `form:"token"`
}

type UserInfoResp struct {
	StatusResponse
	User User `json:"user,omitempty"`
}
