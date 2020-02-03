# fockcache
FockCache  - Minimalized Test Cache Poisoning

[![asciicast](https://asciinema.org/a/WKTlr2ffvL6CAISm7oSzSZCLe.svg)](https://asciinema.org/a/WKTlr2ffvL6CAISm7oSzSZCLe)

Detail For Cache Poisoning : https://portswigger.net/research/practical-web-cache-poisoning

### FockCache #

FockCache tries to make cache poisoning by trying X-Forwarded-Host and X-Forwarded-Scheme headers on web pages.

After successful result, it gives you a poisoned URL.



## To be added soon: 

1 - Page Param Checker

2 - Recursive Checking


## Installation

1 - Install with installer.sh

`chmod +x installer.sh`

`./installer.sh`

2 - Install manual

`go get github.com/briandowns/spinner`

`go get github.com/christophwitzko/go-curl`

`go run main.go --hostname victim.host`

or 

`go build FockCache main.go`

## Run

`./FockCache --hostname victim.host `

