<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->




<!-- PROJECT LOGO -->
<br />
<div align="center">


  <h3 align="center">A gRPC server and client (server streaming) that will listen and
send requests through the gRPC protocol with  requirements</h3>


</div>




## Built With versions

* go1.18.3
* gRPC
* Protocol Buffer : libprotoc 3.21.1
* grpcurl.exe dev build
* GNU Make 4.3




<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.




1. Clone the repo
```sh
   git clone https://github.com/rafiulhc/grpc-blockchain-endpoints.git
```
2. Install required packages mentioned above




## To get executable binaries

```
  make grpc
  ```

## Run local server

```
   MAC or Linux
  ./bin/grpc/server

  for windows

  ./bin/grpc/server.exe
  ```

## make request from client to get block data

```
  MAC or Linux
  ./bin/grpc/client

  for windows

  ./bin/grpc/client.exe
  ```

## response from server with 5 block data and created a JSON file with Block height and Hash

![grpc](https://user-images.githubusercontent.com/68476971/191175281-2816f882-a3a4-46f7-ab70-0bb55266873f.png)

# From CLI using grpcurl

## To list all services exposed by a server, use the "list" verb.

```
  grpcurl -plaintext localhost:50051 list
  ```

## It will list the gRPC service in the project

![grpcurl](https://user-images.githubusercontent.com/68476971/191187897-e981e48f-fcb2-4f9d-a98a-b9ff5bfcbdb2.png)

## The "list" verb also lets us see all methods in a particular service:

```
  grpcurl -plaintext localhost:50051 list grpc.GetLatestBlockService
  ```

## Describe with

```
  grpcurl -plaintext localhost:50051 describe grpc.GetLatestBlockService.GetLatestBlock
  ```
![GETlATESTbLOCK](https://user-images.githubusercontent.com/68476971/191179130-c90a9943-7ada-4224-96cb-d68ea96e0c96.png)

## Invoking gRPC

Invoking an RPC on a trusted server (e.g. TLS without self-signed key or custom CA) that requires no client certs and supports server reflection is the simplest thing to do with grpcurl. This minimal invocation sends an empty request body:

```
  grpcurl -plaintext localhost:50051 grpc.GetLatestBlockService.GetLatestBlock
  ```
![invoke gRPC](https://user-images.githubusercontent.com/68476971/191177329-a52f60bf-54af-40c4-9e1a-c933b6666eb9.png)

<!--Testing-->
## Testing

```
  go test
```


After running the command in terminal if block from gRPC call and direct API call match, should show the tests results....


![Tests](https://user-images.githubusercontent.com/68476971/191172357-e5e74903-196b-4c7d-9652-fe52ef8c8a92.png)



<!-- LICENSE -->
## License

None



<!-- CONTACT -->
## Contact

Rafiul Hasan - [Linkedin](https://www.linkedin.com/in/hrafiul/)
               [Twitter](https://twitter.com/r_hasan_c)
               - rafiul.hasan.chowdhury@gmail.com

Project Link: [https://github.com/rafiulhc/grpc-blockchain-endpoints](https://github.com/rafiulhc/grpc-blockchain-endpoints)

<p align="right">(<a href="#top">back to top</a>)</p>
