package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Pod type
type pod struct {
	Metadata struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		//UID       string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		NodeName string `json:"nodeName"`
	} `json:"spec"`
	Status struct {
		Phase     string    `json:"phase"`
		HostIP    string    `json:"hostIP"`
		PodIP     string    `json:"podIP"`
		//StartTime time.Time `json:"startTime"`
	} `json:"status"`
}

// Podlist type
type podlist struct {
	Items []pod `json:"items"`
}

// Return pertinent pod info
func (c *pod) GetPodInfo() map[string]string {
	return map[string]string{
		"name":      c.Metadata.Name,
		"namespace": c.Metadata.Namespace,
		"node":      c.Spec.NodeName,
		"status":    c.Status.Phase,
		"hostip":    c.Status.HostIP,
		"podip":     c.Status.PodIP,
	}
}

// Generate list of pod info
func (c *podlist) GenerateTable() []map[string]string {
	results := make([]map[string]string, 0)
	for _, p := range c.Items {
		results = append(results, p.GetPodInfo())
	}
	return results
}

// Main
func main() {

	// Parse url flag
	url := flag.String("url", "http://127.0.0.1:10255/pods", "Complete URL to pods endpoint")
	flag.Parse()

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, *url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "kubelet-go")

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	pods := podlist{}
	jsonerr := json.Unmarshal(body, &pods)
	if jsonerr != nil {
		panic(err)
	}

	out, err := json.Marshal(pods.GenerateTable())
	if err != nil {
		panic("Cannot read GenerateTable() result")
	}
	fmt.Println(string(out))

}
