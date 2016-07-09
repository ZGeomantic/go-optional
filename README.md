# go-optional

##Quick Start
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

	var parmInt int = optional.Optional(nilInt).(int)
	var parmStr string = optional.Optional(nilStr).(string)
	var parmFoo Foo = optional.Optional(nilFoo).(Foo)

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
BenchmarkGetIntNil-4      	20000000	        75.2 ns/op
BenchmarkGetInt-4         	20000000	        74.3 ns/op
BenchmarkGetIntNativeNil-4	2000000000	         0.50 ns/op
BenchmarkGetIntNative-4   	2000000000	         1.02 ns/op
BenchmarkGetStrNil-4      	20000000	        95.1 ns/op
BenchmarkGetStr-4         	20000000	        95.7 ns/op
BenchmarkGetStrNativeNil-4	2000000000	         0.50 ns/op
BenchmarkGetStrNative-4   	2000000000	         1.04 ns/op
BenchmarkGetFooNil-4      	20000000	        97.0 ns/op
BenchmarkGetFoo-4         	20000000	        96.4 ns/op
BenchmarkGetStrFooNil-4   	2000000000	         0.50 ns/op
BenchmarkGetStrFoo-4      	2000000000	         1.03 ns/op
```

