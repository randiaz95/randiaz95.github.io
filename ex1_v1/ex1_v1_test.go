package ex1_v1
import "testing"

/*
PS C:\Users\Randy\Desktop\go_dojo\euler\ex1_v1> go test -bench=V1
goos: windows
goarch: amd64
BenchmarkV12-8          300000000           5.67 ns/op
BenchmarkV110-8         100000000           20.4 ns/op
BenchmarkV1100-8        10000000            216 ns/op
BenchmarkV11000-8        1000000            2050 ns/op
BenchmarkV110000-8        100000            20424 ns/op
BenchmarkV1100000-8        10000            204747 ns/op
PASS
ok      _/C_/Users/Randy/Desktop/go_dojo/euler/ex1_v1   13.242s
*/

func benchmarkV1(size int, b *testing.B){
    for n:=0; n< b.N; n++{  
        v1(size)
    }
}


func BenchmarkV12(b *testing.B)      { benchmarkV1(2, b) }  
func BenchmarkV110(b *testing.B)     { benchmarkV1(10, b) }  
func BenchmarkV1100(b *testing.B)    { benchmarkV1(100, b) }  
func BenchmarkV11000(b *testing.B)   { benchmarkV1(1000, b) }  
func BenchmarkV110000(b *testing.B)  { benchmarkV1(10000, b) }  
func BenchmarkV1100000(b *testing.B) { benchmarkV1(100000, b) }