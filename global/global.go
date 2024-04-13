package global

import (
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"kube-auto/config"
)

var (
	CONF             config.Server
	KubeConfigSet    *kubernetes.Clientset
	MysqlDataConnect *gorm.DB
)
