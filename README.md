# Go-Scanner

The following tool is an application to scan a host and port using the `net` package inside Golang


## Testing:
In the root:
- To run the code:
```bash
$ go run .\main.go --host ${host} --port ${port}
```

For example:
```bash
$ go run .\main.go --host 34.30.38.100 --port 3306
```

## Default Arguments:
- host = 127.0.0.1 (localhost)
- port = 3306


## Testing environment:
To avoid firewall issues on a Windows machine:
1) Turned up a virtual machine on GCP (EC2-Micro)
2) Set up firewall rules to accept `TCP connections` on `port 3306`. Also, set `tags` to `mysql` for accepting connections.
3) SSH into the system and run the following commands to set up the VM for using `Docker`
    ```bash
    $ sudo apt install git python3 python3-pip docker.io
    $ sudo docker pull MySQL
    $ vim docker (set up password: admin)
    ```
  
3) Check Docker compose yml for MySQL's latest instance information

4) Run docker-compose in super user:
   ```bash
   sudo docker-compose up -d
   ```

6) Once the image is up and running, `extract the IP address` of the instance from the dashboard or SSH using:
   ```bash
   $ ip address
   ```

7) Pass host and port inside the Go program interface


## Outcome (Based on my testing):
![image](https://github.com/MananDoshi1301/Go-Scanner/assets/65040749/dd112f25-04c4-4331-b481-063271335e41)

![image](https://github.com/MananDoshi1301/Go-Scanner/assets/65040749/7ba3b35c-88df-4144-ab3a-a706cb7a1d2e)



## Program Order
1) Using the `flag` package, input host and port from the user using the command line
2) Using the `DialTimeout` Function, try and establish a connection within 2 seconds using `TCP/IP protocol`
3) If the connection is not established defer with the error command.
4) If the connection is established, send the default `0x0a` initial handshake packet. If this fails, log the error and return
5) If the connection packet is successfully transmitted, read from the connection using a buffer slice of size 1024 [Memory Efficiency Tradeoff limit]. 
6) Once we have the data in byte form, we use the `byte` package to extract the version by using the IndexByte function
7) Finally we log all the information about the address, version number, and the complete data received from handshake 



### Time required to build the scanner: 40 mins with Docker instance
