# Product Images

## Uploading

Note: need to use `--data-binary` to ensure file is not converted to text

```
curl -vv localhost:9090/1/uploadfile.png -X PUT --data-binary @test.png
```

```
curl --location --request POST 'http://localhost:9090/' \
--form 'id="36"' \
--form 'file=@"/path/to/file"'
```
