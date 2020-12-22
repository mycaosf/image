name = img
dllName = $(name).dll
all: $(dllName)

sources = main.go

objsDll =
libsDll =  ./$(name).a -l msvcrt -l ws2_32 -l winmm
defFile = $(name).def

# Debug flag: yes, no
debug = no

ifeq ("$(debug)", "yes")
flags = -gcflags "-N -l"
else
flags = -ldflags "-s -w"
endif

ifndef x64
ifeq ("$(Platform)", "")
x64 = 0
else
ifeq ($(Platform), x64)
x64 = 1
else
x64 = 0
endif
endif
endif

ifeq ($(x64), 1)
	machine := X64
else
	machine := X86
endif

$(defFile):
	@echo "LIBRARY $(dllName)" >$@
	@echo "EXPORTS" >>$@
	@echo "  ImageEncode @1" >> $@

$(name).a: $(sources)
	go build $(flags) -buildmode=c-archive

$(dllName): $(objsDll) $(name).a $(defFile)
	dllwrap -o $@ --no-export-all-symbols --def $(defFile) $(objsDll) $(libsDll)
	lib -out:$(name).lib -def:$(defFile) -machine:$(machine) -nologo
	strip $@

clean:
	rm -f *.a
	rm -f *.lib
	rm -f *.def
	rm -f *.exp
	rm -f $(name).h
	rm -f *.exe

cleanall: clean
	go clean
	rm -f *.dll

