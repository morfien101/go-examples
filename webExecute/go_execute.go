package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	// These are flags that we want to pull in. They have sane defaults.
	addr = flag.String("a", "", "IP that the HTTP engine should listen on.")
	port = flag.Int("n", 8000, "TCP port that the HTTP engine should listen on.")

	// We are being naughty and globally scoping the struct for the jobs list. This is to simplify the code.
	// Currently we just state what it is. In the main function we set it to a value.
	globalJobList *jobList
)

// We type cast the id for jobs. Mainly to show that you can do it.
type id int64

// We then attach a function to the type. This is one reason to type cast.
func (i *id) toInt64() int64 {
	return int64(*i)
}

// jobList is a struct that holds a list of jobs. It also holds a read write mutex.
// This is useful to allow our application to be thread safe. Go's http package makes
// extensive use of go routines. So if we do anything in the handler functions, we
// need to make sure that it uses thread safe objects.
// It also holds the next ID so that we don't need to compute it every time we allocate
// a new job.
type jobList struct {
	sync.RWMutex
	jobs   map[id]*job
	nextID id
}

// job is a struct that holds the information about the job itself.
// It is labled with JSON tags. These are consumed in the json.Marshal function.
// Also notice that the fields are Capitalized. This is required to export them in
// the marshaller.
// The ID is also the type id.
type job struct {
	ID        id     `json:"id"`
	Complete  bool   `json:"finished"`
	Output    string `json:"output"`
	ErrorText string `json:"error"`
}

// This is a constructor function. We use these to help our users not have to
// have intimate knowledge with our structs. Really this should be in its own
// package. But for simplicity we are keeping it all in a single file.
func newJobList() *jobList {
	return &jobList{
		jobs:   make(map[id]*job),
		nextID: 1,
	}
}

// This section is the functions that attach to our job list struct.
// addJob will take a command and its arguments and run it.
// There is scheduling here so if you mash F5 then it will spawn (n * F5 presses) jobs.
// You could make use of worker pools here. Another lesson on that.
// We return an int64 to the user since they can use that.
func (jl *jobList) addJob(cmd string, args ...string) int64 {
	// Collect the next ID which will also increment the number ready for the next request
	jobID := jl.getNextID()
	j := &job{
		ID:        jobID,
		Output:    "No output currently.",
		ErrorText: "No Errors.",
		Complete:  false,
	}
	// lock the stuct to stop other requests from tampering with it and causing a panic
	jl.Lock()
	// add the job to the map
	jl.jobs[jobID] = j
	// unlock now to allow the requst of the requests to process
	jl.Unlock()
	// Start the job and passing on the arguments that we got.
	// args... expands the slice as it calls the function.
	go jl.execute(jobID, cmd, args...)
	// We are using our slighty pointless method here to convert the id to an int64
	// We could also do int64(jobID) to cast it to a int64
	// However you can see the idea of typeing a primitive type.
	return jobID.toInt64()
}

// This function should be called as a go routine as it would block the applcation while
// the command is being executed.
func (jl *jobList) execute(i id, cmd string, args ...string) {
	// exec calls a sub process to run your command. If the go application dies so will this
	// sub process. In this case we combine the STDOUT and STDERR in to a single thing.
	// We collect any error that is passed pack and store it in our job struct.
	out, err := exec.Command(cmd, args...).CombinedOutput()
	// Notice that we lock only AFTER the job has finished executing.
	// If we locked before then all the go routines that need access to jobList would
	// have to wait till this command finished executing and wrote its staatus back.
	// That could casue time outs.
	jl.Lock()
	// Defer statements are executed when you return out the function.
	// They are etreamly useful as you don't need to remember to call it at every exit point.
	defer jl.Unlock()
	// We could re-create the job here also as we have all the information. eg &job{...}
	// Then allocate it to the list again with the same ID. But this reads easier. A single allocation
	// would actually be faster in terms of how long the sturct is locked for.
	jl.jobs[i].Output = string(out)
	jl.jobs[i].Complete = true
	if err != nil {
		jl.jobs[i].ErrorText = fmt.Sprint(err)
	}
}

// This is a very expensive task and will get progressivly slower the more jobs in the list.
// Basically we are converting our map into a list or in go terms a slice.
func (jl *jobList) readAllJobs() []job {
	// Make a list to start pushing jobs into.
	list := make([]job, 0)
	// This job seems like it would just need a read lock but that would leave an opening for a
	// race condition. Consider this, we are cycling through all the current jobs in the map.
	// a read lock doesn't stop writes happening. Therefore we could read a half written job.
	// So we need to lock it completely.
	// Remember to lock at late as possible and unlock as soon as possible.
	jl.Lock()
	defer jl.Unlock()
	// This will not be a sorted list.
	// Maps in Go are made to be random. You should NEVER expect them to be ordered.
	// We get key, value from maps. We don't care for the key so we allocate to the floor. aka the bin.
	for _, v := range jl.jobs {
		// *v copies the value inside the pointer. This means that we give a copy out.
		// The user can do what ever they want to the copy. It is not ours and we
		// don't care about it.
		list = append(list, *v)
	}
	// We could sort the list but that is outside the scope if this demo.
	return list
}

func (jl *jobList) readJob(jobID int64) (job, error) {
	jl.RLock()
	defer jl.RUnlock()
	// We emit a copy of the job here.
	jobData, ok := jl.jobs[id(jobID)]
	if !ok {
		// We can't return a nil in place of a job. So we just return an empty job container.
		return job{}, errors.New("Job ID invalid.")
	}
	// Return a copy of the job and a nil for the error.
	return *jobData, nil
}

func (jl *jobList) getNextID() id {
	jl.Lock()
	defer jl.Unlock()
	i := jl.nextID
	jl.nextID++
	return i
}

func main() {
	// flag.Parse will consume the arguments passed in at run time and set the flags for us.
	flag.Parse()

	// We now set the value of our job list. Before we just told the compiler that it is a thing.
	// Now we are telling it this thing looks like this.
	globalJobList = newJobList()
	// Combine the address and port into something that the http package likes.
	listenAddress := fmt.Sprintf("%s:%d", *addr, *port)

	// We are making use of the Gorilla Mux router because they allow regex in the routes.
	r := mux.NewRouter()
	r.HandleFunc("/", pingPass)
	r.HandleFunc("/jobs", listJobs)
	// The regex here is how we get our job number.
	r.HandleFunc("/jobs/{jobNumber:[0-9]+}", getJobOutput)

	// Start the http server.
	// It is a blocking action so in other programs is often best used in a go routine.
	// Again we can see this in a furture lesson as it is out of the scope for this demo.
	log.Fatal(http.ListenAndServe(listenAddress, r))
}

// PingPass is my sample job. Its not taxing on CPU but takes time and generates a decent output.
// We also need to becareful as the option are different on Windows and unix like platforms.
// Make use of runtime's GOOS constant to see what you are running on.
// In another leason I will show how to do file based os detection.
func pingPass(w http.ResponseWriter, r *http.Request) {
	var c []string
	if runtime.GOOS != "windows" {
		c = []string{"ping", "-c", "10", "127.0.0.1"}
	} else {
		c = []string{"ping", "-n", "10", "127.0.0.1"}
	}
	// The Job ID is a int64 here so we don't need to do anything special to the fmt statement.
	// The struct contains both the command and arguments. So we need to break it up a little.
	// c[1:]... says return c from the second element till the end and expand the output into the function.
	jobid := globalJobList.addJob(c[0], c[1:]...)
	// return a nice page to the user
	fmt.Fprintf(w, `
<html>
	<title>Home</title>
	<body>
		Job started: <a href="/jobs/%d"> Job Status Page</a>
	</body>
</html>`, jobid)
}

// listJobs will list the jobs that are currently available to gather output from
func listJobs(w http.ResponseWriter, r *http.Request) {
	jobslist := globalJobList.readAllJobs()
	if len(jobslist) == 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "There are no jobs to view.")
		return
	}
	// Marshalling returns a json representation of the struct that you pass in.
	// We also decorated our struct with JSON tags so the Marshaller can use that.
	// It throws and error if something goes wrong. We can use that.
	b, err := json.Marshal(jobslist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to Marshal output. Error: %s", err)
	}
	// Set the content type correctly. Some browsers will digest json and pretty print it for you.
	// Firefox. hint, hint. 56+ I think.
	w.Header().Set("content-type", "application/json")
	// We can use the Write function because the Marshaller returns a byte slice.
	w.Write(b)
}

func getJobOutput(w http.ResponseWriter, r *http.Request) {
	// This is how we use the Gorrila Mux router to get the values for us.
	jidString := mux.Vars(r)["jobNumber"]
	// We need to parse it into a base 10 number and represent it as a int64
	// (string, base of number, bit length)
	jidInt, err := strconv.ParseInt(jidString, 10, 64)
	// It is very tricky to get here as the regex for the path doesn't allow it.
	// But we should deal with errors correctly as someone could update the path to allow
	// non positive numeric values to get here.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID %s. Error: %s", jidString, err)
	}
	j, err := globalJobList.readJob(jidInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to red job data. Error: %s", err)
		return
	}
	b, err := json.Marshal(j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to Marshal output. Error: %s", err)
	}
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}
