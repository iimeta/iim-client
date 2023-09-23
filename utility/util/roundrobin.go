package util

type Roundrobin struct {
	CurIndex int
}

// 轮询模式
func (roundrobin *Roundrobin) Roundrobin(values []string) (value string) {

	lens := len(values)
	if roundrobin.CurIndex >= lens {
		roundrobin.CurIndex = 0
	}

	value = values[roundrobin.CurIndex]
	roundrobin.CurIndex = (roundrobin.CurIndex + 1) % lens
	return
}
