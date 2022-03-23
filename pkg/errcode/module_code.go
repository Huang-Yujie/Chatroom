package errcode

var (
	ErrorDuplicatedUserName = NewError(20010001, "用户名重复")
	ErrorUserNameNotFound   = NewError(20010002, "用户名不存在")
	ErrorUserIDNotFound     = NewError(20010003, "用户ID不存在")
	ErrorPassword           = NewError(20010004, "账号或密码错误")
	ErrorRegisterFailure    = NewError(20010005, "注册失败")
	ErrorLoginFailure       = NewError(20010006, "登录失败")

	ErrorSendMessageFail = NewError(20020001, "发送消息失败")
	ErrorGetMessageFail  = NewError(20020002, "获取消息失败")
)
