package contain
import(
	log "github.com/sirupsen/logrus"
	"os/exec"
	"syscall"
	"os"
)

func new_contain_process(it bool,volume string)  (*exec.Cmd,*os.File,string){	
	
	readpipe,writepipe,err:=os.Pipe()
	if err!=nil{
		log.Fatal(err)
	}
	command:=exec.Command("/proc/self/exe","init")
	command.SysProcAttr=&syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if it{
		command.Stdin=os.Stdin
		command.Stdout=os.Stdout
		command.Stderr=os.Stderr
	}
	log.Info(command)
	log.Infof("sucess create a contain process")
	command.ExtraFiles=[]*os.File{readpipe}
	mntURL := "/home/llk/Desktop/llkdocker/commit_docker/merged"
	rootURL := "/home/llk/Desktop/llkdocker/commit_docker"
	command.Dir = mntURL
	prepare_overlays(rootURL) 
	prepare_volume(rootURL,volume)

	log.Info(readpipe)
	return command,writepipe,rootURL
	
}
