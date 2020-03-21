# System Design

Documentation of the system design

## Internal Design

### Server 

Keeps constantly checking for new updates and when there are, emails all clients
Keeps only a list of sources with email ids to them
Expose a subscribe API for clients to push to add email ids

#### Types and Data Structures

- Subscriber
	- Type alias for string
	- Stores the email address of the subscriber
- Source
	- Struct
		- string field Name
		- string field URL
	- Stores the name and URL of the homepage of the sources
- Article
	- Struct
		- string field Title
		- Source field Publisher
		- string field URL
	- Stores the title and URL of the article along with its source
- Subscribers
	- Map
		- key: Source
		- value: Subscriber
	- Mapping of Sources to Subscribers so that each subscriber can receive info from their desired sources

#### Functions

- addSubscriber
- removeSubscriber
- updateSubscribers
- notify

### Client



