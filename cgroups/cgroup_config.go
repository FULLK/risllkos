package cgroups

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	log"github.com/sirupsen/logrus"
)	

type Resource struct{
	Cpu string
	Cpuset string
	Mem  string
}

type Subsystem interface{
	Name() string 
	Set(re *Resource,new_cgroup_path string) error
}

var Subsystemins=[]Subsystem{
	&Cpu{file:"cpu.max"},
	&Cpuset{file: "cpuset.cpus"},
	&memory{file: "memory.max"},
}
type Cgroups struct{
	Cgroups_Name string
	Resour *Resource
	Sub []Subsystem
}


func (c *Cgroups)Name() string{
	return c.Cgroups_Name
}
func (c *Cgroups)Move(pid int,cgroups_path string){
	if err := ioutil.WriteFile(path.Join(cgroups_path, "cgroup.procs"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		log.Info("move fail")
		log.Fatal(err)
	}
}
func (c *Cgroups)Set(cgroups_path string){
	for _ ,sub:=range c.Sub{
		log.Info(sub)
		sub.Set(c.Resour,cgroups_path)
	}

}

func (c *Cgroups)Remove(cgroups_path string){
	err:=os.RemoveAll(cgroups_path)
	if err!=nil{
		log.Fatal(err)
	}
}