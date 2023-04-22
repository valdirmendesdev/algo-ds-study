package entities

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	Name   string
	Weight int
	Value  int
}

type Knapsack struct {
	MaxWeight int
}

type Problem struct {
	Items    []*Item
	Knapsack *Knapsack
}

func NewProblem(items []*Item, knapsack *Knapsack) *Problem {
	return &Problem{
		Items:    items,
		Knapsack: knapsack,
	}
}

func CreateProblemSample() *Problem {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	nItems := r.Intn(299) + 1
	items := []*Item{}
	for i := 1; i <= nItems; i++ {
		items = append(items, &Item{
			Name:   fmt.Sprintf("Item %v", i),
			Weight: r.Intn(99) + 1,
			Value:  r.Intn(49) + 1,
		})
	}
	k := &Knapsack{
		MaxWeight: 1000,
	}
	return NewProblem(items, k)
}

type Solution struct {
	Problem *Problem
	Genes   []bool

	ItemsCount  int
	TotalWeight int
	TotalValue  int
}

func NewSolution(problem *Problem, genes []bool) *Solution {
	s := &Solution{
		Problem:     problem,
		Genes:       genes,
		ItemsCount:  0,
		TotalWeight: 0,
		TotalValue:  0,
	}

	for i, gene := range genes {
		if gene && ((problem.Items[i].Weight + s.TotalWeight) <= problem.Knapsack.MaxWeight) {
			s.ItemsCount++
			s.TotalWeight += problem.Items[i].Weight
			s.TotalValue += problem.Items[i].Value
		}
	}

	return s
}

func (s *Solution) String() string {
	return fmt.Sprintf("Items: %v, TotalWeight: %v, TotalValue: %v", s.ItemsCount, s.TotalWeight, s.TotalValue)
}

// returns the selected items in the solution
func (s *Solution) SelectedItems() []*Item {
	items := []*Item{}
	for i, gene := range s.Genes {
		if gene {
			items = append(items, s.Problem.Items[i])
		}
	}
	return items
}
