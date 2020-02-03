#!/bin/bash
echo "Installing FockCache"
go get github.com/briandowns/spinner
echo "GoLang spinner Pack Installed"
go get github.com/christophwitzko/go-curl
echo "GoLang go-curl Pack Installed"
go build -o FockCache main.go
echo "FockCache Builded."
echo -e "Usage: ./FockCache --hostname victim.host"