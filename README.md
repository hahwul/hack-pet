# hack-pet

## hack-pet
hack-pet is collection of command snippets that are useful to hackers/bug bounty hunters. It is similar to the [recon_profile](https://github.com/nahamsec/recon_profile), but because it uses the pet, it can manage the command set more progressively.

### What is pet?
Simple command-line snippet manager, written in Go => [pet](https://github.com/knqyf263/pet)

### How to use it?
Add the snippet you like from the list below to the snippet toml file(`~/.config/pet/snippet.toml` or `$ pet edit`) in pet. To add all items, you can also paste `hackpet.toml`.

## Tree
```
.
├── hackpet.toml => all snippets
├── snippets     => collection of snippets
```

## Snippets

| Description | Command |
| ----------- | ------- |
| Android set proxy | adb shell settings put global http_proxy <ip>:<port> |
| Android unset proxy | adb shell settings put global http_proxy :0 |
| get url with gau, included parameter | echo <domain> | ~/go/bin/gau | grep "=" | qsreplace -a  |

## Contribute
First, Make snippet DOML with `hack-pet add` 
```
$ hack-pet add
add called
[command]
>>> echo <domain> | ~/go/bin/gau | grep "=" | qsreplace -a
echo <domain> | ~/go/bin/gau | grep "=" | qsreplace -a

[desc]
>>> get url with gau, included parameter
get url with gau, included parameter

[toml filename | e.g nmap_full_scan.toml]
>>> get_url_param.toml
get_url_param.toml

[[snippets]]
command = "echo <domain> | ~/go/bin/gau | grep \"=\" | qsreplace -a "
description = "get url with gau, included parameter"
output = ""
tag = ["hackpet"]
```

Second, move your DOML file to `/snippet` directory
```
$ mv get_url_param.toml ./snippets/
```

Finaly, Send Pull Request! (your DOML file in `./snippets/`)

## Merge (for me)
```
$ hack-pet merge
$ git add hackpet.toml README.md
$ git commit -m "merge and distribute readme"
$ git push -u origin master 
```
