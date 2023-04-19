package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/19 01:29
 * @Desc:
 */

/*
   #cgo CFLAGS: -mmacosx-version-min=10.8
   #cgo LDFLAGS: -framework Carbon -mmacosx-version-min=10.8

   #include <Carbon/Carbon.h>

   WindowRef getWindowByTitle(const char* title) {
       WindowRef window = NULL;
       CFStringRef cfTitle = CFStringCreateWithCString(NULL, title, kCFStringEncodingUTF8);
       if (cfTitle != NULL) {
           CFArrayRef windowList = CGWindowListCopyWindowInfo(kCGWindowListOptionOnScreenOnly, kCGNullWindowID);
           if (windowList != NULL) {
               CFIndex count = CFArrayGetCount(windowList);
               for (CFIndex i = 0; i < count; i++) {
                   CFDictionaryRef windowInfo = CFArrayGetValueAtIndex(windowList, i);
                   CFStringRef windowName = CFDictionaryGetValue(windowInfo, kCGWindowName);
                   if (windowName != NULL && CFStringCompare(windowName, cfTitle, 0) == kCFCompareEqualTo) {
                       int windowNumber = 0;
                       CFNumberRef windowNumberRef = CFDictionaryGetValue(windowInfo, kCGWindowNumber);
                       if (windowNumberRef != NULL) {
                           CFNumberGetValue(windowNumberRef, kCFNumberIntType, &windowNumber);
                           window = GetWindowFromCGWindowID(windowNumber);
                           break;
                       }
                   }
               }
               CFRelease(windowList);
           }
           CFRelease(cfTitle);
       }
       return window;
   }

   WindowRef GetFrontWindowOfClass(OSType windowClass) {
       WindowRef window = NULL;
       ProcessSerialNumber psn = { 0, kCurrentProcess };
       while (GetNextProcess(&psn) == noErr) {
           WindowRef windowList = NULL;
           if (CopyWindowList(psn, &windowList) == noErr) {
               WindowRef currentWindow = windowList;
               while (currentWindow != NULL) {
                   if (GetWindowClass(currentWindow) == windowClass) {
                       window = currentWindow;
                       break;
                   }
                   currentWindow = GetNextWindow(currentWindow);
               }
               DisposeWindowList(windowList);
           }
           if (window != NULL) {
               break;
           }
       }
       return window;
   }

   WindowRef GetWindowFromCGWindowID(int windowID) {
       WindowRef window = NULL;
       CGWindowID cgWindowID = (CGWindowID)windowID;
       CGWindowListOption option = kCGWindowListOptionIncludingWindow;
       CFArrayRef windowList = CGWindowListCopyWindowInfo(option, cgWindowID);
       if (windowList != NULL) {
           CFIndex count = CFArrayGetCount(windowList);
           if (count > 0) {
               CFDictionaryRef windowInfo = CFArrayGetValueAtIndex(windowList, 0);
               CFNumberRef windowNumberRef = CFDictionaryGetValue(windowInfo, kCGWindowNumber);
               if (windowNumberRef != NULL) {
                   int windowNumber = 0;
                   CFNumberGetValue(windowNumberRef, kCFNumberIntType, &windowNumber);
                   window = GetWindowFromPort((CGrafPtr)windowNumber);
               }
           }
           CFRelease(windowList);
       }
       return window;
   }

   void simulateClick(CGPoint point) {
       CGEventRef event = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseDown, point, kCGMouseButtonLeft);
       CGEventPost(kCGHIDEventTap, event);
       CFRelease(event);

       event = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseUp, point, kCGMouseButtonLeft);
       CGEventPost(kCGHIDEventTap, event);
       CFRelease(event);
   }

   void simulateKey(char key) {
       CGEventRef event = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)0, true);
       UniChar ch = (UniChar)key;
       CGEventKeyboardSetUnicodeString(event, 1, &ch);
       CGEventPost(kCGHIDEventTap, event);
       CFRelease(event);
   }

   void simulateShortcut(unsigned short key) {
       CGEventRef event = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)key, true);
       CGEventSetFlags(event, kCGEventFlagMaskCommand);
       CGEventPost(kCGHIDEventTap, event);
       CFRelease(event);

       event = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)key, false);
       CGEventSetFlags(event, kCGEventFlagMaskCommand);
       CGEventPost(kCGHIDEventTap, event);
       CFRelease(event);
   }
*/

import (
	"C"
	"time"
)

func main() {
	// 查找目标窗口
	// 需要注意的是，窗口标题可能会因为不同的应用程序而有所不同，需要根据实际情况进行调整。
	title := "目标窗口标题"
	windowRef := C.getWindowByTitle(C.CString(title))
	if windowRef == nil {
		panic("未找到目标窗口")
	}
	// 将目标窗口置于前台
	C.SelectWindow(windowRef)

	// 模拟点击操作
	point := C.CGPoint{X: 300, Y: 300}
	C.simulateClick(point)

	// 模拟输入操作
	C.simulateKey(C.char('A'))
	C.simulateKey(C.char('B'))

	// 模拟快捷键操作
	C.simulateShortcut(C.kVK_ANSI_C)
	time.Sleep(time.Millisecond * 100)
	C.simulateShortcut(C.kVK_Return)
}
