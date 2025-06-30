package cache

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type ImageCache struct {
	sync.RWMutex
	files map[string][]string
	root  string
	watcher *fsnotify.Watcher
}

// New creates and initializes a new ImageCache instance using the specified assetsRoot directory.
// It loads all image categories from the root directory into the cache.
// Returns the initialized ImageCache or an error if loading categories fails.
func New(assetsRoot string) (*ImageCache, error) {
	cache := &ImageCache{
		files: make(map[string][]string),
		root:  assetsRoot,
	}
	
	if error := cache.loadAllCategories(); error != nil {
		return nil, error
	}

	log.Printf("Image cache initialized with root: %s", assetsRoot)

	return cache, nil
}

// loadAllCategories loads all category directories from the cache root directory.
// It iterates over each directory entry, loading categories found within the root.
// After loading, it logs the number of categories loaded and starts a file watcher
// to monitor changes in the cache. Returns an error if reading the directory or
// starting the watcher fails.
func (c *ImageCache) loadAllCategories() error {
	entries, err := os.ReadDir(c.root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			_ = c.LoadCategory(entry.Name())
		}
	}

	log.Printf("Loaded %d categories from cache", len(c.files))
	if err := c.StartWatching(); err != nil {
		log.Printf("Error starting file watcher: %v", err)
		return err
	}

	return nil
}

// LoadCategory loads the list of file names from the specified category directory into the cache.
// It locks the cache for thread-safe access, reads all non-directory entries from the category folder,
// and updates the cache's file map for the given category. Returns an error if the directory cannot be read.
func (c *ImageCache) LoadCategory(category string) error {
	c.Lock()
	defer c.Unlock()

	folder := filepath.Join(c.root, category)
	entries, err := os.ReadDir(folder)
	if err != nil {
		return err
	}

	var paths []string
	for _, entry := range entries {
		if !entry.IsDir() {
			paths = append(paths, entry.Name())
		}
	}
	c.files[category] = paths

	log.Printf("Loaded %d files for category '%s'", len(paths), category)

	return nil
}

// GetRandom returns a random file path from the cache for the specified category.
// If the category does not exist or contains no files, it returns an error (os.ErrNotExist).
// This method is safe for concurrent use.
func (c *ImageCache) GetRandom(category string) (string, error) {
	c.RLock()
	defer c.RUnlock()

	paths, ok := c.files[category]
	if !ok || len(paths) == 0 {
		return "", os.ErrNotExist
	}

	return paths[rand.Intn(len(paths))], nil
}

func (c *ImageCache) GetFiles(category string) []string {
	c.RLock()
	defer c.RUnlock()

	paths, ok := c.files[category]
	if !ok {
		return nil
	}

	return paths
}

// GetImagePath returns the full file path for an image with the specified category and name.
// It acquires a read lock to ensure thread-safe access to the cache.
// If the image is found in the cache, it returns the full path and true.
// If the image is not found, it returns an empty string and false.
func (c *ImageCache) GetImagePath(category, name string) (string, bool) {
	c.RLock()
	defer c.RUnlock()

	files, ok := c.files[category]
	if !ok {
		return "", false
	}

	for _, f := range files {
		if f == name {
			fullPath := filepath.Join(c.root, category, f)
			return fullPath, true
		}
	}

	return "", false
}