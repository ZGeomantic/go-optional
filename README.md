# go-optional

##Quick Start

if you want to dereference a pointer, usually you neet to do this:

```go
	var a *int
	var b int
	if a != nil {
		b = *a
	} else {
		b = 0
	}

	var c *Foo
	var d Foo
	if c != nil {
		d = *c
	} else {
		d = Foo{}
	}
```
now you can do like this:

```go
	var a *int
	var b int
	b = optional.Deref(a).(int)
	b = optional.DerefMust(a,100).(int)

	var c *Foo
	var d Foo
	d = optional.Deref(c).(Foo)
```

######Download and install

    go get github.com/ZGeomantic/go-optional

######Usage

```go
package main

import "github.com/ZGeomantic/go-optional"


func main() {
	var (
		nilInt *int
		nilStr *string
		nilFoo *Foo
	)

	var parmInt int = optional.Deref(nilInt).(int)
	var parmStr string = optional.Deref(nilStr).(string)
	var parmFoo Foo = optional.Deref(nilFoo).(Foo)

	fmt.Printf("parmInt: %d, parmStr: %s, parmFoo: %+v\n", parmInt, parmStr, parmFoo)
}

```
##Note

Compare to the native way to use ```if...else```, the ```optioinal``` method is slower, but it depends. When the pointer is nill, it's about 150-200 times slower.
When the pointer is not nil, it's about 75-90 times slower. It's your choice if you can take the cost or you want the best performance. After all the time spend isn't really unacceptable, time unit is in ns.

You can run the benchmark on your computer by excute:
```go test -bench="Get.*"```

 the result on my computer is as follow:

```shell
BenchmarkGetIntNil-4      	20000000	       104 ns/op
BenchmarkGetInt-4         	10000000	       119 ns/op
BenchmarkGetIntNativeNil-4	2000000000	         0.50 ns/op
BenchmarkGetIntNative-4   	2000000000	         1.01 ns/op
BenchmarkGetStrNil-4      	10000000	       123 ns/op
BenchmarkGetStr-4         	10000000	       137 ns/op
BenchmarkGetStrNativeNil-4	2000000000	         0.50 ns/op
BenchmarkGetStrNative-4   	2000000000	         1.05 ns/op
BenchmarkGetFooNil-4      	10000000	       123 ns/op
BenchmarkGetFoo-4         	10000000	       137 ns/op
BenchmarkGetStrFooNil-4   	2000000000	         0.51 ns/op
BenchmarkGetStrFoo-4      	2000000000	         1.02 ns/op
```

