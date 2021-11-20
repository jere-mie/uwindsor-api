# UWindsor API
An API for UWindsor data (data not included here)

#### Note: For this API to work, there must be a 'datasets' folder in the 'static' folder, containing all of the necessary sub-folders and json files listed in the `keys` dictionary in `app.py`.

### API Endpoints  

| Dataset | API Endpoint | Parameter | Description |
| ----------- | ----------- | ----------- | ----------- |
| Courses | /api/courses | Course code | Contains course code, title, and description |
| Staff | /api/staff | Staff name | Contains name, email, phone, title, department, and location |
| Buildings | /api/buildings | Building name | Contains building name and location (latitude and longitude) |
| Holidays | /api/holidays | Date (in format YYYY-MM-DD) | Contains date and what the holiday/date is |
| Emergency Services | /api/emergency_services | Service name (ex. Campus Police) | Contains phone, extension, service name, and website |
