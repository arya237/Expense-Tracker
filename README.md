# Expense Tracker

A modern, secure, and scalable expense tracking application built with Go.  Track your expenses efficiently with a RESTful API backend powered by Gin and MongoDB.

## 🎯 Features

- **User Authentication**:  Secure JWT-based authentication with encrypted passwords
- **Expense Management**: Create, read, update, and delete expense records
- **RESTful API**: Clean and intuitive API endpoints built with Gin framework
- **MongoDB Integration**: Reliable NoSQL database for persistent data storage
- **Security**: Industry-standard encryption and security best practices
- **Comprehensive Testing**: Unit tests and integration tests included
- **Production-Ready**: Well-structured codebase following Go best practices

## 🛠 Tech Stack

- **Language**: Go 1.24. 1
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin) - Fast HTTP web framework
- **Database**: [MongoDB](https://www.mongodb.com/) - NoSQL document database
- **Authentication**: [JWT](https://github.com/golang-jwt/jwt) - JSON Web Tokens
- **Encryption**: [golang. org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) - Cryptographic functions
- **Testing**: [Testify](https://github.com/stretchr/testify) - Testing toolkit

## 📁 Project Structure

```
├── cmd/                    # Application commands and entry points
├── internal/
│   ├── db/                 # Database connection and configuration
│   ├── handler/            # HTTP request handlers
│   ├── middleware/         # HTTP middleware (auth, logging, etc.)
│   ├── models/             # Data models and structures
│   ├── repository/         # Data access layer
│   └── service/            # Business logic layer
├── utils/                  # Utility functions and helpers
├── tests/                  # Test files and test utilities
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
└── go.sum                  # Go module checksums
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24.1 or higher
- MongoDB 4.0 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/arya237/Expense-Tracker.git
   cd Expense-Tracker
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the project root:
   ```
   MONGODB_URI=mongodb://localhost:27017
   MONGODB_DB=expense_tracker
   JWT_SECRET=your_secret_key_here
   SERVER_PORT=8080
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`

## 📚 API Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login and receive JWT token

### Expenses
- `GET /api/expenses` - Get all expenses (requires authentication)
- `POST /api/expenses` - Create a new expense (requires authentication)
- `GET /api/expenses/:id` - Get expense by ID (requires authentication)
- `PUT /api/expenses/:id` - Update an expense (requires authentication)
- `DELETE /api/expenses/:id` - Delete an expense (requires authentication)

## 📦 Building

Build the application:
```bash
go build -o expense-tracker main. go
```

Run the compiled binary:
```bash
./expense-tracker
```

## 🔐 Security

- Passwords are encrypted using industry-standard algorithms
- JWT tokens are used for secure API authentication
- All sensitive data is validated before storage
- Database connection uses secure protocols

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. 

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 💬 Support

If you encounter any issues or have questions, please open an issue on GitHub. 

## 🎓 Learning Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Gin Framework Documentation](https://gin-gonic.com/)
- [MongoDB Go Driver](https://pkg.go.dev/go. mongodb.org/mongo-driver)
- [JWT Best Practices](https://tools.ietf.org/html/rfc7519)

---
