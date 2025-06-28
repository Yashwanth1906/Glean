# 🧹 GLEAN - Go Local Environment Auto Node_modules Cleaner

**GLEAN** is a blazing-fast, concurrent CLI tool written in Go that **recursively finds and deletes all `node_modules` folders** (or any other bloated folders) in a given directory.  
It’s perfect for developers who want to **free up gigabytes of disk space** from old, unused projects without hunting folders one by one.

---

## Why GLEAN?

- Recursively scans large directories
- Deletes folders using a **parallel worker pool (goroutines)**
- Cleans your system **faster than traditional shell scripts**
- Supports `--dry-run` mode for safe previews
- Customizable: set your own folder name, worker count

---

## Installation

> Requires **Go 1.16+**

```bash
go install github.com/Yashwanth1906/Glean/cmd/glean@latest
```

## 🛠 Usage

- Dry Run Mode:

```bash
glean --path=dir_you_want_to_delete --target=directory_that_you_want_delete --dry-run=true
```

Output:
![image](https://github.com/user-attachments/assets/6f408d56-938f-4849-bbb8-5f5484504389)

- Deletion Mode:
  
```bash
glean --path=dir_you_want_to_delete --target=directory_that_you_want_delete --workers=no_of_workers_you_want_use
```

Output:
![image](https://github.com/user-attachments/assets/5cbe95a1-15fb-4016-beda-abaa662afc58)

That's it you can use my tool easily...
