# Cloud Desktop Management Platform

The Cloud Desktop Management Platform is a comprehensive solution for managing and accessing cloud desktops, including server-side, web frontend, and client applications. This platform allows users to create, manage, and connect to virtual cloud desktops, supporting GPU allocation, disk management, and user permission control.

## Project Structure

The project contains the following main components:

- **Frontend (front)**: Vue.js-based web management interface
- **Backend (backend)**: Go language and Gin framework-based server
- **Desktop Client (client)**: PyQt5-based Windows desktop application
- **Android Client (client_android)**: uni-app-based Android mobile application
- **Deployment Tools (deploy)**: Scripts and configuration files for service deployment

## Features

- User Management: Create, modify, and delete user accounts
- Cloud Desktop Management: Create, start, shut down, and restart cloud desktops
- Hardware Resource Management: Allocate and monitor GPU resources, manage disk space
- Remote Connection: Connect to cloud desktops via RDP protocol


## Component Description

### Frontend (front)

The frontend is developed based on the Vue.js framework, using Element UI as the UI component library, providing a user-friendly web management interface.


**Development and Deployment:**
- Development environment: `npm run serve`
- Build for deployment: `npm run build`

### Backend (backend)

The backend is developed using Go language and Gin framework, providing RESTful API interfaces.


**Development and Deployment:**
- Development environment: `go run main.go`
- Build for deployment: `go build -o cloud_server_main.exe`

### Desktop Client (client)

The desktop client is developed using Python and PyQt5, providing a Windows desktop application interface.

**Main Files:**
- `main.py`: Main program entry
- `request_operate.py`: Network request handling
- `Ui_login.py`: Login interface
- `Ui_main.py`: Main interface
- `Ui_config.py`: Configuration interface
- `Ui_password.py`: Password modification interface

**Features:**
- Login verification
- Cloud desktop list display
- Cloud desktop operations (power on, power off, restart)
- Remote connection to cloud desktops
- Modify user password

### Android Client (client_android)

The Android client is developed based on the uni-app framework and can be built into an Android application.



## Deployment Instructions

### Frontend Deployment

1. Install nginx on Windows
   - Copy nginx to the specified directory C:\nginx
2. Copy the files in the compiled dist directory to the html directory under the nginx directory
3. Start nginx, open powershell and execute the command: `start nginx`
4. Other commands:
   - Stop nginx: `./nginx.exe -s stop`
   - Reload configuration: `./nginx.exe -s reload`

### Backend Deployment

1. Install nssm on Windows
   - Copy nssm.exe to the specified directory C:\nssm
   - Add the C:\nssm path to the system path Path
2. Install cloud_server on Windows
   - Copy the release files (cloud_server_main.exe, script, settings.yaml) to the specified directory C:\cloud_server
   - Open powershell as administrator and execute the command: `nssm install cloud_server_main C:\cloud_server\cloud_server_main.exe`
   - Start the cloud_server service, execute the command: `nssm start cloud_server_main`
3. Other commands:
   - Stop service: `nssm stop cloud_server_main`
   - Remove service: `nssm remove cloud_server_main`

## System Requirements

- **Frontend**: Modern web browsers (Chrome, Firefox, Edge, etc.)
- **Backend**: Windows operating system
- **Desktop Client**: Windows operating system, supports Python 3.6+
- **Android Client**: Android 5.0+