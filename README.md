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
| Android set proxy | `adb shell settings put global http_proxy <ip>:<port>` |
| Android unset proxy | `adb shell settings put global http_proxy :0` |
| Find subdomain and takeover (with subfinder/amass/assetfinder/subjack) | `subfinder -d <domain> >> domains ; assetfinder -subs-only <domain> >> domains ; amass enum -norecursive -noalts -d <domain> >> domains ; subjack -w domains -t 100 -timeout 30 -ssl -c ~/go/src/github.com/haccer/subjack/fingerprints.json -v  \| tee takeover` |
| Get url with gau, included parameter | `echo <domain> \| ~/go/bin/gau \| grep "=" \| qsreplace -a ` |
| Get urls from urlscanio | `gron "https://urlscan.io/api/v1/search/?q=domain:<domain>"  \| grep 'url' \| gron --ungron` |
| Create a wordlist using param used in the domain | `waybackurls <domain> \|  grep "?" \| unfurl keys  \| sort -u \| tee -a paramlist.txt` |

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

or 

Write TOML Code in `./snippets` directory
```
[[snippets]]
command = "echo <domain> | ~/go/bin/gau | grep \"=\" | qsreplace -a "
description = "Get url with gau, included parameter"
output = ""
tag = ["hackpet"]

```
Please attach a `hackpet` to the tag. This allows you to distinguish between different snippets and hackpet.


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
