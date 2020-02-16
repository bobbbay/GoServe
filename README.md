# GoServe
Quickly serve static sites in Go

[![Bobbbay](https://circleci.com/gh/Bobbbay/GoServe.svg?style=svg)](https://app.circleci.com/github/Bobbbay/GoServe/pipelines)
[![Build Status](https://travis-ci.org/Bobbbay/GoServe.svg?branch=master)](https://travis-ci.org/Bobbbay/GoServe)
[![Build Status](https://app.bitrise.io/app/b0999db5cd64218a/status.svg?token=3krVYrcb8WhnTUEsSOAB8Q)](https://app.bitrise.io/app/b0999db5cd64218a)
![Go Lint](https://github.com/Bobbbay/GoServe/workflows/Go%20Lint/badge.svg)

## What is this?
GoServe is a quick, 1-dep (Golang) serving utility, making anyone serve locally easily and quickly, especially since you have the speed of Go on your side!
Absolutely NO knowledge of Go is needed to use. Just run `make` and it will do the rest!

## Install
### Dependencies
Actually, I lied. You need git to install. Sorry :smirk::grin:

If you don't have Golang installed, you will also need `curl` and `wget`. 

### Just get it
#### I *don't* have Golang installed
Run 
```
wget raw.githubusercontent.com/Bobbbay/GoServe/master/install.sh
bash install.sh
```
And you're done!

#### I *have* Golang installed
Run 
```
git clone https://github.com/Bobbbay/GoServe.git
```
That's Bobbbay with THREE b's

## Start
### *Without* Golang
Just enter the project
```
cd GoServe
```
And then 
```
./GoServe
```

### *With* Golang
To begin, just run 
```
cd GoServe
```
Then
```
make
```

## Update
Just run 
```
make update
```

## HELP!!!
Don't hesitate to open issues and PRs (please)!
We only support `armv7l` and `x86_64` at the moment for pre-compiled binaries. Any additional binaries would be greatly appreciated!
