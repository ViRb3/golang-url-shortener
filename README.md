# Golang URL Shortener

## Main Features

- URL Shortening
- Visitor Counting
- Expirable Links
- URL deletion
- Multiple authorization strategies:
    - Local authorization via OAuth 2.0 (Google, GitHub, Microsoft, and Okta)
    - Proxy authorization for running behind e.g. [Google IAP](https://cloud.google.com/iap/)
- Easy [ShareX](https://github.com/ShareX/ShareX) integration
- Dockerizable
- Multiple supported storage backends
    - High performance local database with [bolt](https://github.com/boltdb/bolt)
    - Persistent non-local storage with [redis](https://redis.io/)

## [Webinterface](https://so.sh0rt.cat)

![Short URLs](https://user-images.githubusercontent.com/17984549/32700384-955d9336-c7c4-11e7-9fab-4141a86a375c.png)

---

![Generate ShareX Configuration](https://user-images.githubusercontent.com/17984549/32700395-cf9f057a-c7c4-11e7-9d2b-7523c8a95a20.png)

## Documentation

- [Installation](https://github.com/mxschmitt/golang-url-shortener/wiki/Installation)
- [Configuration](https://github.com/mxschmitt/golang-url-shortener/wiki/Configuration)
- [Setting up OAuth](https://github.com/mxschmitt/golang-url-shortener/wiki/Setting-up-OAuth)
- [ShareX Usage](https://github.com/mxschmitt/golang-url-shortener/wiki/ShareX)
