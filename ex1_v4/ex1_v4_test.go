package ex1_v4
import "testing"

/*
PS C:\Users\Randy\Desktop\go_dojo\euler\ex1_v4> go test -bench=V4
goos: windows
goarch: amd64
BenchmarkV42-8           1000000              1305 ns/op
testing: BenchmarkV42-8 left GOMAXPROCS set to 2
BenchmarkV410-8          1000000              1193 ns/op
testing: BenchmarkV410-8 left GOMAXPROCS set to 2
BenchmarkV4100-8         1000000              1513 ns/op
testing: BenchmarkV4100-8 left GOMAXPROCS set to 2
BenchmarkV41000-8         300000              4149 ns/op
testing: BenchmarkV41000-8 left GOMAXPROCS set to 2
BenchmarkV410000-8         50000             32029 ns/op
testing: BenchmarkV410000-8 left GOMAXPROCS set to 2
BenchmarkV4100000-8         5000            258985 ns/op
testing: BenchmarkV4100000-8 left GOMAXPROCS set to 2
PASS
ok      _/C_/Users/Randy/Desktop/go_dojo/euler/ex1_v4   8.806s
*/

func benchmarkV4(size int, b *testing.B){
    for n:=0; n< b.N; n++{  
        v4(size)
    }
}


func BenchmarkV42(b *testing.B)      { benchmarkV4(2, b) }  
func BenchmarkV410(b *testing.B)     { benchmarkV4(10, b) }  
func BenchmarkV4100(b *testing.B)    { benchmarkV4(100, b) }  
func BenchmarkV41000(b *testing.B)   { benchmarkV4(1000, b) }  
func BenchmarkV410000(b *testing.B)  { benchmarkV4(10000, b) }  
func BenchmarkV4100000(b *testing.B) { benchmarkV4(100000, b) }