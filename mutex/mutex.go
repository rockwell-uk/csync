package mutex

var (
	c = make(chan int, 1)
)

func init() {
	c <- 1
}

func Lock() {
	<-c
}

func Unlock() {
	c <- 1
}
