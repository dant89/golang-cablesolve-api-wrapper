# Golang CableSolve API Wrapper

A Golang CableSolve API wrapper built to enable concurrent requests to CableSolve SOAP APIs.

## About

The project makes use of the [Gin Web Framework](https://github.com/gin-gonic/gin) for simple API routing and requests.

## Golang WSDL SOAP Client generation

To convert SOAP WSDL to a compatible Golang client [WSDL to Go](https://github.com/hooklift/gowsdl) 
has been used.<br>
You can see example generated clients in `/clients/` directory.

To generate the client you install the command line tool, then it is as simple as running 
`gowsdl https://cablesolve.local/cswebapi/ComponentServices.asmx`. <br>
This generates all the commands and mappings required.

## Example
There is an implementation example with concurrency in the `/examples/` directory.
