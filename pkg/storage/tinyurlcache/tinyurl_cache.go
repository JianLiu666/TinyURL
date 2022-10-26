package tinyurlcache

import (
	"fmt"
	"sync"
	"tinyurl/pkg/storage"

	"github.com/emirpasic/gods/maps/treemap"
)

var once sync.Once
var instance *treemap.Map

// 初始化 tinyurl cache
// key: tinyurl (string)
// Value: *storage.Url
func Init() {
	once.Do(func() {
		instance = treemap.NewWith(func(a, b interface{}) int {
			urla := a.(*storage.Url)
			urlb := b.(*storage.Url)
			if urla.Tiny < urlb.Tiny {
				return -1
			}
			if urla.Tiny == urlb.Tiny {
				return 0
			}
			return 1
		})

		fmt.Println("constructed tinyurl local cache successful.")
	})
}
