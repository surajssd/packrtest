## Demoing packr2 bug

You should have `packr2` installed on the machine, by running `go get -u github.com/gobuffalo/packr/v2/packr2`.

Follow steps below to see how this particular code works:

Build code

```bash
$ make clean
rm -vf bin/*

$ make
go build -o bin/test github.com/schu/test/cmd
cd cmd; packr2
/home/hummer/go/src/github.com/schu/test/cmd
go build -o bin/test-fat github.com/schu/test/cmd
cd cmd; packr2 clean
/home/hummer/go/src/github.com/schu/test/cmd
```

See what file `hello.txt` has

```bash
$ cat templates/hello.txt
hello
```
Both binaries work fine

```bash
$ ./bin/test
hello
Bye.txt
hello.txt

$ ./bin/test-fat 
hello
Bye.txt
hello.txt
```

Let's add some content in the file:

```bash
$ echo hello >>templates/hello.txt

# has picked up changes from local
$ ./bin/test
hello
hello
Bye.txt
hello.txt

# read from memory
$ ./bin/test-fat 
hello
Bye.txt
hello.txt
```

Works even when the binary is copied somewhere else

```
$ cp bin/* /tmp/

$ /tmp/test
hello
hello
Bye.txt
hello.txt

$ /tmp/test-fat 
hello
Bye.txt
hello.txt
```
