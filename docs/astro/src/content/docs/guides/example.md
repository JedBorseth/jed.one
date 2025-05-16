---
title: Getting Started
description: A guide to getting started with my API.
---

This is a guide to getting started with my API. It includes information on how to set up your environment, make requests, and handle responses. This API and its documentation were built as a joke, but they are functional and can be used for various purposes.

## Routes

Many of the routes in this API are available at the following endpoints:

- `/api/v1/welcome`
- `/api/v1/welcome`
- `/api/v1/benny`

## Making Requests

To make a request to the API, you can use any HTTP client.
Here is a one line example using `reqwest` in Rust:

```rust
println!("{}", reqwest::blocking::get("https://api.jed.one/api/v1/welcome").unwrap().text().unwrap());
```

Now here is one using Haskell:

```haskell
import Network.HTTP.Simple; import qualified Data.ByteString.Lazy.Char8 as L8

main = httpLBS "https://api.jed.one/api/v1/welcome" >>= L8.putStrLn . getResponseBody
```

## Handling Responses

The API returns JSON responses. You can use any JSON library to parse the responses.

Here is an example using `javascript`:

```javascript
var data;
var responseText;

try {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/v1/welcome", false);
  xhr.send(null);

  responseText = xhr.responseText;

  data = eval("(" + responseText + ")");
} catch (e) {}

var message = data.message;

document.body.innerHTML = "<h1>" + message + "</h1>";

window.globalData = data;

console.log("Server said: " + responseText);
```

## Further reading

- Read [about how-to guides](https://diataxis.fr/how-to-guides/) in the Diátaxis framework
