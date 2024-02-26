GOC?=go build
TARGET?=virus.exe
BSOD_TARGET?=bsod.exe
SRC?=main.go
BSOD_SRC?=bsod.go
LDFLAGS?=-ldflags -H=windowsgui
THIS_FILE := $(lastword $(MAKEFILE_LIST))
MAKE?=make

.PHONY: clean
all: virus.exe

virus.exe: bsod.exe
	GOOS=windows GOARCH=amd64 $(GOC) -o $(TARGET) $(LDFLAGS) $(SRC)
	$(MAKE) -f $(THIS_FILE) bsod-clean

.PHONY: clean
clean:virus-clean bsod-clean

.PHONY: virus-clean
virus-clean:
	rm -f $(TARGET)

bsod.exe:
	GOOS=windows GOARCH=amd64 $(GOC) -o $(BSOD_TARGET) $(LDFLAGS) $(BSOD_SRC)

.PHONY: bsod-clean
bsod-clean:
	rm -f $(BSOD_TARGET)
