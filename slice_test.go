package utils

import (
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 11:06
 * @Desc:
 */

func TestStringSliceUnique(t *testing.T) {
	tests := []struct {
		x      []string
		expect []string
	}{
		{[]string{"4", "2", "3", "4", "4", "1", "2"}, []string{"1", "2", "3", "4"}},
		{[]string{"2", "4", "6", "4", "1", "5", "6"}, []string{"1", "2", "4", "5", "6"}},
	}

	for _, test := range tests {
		actual := StringSliceUnique(test.x)
		{
			for k, v := range test.expect {
				if v != actual[k] {
					t.Errorf("StringSliceUnique(%v): expect %v, actual %v",
						test.x, test.expect, actual)
					break
				}
			}
		}
	}
}

func TestIntSliceUnique(t *testing.T) {
	tests := []struct {
		x      []int
		expect []int
	}{
		{[]int{4, 2, 3, 4, 4, 1, 2}, []int{1, 2, 3, 4}},
		{[]int{2, 4, 6, 4, 1, 5, 6}, []int{1, 2, 4, 5, 6}},
	}

	for _, test := range tests {
		actual := IntSliceUnique(test.x)
		{
			for k, v := range test.expect {
				if v != actual[k] {
					t.Errorf("IntSliceUnique(%v): expect %v, actual %v",
						test.x, test.expect, actual)
					break
				}
			}
		}
	}
}

func TestIntSliceJoin(t *testing.T) {
	tests := []struct {
		x      []int
		expect string
	}{
		{[]int{4, 2, 3, 4, 4, 1, 2}, "4,2,3,4,4,1,2"},
		{[]int{2, 4, 6, 4, 1, 5, 6}, "2,4,6,4,1,5,6"},
	}

	for _, test := range tests {
		if actual := IntSliceJoin(test.x, ","); actual != test.expect {
			t.Errorf("IntSliceJoin(%v): expect %s, actual %s",
				test.x, test.expect, actual)
		}
	}
}

func BenchmarkIntSliceJoin(b *testing.B) {
	elems := []int{2, 3, 4, 5, 7, 801, 1, 35, 36, 23, 5, 36, 3, 61, 21, 2, 2217, 70}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IntSliceJoin(elems, ",")
	}
	b.StopTimer()

	//打印报告
	b.ReportAllocs()
}

func TestIntSliceDiff(t *testing.T) {
	tests := []struct {
		x      []int
		y      []int
		expect []int
	}{
		{[]int{4, 2, 3, 4, 4, 1, 2}, []int{4, 5, 9, 4, 2}, []int{3, 1}},
		{[]int{2, 4, 6, 4, 1, 5, 6}, []int{7, 8, 6, 1, 1}, []int{2, 4, 4, 5}},
	}

	for _, test := range tests {
		actual := IntSliceDiff(test.x, test.y)
		for k, v := range test.expect {
			if v != actual[k] {
				t.Errorf("IntSliceDiff(%v, %v): expect %v, actual %v",
					test.x, test.y, test.expect, actual)
				break
			}
		}
	}
}

func TestIntSliceIntersect(t *testing.T) {
	tests := []struct {
		x      []int
		y      []int
		expect []int
	}{
		{[]int{4, 2, 3, 4, 4, 1, 2}, []int{4, 5, 9, 4, 2}, []int{4, 2, 4, 4, 2}},
		{[]int{2, 4, 6, 4, 1, 5, 6}, []int{7, 8, 6, 1, 1}, []int{6, 1, 6}},
	}

	for _, test := range tests {
		actual := IntSliceIntersect(test.x, test.y)
		for k, v := range test.expect {
			if v != actual[k] {
				t.Errorf("IntSliceIntersect(%v, %v): expect %v, actual %v",
					test.x, test.y, test.expect, actual)
				break
			}
		}
	}
}

func TestRemoveSliceElement(t *testing.T) {
	sl := make([]int, 0)
	for i := 0; i < 10; i++ {
		sl = append(sl, i)
	}
	sl = RemoveSliceElement(sl, 5)
	t.Log(sl)
}

func BenchmarkRemoveSliceElement(b *testing.B) {
	sl := make([]int, 0)
	for i := 0; i < 100000; i++ {
		sl = append(sl, i)
	}

	for i := 0; i < b.N; i++ {
		RemoveSliceElement(sl, 88888)
	}
}

func BenchmarkRemoveSliceElementParallel(b *testing.B) {
	sl := make([]int, 0)
	for i := 0; i < 100000; i++ {
		sl = append(sl, i)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			RemoveSliceElement(sl, 8888)
		}
	})
}
