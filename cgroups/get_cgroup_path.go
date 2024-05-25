package cgroups

import (
	"bufio"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

const mountPointIndex = 4

func Find_cgroup_path(cgroups_name string ) string{
	mountinfo,err:=os.Open("/proc/self/mountinfo")
	if err!=nil{
		log.Fatal(err)
	}
	defer mountinfo.Close()
	scanner:=bufio.NewScanner(mountinfo);
	for scanner.Scan(){
		text:=scanner.Text()
		split_text:=strings.Split(text," ")
		isfind:=strings.Contains(split_text[mountPointIndex],cgroups_name)
		log.Info(split_text)
		log.Info(split_text[mountPointIndex])
		log.Info(cgroups_name)
		if isfind {
			return split_text[mountPointIndex]
		} else{
			log.Error("error hppen get the mount path")
			
		}
	}
	return " "
}


func Get_cgroups_path(cgroups_name string , new_cgroup_name string) string{

	cgroups_path:=Find_cgroup_path(cgroups_name)
	new_cgroup_path:=path.Join(cgroups_path,new_cgroup_name)
	log.Info(new_cgroup_path)
	_,err:=os.Stat(new_cgroup_path)
	log.Info("os.stat")
	if err!=nil &&os.IsNotExist(err){
		err:=os.Mkdir(new_cgroup_path,0755)
		if err!=nil{
			log.Info("os.mkdir fail ")
		}
		
	}
	
	return new_cgroup_path
}