# Project Description
NexaSatck Backend



# Technologies Used<h2>

* Flask - Python micro webframework.


# Installation
## 
* create virtual environment using virtualenv 
* activate virtual environment
* source venv/bin/activate
* pip install -r requirements.txt

# Next Step is to run py file by this command

* python wsgi.py

# Environment variable

* FLASK_CONFIG = devel/prod/stag


## Already Present Demo User 
* id = 1, email = 'user1@gmail.com', password = 'abcxyz',role = 'admin'
* id = 2, email = 'user2@gmail.com', password = 'abcxyz',role = 'user'

# Endpoints<h2>

## Login and Authentication Routes:
* /auth -  for genrating tokens

###### Command Used:
* ```curl -X POST -H "Content-Type: application/json" http://172.16.0.13:5000/auth -d '{"username": "user1@gmail.com","password":"abcxyz"}'```

###### Expected Output:
* ``` {"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6MSwiaWF0IjoxNDgwNTAzODgzLCJuYmYiOjE0ODA1MDM4ODMsImV4cCI6MTQ4MDUwNDE4M30.4L7Ea-PCsgyaDMO6-Cl3qM29bjzlKTQWgMwdkCaslwE"
}```

## Protected Route:
* /protected - access the route with token only

###### Command Used:
* ```curl -H "Authorization:JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6MSwiaWF0IjoxNDgwNTAzODgzLCJuYmYiOjE0ODA1MDM4ODMsImV4cCI6MTQ4MDUwNDE4M30.4L7Ea-PCsgyaDMO6-Cl3qM29bjzlKTQWgMwdkCaslwE" http://172.16.0.13:5000/protected```

###### Expected Output:
* ```User(id='1')```

## Four Collections And Respective End Points: 
* Projects - /v1/projects
* Nodes - /v1/nodes
* Creds - /v1/creds
* Apps - /v1/apps

## Curl Format Used To Send Request To Server example with one route

### For Post -

###### Command Used:
* ```curl -X POST -H "Content-Type: application/json" -H "project:sharma" -d '{}' http://172.16.0.13:5000/v1/apps```

###### Expected Output:
* if success : post correctly
* else : {"code":0,"message":""}

### For Get-

###### Command Used:
* ```curl -X GET --header 'Accept: application/json' '172.16.0.13:5000/v1/apps?_id=new_app'```
* or ```curl -H "Content-Type: application/json" -H "project:tine" http://172.16.0.13:5000/v1/apps```

###### Expected Output:
* if success: [{"project_id": ["sharma"], "_id": "new_app"}]

### For DELETE -

###### Command Used:
* ```curl -X DELETE --header 'Accept: application/json' '172.16.0.13:5000/v1/projects/arvind'```
* if success: document deleted


### Local : http://0.0.0.0:5000/
### production : http://0.0.0.0:5000/

no connectiving with database for signup only demo users present#1