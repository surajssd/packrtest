## Demoing packr2 bug

You should have `packr2` installed on the machine, by running `go get -u github.com/gobuffalo/packr/v2/packr2`.

#### Working code

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

The code in the repo in the same `GOPATH` is at `f9dae4b780be8e97f707c43bbfd6b694f8b9653a`.

#### Bug reproducing

Now let's vendor the dependency into the repo, which is done in the repo on different branch.

```bash
git checkout vendored-packr2
```

The `Gopkg.toml` file looks like following:

```toml
$ cat Gopkg.toml 
[prune]
  go-tests = true
  non-go = true
  unused-packages = true

[[constraint]]
  name = "github.com/gobuffalo/packr"
  version = "v1.21.9"
```

Build the code

```bash
$ make clean
rm -vf bin/*
removed 'bin/test'
removed 'bin/test-fat'

$ make
go build -o bin/test github.com/schu/test/cmd
cd cmd; packr2
/home/hummer/go/src/github.com/schu/test/cmd
go build -o bin/test-fat github.com/schu/test/cmd
cd cmd; packr2 clean
/home/hummer/go/src/github.com/schu/test/cmd
```

Now running the code like before:

```bash
$ ./bin/test
2019/01/03 16:07:11 stat /home/hummer/go/src/github.com/schu/templates/hello.txt: no such file or directory
```

But it fails while it should run as normal but failed to run after vendoring code.
