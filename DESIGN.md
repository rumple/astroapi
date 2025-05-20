# Repository Design and Explanation

## Overall Purpose

This repository contains a simple web application written in the Go programming language. Its primary function is to provide an API that calculates and returns the sunrise and sunset times for a given geographical latitude and longitude. The project also includes a basic HTML frontend to demonstrate the API's usage.

## File Explanations

### `app.yaml`
This file is the configuration file for deploying the application to Google App Engine. It specifies:
- The application ID (`astroapi`).
- The runtime environment (`go`).
- The API version (`go1`).
- URL handlers:
    - `/favicon.ico`: Served as a static file.
    - `/` (root path): Serves the `index.html` file.
    - `/.*` (all other paths): Handled by the Go application script (`_go_app`).

### `index.html`
This is the main HTML page that serves as the frontend for the application. Its key features are:
- **API Explanation**: It describes what the API does and the expected URL formats for requests.
- **Usage Examples**: Provides concrete examples of how to call the API.
- **JavaScript Functionality**: Includes a JavaScript function (`getastro()`) that:
    - Uses the browser's `navigator.geolocation.getCurrentPosition()` to get the user's current latitude and longitude.
    - Makes asynchronous HTTP requests (XHR) to the backend API (`/v0/sunrise/...` and `/v0/sunset/...`) using these coordinates.
    - Displays the retrieved sunrise and sunset times in a browser alert.
- **Styling**: Basic CSS is included for presentation.

### `sun.go`
This is the core Go application file containing the backend logic.
- **Package**: Belongs to the `main` package.
- **Dependencies**: Imports necessary packages including `fmt`, `net/http`, `strconv`, `strings`, `time`, and importantly `github.com/cpucycle/astrotime` for the astronomical calculations.
- **Constants**: Defines `ver` ("v0") for API versioning and `lenPath` (4) for URL path component count validation.
- **`SunHandlerV0` function**: This is the primary HTTP handler for API requests.
    - It only accepts `GET` requests.
    - It expects a URL path with 4 components (e.g., `v0/sunrise/latXX.XX/lonYY.YY`).
    - It parses the latitude and longitude from the URL path (e.g., `lat38.8895`, `lon77.0352`).
    - If `sunrise` is requested, it calls `astrotime.NextSunrise(time.Now(), latitude, longitude)`.
    - If `sunset` is requested, it calls `astrotime.NextSunset(time.Now(), latitude, longitude)`.
    - It returns the calculated time as a plain text string in Go's default time format.
    - It includes error handling for incorrect request methods, malformed URLs, or parsing errors, returning appropriate HTTP status codes (e.g., `http.StatusNotFound`, `http.StatusBadRequest`).
- **`init()` function**: This function is automatically executed when the package is initialized. It registers `SunHandlerV0` to handle all HTTP requests under the path `/v0/`.

### `README.md`
The existing `README.md` file provides a brief overview of the project, its purpose (a weekend project to learn Go), and basic API usage instructions. This `DESIGN.md` file offers a more in-depth explanation of the codebase and its components.

### `LICENSE`
This file contains the license under which the project is distributed. (The content of the LICENSE file was not inspected, but it's typically MIT, Apache, GPL, etc.)

### `.gitignore`
This is a standard Git file that specifies intentionally untracked files that Git should ignore. This typically includes build artifacts, local configuration files, and operating system-specific files.

## API Functionality

The core functionality of this application is its API for retrieving sunrise and sunset times.

### URL Structure
The API endpoints are structured as follows:
- For sunrise: `GET http://<hostname>/v0/sunrise/lat[latitude]/lon[longitude]`
- For sunset: `GET http://<hostname>/v0/sunset/lat[latitude]/lon[longitude]`

Where:
- `<hostname>` is the domain where the service is hosted (e.g., `astroapi.appspot.com`).
- `[latitude]` is the geographical latitude (e.g., `38.8895`).
- `[longitude]` is the geographical longitude (e.g., `77.0352`).

### Input
- **Latitude**: A floating-point number representing the latitude, prefixed with `lat`.
- **Longitude**: A floating-point number representing the longitude, prefixed with `lon`.

These are extracted directly from the URL path by the `sun.go` handler.

### Output
- The API returns the calculated sunrise or sunset time as a plain text string.
- The format of the time string is Go's default time format: `"2006-01-02 15:04:05.999999999 -0700 MST"`.
- If an error occurs (e.g., malformed URL, invalid parameters), an appropriate HTTP error status and message are returned.

## External Dependencies

The application relies on an external Go package for the core astronomical calculations:

- **`github.com/cpucycle/astrotime`**: This library is used by `sun.go` to calculate the precise times for sunrise and sunset based on the given date, latitude, and longitude. The `NextSunrise()` and `NextSunset()` functions from this package are central to the API's functionality.
