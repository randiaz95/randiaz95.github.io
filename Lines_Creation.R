
#This function transforms 2 points into a line formula. 
#Note The larger the distance the less useful in application, due to the fact that most real phenomena are non-linear.

points_Line <- function(a,b,c,d)
{
 slope = (d-b)/(c-a)
 distance = sqrt((c-a)^2+(d-b)^2)
 yint = b-a*slope
 xint = -1*yint/slope
 
 print("")
 print("Slope")
 print(slope)
 print("")
 print("Distance")
 print(distance)
 print("")
 print("yint")
 print(yint)
 print("")
 print("xint")
 print(xint)
 print("")
}
