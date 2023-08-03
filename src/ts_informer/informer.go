package main

import (
	"fmt"
	"log"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	vcbatch "volcano.sh/apis/pkg/apis/batch/v1alpha1"
	vCllient "volcano.sh/apis/pkg/client/clientset/versioned"
	volcano "volcano.sh/apis/pkg/client/informers/externalversions"
)

func main() {
	// 创建 Kubernetes 配置
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	log.Fatalf("Error creating in-cluster config: %v", err)
	// }
	config, err := clientcmd.BuildConfigFromFlags("", "/home/liuqw/.kube/config")
	if err != nil {
		log.Fatalf("Error creating out-cluster config: %v", err)
	}

	// 创建 Kubernetes 客户端
	config.QPS = 100
	config.Burst = 100
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}
	volcanoClientset, err := vCllient.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Volcano client: %v", err)
	}

	// 创建 SharedInformerFactory
	factory := informers.NewSharedInformerFactory(clientset, 0)
	vFactory := volcano.NewSharedInformerFactory(volcanoClientset, 0)

	// 创建 Pod Informer
	podInformer := factory.Core().V1().Pods().Informer()
	vjInformer := vFactory.Batch().V1alpha1().Jobs().Informer()

	// 等待所有缓存同步
	// factory.WaitForCacheSync(stopCh)

	// 添加事件处理函数
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			start := time.Now()
			pod := obj.(*v1.Pod)
			fmt.Printf("Pod added at %v: %s %s at %v, version=%s\n", start, pod.Namespace, pod.Name, pod.CreationTimestamp, pod.ResourceVersion)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			start := time.Now()
			pod := newObj.(*v1.Pod)
			fmt.Printf("Pod updated at %v: %s %s to %v, version=%s\n", start, pod.Namespace, pod.Name, pod.Status.Phase, pod.ResourceVersion)
		},
		DeleteFunc: func(obj interface{}) {
			start := time.Now()
			pod := obj.(*v1.Pod)
			fmt.Printf("Pod deleted at %v: %s %s, version=%s\n", start, pod.Namespace, pod.Name, pod.ResourceVersion)
		},
	})

	vjInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			start := time.Now()
			vj := obj.(*vcbatch.Job)
			fmt.Printf("VCJob added at %v: %s %s at %v, version=%s\n", start, vj.Namespace, vj.Name, vj.CreationTimestamp, vj.ResourceVersion)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			start := time.Now()
			vj := newObj.(*vcbatch.Job)
			fmt.Printf("VCJob updated at %v: %s %s to %v, version=%s\n", start, vj.Namespace, vj.Name, vj.Status.State.Phase, vj.ResourceVersion)
		},
		DeleteFunc: func(obj interface{}) {
			start := time.Now()
			vj := obj.(*vcbatch.Job)
			fmt.Printf("VCJob deleted at %v: %s %s, version=%s\n", start, vj.Namespace, vj.Name, vj.ResourceVersion)
		},
	})

	// 启动 Informer
	stopCh := make(chan struct{})
	defer close(stopCh)
	factory.Start(stopCh)
	vFactory.Start(stopCh)

	// 运行直到程序终止
	<-stopCh
}
