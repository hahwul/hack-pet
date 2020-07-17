
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
