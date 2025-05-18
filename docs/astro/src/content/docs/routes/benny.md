---
title: Benny Route
description: A reference page for the Benny route in my API.
---

The Benny route (`/api/v1/benny`) provides access to a collection of photos featuring Benny. Each request returns a random image from the collection.

## Endpoint

```
GET /api/v1/benny
```

## Query Parameters

| Parameter | Type    | Required | Description                           |
| --------- | ------- | -------- | ------------------------------------- |
| width     | integer | No       | Desired width of the image in pixels  |
| height    | integer | No       | Desired height of the image in pixels |

## Response

Returns a JPEG image with the following characteristics:

- Content-Type: `image/jpeg`
- Default size: Original image dimensions
- Resized: When width/height parameters are provided

## Examples

### Basic Usage

```bash
curl https://api.jed.one/api/v1/benny
```

### Resized Image

```bash
curl https://api.jed.one/api/v1/benny?width=500&height=300
```

## Notes

- If only one dimension (width or height) is provided, the image maintains its aspect ratio
- The API automatically handles image scaling and optimization
- The Benny route is a fun way to share images of Benny with friends and family
