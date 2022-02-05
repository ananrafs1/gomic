package model


type IProcessor interface {
	OnStart()
	OnFinished()
}