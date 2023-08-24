package main

import (
	informers "dcfenga/cloud-native-practice/crd-custom-vmcontroller/pkg/client/informers/externalversions/vmcontroller/v1"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"time"

	crdv1 "dcfenga/cloud-native-practice/crd-custom-vmcontroller/pkg/apis/vmcontroller/v1"
)

type Controller struct {
	informer  informers.VirtualMachineInformer
	workqueue workqueue.RateLimitingInterface
}

func NewController(informer informers.VirtualMachineInformer) *Controller {
	// 使用client和前面创建的Informer, 初始化了自定义控制器
	controller := &Controller{
		informer: informer,
		// WorkQueue 的实现，负责同步Informer和控制循环之间的数据
		workqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "VirtualMachine"),
	}

	klog.Info("Setting up virtual-machine event handlers")

	// informer 注册了三个Handler(AddFunc、UpdateFunc和DeleteFunc)
	// 分别对应API对象的“添加”“更新”和“删除”事件。
	// 而具体的处理操作，都是将该事件对用的API对象加入到工作队列中
	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueVirtualMachine,
		UpdateFunc: func(old, new interface{}) {
			oldObj := old.(*crdv1.VirtualMachine)
			newObj := new.(*crdv1.VirtualMachine)
			// 如果资源版本相同则不处理
			if oldObj.ResourceVersion == newObj.ResourceVersion {
				return
			}
			controller.enqueueVirtualMachine(new)
		},
		DeleteFunc: controller.enqueueVirtualMachineForDelete,
	})
	return controller
}
func (c *Controller) Run(threadiness int, stopChan <-chan struct{}) error {
	defer runtime.HandleCrash()
	defer c.workqueue.ShutDown()

	// 记录开始日志
	klog.Info("Starting VirtualMachine control loop")
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopChan, c.informer.Informer().HasSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopChan)
	}

	klog.Info("Started workers")
	<-stopChan
	klog.Info("Shutting down workers")
	return nil
}

// runWorker 是一个不断运行的方法，并且一直会调用 c.processNextWorkItem 从workqueue读取和去读取消息
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {

	}
}

// 从workqueue读取和读取消息
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()
	if shutdown {
		return false
	}
	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		if key, ok = obj.(string); !ok {
			c.workqueue.Forget(obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		if err := c.syncHandler(key); err != nil {
			return fmt.Errorf("error syncing '%s' : %s", key, err.Error())
		}
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		runtime.HandleError(err)
		return true
	}
	return true
}

// 尝试从 Informer 维护的缓存中拿到了它对应的 VirtualMachine 对象
func (c *Controller) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	virtualMachine, err := c.informer.Lister().VirtualMachines(namespace).Get(name)

	// 从缓存中拿不到这个对象，那就意味着这个VirtualMachine对象的key是通过前面的"删除"事件添加进工作队列的。
	if err != nil {
		if errors.IsNotFound(err) {
			// 对应的virtualMachine对象已经被删除了
			klog.Warningf("[VirtualMachineCRD] %s/%s does not exist in local cache, will delete it from VirtualMachine ...", namespace, name)
			klog.Infof("[VirtualMachineCRD] deleting virtualMachine: %s/%s ...", namespace, name)
			return nil
		}
		runtime.HandleError(fmt.Errorf("failed to get virtualMachine by: %s/%s", namespace, name))
		return err
	}
	klog.Infof("[VirtualMachineCRD] try to process virtualMachine: %#v ...", virtualMachine)
	return nil
}

func (c *Controller) enqueueVirtualMachineForDelete(obj interface{}) {
	var key string
	var err error

	key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	c.workqueue.AddRateLimited(key)
}

func (c *Controller) enqueueVirtualMachine(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}
	c.workqueue.AddRateLimited(key)
}
