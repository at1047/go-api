## A backend API for my recipes page, built with a MongoDB Collection and a GO application

### Technologies

I wanted to try out a new programming language, GO, and I read that GO is widely used for backend applications and building APIs. Previously I've build a variety of APIs using both Python (specifically serving different content through Flask), and JavaScript. In addition, I've also mainly used relational databases like PostgreSQL, so I wanted to explore NOSQL in this project. I chose MongoDB because I've interacted with a MongoDB database once during a previous internship, and wanted to explore it further.

Creating the MongoDB database could not be easier. MongoDB Atlas offers a free tier that has built in functionalities for network access, and a simple UI for the basic CRUD functions. There was also a tutorial for connecting to the MongoDB database with GO through MongoDB's own library.

To implement routing, I used GO's Gin library. This allows me to call different functions depending on the url of the request and the type of request. I implemented the basic "Read", "Read all", "Create one" functionalities and soon enough I had a working backend API that I could deploy to my virtual machine. I tested the functionality of this API with Postman, a tool I've used extensively at my previous internships, and learnt how to write an OpenAPI document describing my API.

### Deployment

Build this app locally because the vm does not have enough ram to build remotely.

    # Build command, specifiying the output to be amd64 linux since I'm running an ARM Mac
    #   and the resulting output is incompatible
    GOOS=linux GOARCH=amd64 go build -o gin

    # Run application and detatch to keep it running in the background
    nohup ./gin & disown 

    # Find the application code for killing
    ps -ef | grep ./gin

    # Kill the application
    kill -9 your_pid


