package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func benchmarkStringFunction(n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0:
			// strings.Join
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case 1:
			// fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 2:
			// string +
			s = s + "[" + v + "]"
		case 3:
			// temporary bytes.Buffer NewBuffer function
			b := bytes.NewBuffer([]byte(""))
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			s = b.String()
		case 4:
			// temporary bytes.Buffer
			b := bytes.Buffer{}
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			s = b.String()
		case 5:
			// temporary bytes.Buffer new function
			b := new(bytes.Buffer)
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			s = b.String()
		case 6:
			// stable bytes.Buffer
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}

		if i == n-1 {
			// for stable bytes.Buffer
			if index == 6 {
				s = buf.String()
			}
			// consume s to avoid compiler optimization
			fmt.Println(len(s))
		}
	}
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", index, d)
	return d
}

func main() {
	k := 7
	d := [7]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkStringFunction(10000, i)
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}
}
