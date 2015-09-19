package sc

import "fmt"

const (
	IR = 0
	KR = 1
	AR = 2
)

func CheckRate(rate int8) {
	if rate != IR && rate != KR && rate != AR {
		panic(fmt.Errorf("Unsupported rate %d", rate))
	}
}
