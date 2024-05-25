package cgroups
import (
	"io/ioutil"
	"log"
	"path"
)	
type memory struct{
	file string
}

func (c *memory)Name() string{
	return c.file
}
func (c *memory)Set(re *Resource,new_cgroup_path string) error{
	if re.Mem == "" {
		return nil
	}
	err:=ioutil.WriteFile(path.Join(new_cgroup_path,c.file),[]byte(re.Mem),0644)
	if err!=nil{
		log.Fatal(err)
	}
	return err
}