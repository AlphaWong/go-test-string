# go-test-string
Origin url http://studygolang.com/articles/9842

# Objective 
Test the string appending in difference way.

# Expected result
1. fmt.Sprintf should be the same as strings.Join
1. string + faster than method 1 2 times
1. bytes.Buffer faster than method 2 400 to 500 times (define bytes.Buffer outside each looping)
1. Instantly define a bytes.Buffer inside each looping will slow then method 3 50%

# Actual result
0. strings.Join share the worst performance.
0. fmt.Sprintf() needs <strong>50%</strong> from strings.Join.
0. string + (direct append) needs <strong>65%</strong> from fmt.Sprintf().
0. Instantly define a bytes.Buffer inside each loop via <var>bytes.NewBuffer</var> needs <strong>1.675%</strong> from direct append
0.  Instantly define a bytes.Buffer inside each loop via <var>bytes.Buffer{}</var> needs <strong>45.61%</strong> from <var>bytes.NewBuffer</var>
0. <var>new(bytes.Buffer)</var> almost same as <var>bytes.Buffer{}</var>
0. define bytes.Buffer outside each looping only needs <strong>0.589865 ms</strong> (only <strong>43%</strong> from Instantly define)

# Benchmark
<pre>
410000
time of way(0)=573.444884ms
410000
time of way(1)=305.717493ms
410000
time of way(2)=199.182739ms
41
time of way(3)=3.356549ms
41
time of way(4)=1.553018ms
41
time of way(5)=1.357562ms
410000
time of way(6)=589.865µs

way 0 is  972.2 times of way 6
way 1 is  518.3 times of way 6
way 2 is  337.7 times of way 6
way 3 is    5.7 times of way 6
way 4 is    2.6 times of way 6
way 5 is    2.3 times of way 6
</pre>

# Run
```go run main.go```

# Go version
 go version>= 1.8.3
# Code
```
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
```