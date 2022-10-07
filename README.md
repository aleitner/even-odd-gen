# Even Odd Generator

Are you tired of writing if/else for _every_ number to check if those numbers are even or odd? Luckily with this program you can easily generate go code to determine if a number is even or odd.

### Build

```sh
go build main.go
```

### Run

* You can specify `output` which is the location/name of the file being generated.
* You can specify `n` which is the number of even/odd checks to be written.
* You can specify `package` which is the name of the package of the file to be written.
* You can specify `function` which is the name of the even/odd function.


```sh
./main --output="./main2.go" --n=1000 --package="main" --function="main"
```

## Generated code

This section describes how you can build and run the generated code. Amazing!

## Build

```sh
go build main2.go
```

### Run

* You can also specify `n` which is the number to be checked if it is even or odd.

```sh
./main2 --n=69
```