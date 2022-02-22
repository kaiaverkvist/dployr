#!/usr/bin/env bash
version=$(tr -dc A-Za-z0-9 </dev/urandom | head -c 6 | od -A n -t x1 | sed 's/ *//g' ; echo '')
time=$(date +%s)
go build -o dployr -ldflags="-X 'github.com/kaiaverkvist/dployr/version.BuildTime=$time' -X 'github.com/kaiaverkvist/dployr/version.BuildVersion=$version'" cmd/app/app.go