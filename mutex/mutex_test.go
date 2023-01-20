package mutex

import (
	"testing"

	"github.com/rockwell-uk/csync/waitgroup"
)

func TestWaitgroup(t *testing.T) {

	var ctr int
	var loops int = 1000

	for i := 0; i < loops; i++ {
		Lock()
		ctr++
		Unlock()
	}

	if ctr != loops {
		t.Errorf("Expected %v Got %v", loops, ctr)
	}
}

func TestWaitgroupFanout(t *testing.T) {

	var wg *waitgroup.WaitGroup = waitgroup.New()

	var ctr int
	var loops int = 1000

	for i := 0; i < loops; i++ {

		wg.Add(1)
		go func() {
			Lock()
			ctr++
			Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	if ctr != loops {
		t.Errorf("Expected %v Got %v", loops, ctr)
	}
}

func BenchmarkWaitgroup(b *testing.B) {

	var ctr int

	for i := 0; i < b.N; i++ {
		Lock()
		ctr++
		Unlock()
	}
}
