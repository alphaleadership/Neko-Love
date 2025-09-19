package cache

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// StartWatching initializes a new file system watcher for the image cache root directory.
// It adds all subdirectories (categories) within the root directory to the watcher,
// enabling monitoring for file system events. The method starts a background goroutine
// to handle watch events. Returns an error if the watcher cannot be created or if
// reading the root directory fails.
func (c *ImageCache) StartWatching() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	c.watcher = watcher

	entries, err := os.ReadDir(c.root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			categoryPath := filepath.Join(c.root, entry.Name())
			_ = watcher.Add(categoryPath)
		}
	}

	go c.watchLoop()
	return nil
}

// watchLoop continuously listens for filesystem events and errors from the cache's watcher.
// It processes events such as file creation, removal, and renaming within the cache root directory,
// updating the cache for the affected category as needed. If a category directory is detected as missing,
// it attempts to re-add it to the watcher. Any watcher errors are logged. The loop exits when the watcher
// channels are closed.
func (c *ImageCache) watchLoop() {
	for {
		select {
		case event, ok := <-c.watcher.Events:
			if !ok {
				return
			}

			relPath, err := filepath.Rel(c.root, event.Name)
			if err != nil {
				continue
			}
			parts := strings.Split(relPath, string(os.PathSeparator))
			if len(parts) < 1 {
				continue
			}
			category := parts[0]
			fullCategoryPath := filepath.Join(c.root, category)

			if event.Op&(fsnotify.Create|fsnotify.Remove|fsnotify.Rename) != 0 {
				log.Printf("Cache watcher detected change in category '%s': %s", category, event.Name)
				_ = c.LoadCategory(category)

				if _, err := os.Stat(fullCategoryPath); os.IsNotExist(err) {
					_ = c.watcher.Add(fullCategoryPath)
				}
			}
		case err, ok := <-c.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Cache watcher error: %v", err)
		}
	}
}
