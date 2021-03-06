[![Build Status](https://travis-ci.org/balanceit/docker_test.svg?branch=master)](https://travis-ci.org/balanceit/docker_test)

# docker_test

## Requirements

* ~boot2docker installed~
* docker machine is installed
* go build is installed

## Set up
* `boot2docker down` -- shutdown if running
* `brew cask install dockertoolbox` -- installs docker machine and the lastest version of virtual box
* `docker-machine create --driver virtualbox docker-vm` -- create a new VM and machine named *docker-vm*
* `eval $(docker-machine env docker-vm)` -- add needed env variables
* `docker ps` -- just a quick test to ensure it is working

* Setup the VM to forward port 8000 for the test web server
`VBoxManage controlvm "docker-vm" natpf1 "tcp-port8000,tcp,,8000,,8000";`

* Setup the VM to forward port 5432 for postgres
`VBoxManage controlvm "docker-vm" natpf1 "tcp-port5432,tcp,,5432,,5432";`

### to start and stop
* `docker-machine start docker-vm`
* `docker-machine stop docker-vm`

#### The "Something went wrong running an SSH command" error

Too many retries waiting for SSH to be available.  Last error: Maximum number of retries (60) exceeded
Error checking TLS connection: Something went wrong running an SSH command!

see: https://github.com/docker/toolbox/issues/317

* `docker-machine rm docker-vm`
* `docker-machine create --driver virtualbox docker-vm`
* `eval $(docker-machine env docker-vm)`

### building the application
* run the default make task
`make`
 * this will build an executable `main` from the `src\main` package
 * build a docker image defined by `IMAGE_NAME` (defaults to `docker_test`)
 * run a docker container publishing the exposed port 8080 -> 8000 (see `./Dockerfile`) and a container name defined by `CONTAINER_NAME` (this defaults to `go_web_server`)

## building a small container
* see the `.travis.yml`
* from https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

### building static executable
* `GOPATH=/path/to/gb/project/vendor:/path/to/gb/project CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extld ld -extldflags -static' -a -x .``


## aws cli and docker machine
pip install --upgrade pip
pip install awscli

### creating a container host in ec2
docker-machine create --driver amazonec2 --amazonec2-region ap-southeast-2 container-host --amazonec2-vpc-id
docker-machine ls

--- docker-machine env container-host && eval $(docker-machine env container-host) -- if there are any problems connecting to the ec2 instance


docker-machine stop container-host
docker-machine rm container-host
from: https://docs.docker.com/machine/examples/aws/


## creating vpc and network

at the end we will have

* a vpc
`vpcId=vpc-b8d389dd`
* an internet gateway
`internetGatewayId=igw-48a83f2d`
* a subnet (or more for different aws regions)
`subnetId=subnet-4b09373c`
* a routing table
`routeTableId=rtb-83818de6`
* an association
`AssociationId=rtbassoc-8805d3ec` -- not sure if this would have to be deleted
* a security group
`securityGroupId=sg-fcfe4698`
`securityGroupName=test-seav-security-group`

### the shell commands

* create the VPC
`aws ec2 create-vpc --cidr-block 10.0.0.0/22`
  ```
CIDR Range	10.0.0.0/22
Netmask	255.255.252.0
Wildcard Bits	0.0.3.255
First IP	10.0.0.0
Last IP	10.0.3.255
Total Host	1024
  ```
  * there is the option to run something like the below for better scripting
  `vpcId=`aws ec2 create-vpc --cidr-block 10.0.0.0/22 --query 'Vpc.VpcId' --output text``
  * note the use of `--query` and `--output`

* to enable dns (did not do this)
 * `aws ec2 modify-vpc-attribute --vpc-id $vpcId --enable-dns-support "{\"Value\":true}"`
 * `aws ec2 modify-vpc-attribute --vpc-id $vpcId --enable-dns-hostnames "{\"Value\":true}"`


* next to create the way out, an internet gateway
 ```
 internetGatewayId=`aws ec2 create-internet-gateway --query 'InternetGateway.InternetGatewayId' --output text`
 ```
* then we will attach it to the vpc
 ```
 aws ec2 attach-internet-gateway --internet-gateway-id $internetGatewayId --vpc-id $vpcId
 ```
* creating one subnet
`aws ec2 create-subnet --vpc-id $vpcId --cidr-block 10.0.0.0/23`
`aws ec2 create-subnet --vpc-id $vpcId --cidr-block 10.0.0.0/23 --query 'Subnet.SubnetId' --output text`
```
CIDR Range	10.0.0.0/23
Netmask	255.255.254.0
Wildcard Bits	0.0.1.255
First IP	10.0.0.0
Last IP	10.0.1.255
Total Host	512
```
  * if another needs to be created then it should be as such:

`aws ec2 create-subnet --vpc-id $vpcId --cidr-block 10.0.2.0/23 --availability-zone ap-southeast-2b`
`10.0.2.0/23`
```
CIDR Range	10.0.2.0/23
Netmask	255.255.254.0
Wildcard Bits	0.0.1.255
First IP	10.0.2.0
Last IP	10.0.3.255
Total Host	512
```
resulted in `subnet-ee0df38a`

* creating a routing table, this is needed to route traffic
```
aws ec2 create-route-table --vpc-id $vpcId
```
```
routeTableId=`aws ec2 create-route-table --vpc-id $vpcId --query 'RouteTable.RouteTableId' --output text`
```
* then associated it with the subnet
```
aws ec2 associate-route-table --route-table-id $routeTableId --subnet-id $subnetId
```
* and create a rule which will allow traffic from the vpc to the outside world
```
aws ec2 create-route --route-table-id $routeTableId --destination-cidr-block 0.0.0.0/0 --gateway-id $internetGatewayId
```

* lastly a security group is needed to route traffic to any created instances
```
aws ec2 create-security-group --group-name test-seav-security-group --description "Security group for testing" --vpc-id $vpcId
```
securityGroupId=sg-fcfe4698
```
securityGroupId=`aws ec2 create-security-group --group-name test-seav-security-group --description "Security group for testing" --vpc-id $vpcId  --query 'GroupId' --output text`
```
* add the routes to the group, we open 22 for ssh and 8000 for the app
```
aws ec2 authorize-security-group-ingress --group-id $securityGroupId --protocol tcp --port 22 --cidr 0.0.0.0/0
aws ec2 authorize-security-group-ingress --group-id $securityGroupId --protocol tcp --port 2376 --cidr 0.0.0.0/0
aws ec2 authorize-security-group-ingress --group-id $securityGroupId --protocol tcp --port 8000 --cidr 0.0.0.0/0
```

### ec2 docker host instance
* create an instance for holding our docker containers
```
docker-machine create --driver amazonec2 --amazonec2-region ap-southeast-2 --amazonec2-vpc-id $vpcId --amazonec2-subnet-id $subnetId --amazonec2-security-group $securityGroupName test-seav-container-host
```
### hook up to that instance
* set the needed env vars
```
docker-machine env test-seav-container-host
eval $(docker-machine env test-seav-container-host)
```
if there are any connection issues
`docker-machine regenerate-certs test-seav-container-host`

* check it is working
```
docker-machine ls
```

* get the ec2 instances ip address
```
docker-machine ip test-seav-container-host
```

### build and run the docker image/container
* build the image from local `DockerFile`
```
docker build -t testing_image .
```
* run the container
```
docker run -d -p 8000:8080 -e DB_CONNECTION_STRING="dbname=docker_test_developement user=root password=xxx" -e PGHOST="10.0.1.35" --log-driver json-file --name testing_container faroeseav/docker_test
```

### shutdown and clean up
```
docker-machine stop container-host
docker-machine rm container-host
... plus delete the aws stuff ....
```

## RDS creation

### subnet group
```
aws rds create-db-subnet-group \
--db-subnet-group-name test-seav-db-subnet-group \
--db-subnet-group-description "Subnet group for testing database" \
--subnet-ids subnet-4b09373c subnet-ee0df38a
```
### db instance
* create a database instance
```
aws rds create-db-instance \
  --db-name docker_test_developement \
  --db-instance-identifier seav-test-db-instance \
  --allocated-storage 5 \
  --vpc-security-group-ids sg-fcfe4698 \
  --db-instance-class db.m1.small \
  --db-subnet-group-name test-seav-db-subnet-group \
  --engine postgres \
  --engine-version 9.4 \
  --master-username root \
  --master-user-password xxx
```
* have a look
`aws rds describe-db-instances`

### test it
* ssh via docker-machine
`docker-machine ssh test-seav-container-host`
* install psql client
`sudo apt-get update && sudo apt-get upgrade`
`sudo apt-get install postgresql-client`
* attempt to connect
`psql -h test-seav-db.cmtsrxlj8h6c.ap-southeast-2.rds.amazonaws.com -U root -W -p 5432 -d docker_test_developement`

## Database migrations

### Preconditions

### thoughts
http://aranair.github.io/posts/2016/04/27/golang-docker-postgres-digital-ocean/


### Preconditions
* have go-bindata installed
* `gb vendor fetch github.com/rubenv/sql-migrate`
* `create database docker_test;`
* `create database docker_test_test;`
* `create database docker_test_developement;`
* `go-bindata -pkg myapp -o bindata.go db/migrations/`

We are using [goose](https://github.com/ox/goose) to manage db migrations (note, this is not to be confused with https://bitbucket.org/liamstask/goose).

To install:
```
go get github.com/ox/goose/cmd/goose
```
This will add a `goose` executable to your `$GOPATH/bin` directory.

To create the database as defined in the `db/dbconf.yml` file (the `development` environment is the default, to change use the `-env="environment"` switch):
```
 goose create-db
```

To then run outstanding migrations:
```
 goose up
```

# Docker Repositories
docker run -d   --name watchtower   -v /var/run/docker.sock:/var/run/docker.sock   centurylink/watchtower


# Docker tutum
This will allow us to connect docker hub repo with a deployment location
https://docs.docker.com/docker-cloud/infrastructure/link-aws/

## AWS Setup
* link aws account with docker cloud
* create a docker cloud user in IAM
* `dockercloud-user` (save the access key id and secret)
* Create a new policy for this docker cloud user dockercloud-user
* Create your own policy: `dockercloud-policy`
```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:*",
        "iam:ListInstanceProfiles"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
```
* now attach the policy to the user (my goodness aws is tedious)

## Docker cloud setup
* in Cloud settings add the creds from the above user to the "Amazon Web Services" under "Cloud providers"

### Create nodes and node clusters
https://docs.docker.com/docker-cloud/getting-started/your_first_node/
* Create a node in ap-southeast-2
* use `vpc-b8d389dd` created above
* and subnet `subnet-4b09373c` from above
* choose the nano size
* 10 GB (is the smallest)
* 1 node
* Launch the node cluster

### Services
https://docs.docker.com/docker-cloud/getting-started/your_first_service/
 in docker cloud talk a *Service* is a group of containers from the same image
