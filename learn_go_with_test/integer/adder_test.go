package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add test", func(t *testing.T) {
		want := 5
		got := Add(1, 4)
		if want != got {
			t.Errorf("want %v but got %v", want, got)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}
