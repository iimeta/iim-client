package util

type RoundRobin struct {
	CurIndex int
}

// 轮询模式
func (roundRobin *RoundRobin) RoundRobin(values []string) (value string) {

	lens := len(values)

	if roundRobin.CurIndex >= lens {
		roundRobin.CurIndex = 0
	}

	value = values[roundRobin.CurIndex]

	roundRobin.CurIndex = (roundRobin.CurIndex + 1) % lens

	return
}
