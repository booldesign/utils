package tarce

import "testing"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/4/24 08:45
 * @Desc:
 */

func TestTrace(t *testing.T) {
	trace := NewTrace()
	t.Logf("%+v", trace)
}
