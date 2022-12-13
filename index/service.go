package index

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// get files
func Scan(name string) (ret []string) {
	entities, err := os.ReadDir(name)
	if err != nil {
		log.Println("Err: ", err)
		return
	}
	for _, dir := range entities {
		path := name + "/" + dir.Name()
		if dir.IsDir() {
			subs := Scan(path)
			for _, sub := range subs {
				if isValid(sub) {
					ret = append(ret, sub)
				}
			}
		} else {
			if isValid(path) {
				ret = append(ret, path)
			}
		}
	}
	return
}

func isValid(name string) bool {
	_, filename := filepath.Split(name)
	if strings.HasPrefix(filename, ".") || strings.HasPrefix(filename, "~") {
		return false
	}
	return true
}

// add 索引
func Add(name string) {
	fi, err := os.Stat(name)
	if err != nil {
		log.Println("Err: ", err)
		return
	}
	if fi.IsDir() {
		entities, err := os.ReadDir(name)
		if err != nil {
			log.Println("Err: ", err)
			return
		}
		for _, dir := range entities {
			path := name + "/" + dir.Name()
			if dir.IsDir() {
				Add(path)
			} else {
				add(path)
			}
		}
	} else {
		add(name)
	}
}

func add(name string) {
	log.Println("Add to Store: ", name)
}

// remove 索引
func Remove(name string) {
	// 匹配 name 开头的所有文件，从库里删除
	log.Println("匹配 name 开头的所有文件，从库里删除. Remove from Store: ", name)
}
