# Solutions

## Basic Concepts

1. Coding
    - A common example of code smell is a large, bloated class that usually has more than one purpose as it violates the first SOLID principle, which states that a class should have only one job. To prevent such bloated class is by keeping the class small and use a short and clear class name that describes the intent of such class.
    - Dependency Injection (DI) is a technique where an object (client) delegates the creation of its dependencies (services) to another object (injector). DI is an instrument of the fifth SOLID principle, the dependency inversion principle, which states that entities must depend on abstractions, not on concretions. It reduces boilerplates of creating services, enables changing services at runtime, and makes extending the application easier.
2. REST API
    - POST
        - Do's:
            - A response to POST request should return HTTP status 201 and a 'Location' header with a link to the newly created resource on successful resource creation.
        - Don'ts:
            - POST requests should not be used for non-idempotent requests, e.g., for 'update' resource requests as specified in the HTTP specifications.
    - GET
        - Do's:
            - GET requests should be cacheable: if other requests have yet to modify the specified resource, the GET request itself should be cacheable, saving the server a new request.
        - Don'ts:
            - GET requests should not modify the server state: HTTP methods except for POST should be idempotent, which means if you make a request several times, they should have the same effect, i.e., they should not change the state of the server.

## Basic Coding

Imagine you’re working in a company and you’re told to make a design system and transaction flow. How would you make the most suitable for following users and case:

1. User Side
    - User register
    - User input the address
    - User choose the products to be purchased with subscription and/or one-time purchase scheme
    - User pay the bill
    - User skip the delivery due to certain reasons (ex: They have other agenda that prevent them to receive the delivered goods)
    - User cancel the order
2. Supplier Side
    - Supplier register as seller
    - Supplier create the store and complete the address
    - Supplier create products that can be purchased either daily or one-time purchase
    - Supplier determine the price of each product
    - Supplier determine the selling area
3. Additional:
    - If a product can be sold by more than one seller, define the correct algorithm to determine which order to be sent from which seller (assuming the closest mileage and route)!
    - There is a cut-off time everyday, which is the latest time an order can be placed for the next day delivery. All orders placed beyond cut-off time will automatically delivered on the day after tomorrow.

## Algorithm

1. Please see [sock_merchant.go](cmd/sock_merchant/sock_merchant.go).

2. Please see [counting_valleys.go](cmd/counting_valleys/counting_valleys.go).

3. Please see [print_digits.go](cmd/print_digits/print_digits.go).

4. Please see [lamp_switch.go](cmd/lamp_switch/lamp_switch.go).
