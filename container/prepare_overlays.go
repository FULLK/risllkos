package contain
import(
	log "github.com/sirupsen/logrus"
	"os/exec"
	
	"os"
)

func prepare_overlays(rooturl string){
	create_lower(rooturl)
	create_upper(rooturl)
	create_merged(rooturl)
	create_work(rooturl)
	mount_overlays(rooturl)
	log.Infof("prepare_overlays suceess")
}
func create_lower(rooturl string){
	busyboxurl:=rooturl+"/busybox"
	busytarurl:=rooturl+"/busybox.tar"
	_, err := os.Stat(busyboxurl)
    if err != nil {
        if os.IsNotExist(err) {
            log.Infof("文件不存在")
			if err:=os.Mkdir(busyboxurl,0777);err!=nil{
				log.Infof("mkdir error")
			}
			if _,err=exec.Command("tar", "-xvf", busytarurl, "-C", busyboxurl).CombinedOutput(); err!=nil{
				log.Infof("tar busybox.tar error")
			}	
        } else {
			log.Infof("无法获取文件信息: %v\n", err)
        }
    } else{
		log.Infof("busybox exist")
	}
	
}
func create_upper(rooturl string){
	upperurl:=rooturl+"/upper"
	if err:=os.Mkdir(upperurl,0777);err!=nil{
		if os.IsExist(err) {
        log.Infof("Directory %s already exists.", upperurl)
    } else {
        log.Infof("Mkdir upper error: %v", err)
    }
	}
}

func create_merged(rooturl string){
	mergedurl:=rooturl+"/merged"
	if err:=os.Mkdir(mergedurl,0777);err!=nil{
		if os.IsExist(err) {
			log.Infof("Directory %s already exists.", mergedurl)
		} else {
			log.Infof("Mkdir merged error: %v", err)
		}
	}
}
func create_work(rooturl string){
	workurl:=rooturl+"/work"
	if err:=os.Mkdir(workurl,0777);err!=nil{
		if os.IsExist(err) {
			log.Infof("Directory %s already exists.", workurl)
		} else {
			log.Infof("Mkdir work error: %v", err)
		}
	}
}

func mount_overlays(rooturl string){
	mnturl:=rooturl+"/merged"
	dirs:="lowerdir="+rooturl+"/busybox"+",upperdir="+rooturl+"/upper"+",workdir="+rooturl+"/work"
	cmd:=exec.Command("mount","-t","overlay","overlay","-o",dirs,mnturl)//sudo 权限才能运行成功
	if err:=cmd.Run();err!=nil{
		log.Infof("mount overlay error")
		log.Error(err)
	}
}