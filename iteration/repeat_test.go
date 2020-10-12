package iteration

import (
	"fmt"
	"testing"
)

func TestReport(t *testing.T) {
	repeated := Repeat("a", 8)
	expected := "aaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// 基准测试：-bench='.'
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}

// 函数文档：-v
func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)

	// Output: aaaaa
}