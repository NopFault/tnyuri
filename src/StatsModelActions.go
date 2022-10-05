package tnyuri

import "strconv"

func (S *Stats) Increase() {
	Update[Stats](map[string]string{
		"counter": strconv.Itoa(S.Counter + 1),
	}, "id", strconv.Itoa(S.Id))
	S.Counter += 1
}
