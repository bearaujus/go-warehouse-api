# GO-WAREHOUSE-API

Welcome to **GO-WAREHOUSE-API**, a project designed to fulfill the `Senior Back End Engineer` challenge. This repository showcases a simplified warehouse system built with Go, demonstrating my ability to handle real-world warehouse management tasks, such as stock tracking, inventory management, and API design.

## Architecture

### High-Level Design (HLD)

To Be Added (TBA)

### Database Schema

To Be Added (TBA)

## Running the Service

To get the **GO-WAREHOUSE-API** up and running on your local machine, follow the steps below:

### 1. Clone the Repository
Start by cloning the repository:
```bash
git clone https://github.com/bearaujus/go-warehouse-api.git go-warehouse-api
```

### 2. Navigate to the Project Directory
Move into the project directory:
```bash
cd go-warehouse-api
```

### 3. Set Up Environment Variables
You need to create a copy of the `.env.example` file and name it `.env` in the `etc/files/` directory:
```bash
cp etc/files/.env.example etc/files/.env
```

### 4. (Optional) Customize Environment Variables
Feel free to adjust the environment variables to suit your needs. You can modify the `.env` file using your favorite text editor:
```bash
nano etc/files/.env
```

### 5. (Optional) Install Makefile (if not installed)
If you haven’t installed [Make](https://www.gnu.org/software/make/manual/make.html) yet, you can install it using the following commands:
```bash
# Update package list and install make
sudo apt update
sudo apt install make
```

### 6. Run the Application
Finally, you can run the application using `make` commands:

- To start the application in the background:
  ```bash
  make up
  ```

- To stop the application:
  ```bash
  make down
  ```

## Using the Service

You can explore the API via Postman. Full API documentation is available here: [API Spec](https://documenter.getpostman.com/view/37777109/2sAXjNZWt5).


---

Made with ❤️ by [bearaujus](https://www.linkedin.com/in/bearaujus/) © 2024

---
