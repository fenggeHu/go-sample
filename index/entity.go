package index

import "time"

// 基本信息
type Info struct {
	Name     string `json:"name"`
	Summary  string `json:"summary"`
	Category string `json:"category"`
	Path     string `json:"path"` // 绝对路径
	Star     int    `json:"star"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
	ModTime  time.Time
}

// 视频
type Video struct {
	Info
	Poster      string `json:"poster"`      //url
	ReleaseDate string `json:"releaseDate"` //上映日期
	Minute      int    `json:"minute"`      //片长（分钟）
}
