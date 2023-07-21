# Book Tracker

## Description

This is a simple book tracker app that allows you to add books to a list and mark them as read or not read. It also allows you to delete books from the list.

## Running the app

### Build the image

```bash
docker build -t book-tracker .
docker build -t book-tracker-db -f Dockerfile-postgres .
```

### Run the container

```bash
docker run --network host -d --name book-tracker-db-container -p 5432:5432 book-tracker-db
docker run --network host --name book-tracker-app-container -it book-tracker-app 
```
