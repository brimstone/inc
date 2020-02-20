# inc

inc is a simple program to test a few different software development practices

- client/server programs
- multiple architectures
- GitHub Actions
- GoReleaser
- Advanced CLI client

## Make signing key:
```
openssl ecparam -genkey -name secp521r1 -noout -out release.pem
openssl ec  -in release.pem -pubout -out release.pub
```

## TODO
- steal arg parsing from brimstone/jwt
- test selfupdate function
- Add updating to a specific version

## Things I'm not happy with
- The specific cmds (`cmd/inc/cmd`) and the common cmds (`pkg/cmd`).
  1) `cmd` is how I want it, or `cobra` is what it is
