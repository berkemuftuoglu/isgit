# isgit

**isgit** is a lightweight CLI tool to check if a directory is a Git repository and provide key repository details.

---

## Features
- Detect if a directory is a Git repository.
- Show current branch name.
- Display the latest commit details (hash, message, and date).
- Fetch the remote repository URL.

---

## Installation

### From Source
Ensure you have [Go](https://golang.org/dl/) installed. Then:
```bash
git clone git@github.com:berkemuftuoglu/isgit.git
cd isgit
go build -o isgit
```
Move the binary to a directory in your PATH for global usage:
```bash
sudo mv isgit /usr/local/bin/
```

---

## Usage

### Basic Usage
Check if the current directory is a Git repository:
```bash
isgit
```

### Check Another Directory
```bash
isgit --path /path/to/directory
```

### Show Detailed Repository Info
```bash
isgit --details
```

### Display Remote Repository Only
```bash
isgit --remote
```

### Show Help
```bash
isgit --help
```

---

## Example Output

### Details Command
```bash
isgit --details
Current branch: main
Last commit: abc123 Initial commit (2025-01-08 12:00:00 +0000)
Remote: git@github.com:berkemuftuoglu/isgit.git
```

---

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

---

## License
This project is licensed under the [MIT License](LICENSE).
