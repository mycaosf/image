targetName = sample

objs = sample.o

ifdef OS
platform = Windows
CC = cl
LD = link
CFLAGS =
LDFLAGS = -out:
targetName := $(targetName).exe
libs = ../img.lib
objExt = obj
else
platform = Linux
objExt = o
endif
objs := $(subst .o,.$(objExt),$(objs))

%.$(objExt): %.c
	$(CC) -c $<

$(targetName): $(objs)
	$(LD) $(LDFLAGS)$@ $^ $(libs)

clean:
	rm -f *.o
	rm -f *.obj
	rm -f $(targetName)
