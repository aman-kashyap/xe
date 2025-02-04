# NexaDash -go

## Project Description
This project contains nexadash backend APIs in golang

## Technology used
* golang version 1.8
* Postgresql 9.3

## Installation

### first we should download go 1.8 and create a environmental variable for $GOPATH, which points to our workspace directory.
	you shold check the detail [here](https://golang.org/doc/install#install) before using below steps.

### Steps to use

1. $ cd $GOPATH/src/nexadash-backend-go
2. $ go build
3. $ ./nexadash-backend-go

### you can also use Glide for package management for this project. 
### you can install latest glide release on Mac and Linux using script:

```
curl https://glide.sh/get | sh
```
### In case of Ubuntu Precise (12.04), Trusty (14.04), Wily (15.10) or Xenial (16.04) yo can also use PPA:

```
sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
sudo apt-get install glide
```

#### After installing glide use following steps to manage packages

* glide init
   - it will create a .yaml file mentioning packages and sub-packages with their versions
* glide install 
   - it will install the specified versions and lock it which can be found in glide.lock file created
 

 




