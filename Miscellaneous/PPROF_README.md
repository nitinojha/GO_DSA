# PProf in Golang

`pprof` is a powerful profiling tool built into the **Go standard library**. It allows you to analyze various aspects of your program's performance, such as **CPU usage**, **memory allocation**, **goroutines**, and **mutex contention**. By using `pprof`, you can optimize your code, identify bottlenecks, and improve resource utilization.

---

## 📖 What is Profiling?

Profiling is a dynamic program analysis technique that measures various runtime characteristics of your application, such as:
- Execution time for specific functions.
- Memory usage patterns.
- Goroutine states and blocking contention.
- Mutex lock contention.

The data collected by `pprof` can be visualized and analyzed to help you understand where your application is spending time and how resources are utilized.

---

## ✍ PProf Measures

`pprof` provides multiple types of profiling data to help you analyze specific aspects of your application:

### 1. **CPU Profile**
   - Captures **CPU time** spent by your application across different functions.
   - **Use Case**: Identify computational bottlenecks or hotspots in your code.
   - Example: Find which function consumes the most CPU time during execution.

### 2. **Memory Profile**
   - Captures information about **memory allocation**.
   - **Use Case**: Identify memory leaks or areas with excessive memory allocation.
   - Example: Analyze which functions allocate the most heap memory.

### 3. **Mutex Profile**
   - Provides data about **mutex contention** in your program.
   - **Use Case**: Optimize lock contention in concurrent code to improve throughput.
   - Example: Diagnose if too many goroutines are waiting for the same lock.

### 4. **Goroutine Profile**
   - Shows the **state of all goroutines** in your application.
   - **Use Case**: Diagnose deadlocks or analyze goroutine utilization.
   - Example: Check if goroutines are stuck or if you’re running too many.

---

## 🚀 Getting Started with PProf

### 1. **Enable PProf**
Include the `net/http/pprof` package in your application to enable profiling:
```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // Your application logic here
}
 

 ## Few Commands
 go tool pprof http://localhost:6060/debug/pprof/profile    // use the pprof tool to collect profiles

 CPU Profiling
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30  //Collect CPU Profile
go tool pprof -http=:8080 <cpu-profile-file>  //Analyze CPU Usage

Memory Profiling
go tool pprof http://localhost:6060/debug/pprof/heap  //Collect Memory Profile
go tool pprof <heap-profile-file>  // memory allocation hotspots


Goroutine Profiling
curl http://localhost:6060/debug/pprof/goroutine > goroutine.prof  //Capture goroutine states
go tool pprof -http=:8080 goroutine.prof   //Analyze goroutine activity 

