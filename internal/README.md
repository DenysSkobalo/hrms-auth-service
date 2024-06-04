# internal/
The heart of our application — we store all the internal logic of the application here. We do not import /internal into other applications and libraries. The code written here is intended solely for internal use within the codebase.

In /internal, we store the business logic of the project and work with databases. In general, all logic related to this application. The structure within /internal can be organized in different ways, depending on the architecture. However, I will not delve deeply into it, but rather, I will show a superficial view of how it looks. I will provide an example of a three-tier architecture where the application is divided into three layers:

#### Transport Layer:
Responsible for routing requests, handling HTTP requests/responses.
This layer communicates with the business layer.

#### Business Layer:
Contains the core logic of the application, data processing rules.
Interacts with the transport layer to receive requests and send responses.
Interacts with the database layer for storing and retrieving data.

#### Database Layer:

Responsible for database access.
Interacts only with the business layer to perform data operations.
The logic should be structured so that the layers hierarchically refer to each other from top to bottom and vice versa. It is not allowed for one layer to skip over an intermediate layer (for example, the transport layer directly accessing the database), and the lower layer should not communicate with the upper layer (for example, the database accessing the transport layer).


## Analysis internal/:
*/app:*

The point where all our dependencies and logic are assembled to run the application.
Contains a Run method that is called from the /cmd directory.

*/config:*

Contains the initialization of general application configurations that we specified at the root of the project.

*/database (database layer):*

Files containing methods for interacting with databases.

*/models (database layer):*

Structures representing database tables.

*/services (business layer):*

All the business logic of the application.

*/transport (transport layer):*

Here we store server HTTP settings, handlers, ports, etc