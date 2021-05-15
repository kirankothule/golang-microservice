package utils
import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestBubbleSort(t *testing.T) {
	// Initialization
	ele := []int{4,3,2,1}
	// Execution
	BubbleSort(ele)
	// Validation
	assert.NotNil(t, ele)
	assert.EqualValues(t, 4, len(ele))
	assert.EqualValues(t, 1, ele[0])
	assert.EqualValues(t, 2, ele[1])
	assert.EqualValues(t, 3, ele[2])
	assert.EqualValues(t, 4, ele[3])
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization
	ele := []int{1,2,3,4}
	// Execution
	BubbleSort(ele)
	// Validation
	assert.NotNil(t, ele)
	assert.EqualValues(t, 4, len(ele))
	assert.EqualValues(t, 1, ele[0])
	assert.EqualValues(t, 2, ele[1])
	assert.EqualValues(t, 3, ele[2])
	assert.EqualValues(t, 4, ele[3])
}

func getElement(n int) []int {
	result := make([]int, n)
	i := 0
	for j:=n-1; j>=0; j-- {
		result[i]=j
		i++
	}
	return result
}

func TestGetElement(t *testing.T) {
	ele := getElement(4)
	assert.NotNil(t, ele)
	assert.EqualValues(t, 4, ele[3])
	assert.EqualValues(t, 3, ele[2])
	assert.EqualValues(t, 2, ele[1])
	assert.EqualValues(t, 1, ele[0])
}

func BenchmarkBubbleSort10(b *testing.B){
	ele := getElement(10)
	for i:=0; i<b.N; i++ {
		BubbleSort(ele)
	}
}

func BenchmarkBubbleSort1000(b *testing.B){
	ele := getElement(1000)
	for i:=0; i<b.N; i++ {
		BubbleSort(ele)
	}
}

func BenchmarkBubbleSort100000(b *testing.B){
	ele := getElement(100000)
	for i:=0; i<b.N; i++ {
		BubbleSort(ele)
	}
}