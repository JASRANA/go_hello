package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// 使用接口 Shape 达到高度解耦合
	//checkArea := func(t *testing.T, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %.2f want %.2f", got, want)
	//	}
	//}
	//
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{10.0, 10.0}
	//	checkArea(t, rectangle, 100.0)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10.0}
	//	checkArea(t, circle, 314.1592653589793)
	//})

	// 表格驱动测试
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{Width: 12, Length: 6}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			// 使用 %#v 标识结构体结构（或者在表格中增加 name 字段标明；或者是用 t.Run 标明）
			t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
		}
	}
}