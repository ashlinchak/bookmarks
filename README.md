# Bookmarks

Simple bookmark manager with tags written in Golang. It has CLI and uses SQLite database for storing bookmarks.

## Install
You can compile application via `build` command:
```
$ go build
```

## Usage
All available commands available in help interface
```
$ bookmarks --help
```

Implemented commands are:
* **setup** - initial setup database.
* **add** - add bookmark
* **delete** - delete bookmark
* **show** - list all bookmarks or search them by tags
* **tags** - print tags which are used by bookmarks

To get a detail information about each command run `help`:
```
$ bookmarks show --help
```
