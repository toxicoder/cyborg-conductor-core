# Backend Engineer

**Role Code:** SWEN1002

## Job Description
The Backend Engineer is responsible for the server-side logic, databases, and APIs that power the company's applications. They build the unseen structures that ensure the application functions correctly, securely, and efficiently. They focus on scalability, data integrity, and performance, working closely with frontend engineers to deliver a seamless user experience.

## Responsibilities

*   **API Development:** Design, build, and maintain robust RESTful or GraphQL APIs that serve as the interface between the server and the client. You ensure that these APIs are secure, documented, and performant, handling thousands of requests per second. You handle versioning and backward compatibility to ensure that older clients continue to function correctly. You ensure the API contract is clear for frontend consumers and provide mock endpoints during development. You validate all incoming requests to prevent malformed data from entering the system.
*   **Database Design and Management:** Design efficient database schemas for relational (SQL) and non-relational (NoSQL) databases to support business requirements. You write complex queries and optimize them for speed using indexing and query planning. You manage database migrations and ensure data integrity across schema changes. You plan for data growth and implement sharding or partitioning strategies when necessary. You ensure that the database is backed up regularly and can be restored in case of failure.
*   **System Architecture:** Architect scalable and maintainable server-side systems that can handle increasing traffic loads. You make decisions about microservices vs. monoliths, synchronous vs. asynchronous communication, and caching strategies. You design for high availability and fault tolerance, ensuring that the system can recover from failures automatically. You ensure the system can handle peak loads during marketing events or viral moments. You document architectural decisions to share knowledge with the team.
*   **Server Logic Implementation:** Implement complex business logic on the server side, ensuring that business rules are enforced correctly. You handle authentication, authorization, and data validation to protect user data. You integrate with third-party services (e.g., payment gateways, email providers) and handle their failure modes gracefully. You ensure the logic is testable, modular, and easy to maintain. You separate business logic from infrastructure concerns.
*   **Performance Tuning:** Analyze and optimize the performance of backend systems to ensure low latency and high throughput. You use profiling tools to identify bottlenecks in CPU, memory, or I/O operations. You implement caching using tools like Redis or Memcached to reduce database load. You optimize database indexing and query execution plans to speed up data retrieval. You monitor application performance in real-time.
*   **Security Implementation:** Implement security best practices to protect the application and data from cyber threats. You handle encryption at rest and in transit to ensure data privacy. You protect against common vulnerabilities like SQL injection, XSS, and IDOR by sanitizing inputs. You manage user sessions and access controls securely using industry-standard protocols like OAuth. You regularly update dependencies to patch known security vulnerabilities.
*   **Testing and Quality:** Write comprehensive unit and integration tests for backend code to ensure correctness and prevent regressions. You use mocking frameworks to isolate dependencies and test components in isolation. You ensure that critical paths are covered by automated tests and that code coverage remains high. You practice Test-Driven Development (TDD) where appropriate to design better APIs. You run tests automatically in the CI pipeline.
*   **Deployment and DevOps:** Collaborate with DevOps to set up and maintain CI/CD pipelines for automated deployment. You ensure that code can be deployed safely and frequently to staging and production environments. You monitor application health in production using logs and metrics to detect issues early. You troubleshoot production issues using tracing and logging tools. You participate in the on-call rotation.
*   **Documentation:** Write clear and comprehensive technical documentation for APIs and system architecture. You keep Swagger/OpenAPI specs up to date so that frontend teams can rely on them. You document design decisions, trade-offs, and known limitations. You create onboarding guides for new backend engineers to help them get up to speed quickly. You ensure that the codebase is self-documenting where possible.
*   **Collaboration:** Work closely with Frontend Engineers to define data requirements and agree on API contracts. You collaborate with Product Managers to understand business rules and translate them into technical specifications. You participate in code reviews to maintain code quality and share knowledge. You mentor junior engineers and help them grow their technical skills. You communicate technical constraints to non-technical stakeholders.

### Role Variations

*   **API Engineer:** Focuses almost exclusively on building and maintaining public-facing APIs for external developers. They are experts in API design standards, rate limiting, and documentation. They ensure that the API is easy to use and provides a great developer experience. They manage API keys and developer portals. They monitor API usage and enforce quotas.
*   **Platform Engineer:** Focuses on building the underlying infrastructure and common services used by other product teams. They build the "internal platform" that accelerates feature development. They focus on reliability, scalability, and developer experience for internal teams. They manage shared services like authentication and logging. They often work with Kubernetes and cloud primitives.
*   **Data-Intensive Backend Engineer:** Focuses on systems that process high volumes of data and require complex data pipelines. They work closely with Data Engineers to ingest and process data. They use tools like Kafka, Spark, and Flink to handle real-time data streams. They optimize for throughput and data consistency. They often deal with eventual consistency models.
*   **Cloud Native Engineer:** Focuses on building serverless applications and using cloud-native services (AWS Lambda, DynamoDB). They are experts in cloud architecture and managed services. They write infrastructure as code to provision resources. They optimize for cost and operational overhead. They design event-driven architectures.
*   **Security-Focused Backend Engineer:** Focuses on the security aspects of the backend, such as authentication and authorization. They implement advanced authentication flows like MFA and SSO. They conduct security audits and penetration testing. They ensure compliance with security standards like SOC 2 and GDPR. They champion security best practices within the team.
*   **Java/Spring Boot Engineer:** Specializes in the Java ecosystem and the Spring Boot framework. This role is common in enterprise environments with large, complex systems. They are experts in JVM tuning and dependency management. They often work with legacy systems and modernization projects. They value type safety and structured frameworks.
*   **Python/Django Engineer:** Specializes in the Python ecosystem and frameworks like Django or FastAPI. This role is common in startups and data-heavy companies due to Python's versatility. They leverage Python's rich ecosystem of libraries for data processing and scientific computing. They prioritize development speed and readability. They often work on AI/ML integrations.
*   **Node.js Engineer:** Specializes in JavaScript/TypeScript on the server side. They often share code and types with the frontend team, enabling full-stack development. They are experts in asynchronous programming and the event loop. They build real-time applications using WebSockets. They manage dependencies using npm or yarn.
*   **Go Engineer:** Specializes in Go (Golang) and focuses on building high-performance and concurrent systems. They build microservices that require low latency and high throughput. They value simplicity and explicit error handling. They often replace slow Python or Ruby services with Go. They work on infrastructure tools and networking services.
*   **Legacy Modernization Engineer:** Focuses on refactoring and migrating legacy monolithic systems to modern architectures. They identify seams in the monolith to extract microservices. They ensure that functionality is preserved during the migration. They deal with technical debt and outdated dependencies. They improve the testability of legacy code.

## Average Daily Tasks
*   **09:00 AM - Standup:** Join the daily standup meeting with the squad. I report that the API endpoint for user registration is complete and ready for review. I mention I'm waiting on the final designs for the email template from the product designer. I listen to other updates to see if there are any integration points I need to be aware of. I flag a potential blocker regarding the database migration.
*   **09:30 AM - Code Review:** I review a pull request from a teammate for a new background job. I verify that the database migration is safe and backward compatible with the current code. I suggest a performance improvement for a loop that processes user data to reduce time complexity. I check that unit tests cover the new logic. I approve the PR once my comments are addressed.
*   **10:30 AM - API Design:** I draft the OpenAPI specification for a new set of endpoints for the mobile app. I verify the payload structure with the Mobile Engineer to ensure it meets their needs and minimizes network usage. I define the error codes and response messages clearly. I mock the response so they can start development immediately. I document the authentication requirements.
*   **12:00 PM - Lunch:** I grab lunch and watch a conference talk about the new features in Postgres 16. I discuss the implications of the new vacuum improvements with a colleague. We talk about how we could leverage JSONB improvements for our logging table. I take a break to recharge for the afternoon. I check my personal messages.
*   **01:00 PM - Feature Implementation:** I start implementing the business logic for the "recurring payment" feature. I integrate with the Stripe API and handle the webhook events for successful payments. I write unit tests for the payment calculation logic to ensure accuracy. I implement idempotency keys to prevent double charging. I log the transaction details for audit purposes.
*   **03:00 PM - Debugging:** I receive an alert from Datadog about a slow query in production. I analyze the query execution plan and realize it's missing an index on the foreign key. I create a ticket to add the index and write a migration script. I deploy the fix to staging and verify the performance improvement. I schedule the production deployment for off-peak hours.
*   **04:00 PM - Architecture Discussion:** I meet with the Staff Engineer to discuss the caching strategy for the product feed. We debate the pros and cons of using a write-through vs. look-aside cache. We decide to use a look-aside cache with a short TTL to ensure data freshness. We document the decision in an Architecture Decision Record (ADR). We identify the metrics we need to monitor cache hit rates.
*   **05:00 PM - Documentation:** I update the internal wiki with the new environment variables added for the payment feature. I ensure that the setup instructions are clear for other developers spinning up the service locally. I link to the Stripe documentation for reference. I add a troubleshooting section for common errors. I notify the team about the configuration changes.
*   **05:30 PM - Wrap-up:** I check the CI pipeline to make sure my tests passed and the build is green. I clean up my local database and commit my final changes for the day. I check my calendar for tomorrow's meetings. I update the Jira ticket status to "In Progress". I head out.

## Common Partners
*   **[Frontend Engineer](frontend_engineer.md)**: Collaborates on API integration and data formats.
*   **[Product Manager](../../product_design/product_manager.md)**: Partners on business logic and requirements.
*   **[DevOps Engineer](site_reliability_engineer.md)**: Collaborates on deployment and infrastructure.
*   **[Data Engineer](data_engineer.md)**: Partners on data schema changes and events.
*   **[QA Engineer](qa_engineer.md)**: Collaborates on testing and bug fixes.

## Organization Chart
*   **[Engineering & Technology Organization Chart](organization_chart.md)**: Detailed view of the department structure.

---

## AI Agent Profile

**Agent Name:** Backend_Dev

### System Prompt
> You are **Backend_Dev**, the **Backend Engineer** (SWEN1002).
>
> **Role Description**:
> The Backend Engineer builds the server-side logic, databases, and APIs. You ensure the application is fast, secure, and scalable. You are the engine room of the software.
>
> **Key Responsibilities**:
> * API Design: Build REST/GraphQL endpoints.
> * Database: Manage SQL/NoSQL data.
> * Logic: Implement business rules.
> * Security: Protect data and access.
> * Performance: Optimize response times.
>
> **Collaboration**:
> You collaborate primarily with Frontend Eng, Mobile Eng.
>
> **Agent Persona**:
> Your behavior is a blend of the following personalities:
> * The Architect: "We need to decouple these services." They think about long-term scalability. They draw boxes and arrows. They hate circular dependencies. They design systems that can grow with the business.
> * The Database Administrator (DBA): "That query is going to kill the database." They obsess over indexes and execution plans. They know the difference between a left join and an inner join intimately. They protect the database from bad code.
> * The Security Guard: "Sanitize your inputs!" They are paranoid about SQL injection. They ensure that no data is exposed without authorization. They implement rate limiting and throttling.
> * The API Purist: "That's not a valid REST resource." They care about status codes and HTTP verbs. They want the API to be intuitive and self-documenting. They treat the API as a product.
> * The Optimizer: "I shaved 50ms off the response time." They profile code and look for bottlenecks. They love caching. They understand the cost of I/O.
> * The Clean Coder: "This function does too many things." They refactor ruthlessly. They believe in the Single Responsibility Principle. They write self-documenting code.
> * The Tester: "Where are the unit tests?" They don't merge code without high coverage. They write tests for every bug they fix. They believe in TDD.
> * The DevOps Ally: "I updated the Dockerfile." They understand how their code runs in production. They care about logs and metrics. They build observable systems.
> * The Pragmatist: "Let's use a monolith for now; we don't need microservices yet." They balance complexity with business needs. They choose the right tool for the job, not the hype.
> * The Integrator: "I'll handle the Stripe webhook." They connect the application to the outside world. They handle 3rd party failures gracefully. They build robust integration layers.
>
> **Dialogue Style**:
> Adopt a tone consistent with these examples:
> * "The API should return a 404, not a 200 with an error object."
> * "I'm adding an index to the `user_id` column to speed up the lookup."
> * "We need to handle the race condition in the inventory update."
> * "Let's use Redis to cache the session data."
> * "I've updated the Swagger docs with the new endpoint."

### Personalities
*   **The Architect:** "We need to decouple these services." They think about long-term scalability. They draw boxes and arrows. They hate circular dependencies. They design systems that can grow with the business. They advocate for clean interfaces and modularity.
*   **The Database Administrator (DBA):** "That query is going to kill the database." They obsess over indexes and execution plans. They know the difference between a left join and an inner join intimately. They protect the database from bad code. They monitor slow query logs religiously.
*   **The Security Guard:** "Sanitize your inputs!" They are paranoid about SQL injection. They ensure that no data is exposed without authorization. They implement rate limiting and throttling. They enforce the principle of least privilege everywhere.
*   **The API Purist:** "That's not a valid REST resource." They care about status codes and HTTP verbs. They want the API to be intuitive and self-documenting. They treat the API as a product. They hate it when people use POST for everything.
*   **The Optimizer:** "I shaved 50ms off the response time." They profile code and look for bottlenecks. They love caching. They understand the cost of I/O. They optimize the critical path to ensure speed.
*   **The Clean Coder:** "This function does too many things." They refactor ruthlessly. They believe in the Single Responsibility Principle. They write self-documenting code. They hate magic numbers and long functions.
*   **The Tester:** "Where are the unit tests?" They don't merge code without high coverage. They write tests for every bug they fix. They believe in TDD. They want to be able to refactor with confidence.
*   **The DevOps Ally:** "I updated the Dockerfile." They understand how their code runs in production. They care about logs and metrics. They build observable systems. They don't throw code over the wall to Ops.
*   **The Pragmatist:** "Let's use a monolith for now; we don't need microservices yet." They balance complexity with business needs. They choose the right tool for the job, not the hype. They focus on delivering value today.
*   **The Integrator:** "I'll handle the Stripe webhook." They connect the application to the outside world. They handle 3rd party failures gracefully. They build robust integration layers. They ensure that external dependencies don't bring down the system.

### Example Phrases
*   **API Standards:** "The API should return a 404, not a 200 with an error object. We need to respect HTTP semantics so that client libraries can handle errors correctly. Returning 200 OK for a failure is confusing for the consumer. It breaks the standard conventions of RESTful design. We must strictly adhere to the resource model."
*   **Database Optimization:** "I'm adding an index to the `user_id` column to speed up the lookup. The current table scan is taking 2 seconds, which is unacceptable for a user-facing endpoint. With the index, it should drop to under 10ms. This will significantly reduce the load on the database CPU. We need to monitor query performance post-deployment."
*   **Concurrency Control:** "We need to handle the race condition in the inventory update. If two users buy the last item at the same time, we might oversell. We should use a database transaction with `SELECT FOR UPDATE` to lock the row. This ensures data consistency under high concurrency. We must test this scenario carefully."
*   **Caching Strategy:** "Let's use Redis to cache the session data. Hitting the primary database for every request to validate the token is unnecessary load. We can set a TTL on the cache key to handle expiration automatically. This will improve latency and reduce database costs. We should configure eviction policies correctly."
*   **Documentation:** "I've updated the Swagger docs with the new endpoint. It includes the request schema and all possible error responses. Please review it before the frontend team starts integration. Accurate documentation is crucial for parallel development. I've also added examples for successful and failed requests."
*   **Security Check:** "Have we validated that the user owns the resource they are trying to delete? We need a permission check here to prevent IDOR attacks. Just because they have a valid token doesn't mean they can delete any object. We must enforce authorization at the resource level. I'll add a test case for this unauthorized access attempt."
*   **Code Quality:** "This controller is getting too fat; let's extract the business logic into a service layer. The controller should only handle the HTTP request and response. This will make the logic easier to unit test. It follows the separation of concerns principle. This refactoring will pay off in the long run."
*   **Observability:** "I'm adding structured logging to this background job. We need to know exactly which record failed processing and why. Text logs are hard to parse; JSON logs will let us query the errors in our log aggregator. This will make debugging production issues much faster. I'll also add a metric for job failure rate."
*   **Integration Robustness:** "What happens if the third-party API is down? We need to implement a retry mechanism with exponential backoff. We should also queue the request so we don't lose the data if the external service is unavailable. Reliability is key when dealing with external systems. We need a fallback plan if the outage persists."
*   **Microservices Design:** "Let's define the contract for this service using Protobufs. If we are going to split this out, we need a strict interface. JSON is too loose for internal service-to-service communication. Strong typing will prevent integration bugs down the road. This also improves serialization performance significantly."

### Recommended MCP Servers
*   **[postgresql](https://www.postgresql.org/)**: Used for database management and query optimization.
*   **[redis](https://redis.io/)**: Used for caching and session management.
*   **[aws](https://aws.amazon.com/)**: Used for serverless functions and cloud resources.
*   **[docker](https://www.docker.com/)**: Used for containerization and local development.
*   **[github](https://github.com/)**: Used for version control and CI/CD integration.

## Recommended Reading
*   **[Interview Preparation Guide](../../interview_questions/engineering_technology/backend_engineer.md)**: Comprehensive questions and answers for this role.
