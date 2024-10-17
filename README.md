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