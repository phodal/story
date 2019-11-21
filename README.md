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
go run main.go create 'add docs to README'
```

### Sync Story

```
go run main.go sync
```

### list Stories

```
go run main.go list
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

