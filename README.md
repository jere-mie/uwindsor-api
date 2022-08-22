# UWindsor API
An API for UWindsor data (data not included here)

*Note*: For this API to work, you must have a "db.sqlite" file in the root of the project with all of the data in it. If you don't have access to such a file, you probably shouldn't be working on this application...

## API Endpoints  

| Dataset | API Endpoint | Parameter | Description |
| ----------- | ----------- | ----------- | ----------- |
| Courses | /v1/course | Course code | Contains course code, title, and description |
| Staff | /v1/staff | Staff name | Contains name, email, phone, title, department, and location |
| Buildings | /v1/building | Building name | Contains building name and location (latitude and longitude) |

## Parameters

All endpoints have a respective parameter you can include. Parameters are included by adding them after the `/`. For example, `/v1/course/COMP-1000` will return data about COMP-1000.
