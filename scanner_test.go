package scanner

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var names = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}

type NamePrint struct {
	s    *Scanner
	name string
}

func (np *NamePrint) Action() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("has been complete: ", np)
}

func runTask(task Task) {
	reflect.ValueOf(task).MethodByName("Action").Call([]reflect.Value{})
}

func TestPrintName(t *testing.T) {
	s := New(5, runTask)
	defer s.Close()
	for _, name := range names {
		np := NamePrint{s, name}
		s.PushTask(&np)
	}
}
