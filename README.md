# InstaTG

Instagram to Telegram Channel Bot. Can access posts from any public Instagram account or an account that you follow.

### Features

- Keeps track of existing sent posts to tg in an embedded db.
- Fast & Concurrent (at least as much as Telegram's rate limit allows it to be).
- You can send the latest posts at periodic intervals or send all past Instagram posts.
- Tested on Linux, but should also work across platforms (Win/Mac/Linux) and architectures (ARM/x86).

### Prerequisites

- Have a spare Instagram account.
- Create a telegram bot using botfather and obtain its HTTP token.
- Create a telegram channel.
- Add your bot to your channel and make it admin.
- Get your channel chat id using `https://api.telegram.org/botXXX:YYYYY/getUpdates` (replace the XXX:YYYYY with your BOT HTTP API Token you just got from the Telegram BotFather)

### Usage

There are two versions:

* `cmd/app` - Fetch the latest posts from an Instagram channel and send them to a Telegram channel. 
* `cmd/historical` - Fetch all historical/past posts from an Instagram channel and send them to a Telegram channel.

You should probably run `cmd/app` as a background service while `cmd/historical` is a one-off.

Compile using `go build -o ./bin/instatg ./cmd/app`

Then run

```bash
$ ./instatg \
      -tgtoken "XXXXX:YYYYY" \
      -tgchannel "-123456789" \
      -iguser "yourusername" \
      -igpass "yourpassword" \
      -igchan "exampleaccount"
```

#### From CLI

**_Note: You need Go to compile it or maybe you can grab pre-compiled builds from somewhere._**

**_Note: The directory where you place the binaries must be user writeable._**

```
Usage of instatg:
  -dbpath string
        Database File Path (optional) (default "./store.boltdb")
  -igchan string
        Instagram Channel's Username
  -igpass string
        Your Instagram Password
  -iguser string
        Your Instagram Username
  -tgchannel int
        Telegram Channel ID
  -tgtoken string
        Telegram Bot Token
```

#### As SystemD Service

Add service in `~/.config/systemd/user/instatg.service`

```
[Unit]
Description=InstaTg Bot for Channel X

[Service]
WorkingDirectory=/home/icewreck/somefolder
ExecStart=/home/icewreck/somefolder/instatgapp \
                                    -tgtoken "XXXXX:YYYYY" \
                                    -tgchannel "-123456789" \
                                    -iguser "yourusername" \
                                    -igpass "yourpassword" \
                                    -igchan "exampleaccount"

[Install]
WantedBy=default.target


```
And then enable it.

```bash

systemctl --user daemon-reload
systemctl --user start instatg.service
systemctl --user enable instatg.service
# check status
systemctl --user status instatg.service

```

### Thanks To

- [github.com/ahmdrz/goinsta](https://github.com/ahmdrz/goinsta)
- [github.com/boltdb/bolt](https://github.com/boltdb/bolt)
- [github.com/go-telegram-bot-api/telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)
