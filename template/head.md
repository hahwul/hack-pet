<h1 align="center">
  <br>
  <a href=""><img src="https://user-images.githubusercontent.com/13212227/87844645-fd483080-c8f9-11ea-949b-006305de5ac4.png"></a>
  <br>
  <img src="https://img.shields.io/github/languages/top/hahwul/hack-pet?style=flat-square"> <img src="https://api.codacy.com/project/badge/Grade/c3f297eba09342e8b9156c2bc8cbbfbb"> <a href="https://goreportcard.com/report/github.com/hahwul/hack-pet"><img src="https://goreportcard.com/badge/github.com/hahwul/hack-pet"></a> <img src="https://img.shields.io/github/issues-closed/hahwul/hack-pet?style=flat-square"> 
<a href="https://twitter.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/twitter/follow/hahwul?style=flat-square"></a>
</h1>

## hack-pet
hack-pet is collection of command snippets that are useful to hackers/bug bounty hunters. It is similar to the [recon_profile](https://github.com/nahamsec/recon_profile), but it uses the [pet](https://github.com/knqyf263/pet). pet can manage the command set more progressively.

### What is pet
Simple command-line snippet manager, written in Go => [pet](https://github.com/knqyf263/pet)

### How to use hack-pet
Add the snippet you like from the list below to the snippet toml file(`~/.config/pet/snippet.toml` or `$ pet edit`) in pet. To add all items, you can also paste `hackpet.toml`.

1) Copy & Paste hackpet.toml file to pet snippet file
```text
$ cat ./hackpet.toml >> ~/.config/pet/snippet.toml
```

2) Running pet
```text
$ pet exec

or 

$ pet search
```

3) You can find hackpet's snippets. The snippet of the hackpet has the tag `#hackpet`. If you have a lot of snippets in use, you can find them quickly by searching for tags.


## Screenshot
<img src="https://user-images.githubusercontent.com/13212227/87844969-7a28d980-c8fd-11ea-9c08-c96230937b19.png" width="100%">

## Tree
```text
.
├── hackpet.toml => all snippets
├── snippets     => collection of snippets
```

## Snippets

