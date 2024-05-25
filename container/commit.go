package contain

import (

	"os/exec"

	log "github.com/sirupsen/logrus"
)
func Contain_commit(imagename []string){
	mnturl:="/home/llk/Desktop/llkdocker/commit_docker/merged"
	rooturl:="/home/llk/Desktop/llkdocker/commit_docker"
	/*
	var params string
	for i := 0; i < len(imagename); i++ {
	param := rooturl+"/"+imagename[i]+".tar "
	params =params+param   
	}	 //得到参数
	log.Info(params) //类似 路径/1.tar 路径/2.tar 但一次只能生成一个镜像
	*/
	imagetar:=rooturl+"/"+imagename[0]+".tar"
	log.Info(imagetar)

	if len(imagename)>1{
		log.Infof("too many image name ,we only can tar the first")

	}
	if _,err:=exec.Command("tar","-czf",imagetar,"-C",mnturl,".").CombinedOutput();err!=nil{
		log.Info(err)
		log.Fatal("tar error !!!")  //好像一次只能生成一个tar包 
	}
	log.Info("commit finish") //tar需要一定时间，这里做最后结束的回现

}
