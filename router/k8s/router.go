package k8s

import (
	"github.com/gin-gonic/gin"
	"kube-auto/api"
)

type K8sRouter struct {
}

func (*K8sRouter) InitK8s(r *gin.Engine) {
	group := r.Group("/k8s")
	apiGroup := api.ApiGroupApp.K8SApiGroup
	group.GET("/list", apiGroup.GetPodList)
}
