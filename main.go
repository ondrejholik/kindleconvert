package main

import (
  "io/ioutil"
  "log"
  "os"
  "os/exec"
  "strings"
  "sync"
)

func moveToLib(root, file string) {
  err := os.Rename(root+file, root+"mobi/"+file)
      if err != nil {
        log.Fatal(err)
      }
}

func deleteEpub(root, file string) {
  err := os.Remove(root+file)
  if err != nil {
    log.Fatal(err)
  }
}

func getPaths(root string) []string {
  var paths[]string
  formats := map[string]bool  { "epub" : true }
  files, err := ioutil.ReadDir(root)
  if err != nil {
      log.Fatal(err)
  }

  for _, file := range files {
    if !strings.Contains(file.Name(), ".") {
      continue
    }

    f := strings.Split(file.Name(), ".")
    fname := f[0]
    ex := f[1]
    formatOk := formats[ex]
    if !formatOk {
      continue
    }
    // add filename to slice 
    paths = append(paths, fname)
  }

    // return strings

    return paths

}

func main() {

  var wg sync.WaitGroup
  const root = "/home/USERNAME/Downloads"

  for _, x := range getPaths(root) {
    wg.Add(1)
    go func (root, x string) {
      cmd := exec.Command("ebook-convert", root+x+".epub", root+x+".mobi")
      cmd.Run()
      moveToLib(root, x+".mobi")
      deleteEpub(root, x+".epub")
      defer wg.Done()
    }(root,x)
  }

  wg.Wait()

}
