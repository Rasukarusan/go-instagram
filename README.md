# go-instagram

Instagramぶっこ抜き用APIサーバー

## Description

App or cURL ⇄  go-instagram ⇄  Instagram(Scraping)

go-instagram response the following.

- Username(ユーザー名)
- ImageURL(画像URL)
- PostText(投稿文)
- OrgURL(投稿元URL)

## Usage

```bash
# Use port 9000
$ go run main.go

# Post request
$ curl -X POST -H "Content-Type:application/json" http://127.0.0.1:9000/instagram -d '{"URL":"https://www.instagram.com/p/XXXXXX"}' | jq
```

## Todo

- 複数URL
- 非同期
