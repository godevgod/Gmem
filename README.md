# Gmem
Gmem
![Screenshot2024-02-05143040.png](https://github.com/godevgod/Gmem/blob/e80305bcd97f4ba6f1322c3b6634c6e80483ce2a/v2/Screenshot2024-02-05143040.png)



Shared Memory Communication in Go
This repository contains two Go programs demonstrating inter-process communication through shared memory. One program (writer.go) writes data to a shared memory segment, and the other (reader.go) reads data from the same segment.

Writer (writer.go)
The writer creates and writes to a shared memory segment. It repeatedly updates the memory with random bytes and strings, simulating a data production process.

Usage
bash
Copy code
go run writer.go
Details
File Mapping: Creates a named shared memory segment.
Data Structure: Uses a struct Test with byte arrays for storing random bytes and strings.
Random Data Generation: Generates random bytes and strings and writes them to the shared memory.
Logging: Outputs the written data to the console.
Reader (reader.go)
The reader opens the shared memory segment created by the writer and reads data from it, displaying the contents.

Usage
bash
Copy code
go run reader.go
Details
File Mapping: Opens the named shared memory segment created by the writer.
Data Structure: Maps the Test struct to the shared memory segment for reading.
Data Reading: Reads and displays the strings from the shared memory, demonstrating that it accesses the same memory segment as the writer.
Shared Structure (Test)
Both the writer and reader use the Test struct to structure the shared memory. It contains three byte arrays:

Context: A byte array for random context data.
Str1: A byte array for the first string.
Str2: A byte array for the second string.
Requirements
Go installed on your machine.
Sufficient permissions to create and access memory-mapped files.
Notes
Both programs need to be run with appropriate privileges to create and access the shared memory segment.
The shared memory name (ShareMemory) is crucial for the inter-process communication and must be the same in both programs.
Proper error handling is implemented to avoid issues during file mapping and memory mapping operations.
This example provides a basic demonstration of using shared memory for inter-process communication in Go. It can be extended or modified for more complex data structures and use cases.

![Bitcoin QR Code](https://github.com/godevgod/foreverrun/blob/main/1CbE3SsUcvJWZ2YNaDwUj9AQtT8k4AGmLe.png?raw=true)
### Bitcoin:1CbE3SsUcvJWZ2YNaDwUj9AQtT8k4AGmLe
```bash
   1CbE3SsUcvJWZ2YNaDwUj9AQtT8k4AGmLe
