package main

const (
	Lower = 1
	Upper = 300
)

func main() {
	// TODO: Implement FizzBuzz using a for loop from Lower to Upper.
	for i := Lower; i < Upper; i++ {
		if i%5==0 && i%3==0{
			println("FizzBuzz")
		}else if i%5==0{
			println("Buzz")
		}else if i%3==0{
			println("Fizz")
		}else{
			println(i)
		}
	}
}
