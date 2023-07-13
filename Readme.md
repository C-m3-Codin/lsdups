# Duplicate File Finder CLI Tool

A command-line tool for finding and listing duplicate files within a directory.

## Installation

### Prerequisites

- Go 

### Steps

1. Clone the repository:

   ```shell
   git clone https://github.com/C-m3-Codin/lsdups.git
   ```

2. Change into the project directory:

   ```shell
   cd lsdups
   ```

3. Build the project:

   ```shell
   go build
   ```

4. Run the tool:

   ```shell
   ./clean
   ```

## Usage

The tool scans a directory and lists duplicate files within it.

```shell
./clean lsdups
```

<!-- Replace `[directory]` with the path to the directory you want to scan. If no directory is provided, the current directory will be scanned by default. -->

## Example

```shell
# List duplicate files in the current directory
./clean lsdups

# # List duplicate files in a specific directory
# ./clean /path/to/directory
```

## Output

The tool will provide a list of duplicate files found within the specified directory, along with their respective file hashes and the count of duplicates.

For example:

```shell
Directory is /path/to/directory
file hash is abcdef1234567890: {Path: [file1.txt file2.txt], Count: 2}
file hash is 0987654321fedcba: {Path: [file3.txt], Count: 1}
```

## Improvements

1. Implement optional recursive search to include subdirectories in the scan.
2. Optimize the code for better performance by introducing multiple workers for parallel file processing. [done in threadIt branch]


## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request. Make sure to follow the project's code of conduct.

## License

This project is licensed under the [MIT License](LICENSE).

