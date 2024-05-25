package contain
import(
	log "github.com/sirupsen/logrus"
	"os/exec"
	
	"os"
)
func end_overlays(rooturl string){
	unmount_overlays(rooturl)
	delete_upper_work_merged(rooturl)
}
func unmount_overlays(rooturl string){
	mntrul:=rooturl+"/merged"
	cmd:=exec.Command("umount",mntrul)
	if err:=cmd.Run();err!=nil{
		log.Infof("unmount merged error")
	}
}
func delete_upper_work_merged(rooturl string){
	upperurl:=rooturl+"/upper"
	mergedurl:=rooturl+"/merged"
	workurl:=rooturl+"/work"

	if err:=os.RemoveAll(upperurl);err!=nil{
		log.Infof("delete upper error")
	}
	if err:=os.RemoveAll(mergedurl);err!=nil{
		log.Infof("delete merged error")
	}
	if err:=os.RemoveAll(workurl);err!=nil{
		log.Infof("delete work error")
	}
}