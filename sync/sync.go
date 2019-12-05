//author: wongoo
//date: 20191105

package vsync

func IsChanClosed(c chan struct{}) bool {
	if c == nil {
		return true
	}
	select {
	case <-c:
		return true
	default:
		return false
	}
}
