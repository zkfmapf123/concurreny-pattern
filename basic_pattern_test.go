package concurrenypattern

import (
	"testing"
	"zkfmapf123/concurrency/utils"
)

var timer = utils.NewTimeCheck()

func Test_badPattern(t *testing.T) {
	timer.Start()
	basicBadPattern()
	basicBadPattern()
	basicBadPattern()
	basicBadPattern()
	basicBadPattern()
	timer.Finish()
	timer.Result()
}

func Test_goodPattern(t *testing.T) {
	timer.Start()
	basicGoodPattern()
	basicGoodPattern()
	basicGoodPattern()
	basicGoodPattern()
	basicGoodPattern()
	basicGoodPattern()
	timer.Finish()
	timer.Result()
}
