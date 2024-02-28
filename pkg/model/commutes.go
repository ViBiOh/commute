package model

import (
	"fmt"
	"sort"
	"strings"
)

type Commutes map[string]uint8

const (
	HomeArrive = 1 << iota
	WorkLeave
	WorkArrive
	HomeLeave
	Commute

	Home = HomeLeave & HomeArrive
	Work = WorkLeave & WorkArrive
	All  = Home & Work
)

func (c Commutes) String() string {
	output := make([]string, 0, len(c))

	for date, status := range c {
		var confidence string

		if status&Commute != 1 || status&All != 1 {
			confidence = "Nice ride 🚲"
		} else if status&Home != 1 {
			confidence = "Work from home"
		} else if status&Work != 1 {
			confidence = "Not from home"
		}

		item := fmt.Sprintf("%s | %05b | %s", date, status, confidence)

		index := sort.Search(len(output), func(i int) bool {
			return output[i] > item
		})

		output = append(output, item)
		copy(output[index+1:], output[index:])
		output[index] = item
	}

	return fmt.Sprintf("%s\n", strings.Join(output, "\n"))
}
