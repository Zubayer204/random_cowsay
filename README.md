# Random Cowsay

Random Cowsay is a fun command-line tool written in Go that displays random cowsay messages using the output of fortune. It randomly selects one of the available cow files provided by cowsay and pipes a fortune quote through it to generate a delightful message.

## Features

- Generates a randomized cowsay message using a random cowfile.
- Integrates with the classic UNIX programs [fortune](https://en.wikipedia.org/wiki/Fortune_(Unix)) and [cowsay](https://en.wikipedia.org/wiki/Cowsay).
- Easy-to-use command-line tool written in Go.

## Prerequisites

Before using Random Cowsay, ensure you have the following installed on your system:

- [Go](https://golang.org/dl/) (version 1.24.0 or later)
- [cowsay](https://en.wikipedia.org/wiki/Cowsay)  
  Installation examples:
  - Debian/Ubuntu:  
    `sudo apt-get install cowsay`
  - macOS with Homebrew:  
    `brew install cowsay`
- [fortune](https://en.wikipedia.org/wiki/Fortune_(Unix))  
  Installation examples:
  - Debian/Ubuntu:  
    `sudo apt-get install fortune`
  - macOS with Homebrew:  
    `brew install fortune`

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/zubayer204/random_cowsay.git
   cd random_cowsay
   ```

2. Build the project using Go:

   ```
   go build -o random_cowsay ./random_cowsay/main.go
   ```

   This will generate an executable named `random_cowsay` in the current directory.

## Usage

Ensure that both `fortune` and `cowsay` are installed before running the application.

Run the tool using:

```
./random_cowsay
```

On execution, the program will:

- Verify that both `cowsay` and `fortune` are available.
- List available cow files and randomly select one.
- Retrieve a fortune message and display it using cowsay with the random cow.

If any of the required commands are not installed, the program will print an error message guiding you on how to install them.

## Example

Output might look like this:

```
 _____________________
< Your daily fortune >
 ---------------------
   \
    \
        ^__^
        (oo)\_______
        (__)\       )\/\
            ||----w |
            ||     ||
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to check [issues page](https://github.com/zubayer204/random_cowsay/issues) if you want to contribute.

Happy coding and enjoy the randomized cowsay fun!
