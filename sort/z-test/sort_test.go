package z_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func setupTestSorts() {
	fmt.Printf("length: %d\n", length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		testItem[i] = r.Intn(length << 6)
	}
}

func TestSorts(t *testing.T) {
	setupTestSorts()

	for _, st := range subTests {
		t.Run(st.name, func(t *testing.T) {
			actualTestItem := make([]int, length)
			copy(actualTestItem, testItem)

			counter := st.function(actualTestItem)
			t.Log(counter)

			validation(t, actualTestItem)
		})
	}
}

func validation(t *testing.T, result []int) {
	for i := 1; i < length; i++ {
		if result[i] < result[i-1] {
			t.Fatalf("Unexpected value detected at %d", i)
		}
	}
}
