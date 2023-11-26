# MY-LS-1

This project utilizes the Go programming language to implement the ls command with various options. It supports the following flags:

* `-a`: Displays all files, including hidden files, which are typically prefixed with a dot (`.`).

* `-i`: Displays the index number of each file or directory.

* `-l`: Displays detailed information about each file or directory, including permissions, ownership, size, and modification time.

* `-R`: Recursively lists the contents of directories, including nested directories.

* `-r`: reverse order while sorting.

* `-t`: Sorts the listing by modification time, with the most recently modified files appearing first.


## Usage

```bash
go run main.go [options] [path]
```

where `options` is a combination of the supported flags, and `path` is the directory or file to list. For example, to list the contents of the current directory in long format, including hidden files, sorted by modification time, you would use the following command:

bash
go run main.go -l -a -t


## Implementation

The project utilizes the `os` package to access the filesystem and retrieve file information. It parses the provided arguments and applies the specified options to the listing process. For instance, when the `-l` flag is present, it extracts additional details about each file or directory and displays them in a formatted manner. Similarly, when the `-R` flag is used, it recursively traverses the directory structure and lists the contents of each nested directory.

## Testing

The project includes unit tests to ensure the correctness of its functionality. These tests cover various scenarios, including listing directories, hidden files, and sorting by modification time.

## Conclusion

This project demonstrates the practical application of Golang in implementing common command-line tools. It effectively utilizes standard library packages to access the filesystem and manipulate file information. The inclusion of unit tests ensures the reliability and accuracy of the implemented code.