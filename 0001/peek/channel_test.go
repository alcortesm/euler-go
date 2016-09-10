package peek_test

import (
	"testing"
	"time"

	"github.com/alcortesm/euler-go/0001/peek"
)

func TestPeekOK(t *testing.T) {
	c := make(chan int, 3)
	pc := peek.NewChannel(c)

	c <- 1
	c <- 2

	v, ok := pc.Peek()
	if !ok {
		t.Error("peek returned false on a channel with valid data")
	}
	if v != 1 {
		t.Error("peek returned invalid data, expected = 1, obtained = ", v)
	}
}

func TestPeekOnEmpty(t *testing.T) {
	c := make(chan int, 10)
	pc := peek.NewChannel(c)
	close(c)

	_, ok := pc.Peek()
	if ok {
		t.Error("peek returned true on a closed channel")
	}
}

func TestPeekOnClosedButNotEmpty(t *testing.T) {
	c := make(chan int, 10)
	pc := peek.NewChannel(c)
	c <- 0
	close(c)

	_, ok := pc.Peek()
	if !ok {
		t.Error("peek returned false on a closed channel with pending data")
	}
}

func TestPeekOnWaitingChannel(t *testing.T) {
	c := make(chan int, 10)
	pc := peek.NewChannel(c)
	go func() {
		time.Sleep(100 * time.Millisecond)
		c <- 1
		close(c)
	}()

	_, ok := pc.Peek()
	if !ok {
		t.Error("peek returned false on a closed channel with pending data")
	}
}

func TestRecvOK(t *testing.T) {
	c := make(chan int, 3)
	pc := peek.NewChannel(c)

	c <- 1
	c <- 2
	checkRecv(t, pc, 1, true, "test 01")
	checkRecv(t, pc, 2, true, "test 02")
	c <- 0
	checkRecv(t, pc, 0, true, "test 03")
	go func() {
		time.Sleep(100 * time.Millisecond)
		c <- 4
		c <- 5
		close(c)
	}()
	checkRecv(t, pc, 4, true, "test 04")
	_, _ = pc.Peek()
	checkRecv(t, pc, 5, true, "test 05")
	checkRecv(t, pc, 0, false, "test 06")

}

func checkRecv(t *testing.T, pc *peek.Channel, eVal int, eOk bool, comment string) {
	obtained, ok := pc.Recv()
	if ok != eOk {
		t.Errorf("%s\nexpected OK = %t\nobtained OK = %t",
			comment, eOk, ok)
	}
	if obtained != eVal {
		t.Errorf("%s\nexpected value = %d\nobtained value = %d",
			comment, eVal, obtained)
	}
}
