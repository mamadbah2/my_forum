# 💬 Forum Web Application

A complete web forum built with Go featuring user authentication, post management, commenting system, and category-based filtering.

![Go](https://img.shields.io/badge/Go-77.6%25-00ADD8?logo=go)
![CSS](https://img.shields.io/badge/CSS-19.4%25-1572B6?logo=css3)
![JavaScript](https://img.shields.io/badge/JavaScript-1.5%25-F7DF1E?logo=javascript)
![Docker](https://img.shields.io/badge/Docker-1.3%25-2496ED?logo=docker)

## 🎯 Features

### 👤 User Management
- User registration with secure password encryption (bcrypt)
- Login/Logout with session management (UUID-based)
- User profile information

### 📝 Posts & Interactions
- Create posts with multiple categories
- Like/Dislike system for posts
- Category-based filtering
- Post browsing and feed display

### 💬 Comments
- Comment on posts
- Like/Dislike comments
- Nested comment display
- Real-time interaction counters

## 🛠️ Technology Stack

### **Backend**
![Go](https://img.shields.io/badge/Go_1.22.0-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite3-003B57?style=for-the-badge&logo=sqlite)

- **Language**: Go (Golang) 1.22.0
- **Database**: SQLite3
- **Libraries**:
  - `github.com/mattn/go-sqlite3` - Database driver
  - `github.com/gofrs/uuid` - Session token generation
  - `golang.org/x/crypto` - Password hashing (bcrypt)

### **Frontend**
- **HTML5** - Template system with `text/template`
- **CSS3** - Modern responsive design
- **JavaScript** - Dynamic interactions

### **DevOps**
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
- Docker containerization with Alpine Linux
- Automated deployment scripts

## 🚀 Quick Start

**Option 1: Run Locally**
```bash
# Clone the repository
git clone https://github.com/mamadbah2/my_forum.git
cd my_forum

# Run the application
go run ./cmd/web/.

# Access at http://localhost:4000
```

**Option 2: Docker Deployment**
```bash
# Using the provided script
./Run_docker.sh

# Or manually
docker build -t forum .
docker run -p 1234:4000 forum

# Access at http://localhost:1234
```

## 📁 Project Structure

```
my_forum/
├── cmd/web/              # Application core
│   ├── main.go          # Server initialization
│   ├── handlers.go      # HTTP handlers
│   ├── routes.go        # Routing
│   └── templates.go     # Template rendering
├── internal/
│   ├── models/          # Database models
│   ├── filters/         # Business logic
│   └── utils/           # Utilities
├── db/
│   ├── migration.sql    # Database schema
│   └── data.sql         # Seed data
├── ui/
│   ├── static/          # CSS, JS, images
│   └── html/            # HTML templates
├── Dockerfile
├── Run_docker.sh
└── go.mod
```

## 🎓 Skills Acquired

### **Backend Development**
✅ **Go Programming** - HTTP servers, routing, middleware  
✅ **Database Management** - SQLite, SQL queries, migrations  
✅ **Authentication & Security** - Session management, bcrypt encryption  
✅ **MVC Architecture** - Clean code structure and separation of concerns

### **Web Development**
✅ **RESTful APIs** - HTTP methods (GET, POST)  
✅ **Template Engine** - Server-side rendering with Go templates  
✅ **Form Handling** - Data validation and sanitization  
✅ **Session Management** - Cookie-based authentication

### **DevOps**
✅ **Docker** - Containerization and deployment  
✅ **Git** - Version control and collaboration

### **Software Engineering**
✅ **Error Handling** - Robust error management  
✅ **Code Organization** - Modular design patterns  
✅ **Testing & Debugging** - Problem-solving strategies

## 👥 Authors

- **[mamadbah2](https://github.com/mamadbah2)**
- **[Kendisec](https://github.com/Kendisec)** 
- Contributors: belhadjs, msoumare

## 📝 License

Open source project - Available for educational purposes.

---

**Built with ❤️ at Zone01 Dakar**
