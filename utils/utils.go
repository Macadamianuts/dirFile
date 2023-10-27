package utils

import (
	"encoding/json"
	"github.com/wwater/utils/global"
	"github.com/wwater/utils/model"
	"io/fs"
	"os"
	"path/filepath"
)

var workDir string   // 路径地址
var separator string // 分隔符
var node *model.Node // 创建 node 防止空指针

// 定义文件名
const fileJsonName = "dir.json"

// LoadJson 加载 Json
func LoadJson() {
	// 赋值给全局 global
	global.Node = node

	// 切割路径
	separator = string(filepath.Separator)
	// 获取路径
	workDir, _ = os.Getwd()

	// 读取文件
	jsonFileBytes, _ := os.ReadFile(workDir + separator + fileJsonName)
	// json 转换
	err := json.Unmarshal(jsonFileBytes, &global.Node)
	if err != nil {
		panic("json unmarshal is error :" + err.Error())
	}
}

// ParseNode 解析节点
func ParseNode(node *model.Node, parentDir string) {
	if node.Text != "" {
		CreateDir(node, parentDir)
	}

	if parentDir != "" {
		parentDir = parentDir + separator
	}

	if node.Text != "" {
		parentDir = parentDir + node.Text
	}

	for _, v := range node.Children {
		ParseNode(&v, parentDir)
	}
}

// CreateDir 创建文件夹
func CreateDir(node *model.Node, parentDir string) {
	dirPath := workDir + separator

	if dirPath != "" {
		dirPath = dirPath + parentDir + separator
	}
	dirPath = dirPath + node.Text
	// 创建文件
	err := os.MkdirAll(dirPath, fs.ModePerm)
	if err != nil {
		panic("create dir is error:" + err.Error())
	}
}
