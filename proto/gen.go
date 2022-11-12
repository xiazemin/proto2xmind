package gen

import "fmt"

func GenFiles(dst string, files ...string) {
	if len(files) < 1 {
		fmt.Println("files empty")
		return
	}

	xmind := NewXmind()
	box1 := NewSandBox(1, xmind)
	maxNum := box1.Proto2Xmind(files[0])
	for i := 1; i < len(files); i++ {
		box2 := NewSandBox(maxNum, xmind)
		maxNum = box2.Proto2Xmind(files[i])
	}
	xmind.Marshal(dst)
}
