package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitComplexSep(t *testing.T) {
	s := "我爱中国"
	sep := "爱"
	get := Split(s, sep)
	expect := []string{"我", "中国"}
	equal := reflect.DeepEqual(get, expect)

	if !equal {
		t.Errorf("split failed, expect:%v, get:%v", expect, get)
	}

}

func TestSplitSimpleSep(t *testing.T) {
	s := "abcd"
	sep := "b"
	get := Split(s, sep)
	expect := []string{"a", "cd"}
	equal := reflect.DeepEqual(get, expect)
	if !equal {
		t.Errorf("split failed, expect:%v, get:%v", expect, get)
	}

}

func TestSplitAll(t *testing.T) {
	type test struct {
		s      string
		sep    string
		expect []string
	}

	tests := map[string]test{
		"sample1": test{s: "我是中国人", sep: "是", expect: []string{"我", "中国人"}},
		"sample2": test{s: "abcd", sep: "b", expect: []string{"a", "cd"}},
		"sample3": test{s: "我爱中国", sep: "爱", expect: []string{"我", "中国"}},
	}

	for name, sample := range tests {
		t.Run(name, func(t *testing.T) {
			get := Split(sample.s, sample.sep)
			if !reflect.DeepEqual(get, sample.expect) {
				t.Errorf("expect:%v, get:%v", sample.expect, get)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("我爱中国爱你", "爱")
	}
}

func ExampleSplit() {
	fmt.Println(Split("我爱中国", "爱"))
	fmt.Println(Split("abcd", "b"))
	//Output:
	//[我 中国]
	//[a cd]

}
