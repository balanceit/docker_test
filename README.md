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

* Setup the image to forward port 8000
`VBoxManage controlvm "docker-vm" natpf1 "tcp-port8000,tcp,,8000,,8000";`

* run the default make task
`make`
 * this will build an executable `web` from the `src\web` package
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
  `vpcId=`aws ec2 create-vpc --cidr-block 10.0.0.0/28 --query 'Vpc.VpcId' --output text``
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
`10.0.2.0/23`
```
CIDR Range	10.0.2.0/23
Netmask	255.255.254.0
Wildcard Bits	0.0.1.255
First IP	10.0.2.0
Last IP	10.0.3.255
Total Host	512
```

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
```
securityGroupId=`aws ec2 create-security-group --group-name test-seav-security-group --description "Security group for testing" --vpc-id $vpcId  --query 'GroupId' --output text`
```
* add the routes to the group, we open 22 for ssh and 8000 for the app
```
aws ec2 authorize-security-group-ingress --group-id $securityGroupId --protocol tcp --port 22 --cidr 0.0.0.0/0
aws ec2 authorize-security-group-ingress --group-id $securityGroupId --protocol tcp --port 2376 --cidr 0.0.0.0/0
aws ec2 authorize-security-group-ingress --group-id $securityGroupId --protocol tcp --port 8000 --cidr 0.0.0.0/0
```

* create an instance for holding our docker containers
```
docker-machine create --driver amazonec2 --amazonec2-region ap-southeast-2 --amazonec2-vpc-id $vpcId --amazonec2-subnet-id $subnetId --amazonec2-security-group $securityGroupName test-seav-container-host

docker-machine env test-seav-container-host

eval $(docker-machine env test-seav-container-host)

docker-machine ls

docker-machine ip test-seav-container-host

docker build -t testing_image .

docker run -d -p 8000:8080 --log-driver json-file --name testing_container testing_image

docker-machine stop container-host

docker-machine rm container-host


```
