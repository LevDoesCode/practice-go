Write an API to retrieve customer by ID

Acceptance criteria
  - The URL should only accept numeric ids
  - API should return customer as a JSON object
  - In case the customer ID does not exist, API should return http status code 404 (Not Found)
  - In case of an unexpected error, API should return status code 500 (Internal Server Error) along with the error message