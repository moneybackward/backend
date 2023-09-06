## Motivation
To better documentize the API developed, we will use OpenAPI/Swagger.

## Steps
We will use [rswag](https://github.com/rswag/rswag)
1. Add this line to Gemfile
    ```
    gem 'rswag'
    ```
    Also add this line due to error 'rails_helper not found' [solution](https://github.com/rswag/rswag/issues/456#issuecomment-978180827)
    ```
    rails g rswag:install
    ```
2. Run the install generators.
    ```
    rails g rspec:install
    ```
    ```
    rails g rswag:install
    ```

3. Create integration spec
    ```
    rails generate rspec:swagger API::MyController
    ```

4. Generate swagger JSON file(s)
    ```
    rake rswag:specs:swaggerize
    ```
5. Run the server and go to the url (default: `/api-docs`)
