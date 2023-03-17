package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// アメリカ合衆国憲法とか渡してみるといいと思う
	s := `Hello, World!`
	if len(os.Args) > 1 {
		s = os.Args[1]
	}
	fmt.Println(s)

	// A~Zまでの文字の各出現回数をカウント
	m := map[string]int{}
	for _, c := range s {
		if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
			m[strings.ToUpper(string(c))]++
		}
	}

	// 降順にソートしたいのでsliceに詰め替え
	type foo struct {
		char string
		num  int
	}

	count := 0
	fs := make([]foo, 0)
	for k, v := range m {
		count += v
		fs = append(fs, foo{char: k, num: v})
	}

	sort.SliceStable(fs, func(i, j int) bool { return fs[i].num > fs[j].num })

	// output
	for _, f := range fs {
		fmt.Printf("%s: %6.3f%%\n", f.char, float64(f.num)/float64(count)*100)
	}
}
