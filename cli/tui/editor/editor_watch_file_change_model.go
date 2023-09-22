package editor

import (
	"os"
	"time"

	"github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/tui/channel"
	"gopkg.in/fsnotify.v1"
)

func debounce[T any](events <-chan T, duration time.Duration) <-chan T {
	out := make(chan T)
	go func() {
		timer := time.NewTimer(duration)
		var last T
		for {
			select {
			case event, ok := <-events:
				if !ok {
					close(out)
					return
				}
				last = event
				timer.Reset(duration)
			case <-timer.C:
				out <- last
			}
		}
	}()
	return out
}

type fileChangedMsg struct{}

func makeWatchFileChangeModel(
	filePath string,
) channel.Model[fileChangedMsg] {
	return channel.Model[fileChangedMsg]{
		Channel: make(chan fileChangedMsg, 10),
		OnInit: func(c chan fileChangedMsg) func() error {
			stat, err := os.Stat(filePath)
			if err != nil {
				log.I.Error(err.Error())
			}
			lastModTime := stat.ModTime()

			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.I.Fatal(err.Error())
			}

			if err = watcher.Add(filePath); err != nil {
				watcher.Close()
				log.I.Fatal(err.Error())
			}

			debouncedEvents := debounce(watcher.Events, time.Second)

			for range debouncedEvents {
				stat, err := os.Stat(filePath)
				if err != nil {
					log.I.Error(err.Error())
					continue
				}

				if !stat.ModTime().Equal(lastModTime) {
					lastModTime = stat.ModTime()

					select {
					case c <- fileChangedMsg{}:
						// Config sent successfully
					default:
					}
				}
				c <- fileChangedMsg{}
			}

			return func() error {
				return watcher.Close()
			}
		},
	}
}
