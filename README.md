# go-mysql-revel-rest
Example of REST API to create,retrieve,update and delete items with Go,Mysql and Revel framework

## Running

To use this sample you must have Go installed as described [here](https://golang.org/doc/install), and follow the [official Go code organization](https://golang.org/doc/code.html);
MySQL installed and running, if not please check out [MySQL download page](https://dev.mysql.com/downloads/installer/) and follow [these instructions](http://dev.mysql.com/doc/refman/5.7/en/installing.html).

To check your installations, run the following command in the command line:
```
$ go version
go version go1.7.1 linux/amd64 # sample output
$ mysql --version
mysql  Ver 14.14 Distrib 5.5.52, for debian-linux-gnu (x86_64) using readline 6. # sample output
```
# Install dependencies

This project uses Revel framework. You can install it by:

<code>go get github.com/revel/revel</code>

<code>go get github.com/revel/cmd/revel</code>

Set up MySQL database, use -u -p flags to provide username and password:
```
$  mysql < database_backup.sql
```
# Run

Once Revel framework is installed, you can run the server by:

<code>revel run go-revel-rest</code>

Note that the project must be located under <code>$GOPATH/src/go-revel-rest</code>

# Routes

The API routes are defined in <code>conf/routes</code> file:

<code>GET /fruits</code> Retrieve all the fruits

<code>GET /fruits/{id}</code> Retrieve a fruit with id {id}

<code>POST /fruits/</code> Save a new fruit. <code>name</code> and <code>value</code> must be provided as form params.

<code>PUT /fruits/</code> update a fruit with id {id}. <code>name</code> and <code>value</code> optional as form params.

<code>DELETE /fruits/{id}</code> delete a fruit with id {id}


## Technologies
Language - [Go](https://golang.org/)<br />
Web framework - [Revel](https://revel.github.io)<br />
ORM - [Gorp](https://github.com/go-gorp/gorp)<br />
Database - [MySQL](https://www.mysql.com/)<br />