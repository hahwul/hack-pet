<h1 align="center">
  <br>
  <a href=""><img src="https://user-images.githubusercontent.com/13212227/87844645-fd483080-c8f9-11ea-949b-006305de5ac4.png"></a>
  <br>
  <img src="https://img.shields.io/github/languages/top/hahwul/hack-pet?style=flat-square"> <img src="https://api.codacy.com/project/badge/Grade/17cac7b8d1e849a688577f2bbdd6ecd0"> <a href="https://goreportcard.com/report/github.com/hahwul/hack-pet"><img src="https://goreportcard.com/badge/github.com/hahwul/hack-pet"></a> <img src="https://img.shields.io/github/issues-closed/hahwul/hack-pet?style=flat-square"> 
<a href="https://twitter.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/twitter/follow/hahwul?style=flat-square"></a>
</h1>

## hack-pet
hack-pet is collection of command snippets that are useful to hackers/bug bounty hunters. It is similar to the [recon_profile](https://github.com/nahamsec/recon_profile), but it uses the [pet](https://github.com/knqyf263/pet). pet can manage the command set more progressively.

### What is pet?
Simple command-line snippet manager, written in Go => [pet](https://github.com/knqyf263/pet)

### How to use it?
Add the snippet you like from the list below to the snippet toml file(`~/.config/pet/snippet.toml` or `$ pet edit`) in pet. To add all items, you can also paste `hackpet.toml`.

1) Copy & Paste hackpet.toml file to pet snippet file
```
$ cat ./hackpet.toml >> ~/.config/pet/snippet.toml
```

2) Running pet
```
$ pet exec

or 

$ pet search
```

3) You can find hackpet's snippets. The snippet of the hackpet has the tag `#hackpet`. If you have a lot of snippets in use, you can find them quickly by searching for tags.


## Screenshot
<img src="https://user-images.githubusercontent.com/13212227/87844969-7a28d980-c8fd-11ea-9c08-c96230937b19.png" width="100%">

## Tree
```
.
├── hackpet.toml => all snippets
├── snippets     => collection of snippets
```

## Snippets

| Description | Command |
| ----------- | ------- |
| Android set proxy | `adb shell settings put global http_proxy <ip address>:<param>` |
| Android unset proxy | `adb shell settings put global http_proxy :0` |
| Brute forcing for endpoints with dirsearch | `dirsearch -e php,asp,aspx,jsp,py,txt,conf,config,bak,backup,swp,old,db,sql -u <URL>` |
| certprobe / runs httprobe on all the hosts from certspotter | `curl -s https://crt.sh/\?q\=\%.<domain>\&output\=json \| jq -r '.[].name_value' \| sed 's/\*\.//g' \| sort -u \| httprobe \| tee -a ./all.txt` |
| Extract subdomains from IP Range | `nmap <ip range> -sn \| grep "<greping domain>" \| awk '{print $5}'` |
| Find subdomain and takeover (with subfinder/amass/assetfinder/subjack) | `subfinder -d <domain> >> domains ; assetfinder -subs-only <domain> >> domains ; amass enum -norecursive -noalts -d <domain> >> domains ; subjack -w domains -t 100 -timeout 30 -ssl -c ~/go/src/github.com/haccer/subjack/fingerprints.json -v  \| tee takeover` |
| Get url with gau, included parameter | `echo <domain> \| ~/go/bin/gau \| grep "=" \| qsreplace -a ` |
| Get urls from urlscanio | `gron "https://urlscan.io/api/v1/search/?q=domain:<domain>"  \| grep 'url' \| gron --ungron` |
| ipinfo | `curl http://ipinfo.io/<param>` |
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
```toml
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
Oneline
```
$ hack-pet merge
$ git add hackpet.toml README.md ; git commit -m "merge and distribute readme" ; git push -u origin master
```
