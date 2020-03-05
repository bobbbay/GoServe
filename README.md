<h1 align="center"> GoServe </h1>

<p align="center"><img src="https://github.com/Bobbbay/GoServe/raw/master/assets/GoServe-mini.png" height="200" align="center"/></p>

<p align="center">
  <a href="#demo">Demo</a>
  &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#install">Install</a>
  &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#start">Use</a>
  &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#todo">Todo</a>
  &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#contribute">Contribute</a>
</p>

<p align="center">
  <a target="_blank" href="https://app.circleci.com/github/Bobbbay/GoServe/pipelines" title="Circle CI Check"><img src="https://circleci.com/gh/Bobbbay/GoServe.svg?style=svg"/></a>
  <a target="_blank" href="https://travis-ci.org/Bobbbay/GoServe" title="Travis CI Check"><img src="https://travis-ci.org/Bobbbay/GoServe.svg?branch=master"/></a>
  <a target="_blank" href="https://app.bitrise.io/app/b0999db5cd64218a" title="Bitrise Check"><img src="https://app.bitrise.io/app/b0999db5cd64218a/status.svg?token=3krVYrcb8WhnTUEsSOAB8Q"/></a>
</p>

Quickly serve static sites in Go. You don't *need* to have Golang installed, nor do you need to know any Go! Just use the speed of Go to your advantage.

## Demo
[![asciicast](https://asciinema.org/a/307129.svg)](https://asciinema.org/a/307129)

## What is this?
GoServe is a quick, 1-dep (Golang) serving utility, making anyone serve locally easily and quickly, especially since you have the speed of Go on your side!
Absolutely NO knowledge of Go is needed to use. Just run `make` and it will do the rest!

## Install
### Dependencies
Actually, I lied. You need git to install. Sorry :smirk::grin:.

If you don't have Golang installed, you will also need `curl` and `wget`. 

### Just get it
#### I *don't* have Golang installed
NOTE: Your version of GoServe may be old. The most updated precompiled binary is the `x86_64` binary.
Run 
```
wget -q raw.githubusercontent.com/Bobbbay/GoServe/master/install.sh -O - | bash
```
And you're done!

#### I *have* Golang installed
Run 
```
git clone https://github.com/Bobbbay/GoServe.git
```
That's Bobbbay with THREE b's

## Start
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

## TODO
A lot.
- Asciinema at the top => .gif in this README
- WAYYYYY more commenting main.go
- Translate README?
- More Asciinema vids

## Contribute
HELP! Don't hesitate to open issues and PRs (please)!
We only support `armv7l` and `x86_64` at the moment for pre-compiled binaries. Any additional binaries would be greatly appreciated!
