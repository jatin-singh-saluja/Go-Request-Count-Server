/*

 Project: Request Count Server

 Project Description:

 The goal of this project is to deliver simple "RequestCount" server, written in Go. 
 This server:

 a) counts the number of requests handled by the server instance.
 b) counts the number of requests handled by the entire cluster.
 c) serves these numbers over HTTP.
 d) dockerizes the "RequestCount" server.

*/

1) File goserver.go contains the main function,
   and the code for starting up the server and registering
   request handlers.

2) File counter.go contains the source code for requestCounterHandler, 
   which is used to keep track of total number of requests to the server, and the cluster.

3) File redisInstance.go contains the source code for establishing
   connection to the redis server.

4) File nginx.conf contains the configuration for nginx being used as load-balancer.

5) In order to test the functionality of this application, please do the following:
   
   a) Go to the root directory of your project.

   b) Run the following command:
      docker-compose up --build --scale go-server=3

   c) Go to http://localhost:4000/

   d) After you are done with the application, run the following command:
      
      docker-compose down


