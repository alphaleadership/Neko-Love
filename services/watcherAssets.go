package services

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var categories = make(map[string]string)

// WatchAssets sets up a file system watcher on the "assets" directory and its immediate subdirectories.
// It monitors for new directories being created within "assets", adds them to the watcher, and updates
// the global 'categories' map with the new category name and its path. Detected events and errors are
// logged accordingly. This function is intended to run asynchronously and does not block the main thread.
func WatchAssets() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}

	root := "assets"
	_ = watcher.Add(root)

	entries, _ := os.ReadDir(root)
	for _, entry := range entries {
		if entry.IsDir() {
			path := filepath.Join(root, entry.Name())
			_ = watcher.Add(path)
			categories[entry.Name()] = path
		}
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Create != 0 {
					info, err := os.Stat(event.Name)
					if err == nil && info.IsDir() {
						category := filepath.Base(event.Name)
						categories[category] = event.Name
						_ = watcher.Add(event.Name)
						log.Printf("New category detected: %s", category)
					}
				}
			case err, ok := <-watcher.Errors:
				if ok {
					log.Printf("Watcher error: %v", err)
				}
			}
		}
	}()
}

// GetCategoryPath returns the file system path associated with the given category.
// It looks up the category in the categories map and returns the corresponding path
// and a boolean indicating whether the category was found.
func GetCategoryPath(category string) (string, bool) {
	path, ok := categories[category]
	return path, ok
}
