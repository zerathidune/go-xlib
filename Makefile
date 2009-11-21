XLIB_SOURCES=xlib.go xauth.go

xlib-test: xlib.$(GOCHAR) test.$(GOCHAR)
	$(GOCHAR)l -o xlib-test test.$(GOCHAR) xlib.$(GOCHAR)
test.$(GOCHAR): test.go
	$(GOCHAR)g -I. test.go
xlib.$(GOCHAR): $(XLIB_SOURCES)
	$(GOCHAR)g -I. -o xlib.$(GOCHAR) $(XLIB_SOURCES)
clean:
	-rm *.$(GOCHAR) xlib-test
