package testpackage

func AddIntNum(num1, num2 int) int {
  return num1 + num2
}

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
