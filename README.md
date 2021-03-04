<p align="center">
    <img alt="pirate" src="docs/media/gopher_pirate.png"> 
</p>
<p align="center">

captnHook allows users to place trades with Tradingview's webhook alerts. The server is configurable to the most popular broker APIs (Alpaca, Binance, Coinbase, many more)

</p>
<br>
<p align="center"><a href="#">Website Coming Soon</a></p>
<br>
<p align="center">
   <a href="https://github.com/imthaghost/goclone/blob/master/LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg"></a>
   <a href="https://goreportcard.com/report/github.com/bareish/captnHook"><img src="https://goreportcard.com/badge/github.com/bareish/captnHook"></a>
</p>
<br>

<p align="center">
<h3>Directories</h3>
<p>
   • <a href="#">cert</a> A folder for housing a self-signed X.509 TLS certificate <br>
   • <a href="#">cmd</a> Main application for this project <br>
   • <a href="#">docs</a> Design and user documents <br>
   • <a href="#">pkg</a> Library code used by the main command <br>
   </p>
   
<a name="directories"></a>


<a name="installation"></a>
### Installation
Clone the project to your machine

```bash
$ https://github.com/bareish/captnhook.git
```
<a name="envioirnment variables"></a>
### Environment Variables

Copy the contents of .env.example and create a .env file from the contents - 
Then, fill in credentials that are required.
```bash
$ touch .env
```
```bash
$ cp .env.example .env
```
> Directory tree should look as such:
```textmate
  ├── .env
  ├── .env.example
  ├── Dockerfile
  ├── README.md
  ├── cmd
  ├── pkg
  ├── cert
  │   ├── cert.pem
  │   └── key.pem
```
<a name="dot env example"></a>
### Dotenv Example
```dotenv
# general configuration
MODE = dev  # can be dev/prod
BASE_URL=https://0.0.0.0
PORT=8000

# Path to self-signed X.509 TLS certificate
CERT_PATH=./cert/cert.pem
KEY_PATH=./cert/key.pem


# Alpaca Credentials
ALPACA_CLIENT_ID=nicetrygithubscraper
ALPACA_CLIENT_SECRET=nicetrygithubscraper
ALPACA_ACCOUNT_TYPE=Paper

#Binance Credentials
BINANCE_CLIENT_ID=nictrygithubscraper
```

<a name="running locally"></a>
### Running locally
As simple as one command :)
```shell script
$ docker-compose up
```
 Output should look as such:

```go
Successfully tagged captnhook_api:latest
Recreating captnhook_api ... done
Attaching to captnhook_api
api_1  |
api_1  |    ____    __
api_1  |   / __/___/ / ___
api_1  |  / _// __/ _ \/ _ \
api_1  | /___/\__/_//_/\___/ v4.2.0
api_1  | High performance, minimalist Go web framework
api_1  | https://echo.labstack.com
api_1  | ______________________________________________
api_1  |                                     
api_1  | ⇨ http server started on [::]:8000

```




<a name="license"></a>
## 📝 License

By contributing, you agree that your contributions will be licensed under its MIT License.

In short, when you submit code changes, your submissions are understood to be under the same [MIT License](http://choosealicense.com/licenses/mit/) that covers the project. Feel free to contact the maintainers if that's a concern.

<a name="contributors"></a>
