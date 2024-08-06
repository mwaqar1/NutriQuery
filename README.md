# NutriSearch - A Nutritional Product Evaluation System

## Project Description

**NutriSearch** is a nutritional product evaluation application leveraging the Open Food Facts API to provide users with detailed insights into food products. The application is designed to analyze the nutritional content, ingredient composition, and overall healthiness of food items by scanning their barcodes or searching by product names. By accessing a vast database of food products, NutriSearch aims to promote healthier consumer choices by identifying potentially harmful ingredients and presenting NutriScore evaluations.

## Key Features

- **Product Lookup:** Allows users to search for food products, retrieving detailed information about ingredients, nutritional values, and NutriScore grades.

- **NutriScore Evaluation:** Provides a detailed description of the NutriScore, which ranks products from A (excellent choice) to E (unhealthy choice) based on their nutritional content.

- **Ingredient Analysis:** Utilizes a Trie data structure to identify potentially harmful ingredients, providing warnings to users when such ingredients are present.

- **User-Friendly Output:** Displays product information, including ingredient lists, health warnings, and NutriScore descriptions, in a user-friendly format.

## Technical Skills Used

- **Go Programming:** Developed a robust backend using Go, focusing on concurrency, efficient data handling, and error management.

- **API Integration:** Integrated the Open Food Facts API to access a comprehensive database of food products, handling complex JSON responses and API interactions.

- **Data Structures:** Implemented a Trie data structure to efficiently search and identify harmful ingredients within products.

- **JSON Handling:** Managed complex JSON data structures to parse and process product information, leveraging Go's `encoding/json` package.

- **Web Development:** Utilized Gorilla Mux for routing and building RESTful endpoints, enabling dynamic search capabilities and seamless user interactions.

## Technical Challenges

- **Unstructured Data:** The data from the Open Food Facts API is highly unstructured, with fields that can appear as different types across products, such as strings, objects, or integers. This required a dynamic approach to data handling.

- **Type Inconsistencies:** Many fields expected to be strings often appeared as structs or integers. This necessitated changing fields to type `any` to accommodate the variations, which led to runtime issues and the need for custom type conversion functions. A specific function, `convertingIngredientsTextToString`, was developed to convert these fields into strings to facilitate substring searches.

- **Huge Database:** The Open Food Facts database contains approximately 3 million products. Managing such a vast dataset required careful consideration of data structures and performance optimization.

- **Failed Optimization with HashMap:** Initially, a hashmap was considered to optimize data retrieval. However, due to the complex and inconsistent nature of the data, including varying ingredient formats and chemical names, exact matches were impractical. The hashmap's inability to support substring matches highlighted its limitations for this application.

- **Data Complexity and Sanitization:** The database's complexity, funded by the EU, posed challenges in data sanitization. The lack of access to sanitized, uniform data made it difficult to ensure consistent processing, especially with products and ingredients having different names and formats.

- **Language Variability:** Although the application queried for English ingredients, responses frequently included products in other languages, such as French or Spanish. This necessitated additional logic to filter and handle multilingual data effectively.

## Learning Outcomes

### Advanced Go Programming Skills

- **Concurrency:** Leveraged Go's goroutines and channels to handle concurrent data processing, improving performance and responsiveness in handling large datasets. This experience enhanced understanding of Go's concurrency model, enabling efficient multi-threaded programming.

- **Data Structures:** Implemented complex data structures like Tries for substring search capabilities and improved data handling. This required a deep understanding of Go's memory management and efficient use of slices and maps.

- **Error Handling:** Adopted Go's error-handling paradigm to create robust error-checking mechanisms. This involved using custom error types and wrapping errors with context for better debugging and maintenance.

- **Interface and Struct Design:** Mastered the design of flexible and reusable interfaces and structs to accommodate the unstructured data returned by the API. This skill improved the application's extensibility and adaptability to changes.

### API Design and Consumption

- **Integration with External APIs:** Developed expertise in integrating with third-party APIs, understanding how to navigate complex API documentation and implement features such as authentication, pagination, and rate limiting.

- **Handling Complex JSON Structures:** Learned to parse and handle complex, nested JSON structures returned by the Open Food Facts API. This involved using Go's `encoding/json` package to unmarshal data into Go structs and manage dynamic fields with the `any` type.

- **Optimizing API Calls:** Implemented strategies to optimize API calls, such as batching requests and caching results, reducing the number of calls needed and improving application efficiency.

### Data Parsing and Structuring

- **Dynamic Type Handling:** Gained experience in managing dynamic types and unstructured data. Implemented functions to convert various data types to strings, ensuring consistency in data processing and enabling substring searches.

- **Language Localization:** Addressed challenges related to multilingual data by developing logic to handle and filter products in multiple languages, ensuring accurate data representation.

- **Data Normalization:** Learned to normalize and sanitize incoming data to improve consistency and reliability. This involved creating conversion functions and implementing checks for missing or incorrect data fields.

### Nutritional and Environmental Data Analysis

- **Nutritional Data Interpretation:** Developed skills in interpreting nutritional information and environmental scores from the API. This required understanding nutritional standards and translating data into meaningful insights for users.

- **Environmental Impact Assessment:** Integrated environmental data, such as Ecoscore, into the application, providing users with insights into the ecological impact of their food choices.

### Error Handling and Logging

- **Comprehensive Logging:** Implemented a comprehensive logging system to track application behavior and diagnose issues. This involved using Go's logging libraries to capture detailed runtime information.

- **Resilient Error Management:** Designed a robust error-handling framework that gracefully manages API failures and unexpected data conditions, ensuring continuous application operation and user experience.

### Performance Optimization

- **Scalability Considerations:** Developed strategies to handle large datasets efficiently, considering factors like memory usage and processing time. This involved exploring various data structures and algorithms for optimal performance.

- **Optimization Techniques:** Investigated optimization techniques such as indexing and lazy loading, which contribute to faster data retrieval and processing.

### Docker and Deployment

- **Containerization:** Gained experience in containerizing the application using Docker. This involved creating Dockerfiles and managing dependencies to ensure a consistent deployment environment.

- **CI/CD Pipelines:** Explored continuous integration and deployment pipelines to automate testing and deployment processes, ensuring reliable and timely updates.

### Web Development and RESTful Design

- **RESTful API Design:** Designed RESTful endpoints using Gorilla Mux, focusing on principles of statelessness and client-server separation. This improved the modularity and scalability of the application.

- **Routing and Middleware:** Implemented routing and middleware functionalities, such as logging and authentication, enhancing the application's extensibility and maintainability.

## Future Enhancements

- **Docker Integration:** Plan to containerize the application using Docker, simplifying deployment and scaling.

- **Trie vs. Contains Optimization:** Explore optimizing the ingredient search functionality by comparing Trie with alternative data structures for performance improvements.

- **User Interface:** Consider developing a frontend interface to improve user accessibility and experience, possibly leveraging modern web frameworks.

## Conclusion

NutriSearch successfully demonstrates how technology can empower consumers to make healthier food choices. By integrating the Open Food Facts API, the application provides valuable insights into product healthiness and promotes transparency in food labeling. This project not only showcases technical skills in Go and API integration but also contributes to public health awareness through informed dietary decisions.
