```
go get -u github.com/gobuffalo/packr/v2/packr2
make clean
make
cat templates/hello.txt
./bin/test
./bin/test-fat
echo hello >>templates/hello.txt
./bin/test
./bin/test-fat
cp bin/* /tmp/
/tmp/test
/tmp/test-fat
echo hello >>templates/hello.txt
/tmp/test
/tmp/test-fat
rm -rf templates/hello.txt
/tmp/test
/tmp/test-fat
```
