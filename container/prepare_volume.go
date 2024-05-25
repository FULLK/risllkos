package contain

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func prepare_volume(rooturl string,volume string){
	hostvolume,containervolume,err:=volume_extract(volume)
	if err!=nil{
		log.Info("volume_extract error")
	}
	mount_volume(rooturl,hostvolume,containervolume)
	log.Info("prepare_volume success")
}
func volume_extract(volume string) (hostvolume string,containervolume string,err error){
	if volume!=""{
		parts := strings.Split(volume, ":")
		if len(parts) == 2 {
			part1 := parts[0]  //主机上的文件路径
			part2 := parts[1] //容器上的文件路径
			// 使用 part1 和 part2
			log.Infof("get volume success")
			return part1,part2,nil
		} else {
			return "","",fmt.Errorf("invalid volume %s, must split by `:`", volume)
		}
	}else {
		return "","",fmt.Errorf("invalid volume %s, not exist`:`", volume)
	}

}
func mount_volume(rooturl string,hostvolume string,containervolume string) {
	mnturl:=rooturl+"/merged"
	if err:=os.Mkdir(hostvolume,0777);err!=nil{
		log.Infof("mkdir hostvolume error")
	}
	containervolume_inhost:=path.Join(mnturl,containervolume) // 得到容器上的文件在主机上的文件路径
	log.Info(rooturl,hostvolume,containervolume,containervolume_inhost)
	
	if err:=os.Mkdir(containervolume_inhost,0777);err!=nil{
		log.Infof("mkdir containervolume_inhost error")
	}
	cmd:=exec.Command("mount","-o","bind",hostvolume,containervolume_inhost)
	/**/
	cmd.Stdin=os.Stdin
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	if err:=cmd.Run();err!=nil{
		log.Infof("mount -o bind error")
	}

}