# Fontster

Fontster is a generic font server that supports generating `@font-face` types dynamically
It has a "Google fonts" like API

## Features

- Serving from filesystem or from other http servers (GCS, S3, Nginx).
- HTTP/2 Push (TLS must be enabled).
- Only latest web fonts format is supprted (woff2).

## API

`/css?family=Lato|Roboto:700`
`/css?family=Lato:200i`

## Notes

The folder structure for your fonts should be like the following: `Family/Family-Style.woff2`

Examples:

- `Lato/Lato-Regular.woff2`
- `Lato/Lato-BoldItalic.woff2`
