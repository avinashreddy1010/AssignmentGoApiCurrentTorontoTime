# AssignmentGoApiCurrentTorontoTime
# Go API for Current Toronto Time with MySQL Logging

This project provides a simple Go-based web API that returns the current time in **Toronto** and logs each request to a MySQL database. The API exposes two main endpoints:

1. **`/current-time`**: Returns the current time in Toronto and logs the timestamp in the database.
2. **`/time-logs`**: Retrieves all the logged timestamps from the database.

## Features
- Returns the current time in **Toronto** (adjusted for the correct timezone).
- Logs the timestamp to a MySQL database with every request made to the `/current-time` endpoint.
- Allows retrieval of all logged timestamps via the `/time-logs` endpoint.

## Requirements

To run this project, make sure you have the following:

- [Go (v1.18 or higher)](https://golang.org/dl/)
- [MySQL Database](https://dev.mysql.com/downloads/installer/) installed and running
- A MySQL database and table for storing the timestamps

## Setup Instructions

### 1. **Install Go**

Ensure that Go is installed on your system. If it's not, you can download and install it from the official [Go downloads page](https://golang.org/dl/).

### 2. **Install MySQL**

Download and install MySQL from [here](https://dev.mysql.com/downloads/installer/) if you don't already have it installed. Make sure the MySQL service is running before proceeding.

### 3. **Create Database and Table**

After setting up MySQL, execute the following SQL queries to create the database and table where the time logs will be stored:

```sql
CREATE DATABASE avidb;
USE avidb;

CREATE TABLE toronto_time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
);
