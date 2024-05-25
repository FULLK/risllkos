package contain

import (
	"os"
	"os/exec"
	"path"
	log "github.com/sirupsen/logrus"
)
func end_volume(rooturl string,volume string){
	hostvolume,containervolume,err:=volume_extract(volume)
	if err!=nil{
		log.Infof("volume_extract error")
	}
	unmount_volume(rooturl,hostvolume,containervolume)
	
}

func unmount_volume(rooturl string,hostvolume string,containervolume string) {
	mnturl:=rooturl+"/merged"
	containervolume_inhost:=path.Join(mnturl,containervolume)
	cmd:=exec.Command("umount",containervolume_inhost)
	log.Info(containervolume_inhost)
	/**/
	cmd.Stdin=os.Stdin
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	if err:=cmd.Run();err!=nil{
		log.Infof("unmount containervolume_inhost error")
	}
	
	
}