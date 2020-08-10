package service

type NotifyService interface {
	UserRegister()
	UserRestore()
}

type notifyService struct {}