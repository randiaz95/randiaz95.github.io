points_Line <- function(a,b,c,d)
{
 slope = (d-b)/(c-a)
 distance = sqrt((c-a)^2+(d-b)^2)
 yint = b-a*slope
 xint = -1*yint/slope
 
 print("Slope")
 print(slope)
 print("Distance")
 print(distance)
 print("yint")
 print(yint)
 print("xint")
 print(xint)
}
