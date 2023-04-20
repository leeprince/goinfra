package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/spf13/cast"
	"github.com/vcaesar/bitmap"
	"github.com/vcaesar/imgo"
	"strconv"
)

func main() {
	// Mouse()

	// Keyboard()

	// Screen()

	// Bitmap()

	// Opencv()

	Event()

	// Window()
}

func Event() {
	add()
	low()
	event()
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	// evChan := hook.Start()
	// defer hook.End()
	//
	// for ev := range evChan {
	//     fmt.Println("hook: ", ev)
	// }
}

func event() {
	ok := hook.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	keve := hook.AddEvent("k")
	if keve {
		fmt.Println("you press... ", "k")
	}

	mleft := hook.AddEvent("mleft")
	if mleft {
		fmt.Println("you press... ", "mouse left button")
	}
}

/* 报错，暂时不测试
# gocv.io/x/gocv
In file included from calib3d.cpp:1:
calib3d.h:5:10: fatal error: opencv2/opencv.hpp: No such file or directory
#include <opencv2/opencv.hpp>
^~~~~~~~~~~~~~~~~~~~
compilation terminated.
*/
// func Opencv() {
//     name := "test.png"
//     name1 := "test_001.png"
//     /*robotgo.SaveCapture(name1, 10, 10, 30, 30)
//     robotgo.SaveCapture(name)*/
//
//     fmt.Print("gcv find image: ")
//     fmt.Println(gcv.FindImgFile(name1, name))
//     fmt.Println(gcv.FindAllImgFile(name1, name))
//
//     bit := bitmap.Open(name1)
//     defer robotgo.FreeBitmap(bit)
//     fmt.Print("find bitmap: ")
//     fmt.Println(bitmap.Find(bit))
//
//     // bit0 := robotgo.CaptureScreen()
//     // img := robotgo.ToImage(bit0)
//     // bit1 := robotgo.CaptureScreen(10, 10, 30, 30)
//     // img1 := robotgo.ToImage(bit1)
//     // defer robotgo.FreeBitmapArr(bit0, bit1)
//     img := robotgo.CaptureImg()
//     img1 := robotgo.CaptureImg(10, 10, 30, 30)
//
//     fmt.Print("gcv find image: ")
//     fmt.Println(gcv.FindImg(img1, img))
//     fmt.Println()
//
//     res := gcv.FindAllImg(img1, img)
//     fmt.Println(res[0].TopLeft.Y, res[0].Rects.TopLeft.X, res)
//     x, y := res[0].TopLeft.X, res[0].TopLeft.Y
//     robotgo.Move(x, y-rand.Intn(5))
//     robotgo.MilliSleep(100)
//     robotgo.Click()
//
//     res = gcv.FindAll(img1, img) // use find template and sift
//     fmt.Println("find all: ", res)
//     res1 := gcv.Find(img1, img)
//     fmt.Println("find: ", res1)
//
//     img2, _, _ := robotgo.DecodeImg("test_001.png")
//     x, y = gcv.FindX(img2, img)
//     fmt.Println(x, y)
// }

func Bitmap() {
	fmt.Println("------------------------Bitmap------------------------")

	bit := robotgo.CaptureScreen(10, 10, 30, 30)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bit)

	fmt.Println("bitmap...", bit)
	img := robotgo.ToImage(bit)
	// robotgo.SavePng(img, "test_1.png")
	robotgo.Save(img, "test_1.png")

	robotgo.ToCBitmap(robotgo.ImgToBitmap(img))

	/*bit2 := robotgo.ToCBitmap(robotgo.ImgToBitmap(img))
	  fx, fy := bitmap.Find(bit2)
	  // fmt.Println("FindBitmap------ ", fx, fy)
	  robotgo.Move(fx, fy)

	  arr := bitmap.FindAll(bit2)
	  fmt.Println("Find all bitmap: ", arr)

	  fx, fy = bitmap.Find(bit)
	  // fmt.Println("FindBitmap------ ", fx, fy)*/

	bitmap.Save(bit, "test.png")
	fmt.Println("------------------------Bitmap -end------------------------")
}

func Screen() {
	fmt.Println("------------------------Screen------------------------")

	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)

	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color---- ", color)

	sx, sy := robotgo.GetScreenSize()
	fmt.Println("get screen size: ", sx, sy)

	bit := robotgo.CaptureScreen(10, 10, 30, 30)
	defer robotgo.FreeBitmap(bit)

	img := robotgo.ToImage(bit)
	imgo.Save("test.png", img)

	// 未生效
	num := robotgo.DisplaysNum()
	for i := 0; i < num; i++ {
		robotgo.DisplayID = i
		img1 := robotgo.CaptureImg()
		path1 := "save_" + strconv.Itoa(i)
		robotgo.Save(img1, path1+".png")
		robotgo.SaveJpeg(img1, path1+".jpeg", 50)

		img2 := robotgo.CaptureImg(10, 10, 20, 20)
		robotgo.Save(img2, "test_"+strconv.Itoa(i)+".png")
	}
	fmt.Println("------------------------Screen -end------------------------")
}

func Keyboard() {
	fmt.Println("------------------------Keyboard------------------------")
	robotgo.TypeStr("Hello World")
	robotgo.TypeStr("だんしゃり", 0, 1)
	// robotgo.TypeStr("テストする")

	robotgo.TypeStr("Hi, Seattle space needle, Golden gate bridge, One world trade center.")
	robotgo.TypeStr("Hi galaxy, hi stars, hi MT.Rainier, hi sea. こんにちは世界.")
	robotgo.Sleep(1)

	// ustr := uint32(robotgo.CharCodeAt("Test", 0))
	// robotgo.UnicodeType(ustr)

	robotgo.KeySleep = 10
	robotgo.KeyTap("enter")
	// robotgo.TypeStr("en")
	robotgo.KeyTap("i", "alt", "cmd")

	arr := []string{"alt", "cmd"}
	robotgo.KeyTap("i", arr)

	robotgo.MilliSleep(10)
	robotgo.KeyToggle("f")
	robotgo.KeyToggle("Q")
	robotgo.KeyToggle("Q", "up")

	robotgo.WriteAll("Test")
	text, err := robotgo.ReadAll()
	if err == nil {
		fmt.Println(text)
	}
	fmt.Println("------------------------Keyboard -end------------------------")
}

func Mouse() {
	fmt.Println("------------------------Mouse------------------------")

	robotgo.MouseSleep = 10

	robotgo.ScrollMouse(10, "up")
	robotgo.ScrollMouse(20, "right")

	robotgo.Scroll(0, -10)
	robotgo.Scroll(100, 0)

	robotgo.MilliSleep(100)
	robotgo.ScrollSmooth(-10, 6)
	// robotgo.ScrollRelative(10, -100)

	robotgo.Move(10, 20)
	robotgo.MoveRelative(0, -10)
	robotgo.DragSmooth(10, 10)

	robotgo.Click("wheelRight")
	robotgo.Click("left", true)
	robotgo.MoveSmooth(100, 200, 1.0, 10.0)

	robotgo.Toggle("left")
	robotgo.Toggle("left", "up")
	fmt.Println("------------------------Mouse -end------------------------")
}

func Window() {
	fmt.Println("------------------------Window------------------------")

	// 进程名称：windows 系统打开 `资源监视器` 查看
	fpid, err := robotgo.FindIds("chrome.exe")
	if err == nil {
		fmt.Println("robotgo.Window(\"Google\") pids... ", fpid)

		if len(fpid) > 0 {
			robotgo.TypeStr("Hi galaxy!", cast.ToInt(fpid[0]))
			robotgo.KeyTap("a", fpid[0], "cmd")

			robotgo.KeyToggle("a", fpid[0])
			robotgo.KeyToggle("a", fpid[0], "up")

			err = robotgo.ActivePID(fpid[0])
			if err != nil {
				fmt.Println("robotgo.ActivePID(fpid[0]) err:", err)
			}

			err = robotgo.Kill(fpid[0])
			if err != nil {
				fmt.Println("robotgo.Kill(fpid[0]) err:", err)
			}
		}
	}

	err = robotgo.ActiveName("chrome")
	if err != nil {
		fmt.Println("robotgo.ActiveName(\"chrome\") err:", err)
	}

	isExist, err := robotgo.PidExists(100)
	if err == nil && isExist {
		fmt.Println("pid exists is", isExist)

		robotgo.Kill(100)
	}

	abool := robotgo.Alert("test", "robotgo")
	if abool {
		fmt.Println("robotgo.ShowAlert ok")
	}

	title := robotgo.GetTitle()
	fmt.Println("robotgo.GetTitle():", title)
	fmt.Println("------------------------Window -end------------------------")
}
