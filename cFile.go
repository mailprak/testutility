package testpackage

import (
    "io/ioutil"
    "log"
)

func copyFile(srcName, targetName string) {


    bytesRead, err := ioutil.ReadFile(srcName)

    if err != nil {
        log.Fatal(err)
    }

    err = ioutil.WriteFile(targetName, bytesRead, 0644)

    if err != nil {
        log.Fatal(err)
    }
}
