package model

// Node json 文件格式
type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}
