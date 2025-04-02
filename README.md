# **nslookup_go_app**

A simple Go application that reads URLs from a `.txt` file, performs DNS lookups (`nslookup`), and writes the results to an output file.

---

## **Features**
- Reads URLs from a `.txt` file.
- Resolves hostnames to IP addresses.
- Logs errors for invalid or unreachable domains.
- Outputs results to a file.

---

## **Usage**

### **Prerequisites**
- Go installed on your system.
- A `.txt` file with URLs (one per line).

### **Steps to Run**
1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/nslookup_go_app.git
   cd nslookup_go_app
   ```
2. Build and run the app:
   ```bash
   go run main.go
   ```
3. Check the `results.txt` file for the output.

---

## **Example**
Input (`urls.txt`):
```
example.com
google.com
invalid.url
```

Output (`results.txt`):
```
example.com: [93.184.216.34]
google.com: [142.250.190.78 142.250.190.110]
invalid.url: Error resolving invalid.url: lookup invalid.url: no such host
```
