Some notes related to learning of golang.

# gofmt

in [effetive_go](https://golang.org/doc/effective_go.html#formatting)  it says 
`gofmt` use tab for indentation

to use a pre-commit hook before `git commit`, this script is used
```shell
#!/bin/sh
# Copyright 2012 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# git gofmt pre-commit hook
#
# To use, store as .git/hooks/pre-commit inside your repository and make sure
# it has execute permissions.
#
# This script does not handle file names that contain spaces.

gofiles=$(git diff --cached --name-only --diff-filter=ACM | grep '.go$')
[ -z "$gofiles" ] && exit 0

unformatted=$(gofmt -l $gofiles)
[ -z "$unformatted" ] && exit 0

# Some files are not gofmt'd. Print message and fail.

echo >&2 "Go files must be formatted with gofmt. Please run:"
for fn in $unformatted; do
    echo >&2 "  gofmt -w $PWD/$fn"
    done

    exit 1
```

It comes from http://tip.golang.org/misc/git/pre-commit

# Thanks

[Learning Go (original)](https://github.com/miekg/learninggo)

[Learning Go (中文版)](https://github.com/mikespook/Learning-Go-zh-cn)

[Effective Go](https://golang.org/doc/effective_go.html)
