
# Final project golang cohort 3 (Basic trade API)

A repository of my final test assignment in iSwift golang-cohort-3 class.

## Requirements

You need install these things before you can run this project

| Installation | Link     | 
| :----------- | :------- | 
| `Golang`     | [Here](https://go.dev/dl/) |
| `MySQL`      | [Here](https://dev.mysql.com/downloads/) |

Dont forget to register in Cloudinary service because this project using Cloudinary for some feature

| Link     |
| :------- |
| [Here](https://cloudinary.com/) |

## How To Run this project locally

Clone this repository into your local PC

```bash
  git clone https://github.com/LordAvorio/basictrade-gading.git
```

Creating the environment file on the application folder

```bash
DB_HOST="" # INSERT YOUR SERVER HOST
DB_PORT="" # INSERT YOUR SERVER PORT
DB_USER="" # INSERT YOUR DATABASE USER
DB_PASS="" # INSERT YOUR DATABASE PASSWORD
DB_NAME="" # INSERT YOUR DATABASE NAME

# CONFIG APP
APP_PORT="" # INSERT YOUR PORT APPLICATION

# SALT
JWT_SALT="" # INSERT RANDOM STRING FOR SALT PASSWORD
PASSWORD_SALT="" # INSERT RANDOM STRING FOR SALT PASSWORD

# FLAG
AUTO_MIGRATE=false # CHANGE TRUE IF YOU WANT RUN MIGRATION DB

# CLOUDINARY (Please see the API documentation for the infomration)
CLOUDINARY_CLOUD_NAME=""
CLOUDINARY_API_KEY=""
CLOUDINARY_API_SECRET=""
CLOUDINARY_UPLOAD_FOLDER=""
```

Run the main.go file

```bash
  go run main.go
```

