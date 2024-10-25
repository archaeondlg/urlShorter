package initialize

import (
	"project/global"

	"github.com/bwmarrin/snowflake"
)

func Snowflake() *snowflake.Node {
	nodeId := global.Config.System.NodeId
	if nodeId <= 0 {
		nodeId = 1
	}
	node, err := snowflake.NewNode(int64(nodeId))
	if err != nil {
		return nil
	}
	return node
}
