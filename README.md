# Money Backward Backend

## How to Setup
1. Install dependecies
    ```
    make install
    ```
2. Setup environment:

   Copy `.env.example` to `.env` and adjust the content accordingly.

3. Setup database:

    Create a PostgreSQL database with the credentials as specified in your `.env` file.
    

## How to Run
> See Makefile for details
1. Run as developer
    ```
    make dev
    ```

2. Build and run
    ```
    make build-and-run
    ```

\* clean the built binary: ```make clean```


## Layers
> Detailed explanations can be read on the corresponding folders' `README`.
> 
Client <=(Request/Response)=> Route <=(Request/Response)=> Controller <=(DTO)=> Service <=(DTO)=> Repository <=(Models/DAO)=> Database 
1. Route: defines the routes, the HTTP method used, and `Controller`s' method that handles the request.
2. Controller: handles the request and reply with response.
3. Service: handles the input from `Controller` and do the business logic.
4. Repository: read or write to database as called by the `Service`.

:information_source: DB Schema can be seen in [here](https://dbdiagram.io/d/Money-Backward-652ac9bfffbf5169f0af735a)
