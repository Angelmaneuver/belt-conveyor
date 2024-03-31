package watcher

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/Angelmaneuver/belt-conveyor/internal/manager"
	"github.com/fsnotify/fsnotify"
)

const WAIT = 5000 * time.Millisecond

func Start(watchpoint string, cm *manager.ConvertManager) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	go watch(w, cm)

	if err := w.Add(watchpoint); err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}

func watch(w *fsnotify.Watcher, cm *manager.ConvertManager) {
	var (
		mu     sync.Mutex
		timers = make(map[string]*time.Timer)
		event  = func(e fsnotify.Event) {
			if err := cm.Run(e.Name); err != nil {
				log.Panicln("Error ", err)
			}

			mu.Lock()
			delete(timers, e.Name)
			mu.Unlock()
		}
	)

	for {
		select {
		case err, sucess := <-w.Errors:
			if !sucess {
				return
			}

			log.Println("Error", err)

		case e, sucess := <-w.Events:
			if !sucess {
				return
			}

			if !e.Has(fsnotify.Create) && !e.Has(fsnotify.Write) {
				continue
			}

			mu.Lock()
			timer, sucess := timers[e.Name]
			mu.Unlock()

			if !sucess {
				timer = time.AfterFunc(math.MaxInt64, func() { event(e) })
				timer.Stop()

				mu.Lock()
				timers[e.Name] = timer
				mu.Unlock()
			}

			timer.Reset(WAIT)
		}
	}
}
