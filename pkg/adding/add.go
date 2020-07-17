package adding

import (
	"github.com/pelletier/go-toml"
	"fmt"
	"os"
)

// PetSnippet is struct of toml​
type PetSnippet struct {
	Description string `toml:"description"`
	Command     string `toml:"command"`
	Output      string `toml:"output"`
	Tag         []string `toml:"tag"`
}
​
​
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Add is arguments to toml string and file
func Add(command, desc, filename string){
	data := PetSnippet{Description: desc, Command: command, Output: "", Tag: []string{"hackpet"}}
	b, err := toml.Marshal(data)
	check(err)
	buffer := "[[snippets]]\n"+string(b)+"\n"
	fmt.Println(buffer)
	f , err := os.Create(filename)
	check(err)
	_, err = f.Write([]byte(buffer))

