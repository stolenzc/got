# Got

## What's Got

A Version Control System (VCS) tool for the command line. It is very similar to git, but it is not git. It is written in Go. Most of the commands have the same name as their git equivalent, but not all of them do. It's a good project to learn Go, and it's also a good project to learn about VCS.

## Installation

### From source

You need to have Go installed. Then, run the following commands:

```bash
git clone https://github.com/stolenzc/got.git
cd got
go mod download
go build -o got
```

Then, you can run the `got` binary.

### From binary

You can download the binary from the [releases page](https://github.com/stolenzc/got/releases).

## TODO

- [ ] `got init`
- [ ] `got add`
- [ ] `got commit`
- [ ] `got status`
- [ ] `got log`
- [ ] `got branch`
- [ ] `got checkout`
- [ ] `got merge`
- [ ] `got rebase`
- [ ] `got reset`
- [ ] `got stash`
- [ ] `got tag`
- [ ] `got remote`
- [ ] `got fetch`
- [ ] `got pull`
- [ ] `got push`
- [ ] `got clone`
- [ ] `got gc`
- [ ] `got config`
