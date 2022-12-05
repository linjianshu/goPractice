package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"time"

	"golang.org/x/net/context"
	"log"
	"src/4.2GoMicro/blueter"
)

//import "./blueter"

func main() {
	//使用etcd作为注册中心
	reg := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"http://47.244.249.198:2379"}
	})

	micro.NewService()

	//service := service.New(
	//	service.Name("blueter"),
	//	service.Version("lastest"),
	//	service.RegisterTTL(time.Second*30),
	//	service.RegisterInterval(time.Second*10),
	//
	//)

	service := micro.NewService(
		micro.Name("blueter"),
		micro.Version("laster"),
		micro.RegisterTTL(time.Second*30),      //服务发现系统中的生存期
		micro.RegisterInterval(time.Second*10), //重新注册的间隔
		//设置了30s的TTL生存期 并设置每10s一次的重新注册
		micro.Registry(reg))

	/*
		服务通过服务发现功能 在启动时进行服务注册 关闭时进行服务卸载
		有时候这些服务可能会异常挂掉 进程可能会被暂停 可能遇到短暂的网络问题
		这种情况下 节点会在服务发现中被干掉 理想状态是服务被自动移除
		解决方案:
		为了解决这个问题 micro注册机制支持通过TTL time-to-live 和间隔时间注册两种方式
		ttl指定一次注册在注册中心的有效期 过期后便删除 而间隔时间则是定时向注册中心重新注册 以保证服务仍然在线
	*/

	service.Init()

	blueter.RegisterBlueterHandler(service.Server(), new(Blueter))
	/*
		这样就自动注册了一个服务 如果在100台计算机上运行此程序 就注册了100个服务
		客户端根据blueter请求就会自动负载均衡到这100台计算机中的某一台上 服务可以直接调用server.Run运行
		这会让服务监听一个随机端口 这个调用也会让服务将自身注册到注册器上 当服务停止运行时 它会在注册器上注销自己
	*/
	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}

type Blueter struct {
}

func (g *Blueter) Hello(ctx context.Context, req *blueter.HelloRequest, rsp *blueter.HelloResponse) error {
	rsp.Msg = "Hello dddddd " + req.From
	//这里可以做一些数据库新增修改查询操作把结果返回
	return nil
}
