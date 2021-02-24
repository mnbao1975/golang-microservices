# Product Images

## Uploading

Note: need to use `--data-binary` to ensure file is not converted to text
Update through REST

```
curl -vv localhost:9090/1/uploadfile.png -X PUT --data-binary @test.png
```

Upload through multipart form

```
curl --location --request POST 'http://localhost:9090/' \
--form 'id="36"' \
--form 'file=@"/path/to/file"'
```

Get a file from server

```
curl localhost:9090/images/1/hello.png --output hello.png
```

Get compressed file with gzip

```
curl -v --compressed localhost:9090/images/1/hello.png --output hello.png
```
