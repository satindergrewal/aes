
Practice tool for encrypting-decrypting plain text data with go language.

Build using following command:

```bash
git clone https://github.com/satindergrewal/aes
git checkout wasm
go mod tidy
GOOS=js GOARCH=wasm go build -o crypto.wasm main.go
```

`wasm_exec.js` file is already included with this code, but if needed find it with the latest updates from local Go install using following instructions:

Find GOROOT:

```
go env GOROOT
```

Use that path to find wasm_exec.js file, example:

```
find /opt/homebrew/Cellar/go/1.24.0/libexec -name "wasm_exec.js"
```

When found copy that to replace the existing file in the code.

To run it locally can use either Python to serve the page on localhost address:

```
python3 -m http.server 8080
```


For testing use this data:


To encrypt:

```bash
key 1 = hello
key 2 = 
Text = world

# OUTPUT:
c2932347953ad4a4-25f496d260de9c150fc9e4c6-20bc1f8439796cc914eb783b9996a8d9c32d45e2df
```

To decrypt:

```bash
key 1 = hello
key 2 =
Text = c2932347953ad4a4-25f496d260de9c150fc9e4c6-20bc1f8439796cc914eb783b9996a8d9c32d45e2df

# OUTPUT:
world
```
