# 🎯 dirgo

<p align="center">
  <img src="https://github.com/user-attachments/assets/1956bdcd-43cb-4b35-82ee-ae394fc258cb" alt="dirgo" width="400" height="400"/>
  <img src="https://github.com/user-attachments/assets/137f1e94-c304-437b-a521-509b3c2b7e5b" alt="dirgo-hack" width="400" height="400"/>
</p>

<br>
<br>

## ⚡ Introduction

#### **dirgo is a powerful and efficient directory enumeration tool written in Go. It is designed to help security professionals identify hidden files and directories on web servers, providing insights that can assist in security assessments and vulnerability research.**

```bash
GoScan % make run

================================================
                      
        .--|  ||__|.----..-----..-----.
        |  _  ||  ||   _||  _  ||  _  |
        |_____||__||__|  |___  ||_____|
                         |_____|       

© 2024 tr41z

Language used: Golang
Github Repo: https://github.com/tr41z/dirgo
================================================

URL                                                    LENGTH    STATUS_CODE
-----------------------------------------------------------------------------
http://localhost:8080/passwords                        1161      200       
http://localhost:8080/index.html                       4177      200       
http://localhost:8080/.cache                           373       200       
http://localhost:8080/.config                          1031      200       
http://localhost:8080/.mysql_history                   2038      200       
http://localhost:8080/.ssh                             299       200       
```

## 🚀 Features

- High-speed scanning using concurrent requests
- Customizable request flags
- Wordlist-based enumeration
- Simple and intuitive command-line interface
- Supports HTTP and HTTPS

## 🛠️ Installation

#### To run dirgo, you need to have Go installed on your system. You can download Go from the official website: https://golang.org/dl/.

#### Clone the repository and build the tool using the following commands:

```bash
git clone https://github.com/tr41z/dirgo.git
cd dirgo/cmd/main
go build
```

## 📝 Usage

```bash
./dirgo -u <URL> -w <wordlist> [options]
Options:
-u, --url <URL>: Target URL
-w, --wordlist <wordlist>: Path to the wordlist file
-t, --threads <number>: Number of concurrent threads (default: 10)
-s, --status <number>: Target status code (default: 200)
-o, --output <file>: Output results to a file
```

Example:
```bash
./dirgo -u http://example.com:8080 -w wordlist.txt -t 20 -s 403 -o results.txt
```

## ⚠️ Disclaimer

### Important Notice:

#### Running dirgo against targets without explicit permission is illegal. Ensure you have proper authorization before using this tool on any system. The developers of dirgo are not responsible for any misuse or damage caused by this tool. Use it responsibly and ethically.

---

*Thank you for using dirgo! Happy scanning and stay ethical!*
