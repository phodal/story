# story

Kanban as code in projects.


```
go get github.com/phodal/story
```

Usages:

```
Usage:
  story [flags]
  story [command]

Available Commands:
  create      create show
  help        Help about any command
  list        list the stories
  sync        sync the stories

Flags:
      --config string   config file (default is $HOME/.cobra.yaml)
  -h, --help            help for story
  -u, --user string     set author

Use "story [command] --help" for more information about a command.
```

### Create Story

```
story create 'add docs to README'
```

### Sync Story

```
story sync
```

### list Stories

```
story list
```

results:

```

+-----------+--------------------------------+----------------------+--------+--------+
|    ID     |             TITLE              |         DATE         | STATUS | AUTHOR |
+-----------+--------------------------------+----------------------+--------+--------+
| CYJlzObWR | add docs to README             | 2019-11-21T15:33:50Z |        |        |
+-----------+--------------------------------+----------------------+--------+--------+
| Dyp0iOxZg | use cucumber tag as file s     | 2019-11-21T15:45:47Z |        |        |
|           | group                          |                      |        |        |
+-----------+--------------------------------+----------------------+--------+--------+
```

### Update story status

```
story -p CYJlzObWR -u phodal -s done
```



License
---

[![Phodal's Idea](http://brand.phodal.com/shields/idea-small.svg)](http://ideas.phodal.com/)

@ 2019 A [Phodal Huang](https://www.phodal.com)'s [Idea](http://github.com/phodal/ideas).  This code is distributed under the MIT license. See `LICENSE` in this directory.
