package ex1_v1

func v1(max int) int{
    
    sum:=0

    for i:=0; i<max; i++{

        if i%3==0 || i%5==0{
            sum+=i
        }
    }

    return sum
}