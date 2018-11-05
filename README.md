# GO scrips (Vagrant box)

GO scipts assignment + Vagrant box 

(This Vagrant box contains the GO environment installed)


- [x]  A simple 'hello world' script - hello.go
- [ ]  Travis check for our 'hello world' script - check.go
- [ ]  Setup travis to compile a 'hello world' binary on release
- [ ]  Visual check that the compiled binary got uploaded during a release

```
vagrant up
vagrant ssh
go run /vagrant/hello.go
exit
vagrant destroy
```
