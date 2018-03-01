package ex1_v3

func v3(max int) int{
    sum:=0
    for i:=1; 3*i<max; i+=1{
        sum += i*3
    }
    for i:=1; 5*i<max; i+=1{
        if (i*5)%3!=0{
            sum += i*5
        }
    }
    return sum
}