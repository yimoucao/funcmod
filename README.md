# funcmod

functional programming for go

## Create Iterable

```
// use IterList()
iterable1 := funcmod.IterList([]interface{}{1, 2, 3, 4})

// use Channel
ch := make(chan interface{}, 4)
for i := 0; i < 4; i++ {
    ch <- i
}
close(ch)
iterable2 := funcmod.IterChan(ch)
```

## Convert Iterable to list
```
funcmod.List(iterable1)
```

## Map

```
mfunc := func(x interface{}) interface{} {
    v := x.(int)
    v++
    return v
}
res := funcmod.Map(mfunc, funcmod.IterList([]interface{}{1, 2, 3, 4}))
fmt.Println(funcmod.List(res))
```

## Filter

```
f := func(x interface{}) bool { v := x.(int); return v%2 == 0 }
res := funcmod.Filter(f, funcmod.IterList([]interface{}{1, 2, 3, 4}))
fmt.Println(funcmod.List(res))
```

## Reduce

```
rf := func(x, y interface{}) interface{} {
    a := x.(int)
    b := y.(int)
    return a + b
}
res := funcmod.Reduce(funcmod.IterList([]interface{}{1, 2, 3, 4}), rf, 0)
fmt.Println(res)
```

## Chain

```
res = funcmod.Chain(funcmod.IterList([]interface{}{1, 2, 3, 4}), funcmod.IterList([]interface{}{1, 2, 3, 4}))
fmt.Println(funcmod.List(res))
```
