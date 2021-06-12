package topwords

import (
	"reflect"
	"testing"
)


func TestTopWords(t *testing.T) {

	type args struct {
		s string
		n int
	}

	testTable := []struct {
		data args
		expected [][]string
	} {
		{
			data: args{s: "Train Bus Bus Train Taxi Taxi Aeroplane Taxi Bus Bus", n: 2},
			expected: [][]string{{"Bus", "Taxi"}, {"Taxi", "Bus"}},
		},
		{
			data: args{s: "The result of this string should be: string should result", n: 3},
			expected: [][]string{
				{"string", "should", "result"},
				{"string", "result", "should"},
				{"should", "string", "result"},
				{"should", "result", "string"},
				{"result", "string", "should"},
				{"result", "should", "string"},
			},
		},
	}

	for _, testCase := range testTable {
		res := TopWords(testCase.data.s, testCase.data.n)
		isPassed := false
		for _, val := range testCase.expected {
			if reflect.DeepEqual(res, val) {
				isPassed = true
			}
		}
		if isPassed == false {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, res)
		}
	}
}
