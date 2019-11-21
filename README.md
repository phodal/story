# story

Kanban as code in projects.


```
go get github.com/phodal/story
```

Usages:

```
story

Usage:
  story [flags]

Flags:
      --config string    config file (default is $HOME/.cobra.yaml)
  -c, --create string    create a story
  -h, --help             help for story
  -j, --journal string   show user journal
  -l, --list string      list a story
  -p, --pick string      pick a story
  -s, --status string    change status of story
  -y, --sync string      sync story
  -u, --user string      list author
```

Examples:

Create Story:

```
go run main.go --create 'add docs to README'
```
