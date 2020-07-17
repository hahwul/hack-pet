package merging

import (
	"os"
	"io/ioutil"
	"fmt"
	"github.com/pelletier/go-toml"
)

type PetSnippet struct {
	Description string `toml:"description"`
	Command     string `toml:"command"`
	Output      string `toml:"output"`
	Tag         []string `toml:"tag"`
}

type Lists struct {
	Snippet []PetSnippet `toml:"snippets"`
}

// MakeReadme is head + auto write body + foot
func MakeReadme(){
	hackpetToml, err := os.Open("hackpet.toml")
	check(err)
	toml_data, _ := ioutil.ReadAll(hackpetToml)

	snippets := Lists{}

	toml.Unmarshal(toml_data,&snippets)
	readme := "| Description | Command |\n| ----------- | ------- |\n"

	for _,v := range snippets.Snippet {
		readme = readme + "| "+v.Description+" | "+v.Command+" |\n"
	}


	fmt.Println(readme)
	top, err := os.Open("template/head.md")
	check(err)
	head_data, _ := ioutil.ReadAll(top)

	foot, err := os.Open("template/foot.md")
	check(err)
	foot_data, _ := ioutil.ReadAll(foot)


	_=head_data
	_=foot_data
	body := string(head_data) + readme + string(foot_data)
	fmt.Println("======================result====================")
	fmt.Println(body)
	file, err := os.OpenFile(
		"README.md",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,

		os.FileMode(0644))
	check(err)
	defer file.Close()
	_, err = file.Write([]byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}
}
