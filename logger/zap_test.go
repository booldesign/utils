package logger

import (
	"fmt"
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/4/25 23:00
 * @Desc:
 */

func TestLogger(t *testing.T) {
	loggers := NewJSONLogger(
		WithField("domain", fmt.Sprintf("%s[%s]", "testService", "debug")),
		WithInfoLevel(),
		WithLogPath("./micro.log"),
	)
	loggers.Info(111)
}
