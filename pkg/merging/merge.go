package merging

import (
	"fmt"
	"os"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Merge is merging al toml file
func Merge(path string){
	// default path is snippets
	files , err := ioutil.ReadDir(path)
	check(err)
	t , err := os.Create("hackpet.toml")
	for _, f := range files {
		fmt.Println("# "+f.Name())
		data , err := ioutil.ReadFile(path+"/"+f.Name())
		check(err)
		fmt.Println(string(data))
		check(err)
		_, err = t.Write([]byte("# hackpet :: "+f.Name()+"\n"))
		_, err = t.Write([]byte(data))	
		MakeReadme()
	}
}

