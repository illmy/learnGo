# CSP 并发机制

### Actor Model

![Actor Model](../resource/actor.jpg)

### CSP vs. Actor

* 和Actor的直接通讯不同，CSP模式则是通过Channel进⾏通讯的，更松耦合⼀
些。
* Go中channel是有容量限制并且独⽴于处理Groutine，⽽如Erlang，Actor模式
中的mailbox容量是⽆限的，接收进程也总是被动地处理消息。

### Channel

![Channel](../resource/channel.png)