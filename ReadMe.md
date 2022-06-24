
Practice tool for encrypting-decrypting plain text data with go language.

Build using following command:

```bash
git clone https://github.com/satindergrewal/aes
go mod tidy
go build
```

Command help:

```bash
./aes -h
Usage of ./aes:
  -c string
    	ciphertext
  -key1 string
    	passphrase key part 1
  -key2 string
    	passphrase key part 2
```

Extra tools which can be used:

- qrencode
- zbar-tools

```bash
brew install qrencode zbar
```

Encoding QR-Code to an image:

```bash
qrencode -s 6 -l H -o "q.png" "c2932347953ad4a4-25f496d260de9c150fc9e4c6-20bc1f8439796cc914eb783b9996a8d9c32d45e2df"
```

Decoding QR-Code to text:

```bash
$ zbarimg -q --raw q.png
c2932347953ad4a4-25f496d260de9c150fc9e4c6-20bc1f8439796cc914eb783b9996a8d9c32d45e2df
```

Can use a bash shell script to pass arguments to command line by reading input from QR code image with following example bash code:

```bash
# test.sh
export k1="hello"
export k2=""
export c=$(zbarimg -q --raw q.png)
./aes -key1 "$k1" -key2 "$k2" -c "$c"
```

Executing it with bash will output decrypted text:

```bash
$ bash test.sh
decrypted: world
```

Building for cross platform:

```bash
mkdir -p build/{linux,mac}/{amd64,arm64}
GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/aes
GOOS=linux GOARCH=arm64 go build -o build/linux/arm64/aes
GOOS=darwin GOARCH=amd64 go build -o build/mac/amd64/aes
GOOS=darwin GOARCH=arm64 go build -o build/mac/arm64/aes
tar cvf aes.tar.gz build
```
