# prepwd
[![Build Status](https://travis-ci.org/mtib/prepwd.svg?branch=master)](https://travis-ci.org/mtib/prepwd)

This tool written in Go will clone all owned (public) repos and gists of the user
specified, it will also download the starred repos. It will structure the folders
like this:

- user
    - repos
        - ...
    - gists
        - ...
    - stars
        - ...

It will use an escaped version of the gist description as the directory name, because
the random letter string wouldn't make a good descriptor.

By default it will only clone the repositories to a depth of 10 commits.
There is currently no parameter to change this from the command line,
but you could change it in the gclone/pull.go file on your own.

## usage
```
prepwd [https|ssh] <user>
prepwd mtib
prepwd ssh mtib
```

# api links
[<user> Repos](https://api.github.com/users/<user>/repos)
```
https://api.github.com/users/<user>/repos
```

[<user> Gists](https://api.github.com/users/<user>/gists)
```
https://api.github.com/users/<user>/gists
```
