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

## Next steps

Three things I would encourage you to look into next :gopher-dance:
- Write a simple test. The inbuilt low friction test library is one of Go big selling points, so check it out
- You ask yourself in a comment in your code "is it safe to ignore this error" and that's a great question to ask. I would always lean towards the answer being "No", especially for newer gophers
- Your slice based kvstore is cool, but you could also explore a Go primitive called a map, which might be an interesting read