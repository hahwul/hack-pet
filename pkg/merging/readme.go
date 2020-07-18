package merging

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
	"strings"
)

// PetSnippet is struct of pet
type PetSnippet struct {
	Description string   `toml:"description"`
	Command     string   `toml:"command"`
	Output      string   `toml:"output"`
	Tag         []string `toml:"tag"`
}

// Lists is struct of Petsnippet
type Lists struct {
	Snippet []PetSnippet `toml:"snippets"`
}

// MakeReadme is head + auto write body + foot
func MakeReadme() {
	hackpetToml, err := os.Open("hackpet.toml")
	check(err)
	tomlData, _ := ioutil.ReadAll(hackpetToml)

	snippets := Lists{}

	toml.Unmarshal(tomlData, &snippets)
	readme := "| Description | Command |\n| ----------- | ------- |\n"

	for _, v := range snippets.Snippet {
		desc := strings.Replace(v.Description, "|", "\\|", -1)
		comd := strings.Replace(v.Command, "|", "\\|", -1)
		readme = readme + "| " + desc + " | `" + comd + "` |\n"
	}

	fmt.Println(readme)
	top, err := os.Open("template/head.md")
	check(err)
	headData, _ := ioutil.ReadAll(top)

	foot, err := os.Open("template/foot.md")
	check(err)
	footData, _ := ioutil.ReadAll(foot)

	_ = headData
	_ = footData
	body := string(headData) + readme + string(footData)
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
