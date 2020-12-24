package main

import (
	"reflect"
	"testing"
)


func TestMatrix(t *testing.T) {
	mtx := [][]int{
		{1, 2},
		{4, 3},
	}
	slice := []int{1, 2, 3, 4}
	actual := matrix(mtx)
	if !reflect.DeepEqual(slice,actual){
		t.Errorf("Not equal \nexpected - %v\nactual - %v", slice,actual)
	}
	mtx = [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual = matrix(mtx)
	if !reflect.DeepEqual(slice,actual){
		t.Errorf("Not equal \nexpected - %v\nactual - %v", slice,actual)
	}
	mtx = [][]int{
		{1,2,3,4},
		{12,13,14,5},
		{11,16,15,6},
		{10,9,8,7},
	}
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,14,15,16}
	actual = matrix(mtx)
	if !reflect.DeepEqual(slice,actual){
		t.Errorf("Not equal \nexpected - %v\nactual - %v", slice,actual)
	}

	mtx = [][]int{
		{16,15,14,13},
		{5,4,3,12},
		{6,1,2,11},
		{7,8,9,10},
	}
	slice = []int{16, 15, 14, 13, 12, 11, 10, 9, 8,7,6,5,4,3,2,1}
	actual = matrix(mtx)
	if !reflect.DeepEqual(slice,actual){
		t.Errorf("Not equal \nexpected - %v\nactual - %v", slice,actual)
	}

}