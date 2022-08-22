package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"word-count/config"
)

func Cancel(ctx context.Context, c *app.RequestContext) {
	// single file
	config.GlobalDAGConfig.Set("mode", "off")
	config.GlobalDAGConfig.WriteConfig()
	klog.Info(config.GlobalDAGConfig.GetString("mode"))

	c.String(consts.StatusOK, fmt.Sprintf("canceled job seccess!"))
}
