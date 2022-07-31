# Syndio Backend App - Juliana Zlatevski

## How to Run

- clone repo
- in your terminal, run the following commands:
  - `go mod tidy`
  - `go run main.go`
  - `command + t` to open a new terminal tab, and then `export $(cat .env)`
  - `curl localhost:$PORT/employees`
- you should receive a JSON response in your terminal
- to make the same request from your browser, go to `http://localhost:8080/employees`
- note: the PORT can be found and changed within the included .env file if needed

Using the `employees.db` sqlite database in this repository with the following table/data:

```
sqlite> .open employees.db
sqlite> .schema employees
CREATE TABLE employees (id INTEGER PRIMARY KEY, gender TEXT not null);
sqlite> SELECT * FROM employees;
1|male
2|male
3|male
4|female
5|female
6|female
```

Create an api with an endpoint `/employees` that reads from this database and returns the following JSON response (content matters not format/indent):

```
[
    {
        "gender": "male",
        "id": 1
    },
    {
        "gender": "male",
        "id": 2
    },
    {
        "gender": "male",
        "id": 3
    },
    {
        "gender": "female",
        "id": 4
    },
    {
        "gender": "female",
        "id": 5
    },
    {
        "gender": "female",
        "id": 6
    }
]
```

## Requirements

- The api must take an environment variable `PORT` and respond to requests on that port.
- You provide basic setup instructions required to run the api.
- `curl localhost:$PORT/employees` returns the described response.

## Success

- We can run the api from your setup instructions
- The curl returns the described response
- The api is written in Python or Go

## Not Required

- Tests
- Logging, monitoring, or anything more than basic error handling

## Submission

- Respond to the email you received giving you this with:
  - a zip file, or link to a git repo
  - instructions on how to setup and run the code (could be included w/ zip/git)
- We'll follow the setup instructions to test it on a local machine, then we'll get back to you.

## Notes

- Keep it simple
- We expect this to take less than an hour, please try and limit your effort to that window.
- If the api works, and returns what we requested, its a success.
- Anything extra (tests, other endpoints, ...) is not worth bonus/etc.
- We truly value your time and just want a basic benchmark and common piece of code to use in future interviews.
- If we bring you in for in-person interviews we'll expand on this submission.
