package main

// // import (
// // 	"context"
// // 	"crypto/tls"
// // 	"fmt"
// // 	"net"
// // 	"net/http"
// // 	"time"

// // 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// // 	"k8s.io/client-go/kubernetes"
// // 	"k8s.io/client-go/rest"

// // 	tecNotebook "git-plat.xxx.net/ai-platform/notebook-controller/pkg/client/clientset/versioned/typed/xxx.io.ai.platform/v1"
// // )

// // func main() {
// // 	// dev test local k8s cluster.
// // 	tlsConfig := &tls.Config{InsecureSkipVerify: true}
// // 	var transport http.RoundTripper = &http.Transport{
// // 		DialContext: (&net.Dialer{
// // 			Timeout:   30 * time.Second,
// // 			KeepAlive: 30 * time.Second,
// // 		}).DialContext,
// // 		MaxIdleConns:          100,
// // 		IdleConnTimeout:       90 * time.Second,
// // 		TLSHandshakeTimeout:   10 * time.Second,
// // 		ExpectContinueTimeout: 1 * time.Second,
// // 		TLSClientConfig:       tlsConfig,
// // 		DisableCompression:    true,
// // 	}
// // 	config := &rest.Config{
// // 		Host:        "https://10.8.20.2:6443",
// // 		BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImtYVU80b3ZaY0FCSGNwX0pPeGhwc2M2MXZwa1YwSlJiMFphQ2l2aVdXa2MifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJub3RlYm9vay1zeXN0ZW0iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlY3JldC5uYW1lIjoibm90ZWJvb2stc2VydmljZS1hY2NvdW50LXRva2VuLWR6c3RjIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Im5vdGVib29rLXNlcnZpY2UtYWNjb3VudCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImQzZGEyNGQwLTQyZWEtNGI4MC1iMmUwLWNhZmU1NmM4Y2FiZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpub3RlYm9vay1zeXN0ZW06bm90ZWJvb2stc2VydmljZS1hY2NvdW50In0.C70cHYYigUX1W6Q9xVcvaXIcRdHb0U6kBR0sqDD2qPueDlKVXge-eTCdi0Zi1cb1T_acrWZDyQwgnhvXW6yjdf7-b8AssY5jR_R48SamIzPAsr8FZDdjFXYzt6eTl4MdesbtXOo_dsR7vqo343eRt7bF7yM39sY4RLP6iXujEc-loRbXwq9fNtOgVk2K6f-mzwieNKbvcN-kfQxZnhp3Vzy_bvjadfU7cMDGkG8i5JXvPoJ2AC8BZ2k4B52rfO5YPUkoF_V7B0yiGAq0ysCPaPCJ09qj5UrLhKzqcJP3NHBO9W3iYzqGovJy7k3kVYZmXKCRWkkCTdpE9tXfBt_svg",
// // 		Transport:   transport,
// // 	}

// // 	// Create a Kubernetes client.
// // 	client, err := kubernetes.NewForConfig(config)
// // 	if err != nil {
// // 		// Handle error.
// // 		fmt.Println(err)
// // 		return
// // 	}

// // 	// List all Pods across all namespaces.
// // 	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
// // 	if err != nil {
// // 		// Handle error.
// // 		fmt.Println(err)
// // 		return
// // 	}

// // 	// Print the name of each Pod.
// // 	fmt.Printf("Found %d pods.\n ", len(pods.Items))

// // 	nbClient, err := tecNotebook.NewForConfig(config)
// // 	if err != nil {
// // 		fmt.Printf("get tec notebook client set fail, err=%v", err)
// // 		return
// // 	}
// // 	nbs, err := nbClient.Notebooks("").List(context.TODO(),
// // 		metav1.ListOptions{
// // 			LabelSelector: "tenant_id=12",
// // 		})
// // 	if err != nil {
// // 		// Handle error.
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	fmt.Printf("Found %d nbs.\n ", len(nbs.Items))
// // }
