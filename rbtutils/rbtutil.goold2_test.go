package rbtutils

// go test
// go test -v
// go test -cover 测试代码覆盖率
// go test -bench=Reverse  性能基准测试
//
//import "testing"
//
//func TestReverse(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
//	got := Reverse("1234567890") // 调用程序，会得到一个结果
//	want := "0987654321"         // 我希望返回的结果
//	if got != want {             // 如果不相等。对于不能直接比较的类型，要用reflect.DeepEqual(want, got)
//		t.Errorf("错误！期望得到：%v ；实际的到：：%v 。\n", want, got)
//	}
//	got = Reverse("12我是34我的5678心情90可以") // 调用程序，会得到一个结果
//	want = "以可09情心8765的我43是我21"         // 我希望返回的结果
//	if got != want {                    // 如果不相等。对于不能直接比较的类型，要用reflect.DeepEqual(want, got)
//		t.Errorf("错误！期望得到：%v ；实际的到：：%v 。\n", want, got)
//	}
//}
//
//// 性能基准测试
//// go test -bench=Reverse
//func BenchmarkReverse(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Reverse("12我是34我的5678心情90可以")
//	}
//}
