package test

import "fmt"

type ISubjectProxy interface {
	Process()
}

type RealSubject struct {

}

func (this *RealSubject)Process()  {
	fmt.Println("RealSubject 接口")
}

type ProxySubject struct {
	realSubject RealSubject
}

func (this *ProxySubject)Process()  {
	fmt.Println("Proxy 接口")
	this.realSubject.Process()
}

type ClientApp struct {

}

func (this *ClientApp)ProcessProxy()  {
	fmt.Println("客户端client")
	isubject := &ProxySubject{}
	isubject.Process()
}