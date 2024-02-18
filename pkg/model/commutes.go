package model

import (
	"fmt"
	"sort"
	"strings"
)

type Commutes map[string]uint8

func (c Commutes) String() string {
	output := make([]string, 0, len(c))

	for date, status := range c {
		item := fmt.Sprintf("%s | %04b", date, status)

		index := sort.Search(len(output), func(i int) bool {
			return output[i] > item
		})

		output = append(output, item)
		copy(output[index+1:], output[index:])
		output[index] = item
	}

	return fmt.Sprintf("%s\n", strings.Join(output, "\n"))
}
