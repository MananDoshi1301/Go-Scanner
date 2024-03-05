# Go-Scanner

## The following tool is a one time application to scan a host and port using the `net` package inside Golang


## Testing:
In the root:
To run the code:
$ go run .\main.go --host ${host} --port ${port}

For example:
$ go run .\main.go --host 34.30.38.100 --port 3306

By default:
host = 127.0.0.1 (localhost)
port = 3306

Outcome (Based on my testing):


Testing environment:
For avoiding firewall issues on a Windows machine:
1) Turned up a virtual machine on GCP (EC2-Micro)
2) Set up firewall rules to accept TCP connections on port 3306. Also set tags to mysql for accepting connections.
3) SSH into the system and run the following commands to set up the VM for using Docker
    a)  sudo apt install git python3 python3-pip docker.io
    b)  sudo docker pull mysql
    c) vim docker (set up password: admin)    
  
3) Check Docker compose yml for mysql latest instance information

4) Run docker compose in super user: sudo docker-compose up -d 

5) Once the image is up and running, extract the ip address of the instance from dashboard or ssh using ip address command

6) Pass host and port inside Go program interface


The program order is as follows:
1) Using `flag` package, input host and port from user using command line
2) Using DialTimeout Function, try and establish a connection within 2 seconds using `tcp/ip protocol`
3) If the connection is not established defer with the error command.
4) If the connection is established, send default `0x0a` initial handshake packet. If this fails, log error and return
5) If the connection packet is successfully transmitted, read from the connection using buffer slice of size 1024 [Memory Efficiency Tradeoff limit]. 
6) Once we have the data in byte form, we use `byte` package to extract version by using IndexByte function
7) Finally we log all the information about the address, version number and the complete data recieved from handshake 



Time required to build the scanner: 40 mins with Docker instance