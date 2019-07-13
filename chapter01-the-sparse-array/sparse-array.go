package main

import "fmt"

type SparseArray struct {
	sparseArr []*node
}

type node struct {
	Row int
	Col int
	Val int
}

func (s *SparseArray) Convert(a [][]int, defaultVal int) {
	var (
		length     = len(a)
		width      = len(a[0])
		sparseArr  = make([]*node, 0)
		sparseNode *node
	)
	// first row: default value
	sparseNode = &node{
		Row: length,
		Col: width,
		Val: defaultVal,
	}
	// others
	sparseArr = append(sparseArr, sparseNode)
	for i := range a {
		for j := range a[i] {
			if a[i][j] != defaultVal {
				sparseNode = &node{
					Row: i,
					Col: j,
					Val: a[i][j],
				}
				sparseArr = append(sparseArr, sparseNode)
			}
		}
	}
	s.sparseArr = sparseArr
}

func (s *SparseArray) Restore() (a [][]int) {
	a = make([][]int, 0)
	firstNode := s.sparseArr[0]
	// init as all default
	for i := 0; i < firstNode.Row; i ++ {
		subArr := make([]int, 0)
		for j := 0; j < firstNode.Col; j ++ {
			subArr = append(subArr, firstNode.Val)
		}
		a = append(a, subArr)
	}
	//
	for i := 1; i < len(s.sparseArr); i ++ {
		a[s.sparseArr[i].Row][s.sparseArr[i].Col] = s.sparseArr[i].Val
	}
	return a
}

func (s *SparseArray) print() {
	for i := range s.sparseArr {
		fmt.Printf("%+v\n", s.sparseArr[i])
	}
}
