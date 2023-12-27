package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type ByStart []Interval

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

func optimizeIntervals(intervals []Interval) []Interval {
	// Paso 1: Ordenar los intervalos según el primer elemento
	sort.Sort(ByStart(intervals))

	// Paso 2: Fusionar intervalos superpuestos
	var mergedIntervals []Interval
	currentInterval := intervals[0]

	for i := 1; i < len(intervals); i++ {
		nextInterval := intervals[i]

		// Verificar si hay superposición
		if currentInterval.End >= nextInterval.Start {
			// Fusionar intervalos
			currentInterval.End = max(currentInterval.End, nextInterval.End)
		} else {
			// No hay superposición, agregar intervalo actual a la lista
			mergedIntervals = append(mergedIntervals, currentInterval)
			currentInterval = nextInterval
		}
	}

	// Agregar el último intervalo a la lista
	mergedIntervals = append(mergedIntervals, currentInterval)

	return mergedIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

	intervals := []Interval{
		{5, 8},
		{2, 7},
		{3, 4},
	} // [[2, 8]]

	fmt.Println(optimizeIntervals(intervals))
	intervals = []Interval{
		{1, 3},
		{8, 10},
		{2, 6},
	} // [[1, 6], [8, 10]]
	fmt.Println(optimizeIntervals(intervals))
	intervals = []Interval{
		{3, 4},
		{1, 2},
		{5, 6},
	} // [[1, 2], [3, 4], [5, 6]]
	fmt.Println(optimizeIntervals(intervals))

}
