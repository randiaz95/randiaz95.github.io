package ex1_v2
import "sync"

var wg sync.WaitGroup
func sum_multiples(multiple int,
                   end int,
                   sum *int,
                   prev_multiple int){
    defer wg.Done()

    if prev_multiple==-1{
        for i:=multiple; i<end; i+=multiple{
            if i%multiple==0 {
                *sum += i
            }
        }
    }else{
        for i:=multiple; i<end; i+=multiple{
            if i%multiple==0 && i%prev_multiple!=0{
                *sum += i
            }
        }
    }
}

func v2(max int) int{
    sum3:=0
    sum5:=0
    wg.Add(1)
    go sum_multiples(3, 
                     max, 
                     &sum3,
                     -1)
    wg.Add(1)
    go sum_multiples(5, 
                     max, 
                     &sum5,
                     3)
    wg.Wait()
    return sum3+sum5
}