package scanner


import (
  "errors"
  "io/fs"
  "os"
  "path/filepath"
)


func FindTargetDirs(rootPath string, target string) ([]string, error) {
  var matches []string

  info, err := os.Stat(rootPath) // Check if the root path exists
  if err != nil {
    return nil, err
  }
  if !info.IsDir() {
    return nil, errors.New("provided path is not a directory")
  }

  // it is single threaded, no go routines are used here to maintatin the search process simple because
  // using go routines here may cause several issues
  // Already walkDir is optmized for performance
  
  err = filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
    if err != nil {
      return err
    }

    if d.IsDir() && d.Name() == target {
      matches = append(matches, path)
      return filepath.SkipDir
    }

    return nil
  })

  return matches, err
}
