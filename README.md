# GO scrips (Vagrant box)

GO scipts assignment + Vagrant box 

(This Vagrant box contains the GO environment installed)


- [x]  A simple 'hello world' script - hello.go
- [x]  Travis check for our 'hello world' script - check.go
- [x]  Setup travis to compile a 'hello world' binary on release
- [x]  Visual check that the compiled binary got uploaded during a release

```
vagrant up
vagrant ssh
go run go/src/hello/hello.go
go install -v hello && go test -v hello
exit
vagrant destroy
```

