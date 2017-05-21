
## Installing GO 

Download from 

`https://golang.org/dl/`


`sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`

Create Paths for your go projects in your home directory 

`mkdir goprojects`



Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your 

Add these two lines 

`export PATH=$PATH:/usr/local/go/bin`

` export GOPATH=$HOME/goprojects`


At the end of this file 

` sudo vim /etc/profile `

The 

`source /etc/profile`



## Cassandra Go Driver 

go get github.com/gocql/gocql

## ArangoDb Driver

go get github.com/arangodb/go-driver
 
 
## Redis

go get -u github.com/go-redis/redis

## MongoDB

go get gopkg.in/mgo.v2

## Install Web API

go get gopkg.in/kataras/iris.v6

## RUNNING THE TEST 

`http://localhost:8080/arango/write/1000`

for creating 1000 object 

`http://localhost:8080/arango/write/1`

For creating 1 object 

