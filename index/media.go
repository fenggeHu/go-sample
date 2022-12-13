package index

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	VideoType = []string{".mkv", ".mp4"}
)

// 读video信息
func VideoInfo(name string, root string) *Video {
	dir, filename := filepath.Split(name)
	li := strings.LastIndex(filename, ".")
	filetype := filename[li:]
	if !isVideo(filetype) {
		return nil
	}

	stats, err := os.Stat(name)
	if err != nil {
		log.Println(name, err)
		return nil
	}

	category := ""
	if root != "" && strings.HasPrefix(dir, root) {
		category = dir[len(root) : len(dir)-1]
		fs := string(filepath.Separator)
		if strings.HasPrefix(category, fs) {
			category = category[len(fs):]
		}
	}
	info := Info{
		Name:     filename,
		Path:     name,
		Category: category,
		Size:     stats.Size(),
		Type:     filetype,
		ModTime:  stats.ModTime(),
	}

	return &Video{
		Info:        info,
		Poster:      "",
		ReleaseDate: "",
		Minute:      0,
	}
}

func isVideo(fileType string) bool {
	for _, v := range VideoType {
		if strings.EqualFold(v, fileType) {
			return true
		}
	}
	return false
}
