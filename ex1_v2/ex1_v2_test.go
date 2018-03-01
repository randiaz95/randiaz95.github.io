package ex1_v2
import "testing"

/*
PS C:\Users\Randy\Desktop\go_dojo\euler\ex1_v2> go test -bench=V2
goos: windows
goarch: amd64
BenchmarkV22-8           1000000            1866 ns/op
BenchmarkV210-8          1000000            1852 ns/op
BenchmarkV2100-8         1000000            2411 ns/op
BenchmarkV21000-8         200000            7270 ns/op
BenchmarkV210000-8         50000            32070 ns/op
BenchmarkV2100000-8         5000            295844 ns/op
PASS
ok      _/C_/Users/Randy/Desktop/go_dojo/euler/ex1_v2   11.413s
*/

func benchmarkV2(size int, b *testing.B){
    for n:=0; n< b.N; n++{  
        v2(size)
    }
}


func BenchmarkV22(b *testing.B)      { benchmarkV2(2, b) }  
func BenchmarkV210(b *testing.B)     { benchmarkV2(10, b) }  
func BenchmarkV2100(b *testing.B)    { benchmarkV2(100, b) }  
func BenchmarkV21000(b *testing.B)   { benchmarkV2(1000, b) }  
func BenchmarkV210000(b *testing.B)  { benchmarkV2(10000, b) }  
func BenchmarkV2100000(b *testing.B) { benchmarkV2(100000, b) }