# named-pipe-ipc

Golang parent/child process IPC via named pipes

### What does it do

Once `process_executor` executable is ran, it launches `spawned_process` executable and then it listens for messages from named pipe.
The `spawned_process` process sends one message via named pipe.
Those messages are received by `process_executor` process which prints them to console. 

### How to use it

* `make build` compiles all executable into bin directory
* `make run` runs `process_executor` which in turn runs `spawned_process`

### Sample execution

Note, logs prefixed with timestamp are from `process_executor`, the others are from `spawned_process`.

```
$> make run
2019/08/27 09:42:35 Create named pipe /var/folders/wj/b2xcxqhs4js3t7qf288h5wyjk5cbnl/T/named-pipes785412834
2019/08/27 09:42:35 Opening named pipe for reading
Opening named pipe for writing
Writing
2019/08/27 09:42:35 Reading
2019/08/27 09:42:35 Waiting for someone to write something
2019/08/27 09:42:35 Data: hello
2019/08/27 09:42:35 Successfully deleted named pipe /var/folders/wj/b2xcxqhs4js3t7qf288h5wyjk5cbnl/T/named-pipes785412834/stdout
```