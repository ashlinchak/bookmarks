# Bookmarks

Simple bookmark manager with tags written in Golang. It has CLI and uses SQLite database for storing bookmarks.

## Install

```sh
$ go get github.com/ashlinchak/bookmarks
```

## Configuration

By default, SQLite database will be located in data/bookmarks.db path. You can configure it by specifying environment variable:
```
export BOOKMARKS_DB_PATH="$HOME/opt/bookmarks/bookmarks.db"
```

## Usage

All available commands available in help interface

```sh
$ bookmarks --help
```

Implemented commands are:
* **setup** - initial setup database.
* **add** - add bookmark
* **update** - update bookmark
* **delete** - delete bookmark
* **show** - list all bookmarks or search them by tags
* **tags** - print tags which are used by bookmarks

To get a detail information about each command run `help`:

```sh
$ bookmarks show --help
```
