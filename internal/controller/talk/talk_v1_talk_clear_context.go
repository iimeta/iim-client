package talk

import (
	"context"
	"fmt"
	"github.com/iimeta/iim-client/internal/consts"
	"github.com/iimeta/iim-client/internal/service"
	"github.com/iimeta/iim-client/utility/redis"

	"github.com/iimeta/iim-client/api/talk/v1"
)

func (c *ControllerV1) TalkClearContext(ctx context.Context, req *v1.TalkClearContextReq) (res *v1.TalkClearContextRes, err error) {

	_, err = redis.Del(ctx, fmt.Sprintf(consts.CHAT_MESSAGES_PREFIX_KEY, service.Session().GetUid(ctx), req.ReceiverId))

	return
}
