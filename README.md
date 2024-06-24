# NOT ALL FEATURES ARE IMPLEMENTED (PROJECT IN PROGRESS)

# 🎯 GoScan

## ⚡ Introduction

#### **GoScan is a powerful and efficient directory enumeration tool written in Go. It is designed to help security professionals identify hidden files and directories on web servers, providing insights that can assist in security assessments and vulnerability research.**

## 🚀 Features

- High-speed scanning using concurrent requests
- Customizable request headers
- Wordlist-based enumeration
- Simple and intuitive command-line interface
- Supports HTTP and HTTPS

## 🛠️ Installation

#### To install GoScan, you need to have Go installed on your system. You can download Go from the official website: https://golang.org/dl/.

#### Clone the repository and build the tool using the following commands:

```bash
git clone https://github.com/tr41z/GoScan.git
cd GoScan/cmd/main
go build main.go
```

## 📝 Usage

```bash
./GoScan -u <URL> -w <wordlist> [options]
Options:
-u, --url <URL>: Target URL
-w, --wordlist <wordlist>: Path to the wordlist file
-t, --threads <number>: Number of concurrent threads (default: 10)
-H, --header <header>: Custom header to include in requests (e.g., "Authorization: Bearer token")
-o, --output <file>: Output results to a file
```

Example:
```bash
./GoScan -u http://example.com -w wordlist.txt -t 20 -H "Authorization: Bearer token" -o results.txt
```

## ⚠️ Disclaimer

### Important Notice:

#### Running GoScan against targets without explicit permission is illegal. Ensure you have proper authorization before using this tool on any system. Responsibility: The developers of GoScan are not responsible for any misuse or damage caused by this tool. Use it responsibly and ethically.

---

*Thank you for using GoScan! Happy scanning and stay ethical!*
