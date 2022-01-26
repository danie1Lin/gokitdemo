
```
$kit new service hello
$kit g s hello --dmw 
$kit g s hello -t grpc
WARN[0000] ===============================================================
WARN[0000] The GRPC implementation is not finished you need to update your
WARN[0000]  service proto buffer and run the compile script.
WARN[0000] ---------------------------------------------------------------
WARN[0000] You also need to implement the Encoders and Decoders!
WARN[0000] ===============================================================

kit g m hi -s hello

// For endpoint
kit g m hi -s hello -e
INFO[0000] Do not forget to append your endpoint middleware to your service middlewares
INFO[0000] Add it to cmd/service/service.go#getEndpointMiddleware()


// For client
kit g c hello

// docker
kit g d
```
