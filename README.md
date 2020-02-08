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
```
sudo apt-get install git
```

### Get the code
#### I *don't* have Golang installed
Run the installer, it'll do everything for you!
```
curl https://raw.githubusercontent.com/Bobbbay/GoServe/master/install.sh | bash
```

Currently only support `x86_64`


#### I *do* have Golang installed
Just run 
```
git clone https://github.com/Bobbbay/GoServe.git
```
That's Bobbbay with THREE b's

## Start
To start, just run 

```
cd GoServe
```
Then
```
make
```

And your server will start! Happy serving!

## Update
Just run 
```
make update
```

## GoServe is still in Beta!
Don't hesitate to open issues and PRs!
