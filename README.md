# booking-app-go-lang

go mod init <module-name>: it generates module and version of go that is being used
 ## executing application 
 `go run <filename>`
 ## supported datatypes
 - Strings, Integers, Booleans, Maps, Arrays 
 `%v` let's you print the value of given variable and `%T` let you print the type of given datatype eg: `fmt.Printf("datatype is %T and value is %v", myVariable, myVariable)`

*** Explicity declaration ***
``` var myNumber as uint ``` explicity means leaving no doubt
### Array and slices;
```
Syntax: var bookings = [limit]data_type{element1, element2,...elementn}
ex: var bookings = [11]string{"joseph", "janet", "josephine", "joshua"}
```
Slices uses array under the hood but it uses on array with dynamic size, that's differenciate it from array.
## conditions, loop,
 - switch statement
 ``` 
 // syntax
  city := "London"
switch city {
    case "New York":
       // logic for above value
    case "Mexico city" ,"Alberta":
        // logics for "Mexico city" ,"Alberta"
    case "London", "Berlin":
        // logics for "London" ,"Berlin"
    default:
      //  code goes here
}
 ```
 ## Package
 A package is folder that keeps one or more files
  this is how you run code from package `go run file.go. file2.go` or `go run .` dot (.) means execure all go file in current directory 
  `multiple package` export from multiple package you need to
  1. save a file under the folder which has name similar to package name `helper/helper.go`, 
  2. the first line of code should be package package_name you chose `package my_package_name`
  3.  capitalize function name  and  `func MyFunction(){}`
  4. import  by specifying module name  `module_name/packageName` i.e module name is found from `go.mod file` on the first line.
  5. to use the function from other package type `package dot functionName` ex: helper.MyFunction()

 