package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	state string
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)
		prevAction := rune(0)

		task := Task{state: "NotStarted"}
		valid := true

		for _, action := range s {
			if prevAction == action {
				valid = false
				break
			}
			prevAction = action

			switch action {
			case 'M':
				valid = task.M()
			case 'R':
				valid = task.R()
			case 'C':
				valid = task.C()
			case 'D':
				valid = task.D()
			default:
				valid = false
			}

			if !valid {
				break
			}
		}

		if valid && task.state == "Done" {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func (t *Task) M() bool {
	switch t.state {
	case "NotStarted", "Cancelled", "Done":
		t.state = "Started"
		return true
	default:
		return false
	}
}

func (t *Task) R() bool {
	if t.state == "Started" {
		t.state = "Restarted"
		return true
	}
	return false
}

func (t *Task) C() bool {
	switch t.state {
	case "Started", "Restarted":
		t.state = "Cancelled"
		return true
	default:
		return false
	}
}

func (t *Task) D() bool {
	if t.state == "Started" {
		t.state = "Done"
		return true
	}
	return false
}
