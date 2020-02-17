# inc

inc is a simple program to test a few different software development practices

- client/server programs
- multiple architectures
- GitHub Actions
- GoReleaser
- Advanced CLI client

## Make signing key:
```
openssl ecparam -genkey -name secp160k1 -noout -out release.pem
openssl ec  -in release.pem -pubout -out release.pub
```

