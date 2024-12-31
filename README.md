# IrNotifier API SDK Documentation

## Overview

IrNotifier is a Go library for sending text SMS messages. This documentation provides an overview of the functions and interfaces available in the library, along with examples to help you get started.

## Installation

To install the library, use the following command:

```sh
go get github.com/mekramy/irnotifier
```

## Creating a Notifier Instance

To create a new notifier instance, use the `NewNotifier` function:

```go
package main

import "github.com/mekramy/irnotifier"

func main() {
    client := irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)
}
```

## Interfaces

### Notifier Interface

The `Notifier` interface defines the methods available for interacting with the IrNotifier API.

```go
type Notifier interface {
    IsValidRequest(authHeader string) bool
    ParseReport(body []byte) (*Report, error)
    Information() (*Information, error)
    Statistic(metadata *string) (*Statistics, error)
    Inquiry(id string) (MessageStatus, error)
    Queue(parameter *QueueParams) (string, error)
    Requeue(id string, parameter *QueueParams) (bool, error)
    Dequeue(id string) (bool, error)
    Suspend(metadata string, force bool) (int64, error)
    Resume(metadata string) (int64, error)
    DequeueAll(parameter *DequeueParams) (int64, error)
    FailList(parameter *FailParams) (*SearchResult[FailMessage], error)
    SentList(parameter *SentParams) (*SearchResult[SentMessage], error)
}
```

## Functions

### IsValidRequest

Checks if the request authorization header is valid.

```go
func (notifier irNotifier) IsValidRequest(authHeader string) bool
```

### ParseReport

Parses a delivery report sent by POST method from irnotifier.ir.

```go
func (notifier irNotifier) ParseReport(body []byte) (*Report, error)
```

### Information

Gets client information.

```go
func (notifier irNotifier) Information() (*Information, error)
```

### Statistic

Gets client statistics.

```go
func (notifier irNotifier) Statistic(metadata *string) (*Statistics, error)
```

### Inquiry

Gets the status of a message.

```go
func (notifier irNotifier) Inquiry(id string) (MessageStatus, error)
```

### Queue

Sends a new message.

```go
func (notifier irNotifier) Queue(parameter *QueueParams) (string, error)
```

### Requeue

Updates a queued message.

```go
func (notifier irNotifier) Requeue(id string, parameter *QueueParams) (bool, error)
```

### Dequeue

Deletes a queued message.

```go
func (notifier irNotifier) Dequeue(id string) (bool, error)
```

### Suspend

Suspends queued messages.

```go
func (notifier irNotifier) Suspend(metadata string, force bool) (int64, error)
```

### Resume

Resumes suspended messages.

```go
func (notifier irNotifier) Resume(metadata string) (int64, error)
```

### DequeueAll

Deletes all queued messages.

```go
func (notifier irNotifier) DequeueAll(parameter *DequeueParams) (int64, error)
```

### FailList

Gets a list of failed messages.

```go
func (notifier irNotifier) FailList(parameter *FailParams) (*SearchResult[FailMessage], error)
```

### SentList

Gets a list of sent messages.

```go
func (notifier irNotifier) SentList(parameter *SentParams) (*SearchResult[SentMessage], error)
```

## Examples

### Handle Message Report

```go
package controllers

import (
    "github.com/mekramy/irnotifier"
    "net/http"
)

var client = irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)

// POST /callback/sms
func HandleSMSReport(w http.ResponseWriter, r *http.Request) {
    if !client.IsValidRequest(r.Header.Get("Authorization")) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    body, _ := io.ReadAll(r.Body)
    if report, err := client.ParseReport(body); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    } else {
        // Handle report result ...
        w.WriteHeader(http.StatusOK)
    }
}
```

### Get Account Balance and Information

```go
package main

import (
    "github.com/mekramy/irnotifier"
    "log"
)

func main() {
    client := irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)
    if info, err := client.Information(); err != nil {
        log.Fatal(err)
    } else {
        log.Printf("Your Balance Is: %d, %d Queued, %d Pendings\n", info.Balance, info.Queued, info.Pendings)
    }
}
```

### Get Message Status

```go
package main

import (
    "github.com/mekramy/irnotifier"
    "log"
)

func main() {
    client := irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)
    if res, err := client.Inquiry("6640558a91bfe333529bef7a"); err != nil {
        if irnotifier.IsNotFoundError(err) {
            log.Println("Message not found!")
        } else {
            log.Fatal(err)
        }
    } else {
        log.Printf("Message Status: %s\n", res)
    }
}
```

### Get Sent Message List

```go
package main

import (
    "github.com/mekramy/irnotifier"
    "log"
)

func main() {
    client := irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)
    if res, err := client.SentList(nil); err != nil {
        log.Fatal(err)
    } else {
        log.Println(res)
    }
}
```

### Schedule Message Send

```go
package main

import (
    "github.com/mekramy/irnotifier"
    "log"
    "time"
)

func main() {
    client := irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)
    params := irnotifier.QueueParameter().
        To("09120001122").
        Metadata("test").
        Pattern("test").
        SendAt(time.Now()).
        Expiration(time.Now().Add(5 * time.Minute)).
        AddParameter("code", "12345")
    if id, err := client.Queue(params); err != nil {
        log.Fatal(err)
    } else {
        log.Printf("Message scheduled with %s id\n", id)
    }
}
```

### Update Scheduled Message Send

```go
package main

import (
    "github.com/mekramy/irnotifier"
    "log"
    "time"
)

func main() {
    client := irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)
    params := irnotifier.QueueParameter().
        To("09120003340").
        Metadata("test").
        Pattern("test").
        SendAt(time.Now()).
        Expiration(time.Now().Add(5 * time.Minute)).
        AddParameter("code", "12345")
    if ok, err := client.Requeue("6640558a91bfe333529bef7a", params); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Message scheduled updated")
    }
}
```

### Delete Scheduled Message

```go
package main

import (
    "github.com/mekramy/irnotifier"
    "log"
)

func main() {
    client := irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)
    if ok, err := client.Dequeue("6640558a91bfe333529bef7a"); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Message scheduled deleted")
    }
}
```

## Error Handling

IrNotifier provides several helper functions to handle different types of errors.

```go
package main

import (
    "github.com/mekramy/irnotifier"
    "log"
)

func main() {
    client := irnotifier.NewNotifier("your-api-key", "https://irnotifier.ir", irnotifier.V1)
    err := client.DoSomeThing()

    if irnotifier.IsAPIError(err) {
        log.Println("Something wrong on irnotifier.ir API server!")
    }

    if irnotifier.IsServerError(err) {
        log.Println("Error 500 on server!")
    }

    if irnotifier.IsUnavailableError(err) {
        log.Println("Under maintenance!")
    }

    if irnotifier.IsAuthError(err) {
        log.Println("You are not authorized or your account is deactivated!")
    }

    if irnotifier.IsCreditError(err) {
        log.Println("Not enough credit!")
    }

    if irnotifier.IsNotFoundError(err) {
        log.Println("Record not found!")
    }

    if irnotifier.IsRequestLimitError(err) {
        log.Println("Too many requests on server. Please try again later!")
    }

    if errors := irnotifier.ValidationErrors(err); errors != nil {
        log.Printf("Invalid input data: %+v\n", errors)
    }
}
```
