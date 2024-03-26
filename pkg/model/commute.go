package model

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Day struct {
	Date   time.Time
	IDs    []string
	Status uint8
}

type Commutes []Day

// CommutesByDate sort Commutes by Date
type CommutesByDate Commutes

func (a CommutesByDate) Len() int      { return len(a) }
func (a CommutesByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a CommutesByDate) Less(i, j int) bool {
	return a[i].Date.Before(a[j].Date)
}

const (
	Home = HomeLeave | HomeArrive
	Work = WorkLeave | WorkArrive
	All  = Home | Work
)

func (c Commutes) String() string {
	output := make([]string, 0, len(c))

	for _, day := range c {
		item := fmt.Sprintf("%s | %05b", day.Date, day.Status)

		index := sort.Search(len(output), func(i int) bool {
			return output[i] > item
		})

		output = append(output, item)
		copy(output[index+1:], output[index:])
		output[index] = item
	}

	return fmt.Sprintf("%s\n", strings.Join(output, "\n"))
}
