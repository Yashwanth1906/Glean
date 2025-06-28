package main

import (
  "fmt"
  "flag"
  "log"
  "path/filepath"
  "time"
  "github.com/Yashwanth1906/Glean/pkg/scanner"
  "github.com/Yashwanth1906/Glean/pkg/deleter"
)


func main() {
  start := time.Now()

  path := flag.String("path", ".", "Root path to start scanning")
  target := flag.String("target", "node_modules", "Target directory to delete")
  workers := flag.Int("workers", 8, "Number of concurrent deletion workers")
  dryRun := flag.Bool("dry-run", false, "If it is true doesn't delete anything, just prints the directories to be deleted")

  flag.Parse()

  absPath, err := filepath.Abs(*path)

  if err != nil {
    log.Fatalf("❌ Invalid path: %v", err)
  }

  fmt.Printf("🔍 Scanning path: %s for target: %s\n", absPath, *target)

  targets, err := scanner.FindTargetDirs(absPath, *target)

  if err != nil {
    log.Fatalf("❌ Error scanning directories: %v", err)
  }

  if len(targets) == 0 {
    fmt.Println("✅ No target directories found.")
    return
  }

  if *dryRun {
    fmt.Println("🔍 Dry run mode enabled. The following directories would be deleted:")

    for _, dir := range targets {
      fmt.Println(dir)
    }

    fmt.Printf("✅ Found %d target directories to delete.\n", len(targets))
  
  } else {
    
    fmt.Printf("🗑️ Deleting %d target directories with %d workers...\n", len(targets), *workers)
  
    deleter.DeleteTargets(targets, *workers)
  }

  elapsed := time.Since(start)
  fmt.Printf("⏱️ Execution time: %s\n", elapsed)
  fmt.Println("👋 Goodbye!")

}
