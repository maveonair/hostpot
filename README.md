# Hotspot

Hotspot shows you the hotspots of your code base by reading the churn rate from
the git history and the lines of code for each file.

## Requirements
A working GO environment: https://golang.org/doc/code.html#GOPATH

## Install

```
$ git clone git@github.com:maveonair/hostpot.git
$ go install
```

## Usage

CSV output:

```
$ cd your-git-project
$ hotspot
```

or 

```
$ cd your-git-project
$ hotspot -format=csv`
```

JSON output:

```
$ cd your-git-project
$ hotspot -format=json
```
