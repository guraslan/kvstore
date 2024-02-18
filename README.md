# KV Store
- kvstore foo bar (set key foo to bar)
- kvstore foo (get value of key foo)
- kvstore (dump all keys stored)

## How to store
- Two simple approaches to store kv pairs in a file

  - Save key-value pairs in a file in the format key=value per line
    Use the first = as the token to parse

  - How can we do it without parsing?

  key1<br>
  value1<br>
  key2<br>
  value2<br>
  ...

- Add resource version i.e an integer counter that represents the current version of the key

```
$ make clean
$ make
```