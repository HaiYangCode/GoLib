package GoLib

import "time"

// CurrentTime https://zhuanlan.zhihu.com/p/145009400
func CurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006.01.02 15:04:05")
}
