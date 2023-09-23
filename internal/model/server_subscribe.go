package model

type AuthConn struct {
	Uid     int    `json:"uid"`
	Channel string `json:"channel"`
}

type Authorize struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}

type SubscribeContent struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
