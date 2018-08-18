package main
import "fmt"

//func to calculate all the sum of squared of numbers
func arrsquaredsum(arr []int,m int) int {
  
  if (m == 0) {
        return 0;
    }
  var temp int;
  temp = 0
  if (arr[m-1] > 0) {
  temp = arr[m-1]*arr[m-1];
  }
  if (m > 0) {// when array has more than 0 values
    temp += arrsquaredsum(arr,m-1);
  }
  return temp;
}
  

var a []int;
// function to read all integers from a line. Enter digits with a space
func read(x int) {
    if (x == 0) {
        return
    }
    var i int
    _, err := fmt.Scanf("%d",&i)
    if err != nil {
        return
    }
    a = append(a, i)
    //fmt.Println(a)
    read(x - 1)
  }

// Function to specify number of digits to utilize
func challenge (n int) {
  fmt.Println("Enter the number of inputs you would like");
  var x int;
  fmt.Scanf("%d\n",&x);
  a = nil;
  fmt.Println("enter the number(s) whose sum of square you would like");
  read(x)
  fmt.Println("The sum of squared ofpositive integers is",arrsquaredsum(a,x));
  if (n != 0) {
    challenge(n-1);
  }
}



func main() {
  
  fmt.Println("How many lines of input do you want");
  var n int;
  fmt.Scanf("%d\n", &n);
  challenge(n-1);
}