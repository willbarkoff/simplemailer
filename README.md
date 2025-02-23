# SimpleMailer

![Go version](https://img.shields.io/github/go-mod/go-version/willbarkoff/simplemailer?logo=go&style=flat-square)
[![Latest release](https://img.shields.io/github/v/tag/willbarkoff/simplemailer?label=latest%20release&sort=semver&style=flat-square)](https://github.com/willbarkoff/simplemailer/releases)
[![License](https://img.shields.io/github/license/willbarkoff/simplemailer?style=flat-square)](./LICENSE.md)

- [⬇️ **Download**](https://github.com/willbarkoff/simplemailer/releases) 
- [🐛 **Report a bug**](https://github.com/willbarkoff/simplemailer/issues/new)

---

SimpleMailer is a program that is used to send emails with SMTP. To use it create a config.toml file, an addresses.txt file, and a templates/ directory.

The `config.toml` file should have the following format:
```toml
FromAddress = "barry@starlabs.com" # the email address to send from
FromDisplay = "Barry Allen <barry@starlabs.com>" # the email address and name to display in sent mail
SMTPHost = "smtp.starlabs.com" # the SMTP server
SendDomain = "starlabs.com" # the domain to send from
SMTPPort = 587 # the port to use with SMTP
SMTPUsername = "barry@starlabs.com" # the SMTP username
SMTPPassword = "superSecret1!" # the SMTP password
```

The `addresses.txt` file should be a list of email addresses, with one per line. For example:
```
caitlin@starlabs.com
cisco@starlabs.com
ralph@starlabs.com
iris@cccitizen.com
joe_west@ccpd.centralcity.gov
ceclile.horton@centralcity.gov
```

The `templates/` directory should contain a folder for each message template to send. Each of those folders should contain a `subject.txt` file used for the subject of the email, a `template.txt` file for the plaintext version of the email, and a `template.html` file for the HTML version of the email. Any files included in the `templates/` directory will be included in the templates. You can use go's [`text/template`](https://golang.org/pkg/text/template/) templating language to write templates. As of right now, there is no way to pass values to the templates, but I hope to add that soon.

## SimpleMailer usage
```shell
$ simplemailer --help
Usage of simplemailer:
  -addresses string
        The list of email addresses to send to (default "addresses.txt")
  -config string
        The configuration file to use for sending mail (default "config.toml")
  -stopOnFail
        Whether to stop sending on a send failure
  -template string
        The template to send (default "message")
  -templatePath string
        The path to the template directory to use when sending (default "templates/")
```

By default, SimpleMailer will send each email and continue sending each remaining email if it runs into an issue. If you want it to stop after an issue, you can use the ``--stopOnFail`` flag, for example:

```
$ simplemailer
go run main.go
An error occured sending to caitlin@starlabs.com: 550 Invalid sender eva@mccullochtech.com
An error occured sending to cisco@starlabs.com: 550 Invalid sender eva@mccullochtech.com
An error occured sending to ralph@starlabs.com: 550 Invalid sender eva@mccullochtech.com
An error occured sending to iris@cccitizen.com: 550 Invalid sender eva@mccullochtech.com
An error occured sending to thawne@starlabs.com: 550 Invalid sender eva@mccullochtech.com
An error occured sending to joe_west@ccpd.centralcity.gov: 550 Invalid sender eva@mccullochtech.com
An error occured sending to ceclile.horton@centralcity.gov: 550 Invalid sender eva@mccullochtech.com
$ simplemailer --stopOnFail
An error occured sending to caitlin@starlabs.com:
panic: 550 Invalid sender eva@mccullochtech.com

goroutine 1 [running]:
main.main()
        /Users/williambarkoff/go/src/github.com/willbarkoff/simplemailer/main.go:49 +0x975
exit status 2
```

## Building from source
1. Install go.
2. Install dependencies (SimpleMailer uses [Go Modules](https://blog.golang.org/using-go-modules)):
```shell
$ go get
```
3. Run
```shell
$ go run main.go
```

## Contributing
Pull requests are welcome!