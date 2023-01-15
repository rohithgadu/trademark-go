# Trademark APIs

A Restful API to search Trademark Details using Application Number or Trademark Name.

The data is taken from this link:

  https://drive.google.com/drive/folders/1SUStBxmW4-0qVqW7Zk3cGHyqHBjNnG-f

Deployed link:

  https://trademark-go.fly.dev/

Postman-Docs:

https://documenter.getpostman.com/view/19762721/2s8ZDSd62N


## Tech Stack

- `Backend Framework:` `GoLang`
- `Databases:` `Postgresql`
- `Libraries` `Fiber, GORM`

## API Reference 

#### Search by Application Number
```javascript
  GET https://trademark-go.fly.dev/api/number/300000105812500
```
#### Example Response
```javascript
  200 OK

```

```json 
{
  "trademark_name": "PPG INDUSTRIES OHIO, INC.",
  "application_number": 300000105812500,
  "status": "Registered",
  "class": "Paints, varnishes and lacquers",
  "class_description": "Commercial and industrial exterior paints for signage.",
  "registration_number": "TMA593169",
  "registration_date": "2003-10-27",
  "renewal_date": "2018-10-27",
  "abandoned_date": ""
}
```

#### For all API endpoints please refer the Postman Docs 
