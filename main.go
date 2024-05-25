package main

import(
	"os"
	log"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)
func main(){
	//定义相关命令
	app:=cli.NewApp()
	app.Name="llkdocker"
	app.Usage="my simple docker -llkdocker "
	app.Commands=[]cli.Command{
		runcommand,
		initcommand,
		commitcommand,
	}//Commands 属性是一个 []cli.Command 类型的切片

	app.Before=func(context *cli.Context)error{

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return nil
	} //在处理命令参数之前先进行的函数

	if err:=app.Run(os.Args);err!=nil{
		log.Fatal(err)
	}
}