# Map声明、元素访问及遍历

### Map 声明

```Go
m := map[string]int{"one": 1, "two": 2, "three": 3}
m1 := map[string]int{}
m1["one"] = 1
m2 := make(map[string]int, 10 /*Initial Capacity*/)
//为什么不初始化len？
```

### Map 元素的访问

与其他主要编程语⾔的差异

* 在访问的 Key 不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在

```Go
if v, ok := m["four"]; ok {
 t.Log("four", v)
} else {
 t.Log("Not existing")
}
```

[代码地址](../code/go_learning/src/ch7/map_test/map_test.go)

### Map 遍历

```Go
m := map[string]int{"one": 1, "two": 2, "three": 3}
for k, v := range m {
 t.Log(k, v)
}
```