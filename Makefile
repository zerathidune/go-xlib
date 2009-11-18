xlib-test: xlib.6 test.6 xauth.6
	6l -o xlib-test test.6 xlib.6
test.6: test.go
	6g -I. test.go
xlib.6: xlib.go
	6g xlib.go
xauth.6:
	6g xauth.go
clean:
	-rm *.6 xlib-test
