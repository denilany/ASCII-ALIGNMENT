package functions

import (
	"syscall"
	"unsafe"
)

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getTerminalSize() (int, int, error) {
	ws := &Winsize{}
	// TIOCGWINSZ is the ioctl request code to get terminal window size
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if err != 0 {
		return 0, 0, err
	}
	return int(ws.Col), int(ws.Row), nil
}
