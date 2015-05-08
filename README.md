toolbox
=======

### Requirements

On Linux

```
$ sudo apt-get install golang
$ echo 'GOPATH=~/go' >> ~/.bashrc
$ echo 'PATH=$GOPATH/bin:$PATH' >> ~/.bashrc
$ source ~/.bashrc
```

On OSX

```
$ brew install go
$ echo 'GOPATH=~/go' >> ~/.bash_profile
$ echo 'PATH=$GOPATH/bin:$PATH' >> ~/.bash_profile
$ source ~/.bash_profile
```

### Tools

* [httpserver](httpserver) A simple http server
* [sendmail](sendmail) A tool for sending email
* [sshexec](sshexec) Execute remote commands through ssh
