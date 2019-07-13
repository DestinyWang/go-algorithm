package main

import (
	"github.com/coreos/etcd/pkg/testutil"
	"testing"
)

func TestSparseArray_Convert(t *testing.T) {
	arr := make([][]int, 0)
	for i := 0; i < 11; i++ {
		subArr := make([]int, 0)
		arr = append(arr, subArr)
		for j := 0; j < 11; j++ {
			arr[i] = append(arr[i], 0)
		}
	}
	arr[2][3] = 1
	arr[2][4] = 2
	
	sa := &SparseArray{
	
	}
	
	sa.Convert(arr, 0)
	sa.print()
	
	restoreArr := sa.Restore()
	testutil.AssertEqual(t, arr, restoreArr)
	
}
