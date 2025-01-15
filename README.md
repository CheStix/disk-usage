### **Monitor Disk Usage** 
for Windows

![Static Badge](https://img.shields.io/badge/CheStix-diskUsage-diskUsage)
![GitHub top language](https://img.shields.io/github/languages/top/CheStix/disk-usage)
![GitHub](https://img.shields.io/github/license/CheStix/disk-usage)
![GitHub Repo stars](https://img.shields.io/github/stars/CheStix/disk-usage)
![GitHub issues](https://img.shields.io/github/issues/CheStix/disk-usage)
- **Description**: A CLI tool that displays usage statistics for a specified disk, and if no disk is specified, then for all disks in the system, including total, used, and free space.
- **Key Features**:
    - Accepts the disk name as a command line argument.
    - Outputs disk/disks usage details in a human-readable format.
    - Includes error handling for invalid paths.
- **Usage**:
  - By running the source code
    ```cmd
    go run disk_usage.go <disk letter, like C:\>
    ```
    ```cmd
    go run disk_usage.go
    ```
  - By running the executable file
    ```cmd
    disk_usage.exe <disk letter, like C:\>
    ```
    ```cmd
    disk_usage.exe
    ```
