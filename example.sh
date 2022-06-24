export k1="main key one"
export k2="key 2"
export c=$(zbarimg -q --raw k.png)
./aes -key1 "$k1" -key2 "$k2" -c "$c"
