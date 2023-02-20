// Code generated by hertz generator.

package api

import (
	"context"
	"errors"
	"github.com/TremblingV5/DouTok/applications/api/biz/handler"
	"github.com/TremblingV5/DouTok/applications/api/initialize"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/TremblingV5/DouTok/pkg/errno"

	api "github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
)

// MessageChat .
// @router /douyin/message/chat [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildMessageChatResp(errno.ErrBind))
		return
	}

	resp, err := rpc.MessageChat(ctx, rpc.MessageClient, &message.DouyinMessageChatRequest{
		ToUserId: req.ToUserId,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildMessageChatResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}

// MessageAction .
// @router /douyin/message/action [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinMessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildMessageActionResp(errno.ErrBind))
		return
	}

	userId, e := c.Get(initialize.AuthMiddleware.IdentityKey)
	if !e {
		err = errors.New("identity not exist")
	}
	resp, err := rpc.MessageAction(ctx, rpc.MessageClient, &message.DouyinMessageActionRequest{
		UserId:     userId.(int64),
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildMessageActionResp(errno.ConvertErr(err)))
		return
	}
	handler.SendResponse(c, resp)
	//chat.ServeWs(ctx, c)
}
