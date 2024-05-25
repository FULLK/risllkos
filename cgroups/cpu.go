package cgroups
import (
	"io/ioutil"
	"path"
	"strconv"
	log"github.com/sirupsen/logrus"
)	
type Cpu struct{
	file string
}
const(
	max=100000
	multiple=1000
)

func (c *Cpu)Name() string{
	return c.file
}
func (c *Cpu)Set(re *Resource,new_cgroup_path string) error{
	if re.Cpu==""{
		return nil;
	}

	cpu,err:=(strconv.Atoi(re.Cpu))
	log.Info(re.Cpu)
	log.Info(cpu*multiple/max)
	
	err=ioutil.WriteFile(path.Join(new_cgroup_path,c.file),[]byte( strconv.Itoa( (cpu*multiple)) ),0644)
	log.Info([]byte( strconv.Itoa( (cpu*multiple)/max) ))
	if err!=nil{
		log.Fatal(err)
	}
	return err
}