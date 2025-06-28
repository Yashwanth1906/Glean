# 🧹 GLEAN - Go Local Environment Auto Node_modules Cleaner

**GLEAN** is a blazing-fast, concurrent CLI tool written in Go that **recursively finds and deletes all `node_modules` folders** (or any other bloated folders) in a given directory.  
It’s perfect for developers who want to **free up gigabytes of disk space** from old, unused projects without hunting folders one by one.

---

## ⚡ Why GLEAN?

- 🔍 Recursively scans large directories
- ⚙️ Deletes folders using a **parallel worker pool (goroutines)**
- 💨 Cleans your system **faster than traditional shell scripts**
- 🧪 Supports `--dry-run` mode for safe previews
- 🔧 Customizable: set your own folder name, worker count

---

## 🚀 Installation

> Requires **Go 1.16+**

```bash
go install github.com/Yashwanth1906/Glean/cmd/glean@latest


This installs the `glean` binary to your `$GOPATH/bin` directory.


# Make sure to add $GOPATH/bin to your PATH

export PATH=$PATH:$GOPATH/bin
``` 

## 🛠 Usage

```bash
- Dry Run Mode:

glean --path=/home/user/projects --target=node_modules --dry-run=true

- Deletion Mode:

glean --path=/home/user/projects --target=node_modules --workers=8
```
