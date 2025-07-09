# Fallout NPC Retriever

## Overview

The Fallout NPC Retriever is a web application that functions as both a web scraper and a RESTful API, designed to gather and provide information about non-playable characters (NPCs) from the Fallout video game series (only Fallout 4 for now). It scrapes data from the Fallout fandom wiki, collecting details such as character names, locations, and special traits (like whether they are doctors or merchants).

## Features

- **Web Scraping**: Automatically collects NPC data from the Fallout fandom wiki.
- **RESTful API**: Access character information through a simple web address.
- **Structured Response**: Returns character details in a structured format (JSON).
- **Error Handling**: Provides clear error messages for invalid requests.
- **Rate Limiting**: Manages the number of requests to prevent abuse.
- **CORS Support**: Allows safe access from different web applications.

## Directory Structure

```
fallout-npc-retriever/
├── README.md
├── scrapper/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
└── webpage/
├── index.html
├── index.css
└── index.js
```

### `scrapper/`

- Contains the Go application that scrapes NPC data and serves the RESTful API.

### `webpage/`

- Contains the front-end files (HTML, CSS, and JavaScript) for displaying NPC information.

## Usage

To use the API, send a GET request to the following endpoint:

`GET /fallout-npc-scrapper/:name`


Replace `:name` with the name of the NPC you want to retrieve information about.


### Running the Scraper

1. Navigate to the `scrapper` directory:

`cd scrapper`


2. Install the required dependencies: 

`go mod tidy`


3. Run the application: 

`go run main.go`


The server will start and listen on localhost:12300.

### Frontend

To view the NPC information in a web interface, open the index.html file in your web browser. The front-end files in the webpage directory provide a user-friendly way to interact with the API.
Contributing

### Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.
License

### License

This project is licensed under the MIT License.