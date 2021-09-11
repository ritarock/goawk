# goawk
## install
```
$ git clone https://github.com/ritarock/goawk.git
$ cd goawk/
$ go install
```

## Usage
```
NAME:
   goawk - scanning word

USAGE:
   [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --command value, -c value  set command
   --file value, -f value     set file
   --field value, -F value    set file (default: " ")
   --help, -h                 show help (default: false)
```

## Sample
```bash
goawk -c '{print $1}' -f test.txt -F '[ ]'
```
