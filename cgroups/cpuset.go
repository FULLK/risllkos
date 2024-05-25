package cgroups
import (
	"io/ioutil"
	"log"
	"path"
)	
type Cpuset struct{
	file string
}

func (c *Cpuset)Name() string{
	return c.file
}
func (c *Cpuset)Set(re *Resource,new_cgroup_path string) error{
	if re.Cpuset == "" {
		return nil
	}
	err:=ioutil.WriteFile(path.Join(new_cgroup_path,c.file),[]byte(re.Cpuset),0644)
	if err!=nil{
		log.Fatal(err)
	}
	return err
}