# Roku

https://justwatch.com/us/provider/the-roku-channel

## Android client

just a remote control:

https://play.google.com/store/apps/details?id=com.roku.remote

## Web premium

https://therokuchannel.roku.com/watch/32c95b576307502b98f7fe32c4aa0a22

We can create free account, but login uses CAPTCHA:

~~~
POST https://my.roku.com/api/auth/1/login HTTP/2.0
csrf-token: sVzM79JV-kKC0kN2Jlz_PlI6vHLZ3NhqOqFk
content-type: application/json
cookie: _csrf=LHBM-wxg8GRExB8JboGxeJCC

{"email":"EMAIL","password":"PASSWORD","rememberMe":false}

{
  "error": "captcha"
}
~~~
