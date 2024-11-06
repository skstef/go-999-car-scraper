# Car Listings Scraper in Go

This project is a web scraper written in Go that fetches car listings from the website [999.md](https://999.md). It scrapes data such as the car title, price, and image URL from multiple pages and stores the results in a JSON file.

## Features

- Scrapes car listings from [999.md](https://999.md).
- Extracts data like the car title, price, and image URL.
- Saves the extracted data into a JSON file named based on the current date (e.g., `scrapped_2024_11_06.json`).
- Iterates through multiple pages of car listings until no more listings are found.

## Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.16+ recommended).
- Internet access to fetch the car listing pages.

## Setup and Installation

### 1. Clone or Download the Repository

Clone this repository to your local machine or download the files directly.
```bash
git clone https://github.com/skstef/go-999-car-scraper
cd go-999-car-scraper
```

### 2. Install Dependencies
The required Python packages are listed in `requirements.txt`. Install them by running:
```bash
go run scrapper.go
```

### 3. Run the Script
After installing the dependencies, you can run the script:
```bash
python scrapper.py
```

## Example Output
Here is an example of what the JSON output might look like:

```json
[
    {
        "id": "12345678",
        "title": "Mercedes-Benz E-Class",
        "price": "15,000 €",
        "image": "https://example.com/image1.jpg"
    },
    {
        "id": "87654321",
        "title": "BMW 3 Series",
        "price": "12,500 €",
        "image": "https://example.com/image2.jpg"
    }
]
```

## License
This project is open-source and available under the [MIT License](LICENSE).