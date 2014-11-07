Coroutines and Channels
===

## Go Routines

A goroutine is a function that is capable of running concurrently with other
functions. To create a goroutine we use the keyword go followed by a function
invocation:

```go
func test() {}

go test()
```

## Channel

Channels provide a way for two goroutines to communicate with each other and
synchronize their execution.

A channel type is represented with keyword `chan` followed by the type of the
things that will be passed by the channel. `<-` is used to send and receive
messages on channel.

```go
func sender(c chan string) {
  for i := 1; i < 10; i++ {
     c <- "some message"
  }
}
func reader(c chan string) {
  for {
    x := <- c  //read message from channel
    fmt.Println(x)
    time.Sleep(100*time.Millisecond)  
  }
}

func main(){
  var c chan string = make(chan string)
  go sender(c)
  go reader(c)

  var input string
  fmt.Scanln(&input)
}
```
When sender want to send a message through channel, it will wait until the
reader is ready to receive the message(this is known as blocking).

With a goroutine we return immediately to the next line and don't wait for the
function to complete. This is why the call to the Scanln function has been
included; without it the program would exit before being given the opportunity
to print all the numbers.

We can also set direction for a channel to use. If no direction set, the channel
can be used to both read and write.

```go
func sender(c chan<- string) {} //only can send
func sender(c <-chan string) {} //only can receive
```
### Select for Channel

You can use `select` to work with multiple channels.

```go
select {
case msg1 := <- c1:
    fmt.Println("Message 1", msg1)
case msg2 := <- c2:
    fmt.Println("Message 2", msg2)
case <- time.After(time.Second):
    fmt.Println("timeout")
}
```
