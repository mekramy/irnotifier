# IrNotifier API SDK

## Create Notifier Instance

```go
package main
import "github.com/mekramy/irnotifier"

func main() {
    client := irnotifier.NewNotifier(myApiKey)
}
```

### Handle Message Report

```go
package controllers

// POST /callback/sms
func HandleSMSReport(ctx http.RequestContext) error {
    // Check if request is valid
    if !client.IsValidRequest(ctx.Headers.Get("Authorization")) {
        return ctx.SendStatus(401)
    }

    if report, err := client.ParseReport(ctx.Body); err != nil {
        log.Log(err)
        return ctx.SendStatus(500)
    } else {
        // Handle report result ...
        return ctx.SendStatus(200)
    }
}
```

### Get Account Balance And Information

```go

func main() {
    if info, err := client.Info(); err != nil {
        log.Fatal(err)
    } else {
        log.Logf("Your Balance Is: %d, %d Queued, %d Pendings\n", info.Balance, info.Queued, info.Pendings)
    }
}
```

### Get Message Status

```go

func main() {
    if res, err := client.Inquiry("6640558a91bfe333529bef7a"); err != nil {
        if irnotifier.IsNotFoundErr(err){
            log.Log("Message not found!")
        }else{
            log.Fatal(err)
        }
    } else {
        log.Logf("Message Status: %s\n", res)
    }
}
```

### Get Message Sent List

```go

func main() {
    if res, err := client.Sent(
        1,
        irnotifier.PerPage50,
        irnotifier.SortSentAt,
        irnotifier.OrderAsc,
        "0912",
        "My Meta Filter",
        "1402-01-01",
        "1402-03-31",
    ); err != nil {
        log.Fatal(err)
    } else {
        log.Log(res)
    }
}
```

### Schedule Message Send

```go

func main() {
    if id, err := client.Queue(
        "login",
        "",
        "09120001122",
        "metadata",
        "https://mysite.com/sms-callback",
        time.Now(),
        time.Now().Add(5*time.Minute),
        map[string]string{"code": "12345"},
    ); err != nil {
        log.Fatal(err)
    } else {
        log.Logf("Message scheduled with %s id\n", id)
    }
}
```

### Update Scheduled Message Send

```go

func main() {
    if ok, err := client.ReQueue(
        "6640558a91bfe333529bef7a",
        "login",
        "",
        "09120003340",
        "",
        "https://mysite.com/sms-callback",
        time.Now(),
        time.Now().Add(5*time.Minute),
        map[string]string{"code": "12345"},
    ); err != nil {
        log.Fatal(err)
    } else {
        log.Log("Message scheduled updated")
    }
}
```

### Delete Scheduled Message

```go

func main() {
    if ok, err := client.UnQueue("6640558a91bfe333529bef7a"); err != nil {
        log.Fatal(err)
    } else {
        log.Log("Message scheduled deleted")
    }
}
```

## Error Handling

```go

func main() {
    err := client.DoSomeThing()

    if irnotifier.IsAPIError(err) {
        fmt.Println("Something wrong on irnotifier.ir api server!")
    }

    if irnotifier.IsServerError(err) {
        fmt.Println("Error 500 on server!")
    }

    if irnotifier.IsUnavailableError(err) {
        fmt.Println("Under maintenance!")
    }

    if irnotifier.IsAuthError(err) {
        fmt.Println("You are not authorized or your account is deactivated!")
    }

    if irnotifier.IsCreditError(err) {
        fmt.Println("Not enough credit!")
    }


    if irnotifier.IsNotFoundError(err) {
        fmt.Println("Record not found!")
    }

    if irnotifier.IsRequestLimitError(err) {
        fmt.Println("Too many request on server. Please try again later!")
    }

    if errors := irnotifier.ValidationErrors(err); errors != nil {
        fmt.Printf("Invalid input data: %+v\n", errors)
    }
}
```
