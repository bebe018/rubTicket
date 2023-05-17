package WD

import (
	"sync"
	"github.com/tebeka/selenium"
)

var mu sync.Mutex

func safeGet(wd selenium.WebDriver, url string) {
	mu.Lock()
	wd.Get(url)
	mu.Unlock()
}