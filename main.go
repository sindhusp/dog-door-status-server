package main
import (
	"fmt"
	"log"
	"net/http"
        "strings"
	"encoding/json"
	"os"
	"os/exec"
	
)

type Data struct {
  Status string `json:"status"`
}

var dogStatus string = "Unknown"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello from your Go server!")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Data
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	dogStatus = req.Status

        if strings.EqualFold("BY_THE_DOOR", dogStatus) {
      		cmd := exec.Command("sh", "-c", " eips -g /mnt/us/dog-inside.png")
	
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		
		err := cmd.Run()
		if err != nil {
			log.Printf("kindle display command failed: %v", err)
			http.Error(w, "Failed to update display", http.StatusInternalServerError)
			return
		}
	
	}

        if strings.EqualFold("NOT_BY_THE_DOOR", dogStatus) {
                cmd := exec.Command("sh", "-c", " eips -g /mnt/us/dog-unknown.png")

                cmd.Stdout = os.Stdout
                cmd.Stderr = os.Stderr


                err := cmd.Run()
                if err != nil {
                        log.Printf("kindle display command failed: %v", err)
                        http.Error(w, "Failed to update display", http.StatusInternalServerError)
                        return
                }

        }
	fmt.Fprintf(w, "Variable updated to: %s\n", dogStatus)
}

func handleClear(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}


	cmd := exec.Command("sh", "-c", " eips -c")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr


	err := cmd.Run()
	if err != nil {
		log.Printf("kindle display command failed: %v", err)
		http.Error(w, "Failed to clear display", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/dog", handlePost)
	http.HandleFunc("/clear", handleClear)

	fmt.Println("Starting server at port :8080")
	
	// Start the server on port 8080 and log if something goes wrong
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
