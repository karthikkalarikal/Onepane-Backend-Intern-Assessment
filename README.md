# Onepane-Backend-Intern-Assessment

This repository contains the solution for the OnePane REST API assignment. The goal of the assignment is to create a REST API using Golang that fetches data from three different endpoints, combines the results, and returns a formatted JSON response.

## Solution Overview

### Packages Used

The solution utilizes the following Golang packages:

- **net/http:** Handles HTTP requests and responses.
- **encoding/json:** Used for marshaling and unmarshaling JSON data.

## Logic Used

I used net/http for router, encoding/json: for marshaling JSON data, goroutines for consuming the api concurrently. The final JSON format was achieved by using for loops.

### Expected Result Format

The API is designed to return a combined result in the following JSON format:

```json
[
  {
    "postId": "",
    "postName": "",
    "commentsCount": "",
    "userName": "",
    "body": ""
  }
]