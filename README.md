# go-static-file-server
Static file server with dead time based on Golang

If you want the most simple way to share some local directories for a specific period of time, you have to use it.

### How to use it:

```commandline
git clone https://github.com/alknopfler/go-static-file-server.git
```

After that, you could build the project if you want change something like the port to use, or build for another distro (based on linux right now):
```commandline
go build -o <binary-name> .
```

or maybe you could use by default (linux binary pre-built):

```commandline
$ ./go-static-fileserver 
```

### Options available:

- Port: By default is 8886 but you could change using option p=xxxx
- Directory: By default ./ but you could change using option d=/xx/xx
- TimeToShare: By deault 10min but you could change using option t=10m
    - Just to keep in mind, the time should be a number + "s, m, h" depending on the period of time you want to use (seconds, minutes, or hours)
    
### Examples:

```commandline
$ ./go-static-fileserver -p=8888 -d=/home/xxxx -t=1h
```