# keil-cli

A command-line to merge boards data from different files into a single file and build metadata.

## Usage
```
NAME:
   keil merge - Merges boards metadata from json files into one file

USAGE:
   keil merge [command [command options]]

OPTIONS:
   --dir value, -d value  Single directory with all files to be merged
   --out value, -o value  Output file name (default: "out.json")
   --enableIndentation    Enables Indentation in output (default: false)
   --help, -h             show help    
```

## Installation

* Run the command `make dist` and use the file that suits your system.
* Rename the file to keil
* Make the file executable `chmod +x keil`
* Optional - Move the executable under a directory in your $PATH (e.g. /usr/local/go/bin)
* Run keil help

## Missing items

* Consider reading files from a provided array of files
* Add releases support
* Use buffered reading/writing into files, instead of the simple os.WriteFile/os.ReadFile functions. For better memory usage and for easier testing.