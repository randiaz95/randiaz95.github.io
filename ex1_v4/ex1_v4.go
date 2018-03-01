package ex1_v4
import ("sync"; "runtime")

var wg sync.WaitGroup

func sum_multiples(multiple int,
                   end int,
                   sum *int,
                   prev_multiple int){

    if prev_multiple==-1{
        for i:=1; multiple*i<end; i+=1{
                *sum += i*multiple
        }
    }else{
        for i:=1; multiple*i<end; i+=1{
            if (i*multiple)%prev_multiple!=0{
                *sum += i*multiple
            }
        }
    }
    wg.Done()
}

func v4(max int) int{
    sum:=0
    runtime.GOMAXPROCS(2)
    wg.Add(2)

    go sum_multiples(3, 
                     max, 
                     &sum,
                     -1)
    go sum_multiples(5, 
                     max, 
                     &sum,
                     3)

    wg.Wait()
    return sum
}
