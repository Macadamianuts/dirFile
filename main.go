package main

import (
	"github.com/wwater/utils/global"
	"github.com/wwater/utils/utils"
)

func main() {
	utils.LoadJson()
	utils.ParseNode(global.Node, "")
}
