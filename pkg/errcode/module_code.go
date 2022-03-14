package errcode

var (
	ErrorDuplicatedUserName = NewError(20010001, "用户名重复")
	ErrorPassword           = NewError(20010002, "账号或密码错误")

	ErrorSendMessageFail = NewError(20020001, "发送消息失败")
	ErrorGetMessageFail  = NewError(20020002, "获取消息失败")
)
