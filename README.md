# go-disposable

go-disposable can be used to detect if a domain or email address is potentially fake or temporary based on it being registered using a disposable email provider (identified by the domain name).

It is based on a great project called [disposable-email-domains](https://github.com/ivolo/disposable-email-domains) by [ivolo](https://github.com/ivolo). In fact, the list of domains are extracted from this project. Big thanks!

## Examples
Check if a specific domain is a known disposable email provider:
```go
fmt.Println(disposable.IsDomainDisposable("quant2new.pro")) // true
fmt.Println(disposable.IsDomainDisposable("gmail.com")) // false
```

Check if a specfic domain is subdomain of a known wildcard disposable email provider
```go
fmt.Println(disposable.IsDomainWildcard("some.subdomain.anonbox.net")) // true
```

Check if an email address is considered disposable
```go
fmt.Println(disposable.IsEmailAddressDisposable("foo@sub.33m.co")) // true
```

## Keep up to date
If you want the most recent up to date version of the list of disposable email domains, you can clone this repository and run the `go generate` command in the root.