set CGO_CFLAGS=-I%IUPHOME%\include
set CGO_LDFLAGS=-L%IUPHOME%
path=%IUPHOME%;%path%

go install github.com/leonrbaker/iup...
