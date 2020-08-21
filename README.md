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

| Description | Command |
| ----------- | ------- |
| Find All Allocated IP ranges for ASN given an IP address | `whois -h whois.radb.net -i origin -T route $(whois -h whois.radb.net <Organization> \| grep origin: \| awk '{print $NF}' \| head -1) \| grep -w "route:" \| awk '{print $NF}' \| sort -n` |
| Android set proxy | `adb shell settings put global http_proxy <ip address>:<param>` |
| Android unset proxy | `adb shell settings put global http_proxy :0` |
| Brute forcing for endpoints with dirsearch | `dirsearch -e php,asp,aspx,jsp,py,txt,conf,config,bak,backup,swp,old,db,sql -u <URL>` |
| certprobe / runs httprobe on all the hosts from certspotter | `curl -s https://crt.sh/\?q\=\%.<domain>\&output\=json \| jq -r '.[].name_value' \| sed 's/\*\.//g' \| sort -u \| httprobe \| tee -a ./all.txt` |
| Extract subdomains from IP Range | `nmap <ip range> -sn \| grep "<greping domain>" \| awk '{print $5}'` |
| Find subdomain and takeover (with subfinder/amass/assetfinder/subjack) | `subfinder -d <domain> >> domains ; assetfinder -subs-only <domain> >> domains ; amass enum -norecursive -noalts -d <domain> >> domains ; subjack -w domains -t 100 -timeout 30 -ssl -c ~/go/src/github.com/haccer/subjack/fingerprints.json -v  \| tee takeover` |
| Find LFI with gau | `~/go/bin/gau <domain> \| ~/go/bin/gf lfi \| ~/go/bin/qsreplace "/etc/passwd" \| xargs -I % -P 25 sh -c 'curl -s "%" 2>&1 \| grep -q "root:x" && echo "VULN! %"'` |
| Find OpenRedirect with gau | `export LHOST="http://localhost"; gau <domain> \| gf redirect \| qsreplace "$LHOST" \| xargs -I % -P 25 sh -c 'curl -Is "%" 2>&1 \| grep -q "Location: $LHOST" && echo "VULN! %"'` |
| Get bugcrowd programs | `curl -sL https://github.com/arkadiyt/bounty-targets-data/raw/master/data/bugcrowd_data.json \| jq -r '.[].targets.in_scope[] \| [.target, .type] \| @tsv'` |
| one \| uniq); doneGet CIDR and Orgz from target lists | `for DOMAIN in $(cat <FILE NAME>);do echo $(for ip in $(dig a $DOMAIN +short); do whois $ip \| grep -e "CIDR\\|Organization" \| tr -s " " \| paste - -; d` |
| Get hackerone programs | `curl -sL https://github.com/arkadiyt/bounty-targets-data/blob/master/data/hackerone_data.json?raw=true \| jq -r '.[].targets.in_scope[] \| [.asset_identifier, .asset_type] \| @tsv'` |
| Get intigriti programs | `curl -sL https://github.com/arkadiyt/bounty-targets-data/raw/master/data/intigriti_data.json \| jq -r '.[].targets.in_scope[] \| [.endpoint, .type] \| @tsv'` |
| Get Subdomains from Archive | `curl -s "http://web.archive.org/cdx/search/cdx?url=*.<domain>/*&output=text&fl=original&collapse=urlkey" \| sed -e 's_https*://__' -e "s/\/.*//" \| sort -u` |
| Get Subdomains from BufferOverRun | `curl -s https://dns.bufferover.run/dns?q=.<domain> \|jq -r .FDNS_A[]\|cut -d',' -f2\|sort -u` |
| Get Subdomains from CertSpotter | `curl -s "https://certspotter.com/api/v0/certs?domain=<domain>" \| grep -Po "((http\|https):\/\/)?(([\w.-]*)\.([\w]*)\.([A-z]))\w+" \| sort -u` |
| Get Subdomains from crt.sh | `curl -s "https://crt.sh/?q=%25.<domain>&output=json" \| jq -r '.[].name_value' \| sed 's/\*\.//g' \| sort -u` |
| Get Subdomains from JLDC | `curl -s "https://jldc.me/anubis/subdomains/<domain>?" \| grep -Po "((http\|https):\/\/)?(([\w.-]*)\.([\w]*)\.([A-z]))\w+" \| sort -u` |
| Get Subdomains from RapidDNS.io | `curl -s "https://rapiddns.io/subdomain/<domain>?full=1#result" \| grep "<td><a" \| cut -d '"' -f 2 \| grep http \| cut -d '/' -f3 \| sed 's/#results//g' \| sort -u` |
| Get Subdomains from Riddler.io | `curl -s "https://riddler.io/search/exportcsv?q=pld:<domain>" \| grep -Po "(([\w.-]*)\.([\w]*)\.([A-z]))\w+" \| sort -u` |
| Get Subdomains from VirusTotal | `curl -s "https://www.virustotal.com/ui/domains/<domain>/subdomains?limit=40" \| grep -Po "((http\|https):\/\/)?(([\w.-]*)\.([\w]*)\.([A-z]))\w+" \| sort -u` |
| Get url with gau, included parameter | `echo <domain> \| ~/go/bin/gau \| grep "=" \| qsreplace -a ` |
| Get all the urls out of a sitemap.xml | `curl -s <sitemap URL> \| xmllint --format - \| grep -e 'loc' \| sed -r 's\|</?loc>\|\|g'` |
| Get urls from urlscanio | `gron "https://urlscan.io/api/v1/search/?q=domain:<domain>"  \| grep 'url' \| gron --ungron` |
| Find XSS with gospider | `gospider -S <TARGET URLS FILE> -c 10 -d 5 --blacklist ".(jpg\|jpeg\|gif\|css\|tif\|tiff\|png\|ttf\|woff\|woff2\|ico\|pdf\|svg\|txt)" --other-source \| grep -e "code-200" \| awk '{print $5}'\| grep "=" \| qsreplace -a \| dalfox pipe -o result.txt` |
| ipinfo | `curl http://ipinfo.io/<param>` |
| Create a wordlist using param used in the domain | `waybackurls <domain> \|  grep "?" \| unfurl keys  \| sort -u \| tee -a paramlist.txt` |
| Ports Scan without CloudFlare | `subfinder -silent -d <domain> \| filter-resolved \| cf-check \| sort -u \| naabu -rate 40000 -silent -verify \| httprobe` |
| Sort & Tested Domains from Recon.dev | `curl "https://recon.dev/api/search?key=<API Key>&domain=<domain>" \|jq -r '.[].rawDomains[]' \| sed 's/ //g' \| sort -u \|httpx -silent` |
| Find Subdomains TakeOver | `subfinder -d <target> >> domains ; assetfinder -subs-only <target> >> domains ; amass enum -norecursive -noalts -d <target> >> domains ; subjack -w domains -t 100 -timeout 30 -ssl -c ~/go/src/github.com/haccer/subjack/fingerprints.json -v 3 >> takeover ;` |
| Get multiple target's Custom URLs from ParamSpider | `cat <domains file> \| xargs -I % python3 ~/tool/ParamSpider/paramspider.py -l high -o ./spidering/paramspider/% -d % ;` |
| URLs Probing with cURL + Parallel | `cat <domains file> \| parallel -j50 -q curl -w 'Status:%{http_code}\t  Size:%{size_download}\t %{url_effective}\n' -o /dev/null -sk` |

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
