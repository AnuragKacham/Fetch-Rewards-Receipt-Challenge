# **Fetch-Rewards: Receipt Processor Challenge**

**Author: Anurag Kacham**

## **Prerequisites**
Docker: To build and run the application in a containerized environment.
Git: To clone the repository and manage source code versions.

## **Project Structure**
The project consists of the following components:

### 1. Handlers
HTTP request handlers for processing receipts and retrieving points.

### 2. Point Calculation
Logic for calculating points based on receipt data.

### 3. Storage
In-memory storage for storing receipt points temporarily during application runtime.

### 4. API Endpoints
Exposes endpoints for receipt processing (POST /process) and points retrieval (GET /points/{id}).

## **Build and Run**
1. Make sure current terminal session is in the working directory.

2. Run build_image.sh
Commands:
**chmod +x build_image.sh**
**./build_image.sh**

3. Run run_app.sh
Commands:
**chmod +x run_app.sh**
**./run_app.sh**

## **Store data**
Store data in a bash file as shown in data_add.sh and run using:
**chmod +x ./data_add.sh**
**./data_add.sh**

Execution of bash file will output an id of the record. For ex: {"id":"b82eb836-0404-47aa-92d6-83c390203c13"}

Retrieving points:
2. Commands:
**Browser: http://localhost:8080/receipts/b82eb836-0404-47aa-92d6-83c390203c13/points**
**Terminal: curl http://localhost:8080/receipts/b82eb836-0404-47aa-92d6-83c390203c13/points**
