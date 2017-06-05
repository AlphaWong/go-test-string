# go-test-string
Origin url http://studygolang.com/articles/9842

# Objective 
Test the string appending in difference way.

# Expected result
1. fmt.Sprintf should same as strings.Join
1. string + faster than method 1 2 times
1. bytes.Buffer faster than method 2 400 to 500 times (define bytes.Buffer outside each loop)
1. Instantly define a bytes.Buffer inside each loop will slow then method 3 50%

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
            // fmt.Sprintf
            s = fmt.Sprintf("%s[%s]", s, v)
        case 1: 
            // string +
            s = s + "[" + v + "]"
        case 2: 
            // strings.Join
            s = strings.Join([]string{s, "[", v, "]"}, "")
        case 3: 
            // temporary bytes.Buffer
            b := bytes.Buffer{}
            b.WriteString("[")
            b.WriteString(v)
            b.WriteString("]")
            s = b.String()
        case 4: 
            // stable bytes.Buffer
            buf.WriteString("[")
            buf.WriteString(v)
            buf.WriteString("]")
        }

        if i == n-1 {
            if index == 4 { // for stable bytes.Buffer
                s = buf.String()
            }
            fmt.Println(len(s)) // consume s to avoid compiler optimization
        }
    }
    t1 := time.Now()
    d = t1.Sub(t0)
    fmt.Printf("time of way(%d)=%v\n", index, d)
    return d
}
```
