package testmath

import "testing"

func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//文件名必须以_test为后缀

//方法名必须以Test开头
func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

/*************
可采用工具查看影响性能的部分
go test -bench . -cpuprofile cpu.out  #会生成二进制文件
go tool pprof cpu.out #然后输入 web， 需要安装 Graphviz
brew install graphviz
******************/
//性能测试方法名必须以Benchmark开头
func BenchmarkOne(b *testing.B) {
	//s := "当你的声母韵母声调都没有问题了，只是运用得不够娴熟时，可以加强绕口令的训练。声母韵母声调一点都不准，绕口令练得再多，也白搭"
	s := "当你的声母韵母声调都没有"
	ans := 9
	for i := 0; i < b.N; i++ {
		actual := lengthOfLongestSubstring(s)
		if ans != actual {
			b.Errorf("性能测试：：：：：")
		}
	}
}
