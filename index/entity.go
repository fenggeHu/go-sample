package index

import "time"

// 基本信息
type Info struct {
	Name     string    `json:"name"`
	Summary  string    `json:"summary"`
	Category string    `json:"category"`
	Path     string    `json:"path" gorm:"uniqueIndex"` // 绝对路径
	Star     int       `json:"star"`
	Size     int64     `json:"size"`
	Type     string    `json:"type"`
	ModTime  time.Time `json:"modTime"`           //文件的最近修改时间
	Root     string    `json:"root" gorm:"index"` //配置的根路径
}

// 视频
type Video struct {
	Info
	Poster      string `json:"poster"`      //url
	ReleaseDate string `json:"releaseDate"` //上映日期
	Minute      int    `json:"minute"`      //片长（分钟）
}
