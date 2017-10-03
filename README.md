# Time-tracker
Simple time tracker (tt). Writes current time and time since last track to terminal (console) every time you press enter.  
Shutdown with `Ctrl+C` (terminate signal).  
Format: `<current time> | <hours*>h <minutes*>m <seconds*>s` __*__ Since last track.  
Example output:
```
Time-tracker 1.0
22:52:44 | 0h 0m 0s
22:52:47 | 0h 0m 3s
Tracker works for 22:52:47 | 0h 0m 3s
```

You can write some text before log, so it will be displayed in terminal.
Example:
```
Time-tracker 1.0
23:13:27 | 0h 0m 0s eat twix
23:13:59 | 0h 0m 32s think
23:14:51 | 0h 0m 52s think again
23:15:21 | 0h 0m 30s fix mistake
23:15:35 | 0h 0m 14s 
23:15:37 | 0h 0m 1s 
Tracker works for 0h 2m 10s
```

## Install
```
go get github.com/vetcher/tt
```

## Run
```
$ tt
```