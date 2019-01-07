# redirect

<a href="https://www.buymeacoffee.com/virtcanhead" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>
[![Build Status](https://travis-ci.org/virtcanhead/redirect.svg?branch=master)](https://travis-ci.org/virtcanhead/redirect)

Golang based redirect tool

## Environment Variable

* `PORT`, port to listen, default to `80`

* `REDIRECT_TARGET`, the target for redirection, for example, `https://example.com`

* `REDIRECT_PERMANENTLY`, if set, will use 301 instead of 302 for redirection

* `REDIRECT_NO_PATH`, if set, will ignore path while redirecting

## License

canhead <hi@canhead.xyz> MIT License
