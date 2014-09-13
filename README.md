Gundam
======

Gundam is the most destructive robot! It conquers the world with [Go](http://golang.org) and [Sphero](http://www.gosphero.com/).

This is forked from [jingweno/gundam](https://github.com/jingweno/gundam) ( [YouTube Video](https://www.youtube.com/watch?v=BcaqdXh566E) )

Updated this repo to latest Sphero library.

###Latest library used from [gobot.io](http://gobot.io)

```
github.com/hybridgroup/gobot
github.com/hybridgroup/gobot/platforms/sphero
```

###Run

```
$ go build
$ PORT=/dev/tty.Sphero-RGG-AMP-SPP ./gundam #This is my device port, change to yours
2014/01/31 13:51:09 Initializing connections...
2014/01/31 13:51:09 Initializing connection  sphero ...
2014/01/31 13:51:09 Initializing devices...
2014/01/31 13:51:09 Initializing device  Gundam ...
2014/01/31 13:51:09 Starting connections...
2014/01/31 13:51:09 Starting connection sphero...
2014/01/31 13:51:09 Connecting to sphero on port /dev/tty.Sphero-ORY-AMP-SPP...
[martini] listening on host:port :3000
Starting connections...
Starting connection sphero...
Connecting to sphero on port /dev/tty.Sphero-ORY-AMP-SPP...
Starting devices...
Starting device sphero...
Device sphero started

$ curl -X PUT localhost:3000/rgb/255,0,0 #change color
ok
```

###Scripts - change colors every 1 second

```
$ ./ranbow.sh
or 
$ ruby infinity_ranbow.rb
```

###TODO 

* Add motions