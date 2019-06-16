# go-instagram

Instagramぶっこ抜き用APIサーバー

## Demo(Client)

![Demo](https://user-images.githubusercontent.com/17779386/58763844-7fad4e00-859b-11e9-958c-de0cedc60b72.gif)

これのサーバー側をGoでやる。

## Description

App or cURL ⇄  go-instagram ⇄  Instagram(Scraping)

![理想](https://user-images.githubusercontent.com/17779386/59564387-27894800-9081-11e9-84f3-30b97b3284a5.png)


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

## Example

```bash
$ curl -X POST -H "Content-Type:application/json" http://127.0.0.1:9000/instagram -d '{"URL":"https://www.instagram.com/p/ByNfe47AvR7/"}' | jq
{
  "Username": "https://instagram.fkix2-1.fna.fbcdn.net/vp/d95064d8a2957c759d1a7ec827f5f8df/5D7E4B69/t51.2885-15/e35/61400279_2257498447913426_8272622044097853955_n.jpg?_nc_ht=instagram.fkix2-1.fna.fbcdn.net",
  "ImageURL": "buky0907",
  "PostText": "呼！#karate#奥義#刃牙#japan#manga#street#愚地独歩#comics#ワンシーン#漫画#アニメ#コミック#ミリオンロック#奥義#炸裂#三戦#技#必殺技#"
}
```

## Todo

- 複数URL
- 非同期

