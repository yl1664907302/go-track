package k8s

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube-auto/global"
	"net/http"
)

type PodApi struct {
}

func (*PodApi) GetPodList(c *gin.Context) {

	//不懂
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
	}
	for _, item := range list.Items {
		fmt.Println(item.Namespace, item.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"正在读取k8s": "数据",
	})
}
