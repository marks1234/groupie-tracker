# Groupie Trackers

Groupie Trackers is a web application that utilizes a Go-based backend to retrieve and manipulate data from a given API, enabling the creation of a user-friendly website to display information about bands and artists, their concert locations, dates, and their relations. This project also focuses on implementing client-server communication through various user-triggered actions.

## Authors

- Marco Kuzmina ([Gitea](https://01.kood.tech/git/mkuzmina))
- Aburnaz ([Gitea](https://01.kood.tech/git/aburnaz))

## Features

- **Data Display**: Groupie Trackers allows users to view information about bands and artists, including their names, images, founding years, the date of their first album, and their members.

- **Concert Information**: Users can also access details about concert locations and dates for both past and upcoming events.

- **Visualization**: The application provides multiple data visualization options such as blocks, cards, tables, lists, pages, and graphics to present the information in a user-friendly manner.

- **Client-Server Communication**: Groupie Trackers implements client-server communication through events and actions, enabling users to trigger actions that request information from the server and receive responses.

## Backend

The backend of Groupie Trackers is written in Go, utilizing only standard Go packages. It ensures robust and error-free operation, preventing crashes, and adheres to best practices in coding. The backend also includes unit tests to verify the correctness of the code.

## Usage

1. Clone the Groupie Trackers repository to your local machine.
2. Set up the Go environment if not already done.
3. Build and run the backend server.
4. Access the web application through a web browser.

## Installation

To run Groupie Trackers, follow these steps:

1. Clone the repository:

```bash
git clone https://01.kood.tech/git/mkuzmina/groupie-tracker.git
```

2. Navigate to the project directory:

```bash
cd groupie-trackers
```

3. Run:

```bash
go run .
```


5. Open your web browser and access the Groupie Trackers website at `http://localhost:8080`.

## API

The backend of Groupie Trackers consumes a RESTful API to fetch data about bands, artists, concert locations, and dates. The API consists of the following endpoints:

- `/api/artists`: Retrieve information about bands and artists.
- `/api/locations`: Get details about concert locations.
- `/api/dates`: Access information about concert dates.
- `/api/relation`: Retrieve relations between artists, locations, and dates.

