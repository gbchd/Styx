# Styx 

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/guillaumebchd/styx.svg)](https://github.com/guillaumebchd/styx)
[![GoReportCard example](https://goreportcard.com/badge/github.com/guillaumebchd/styx?update)](https://goreportcard.com/report/github.com/guillaumebchd/styx)

Styx is a HTTP reverse proxy and load balancer with anti-ddos functionality. It was built as the session project of the course IFT729 - Real-time system design at the Sherbrooke University.

## Overview

Imagine that you own a virtual private server (VPS) and you want to host your website and your api. You could host the web application on the port 80 and the api on another port, but this architecture will gradually become harder overtime. Let's say you want to host another website, the port 80 being already taken, you can't use it again so you'll need to use another port and that become annoying. 

That's when a reverse proxy like Styx comes in handy. 

You simply configure your reverse proxy on the port 80 and it will redirect the user to the correct service using the url used to access your vps.


## Features

- Performant Reverse Proxy.
- Basic mapping of HTTP entrypoint to HTTP destinations.
- Basic Load Balancing with round robin algorithm.
- Anti-DDOS functionnality.
- Simple configuration file.

## How it works

- To get started, simply clone the project:
```
git clone https://github.com/GuillaumeBchd/Styx.git
```
- Then build the executable:
```
cd <path>/styx
go build ./cmd/styx
```
- Edit the configuration file to your liking and then launch the executable 😄:
```
./styx
```


## Contributors

- [Adrien](https://github.com/AdrienVerdier) 😉
- [Alexandre](https://github.com/TurpinA) 🦍
- [Guillaume](https://github.com/GuillaumeBchd) 😎


## License

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](../master/LICENSE)
