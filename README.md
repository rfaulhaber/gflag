# gflag [![Build Status](https://travis-ci.org/rfaulhaber/gflag.svg?branch=master)](https://travis-ci.org/rfaulhaber/gflag)

Go package that allows for GNU style flags. 

This package largely works the same way as the standard Go Flag package does. The biggest difference is that flags can
be declared with "short" or "long" options.

This means that if you declare a flag as so:
```
enableLogging := gflag.Bool("l", "logging", false, "set to enable logging")
```

you can run it in the command like like:
```
go run program.go -l
```

or:

```
go run program.go --logging
```

and `*enableLogging` will be `true`.

If you leave either the short or long option blank, the only valid flag will be the one you specify.

*NOTE* this project will no longer receive updates and is deprecated in favor of spf13's [pflag](https://github.com/spf13/pflag) library, which is much better than this repo.
