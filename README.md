# UWindsor API
An API for UWindsor data (data not included here)

## API Endpoints  

| Dataset | API Endpoint | Parameter | Description |
| ----------- | ----------- | ----------- | ----------- |
| Courses | /v1/course | Course code | Contains course code, title, and description |
| Staff | /v1/staff | Staff name | Contains name, email, phone, title, department, and location |
| Buildings | /v1/building | Building name | Contains building name and location (latitude and longitude) |

## Parameters

All endpoints have a respective parameter you can include. Parameters are included by adding them after the `/`. For example, `/v1/course/COMP-1000` will return data about COMP-1000.

## Data Source

This API expects a `db.sqlite` file in the root of the project to serve as a datasource. A sample database file is provided for development purposes, named `sample.db.sqlite`. Simply make a copy of this file and name it `db.sqlite` to get up and running.
