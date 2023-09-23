package robot

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"strings"
)

func IsNeedRobotReply(ctx context.Context, senderId int, mentionUids []int) (*model.Robot, bool) {

	// todo 需要改成查缓存
	robotInfo, _ := service.Robot().GetRobotByUserId(ctx, senderId)

	if robotInfo != nil {
		if robotInfo.UserId == 1 {
			return nil, false
		}
	}

	if robotInfo == nil && len(mentionUids) > 0 {
		robotInfo, _ = service.Robot().GetRobotByUserId(ctx, mentionUids[0])
	}

	if robotInfo == nil {
		return nil, false
	}

	return robotInfo, true
}

func RobotReply(ctx context.Context, robotInfo *model.Robot, senderId, receiverId, talkType int, text string, mentions ...string) {

	logger.Info(ctx, gjson.MustEncodeString(robotInfo))

	text = strings.TrimSpace(text)

	switch robotInfo.Company {
	case "OpenAI":
		switch robotInfo.ModelType {
		case "chat":
			OpenAI.Chat(ctx, senderId, receiverId, talkType, text, robotInfo.Model, mentions...)
		case "image":
			OpenAI.Image(ctx, senderId, receiverId, talkType, text, mentions...)
		}
	case "Baidu":
		switch robotInfo.ModelType {
		case "chat":
			ErnieBot.Chat(ctx, senderId, receiverId, talkType, text, robotInfo.Model, mentions...)
		}
	case "Xfyun":
		switch robotInfo.ModelType {
		case "chat":
			Spark.Chat(ctx, senderId, receiverId, talkType, text, robotInfo.Model, mentions...)
		}
	case "Aliyun":
		switch robotInfo.ModelType {
		case "chat":
			Aliyun.Chat(ctx, senderId, receiverId, talkType, text, robotInfo.Model, mentions...)
		}
	case "Midjourney":
		switch robotInfo.ModelType {
		case "image":
			Midjourney.Image(ctx, senderId, receiverId, talkType, text, robotInfo.Proxy)
		}
	}
}
