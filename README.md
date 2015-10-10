# Hotspot

Hotspot shows you the hotspots of your code base by reading the churn rate from
the git history and the lines of code for each file.

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
