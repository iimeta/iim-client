package robot

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/dao"
	"github.com/iimeta/iim-client/internal/errors"
	"github.com/iimeta/iim-client/internal/model"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/logger"
	"github.com/iimeta/iim-sdk/sdk"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type sRobot struct{}

func init() {
	service.RegisterRobot(New())
}

func New() service.IRobot {
	return &sRobot{}
}

func (s *sRobot) GetRobotByUserId(ctx context.Context, userId int) (*model.Robot, error) {

	robot, err := dao.Robot.GetRobotByUserId(ctx, userId)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
		return nil, err
	}

	if robot == nil {
		return nil, nil
	}

	return &model.Robot{
		UserId:    robot.UserId,
		IsTalk:    robot.IsTalk,
		Status:    robot.Status,
		Type:      robot.Type,
		Corp:      robot.Corp,
		Model:     robot.Model,
		ModelType: robot.ModelType,
		Role:      robot.Role,
		Prompt:    robot.Prompt,
		Proxy:     robot.Proxy,
		CreatedAt: robot.CreatedAt,
		UpdatedAt: robot.UpdatedAt,
	}, nil
}

func (s *sRobot) RobotReply(ctx context.Context, uid int, textMessageReq *model.TextMessageReq) {

	talkType := textMessageReq.Receiver.TalkType
	receiverId := uid
	mentionNickname := ""

	if talkType == 2 {

		if textMessageReq.Mention.Type != 1 && len(textMessageReq.Mention.Uids) == 0 {
			return
		}

		receiverId = textMessageReq.Receiver.ReceiverId

		if textMessageReq.Mention.Type == 1 && dao.GroupMember.IsLeader(ctx, receiverId, uid) {
			ids := dao.GroupMember.GetMemberIds(ctx, receiverId)
			if ids != nil && len(ids) > 0 {
				textMessageReq.Mention.Uids = ids
			}
		}
	}

	robots, isNeed := sdk.Robot.IsNeedRobotReply(ctx, append(textMessageReq.Mention.Uids, textMessageReq.Receiver.ReceiverId)...)
	if isNeed {

		prompt := textMessageReq.Text

		if len(prompt) == 0 {
			return
		}

		if talkType == 2 {
			user, err := dao.User.FindUserByUserId(ctx, uid)
			if err != nil {
				logger.Error(ctx, err)
			}
			mentionNickname = user.Nickname
		}

		session, err := service.TalkSession().FindBySession(ctx, uid, textMessageReq.Receiver.ReceiverId, talkType)
		if err != nil {
			logger.Error(ctx, err)
		}

		// 机器人回复
		for _, robot := range robots {
			robot := robot
			_ = grpool.AddWithRecover(ctx, func(ctx context.Context) {

				switch robot.ModelType {
				case sdk.MODEL_TYPE_TEXT:

					message := sdk.NewMessage()
					message.Prompt = prompt
					message.Stype = talkType
					message.IsWithContext = session.IsOpenContext == 0

					if talkType == 2 {
						message.Sid = textMessageReq.Receiver.ReceiverId
					} else {
						message.Sid = session.Id
					}

					content := ""
					text, err := sdk.Robot.Text(ctx, robot, message)
					if err != nil {
						logger.Error(ctx, err)
						content = err.Error()
					} else {
						content = text.Content
					}

					if talkType == 2 {
						content += "\n@" + mentionNickname
					}

					if err = service.TalkMessage().SendMessage(ctx, &model.Message{
						MsgType:  consts.MsgTypeText,
						TalkType: talkType,
						Text: &model.Text{
							Content: content,
						},
						Sender: &model.Sender{
							Id: robot.UserId,
						},
						Receiver: &model.Receiver{
							TalkType:   talkType,
							Id:         receiverId,
							ReceiverId: receiverId,
						},
					}); err != nil {
						logger.Error(ctx, err)
						return
					}

				case sdk.MODEL_TYPE_IMAGE:

					content := "您的请求已收到, 请耐心等待1-5分钟, 精彩马上为您呈献..."

					if talkType == 2 {
						content += "\n@" + mentionNickname
					}

					if err := service.TalkMessage().SendMessage(ctx, &model.Message{
						MsgType:  consts.MsgTypeText,
						TalkType: talkType,
						Text: &model.Text{
							Content: content,
						},
						Sender: &model.Sender{
							Id: robot.UserId,
						},
						Receiver: &model.Receiver{
							TalkType:   talkType,
							Id:         receiverId,
							ReceiverId: receiverId,
						},
					}); err != nil {
						logger.Error(ctx, err)
					}

					if robot.Corp == sdk.CORP_MIDJOURNEY {
						prompt = gstr.Replace(prompt, "\n", "")
						prompt = gstr.Replace(prompt, "\r", "")
						prompt = gstr.TrimLeftStr(prompt, "/mj")
						prompt = gstr.TrimLeftStr(prompt, "/imagine")
						prompt = strings.TrimSpace(prompt)
					}

					message := sdk.NewMessage()
					message.Prompt = prompt
					message.Stype = talkType
					message.IsSave = true

					if talkType == 2 {
						message.Sid = textMessageReq.Receiver.ReceiverId
					} else {
						message.Sid = session.Id
					}

					image, err := sdk.Robot.Image(ctx, robot, message)
					if err != nil {
						logger.Error(ctx, err)

						content = err.Error()

						if talkType == 2 {
							content += "\n@" + mentionNickname
						}

						if err = service.TalkMessage().SendMessage(ctx, &model.Message{
							MsgType:  consts.MsgTypeText,
							TalkType: talkType,
							Text: &model.Text{
								Content: content,
							},
							Sender: &model.Sender{
								Id: robot.UserId,
							},
							Receiver: &model.Receiver{
								TalkType:   talkType,
								Id:         receiverId,
								ReceiverId: receiverId,
							},
						}); err != nil {
							logger.Error(ctx, err)
							return
						}

						return
					}

					items := make([]*model.MixedItem, 0)
					items = append(items, &model.MixedItem{
						MsgType: consts.MsgTypeImage,
						Image: &model.Image{
							Url:    image.Url,
							Width:  image.Width,
							Height: image.Height,
							Size:   image.Size,
						},
					})

					if talkType == 2 {
						items = append(items, &model.MixedItem{
							MsgType: consts.MsgTypeText,
							Text: &model.Text{
								Content: "@" + mentionNickname,
							},
						})
					}

					if err := service.TalkMessage().SendMessage(ctx, &model.Message{
						MsgType:  consts.MsgTypeMixed,
						TalkType: talkType,
						Mixed:    &model.Mixed{Items: items},
						Sender: &model.Sender{
							Id: robot.UserId,
						},
						Receiver: &model.Receiver{
							TalkType:   talkType,
							Id:         receiverId,
							ReceiverId: receiverId,
						},
					}); err != nil {
						logger.Error(ctx, err)
					}

					if robot.Corp == sdk.CORP_MIDJOURNEY && !gstr.HasPrefix(prompt, "UPSCALE") {

						taskId := image.TaskId
						content := fmt.Sprintf("Prompt: %s\n", prompt)
						content += fmt.Sprintf("Result: 任务ID: %s, 您可以回复以下内容对图片进行相应操作:\n", taskId)
						content += fmt.Sprintf("第一张: UPSCALE::1::%s , VARIATION::1::%s\n", taskId, taskId)
						content += fmt.Sprintf("第二张: UPSCALE::2::%s , VARIATION::2::%s\n", taskId, taskId)
						content += fmt.Sprintf("第三张: UPSCALE::3::%s , VARIATION::3::%s\n", taskId, taskId)
						content += fmt.Sprintf("第四张: UPSCALE::4::%s , VARIATION::4::%s\n", taskId, taskId)
						content += fmt.Sprintf("操作说明: UPSCALE为放大, VARIATION为微调, 示例: UPSCALE::1::%s", taskId)

						if talkType == 2 {
							content += "\n@" + mentionNickname
						}

						if err = service.TalkMessage().SendMessage(ctx, &model.Message{
							MsgType:  consts.MsgTypeText,
							TalkType: talkType,
							Text: &model.Text{
								Content: content,
							},
							Sender: &model.Sender{
								Id: robot.UserId,
							},
							Receiver: &model.Receiver{
								TalkType:   talkType,
								Id:         receiverId,
								ReceiverId: receiverId,
							},
						}); err != nil {
							logger.Error(ctx, err)
							return
						}
					}
				}
			}, nil)
		}
	}
}
