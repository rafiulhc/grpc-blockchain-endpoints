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




### Built With versions

* go version go1.18.3
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
  ./bin/grpc/server

  for windows user

  ./bin/grpc/server.exe
  ```

## Run below command to make a request from client to get block data
# It will
```
  ./bin/grpc/client

  for windows user

  ./bin/grpc/client.exe
  ```

## It will get response from server with 5 block data and will create a JSON file will Block height and Hash
![grpc](https://user-images.githubusercontent.com/68476971/191175281-2816f882-a3a4-46f7-ab70-0bb55266873f.png)

## Frpm CLI using grpcurl

## To list all services exposed by a server, use the "list" verb.

```
  grpcurl -plaintext localhost:50051 list
  ```
# It will list the gRPC service in the project

![grpc](https://user-images.githubusercontent.com/68476971/191175281-2816f882-a3a4-46f7-ab70-0bb55266873f.png)

## The "list" verb also lets us see all methods in a particular service:

```
  grpcurl -plaintext localhost:50051 list grpc.GetLatestBlockService
  ```
## Describe with --
```
  grpcurl -plaintext localhost:50051 describe grpc.GetLatestBlockService.GetLatestBlock
  ```
![Secret Variables](https://user-images.githubusercontent.com/68476971/169951589-da24b489-0cb6-44f8-a1fb-f9f02afca154.png)

## Invoking gRPC

# Invoking an RPC on a trusted server (e.g. TLS without self-signed key or custom CA) that requires no client certs and supports server reflection is the simplest thing to do with grpcurl. This minimal invocation sends an empty request body:
```
  grpcurl -plaintext localhost:50051 grpc.GetLatestBlockService.GetLatestBlock
  ```
![Secret Variables](https://user-images.githubusercontent.com/68476971/169951589-da24b489-0cb6-44f8-a1fb-f9f02afca154.png)

<!--Testing-->
## Testing

  ```
  go test
  ```


After running the command in terminal if block from gRPC call and direct API call match, should show the tests results....


![Tests](https://user-images.githubusercontent.com/68476971/191172357-e5e74903-196b-4c7d-9652-fe52ef8c8a92.png)






1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request





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
