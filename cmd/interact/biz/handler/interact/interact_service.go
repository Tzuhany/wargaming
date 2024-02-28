// Code generated by hertz generator.

package interact

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"wargaming/cmd/api/biz/pack"
	"wargaming/cmd/interact/biz/model/interact"
	"wargaming/cmd/interact/biz/ws"
)

// Interact .
// @router /interact [GET]
func Interact(ctx context.Context, c *app.RequestContext) {

	var err error
	var req interact.InteractReq
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	ws.Interact(ctx, c, &req)
}
