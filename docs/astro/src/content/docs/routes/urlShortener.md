---
title: Url Shortener
description: A page for the URL shortener route in my API.
---

The URL shortener route (`/api/v1/url`) is a simple API that takes a long URL and returns a shortened version of it. This is useful for sharing links in a more manageable format.

## Endpoint

```
GET /api/v1/url
```

## Query Parameters

| Parameter | Type   | Required | Description                  |
| --------- | ------ | -------- | ---------------------------- |
| url       | string | Yes      | The long URL to be shortened |

## Response

Returns a JSON object with the following structure:

```json
{
  "shortened_url": "https://api.jed.one/url/abc123"
}
```

## Examples

### Basic Usage

```bash
curl "https://api.jed.one/api/v1/url?url=https://example.com/very/long/url"
```

### Shortened URL

```json
{
  "shortened_url": "https://api.jed.one/url/abc123"
}
```

## Notes

- The shortened URL will redirect to the original long URL when accessed.
- The API uses a hashing algorithm to generate unique shortened URLs.
