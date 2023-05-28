package channels_test

import (
	"sync"
	"testing"

	"github.com/andygeiss/channels"
)

func process(x int) int {
	return x + 1
}

func TestChannels(t *testing.T) {

	stageOne := channels.Generate(1, 2)
	stageTwo := channels.Split(stageOne, 3)
	stageThree0 := channels.Process(stageTwo[0], process)
	stageThree1 := channels.Process(stageTwo[1], process)
	stageThree2 := channels.Process(stageTwo[2], process)
	stageFour := channels.Merge(stageThree0, stageThree1, stageThree2)
	stageFive := channels.Multiplex(stageFour, 3)
	stageSix0 := channels.Process(stageFive[0], process)
	stageSix1 := channels.Process(stageFive[1], process)
	stageSix2 := channels.Process(stageFive[2], process)
	stageSeven := channels.Merge(stageSix0, stageSix1, stageSix2)

	sum := 0
	mutex := sync.Mutex{}

	<-channels.Drain(stageSeven, func(x int) {
		mutex.Lock()
		defer mutex.Unlock()
		sum += x
	})

	if sum != 21 {
		t.Fatalf("sum should be 21 but is %d", sum)
	}
}
