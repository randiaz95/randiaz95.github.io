package ex1_v3
import "testing"

/*
PS C:\Users\Randy\Desktop\go_dojo\euler\ex1_v3> go test -bench=V3
goos: windows
goarch: amd64
BenchmarkV32-8          500000000            3.47 ns/op
BenchmarkV310-8         200000000            6.61 ns/op
BenchmarkV3100-8        30000000             44.8 ns/op
BenchmarkV31000-8        3000000             466 ns/op
BenchmarkV310000-8        300000             4449 ns/op
BenchmarkV3100000-8        30000             44029 ns/op
PASS
ok      _/C_/Users/Randy/Desktop/go_dojo/euler/ex1_v3   10.621s
*/

func benchmarkV3(size int, b *testing.B){
    for n:=0; n< b.N; n++{  
        v3(size)
    }
}


func BenchmarkV32(b *testing.B)      { benchmarkV3(2, b) }  
func BenchmarkV310(b *testing.B)     { benchmarkV3(10, b) }  
func BenchmarkV3100(b *testing.B)    { benchmarkV3(100, b) }  
func BenchmarkV31000(b *testing.B)   { benchmarkV3(1000, b) }  
func BenchmarkV310000(b *testing.B)  { benchmarkV3(10000, b) }  
func BenchmarkV3100000(b *testing.B) { benchmarkV3(100000, b) }