# Mphasis
## "go run app.go" to start the server

## End Point(Image Get) - /image
### Request Method - GET
### Query Parameters - id=1&page=1&size=10&sort=image_name&order=asc
### Response - 
`{
    "result": [
        {
            "image_id": 2,
            "image_name": "img1.png",
            "created_at": "2020-12-14T23:57:56.892302Z"
        },
        {
            "image_id": 1,
            "image_name": "img2.jpg",
            "created_at": "2020-12-14T23:44:22.690508Z"
        }
    ]
}`

## End Point(Upload Get) - /image
### Request Method - POST
### Request Body - form-data with "file" key and image as a value
### Response - 
`{
    "message": "image uploaded successfully"
}`
