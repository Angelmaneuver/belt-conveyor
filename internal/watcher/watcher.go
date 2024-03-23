package watcher

import (
	"log"

	"github.com/Angelmaneuver/belt-conveyor/internal/manager"
	"github.com/fsnotify/fsnotify"
)

func Start(watchpoint string, cm *manager.ConvertManager) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, sucess := <-watcher.Events:
				if !sucess {
					return
				}

				log.Println("Event ", event)

				if event.Has(fsnotify.Create) {
					if err := cm.Run(event.Name); err != nil {
						log.Panicln("Error ", err)
					}
				}
			case err, success := <-watcher.Errors:
				if !success {
					return
				}

				log.Println("Error ", err)
			}
		}
	}()

	if err := watcher.Add(watchpoint); err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}
