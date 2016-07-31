# awslog

This is a command line tool to use cloudwatch logs written in golang.

## NOTE

My reference implementation is [jorgebastida/awslogs](https://github.com/jorgebastida/awslogs).

I recommend to use [jorgebastida/awslogs](https://github.com/jorgebastida/awslogs) if you're searching any Cloudwatch Logs tool. :smile:

# How to use

```
$ awslog -h
```

```
$ awslog groups -h
```

```
$ awslog stream -h
```

```
$ awslog export -h
```

## Supported format in start/end options.

### time format

Use this format such like '2016-07-01 00:00:00 +0900' if you want to give the time.

And use 'now' to give the now.

```
$ awslog export -s '2016-07-01 00:00:00 +0900' -e 'now' sample_group sample_stream
```

### m, minutes

Use 'N m', 'N minutes' if you want to give N minutes ago from now.

```
$ awslog export -s 30m -e 2minutes sample_group sample_stream
```

### h, hours

Use 'N h', 'N hours' if you want to give N hours ago from now.

```
$ awslog export -s 4h  -e 1hours sample_group sample_stream
```

### d, days

Use 'N d', 'N days' if you want to give N days ago from now.

```
$ awslog export -s 20d -e 1days sample_group sample_stream
```

### w  weeks

Use 'N w', 'N weeks' if you want to give N weeks ago from now.

```
$ awslog export -s 3w  -e 1weeks sample_group sample_stream
```

## License

MIT License

## Contributing

1. Fork it ( https://github.com/[my-github-username]/awslog/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
