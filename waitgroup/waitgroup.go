package waitgroup

type WaitGroup struct {
	ops chan op
}

type op struct {
	_type string
	res   chan int
	n     int
}

const (
	opCheck = "CHECK"
)

func New() *WaitGroup {

	var n int
	var wg WaitGroup = WaitGroup{
		ops: make(chan op),
	}

	go func(w *WaitGroup) {
		for o := range w.ops {
			n += o.n
			if o._type == opCheck {
				o.res <- n
			}
		}
	}(&wg)

	return &wg
}

func (wg *WaitGroup) Add(i int) {
	wg.ops <- op{n: i}
}

func (wg *WaitGroup) Done() {
	wg.ops <- op{n: -1}
}

func (wg *WaitGroup) Wait() {

	for {
		res := make(chan int)
		wg.ops <- op{opCheck, res, 0}
		if <-res == 0 {
			return
		}
	}
}
