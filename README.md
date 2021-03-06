# go-firebase-admin

[![Go Report Card](https://goreportcard.com/badge/github.com/acoshift/go-firebase-admin)](https://goreportcard.com/report/github.com/acoshift/go-firebase-admin)
[![GoDoc](https://godoc.org/github.com/acoshift/go-firebase-admin?status.svg)](https://godoc.org/github.com/acoshift/go-firebase-admin)

> Work in Progress

Firebase Admin SDK for Golang

## Usage

See godoc for more detail.

```go
package main

import (
  "io/ioutil"

  "github.com/acoshift/go-firebase-admin"
)

func main() {
  // Init App
  serviceAccount, _ := ioutil.ReadFile("service_account.json")
  firApp, err := admin.InitializeApp(admin.AppOption{
    ServiceAccount: serviceAccount,
    ProjectID: "YOUR_PROJECT_ID",
    DatabaseURL: "YOUR_DATABASE_URL",
  })
  if err != nil {
    panic(err)
  }
  firAuth := firApp.Auth()
  firDatabase := firApp.Database()
  // ...
}
```
