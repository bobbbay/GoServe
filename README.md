# GoServe
Quickly serve static sites in Go

[![Bobbbay](https://circleci.com/gh/Bobbbay/GoServe.svg?style=svg)](https://app.circleci.com/github/Bobbbay/GoServe/pipelines)
[![Build Status](https://travis-ci.org/Bobbbay/GoServe.svg?branch=master)](https://travis-ci.org/Bobbbay/GoServe)

## GoServe is still in Beta, v0.1
Don't hesitate to open issues and PRs!

## What is this?
GoServe is a quick, 1-dep (Golang) serving utility, making anyone serve locally easily and quickly, especially since you have the speed of Go on your side!
Absolutely NO knowledge of Go is needed to use. Just run make start and will do the rest!

## Install
### Dependencies
Actually, I lied. You need git to install. Sorry :smirk::grin:
```
sudo apt-get install git
```
And here's Go's [installation docs](https://golang.org/doc/install).

### Get the code
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

Enter your creds for your server (only asks for port, source directory for your files and ignore directory).

Then, run 
```
make serve
```
And your server will start! Happy serving!
