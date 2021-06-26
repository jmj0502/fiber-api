# A Simple REST API Written in Fiber
I built this whole API with the sole pourpose of experimenting with fiber, a good friend of mine told me great things about
it and I was really eager to give it a try. In general, I ended up liking it a lot since it's somehow "intuitive" and its docs are great. 

The fact that I wanted to try and use fiber in a somehow "breazy" project doesn't mean that I just wanted to build some harcoded project, so I ended up using GORM in order to stablish a connection with a MySQL db. Initially I was following an article regarding the use of the framework, but after a while I ended up deviating from it and doing things my way.

## Project Structure
The structure of the project is pretty straight forward, it is composed of three packages (**main**, **book**, **database**); each one of them contains well distribuited logic that is pretty close to what a realworld app is. Now I'll dive a little bit deeper in them:
* **main**: Contains the **main** entry point of the app, and executes the logic regarding file reading, server initialization, database connection, and request mapping.
* **book**: Contains the necessary logic to run this micro library CRUD. It contains the **Book** (gorm model) and the **Response** structs, and request handler.
* **database**: Contains the initialization of our database connection instance. This connection is shared across every other module as it allow us to perform every CRUD operation over a database each time a specific handler is triggered. 