package config

import "github.com/fsnotify/fsnotify"

var (
	watchFiles = make(map[string]func(fileName string))
	watcher    *fsnotify.Watcher
)

func AddWatch(file string) {

}
