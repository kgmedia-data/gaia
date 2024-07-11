package handler

type IJob interface {
	Run() error
}
