# go-instagram

Instagramぶっこ抜き用APIサーバー

## Demo(Client)

![Demo](https://user-images.githubusercontent.com/17779386/59974954-5bb6b880-95ed-11e9-9eb3-56627862e1c8.gif)


これのサーバー側をGoでやる。

## Description

App or cURL ⇄  go-instagram ⇄  Instagram(Scraping)  

イメージ(理想)
![理想](https://user-images.githubusercontent.com/17779386/59564387-27894800-9081-11e9-84f3-30b97b3284a5.png)


go-instagram response the following.

- Username(ユーザー名)
- ImageURL(画像URL)
- PostText(投稿文)
- OrgURL(投稿元URL)

## Usage

```bash
$ export PORT=9000
$ go run main.go

# Post request
$ curl -X POST \
        -H "Origin:https://rasukarusan.github.io/instagram-client" \
        -H "Content-Type:application/json" \
        http://127.0.0.1:9000/instagram \
        -d '{"URLs":["https://www.instagram.com/p/XXXX/","https://www.instagram.com/p/YYYYY/"]}' | jq

```

## Example

```bash
curl -s -X POST \
         -H "Origin:https://rasukarusan.github.io/instagram-client" \
         -H "Content-Type:application/json" \
         http://127.0.0.1:9000/instagram \
         -d '{"URLs":["https://www.instagram.com/p/By2CAExjikt/","https://www.instagram.com/p/ByNfe47AvR7/"]}' | jq


[
  {
    "Username": "buky0907",
    "ImageURL": "https://instagram.ffuk4-2.fna.fbcdn.net/vp/6266b086bf1301741507de86908431c9/5DA5D869/t51.2885-15/e35/61400279_2257498447913426_8272622044097853955_n.jpg?_nc_ht=instagram.ffuk4-2.fna.fbcdn.net",
    "PostText": "呼！#karate#奥義#刃牙#japan#manga#street#愚地独歩#comics#ワンシーン#漫画#アニメ#コミック#ミリオンロック#奥義#炸裂#三戦#技#必殺技#",
    "OrgURL": "https://www.instagram.com/p/ByNfe47AvR7/",
    "Err": ""
  },
  {
    "Username": "shimokitapiapia",
    "ImageURL": "https://instagram.ffuk4-2.fna.fbcdn.net/vp/e2b1a560456307e35c9aeb4d6291661a/5D8D4DF8/t51.2885-15/fr/e15/s1080x1080/64210746_2319604821587559_3466696255239768721_n.jpg?_nc_ht=instagram.ffuk4-2.fna.fbcdn.net",
    "PostText": "強くなりたくばぴあ＆ぴあで喰らえ!! #下北沢#下北#シモキタ#居酒屋#ぴあぴあ#ぴあ＆ぴあ#刃牙#範馬勇次郎#鬼#japan#setagaya#shimokitazawa#izakaya#piapia#pia&pia",
    "OrgURL": "https://www.instagram.com/p/By2CAExjikt/",
    "Err": ""
  }
]
```

## Todo

- 複数URL
- 非同期

