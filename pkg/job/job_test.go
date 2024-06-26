package job

import (
	"github.com/bakito/batch-job-controller/pkg/config"
	"github.com/bakito/batch-job-controller/pkg/controller"
	"github.com/ghodss/yaml"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Job", func() {
	Context("New", func() {
		var (
			cfg              *config.Config
			name             string
			namespace        string
			nodeName         string
			id               string
			serviceIP        string
			sacc             string
			imagePullSecrets []corev1.LocalObjectReference
		)
		BeforeEach(func() {
			name = uuid.New().String()
			namespace = uuid.New().String()
			sacc = uuid.New().String()
			imagePullSecrets = []corev1.LocalObjectReference{{Name: "mySecret"}}
			cfg = &config.Config{
				Name:                name,
				Namespace:           namespace,
				JobServiceAccount:   sacc,
				JobImagePullSecrets: imagePullSecrets,
				JobPodTemplate:      "kind: Pod",
				CallbackServicePort: 12345,
			}
			nodeName = uuid.New().String()
			id = uuid.New().String()
			serviceIP = "1.1.1.1"
		})
		It("should set default fields", func() {
			pod, err := New(cfg, nodeName, id, serviceIP, nil)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(pod).ShouldNot(BeNil())

			Ω(pod.Name).Should(Equal(name + "-job-" + nodeName + "-" + id))
			Ω(pod.Namespace).Should(Equal(namespace))
			Ω(pod.Spec.RestartPolicy).Should(Equal(corev1.RestartPolicyNever))
			Ω(pod.Spec.NodeName).Should(Equal(nodeName))
			Ω(pod.Spec.ServiceAccountName).Should(Equal(sacc))
			Ω(pod.Spec.ImagePullSecrets).Should(Equal(imagePullSecrets))

			Ω(pod.Labels[controller.LabelExecutionID]).Should(Equal(id))
			Ω(pod.Labels[controller.LabelOwner]).Should(Equal(name))
		})

		Context("Env vars", func() {
			BeforeEach(func() {
				pod := &corev1.Pod{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name: "c1",
								Env: []corev1.EnvVar{
									{Name: envExecutionID, Value: "myID"},
									{Name: "FOO", Value: "bar"},
								},
							},
						},
						InitContainers: []corev1.Container{
							{
								Name: "ic1",
								Env: []corev1.EnvVar{
									{Name: envExecutionID, Value: "myID"},
									{Name: "BAR", Value: "foo"},
								},
							},
						},
					},
				}
				b, _ := yaml.Marshal(pod)
				cfg.JobPodTemplate = string(b)
			})
			It("should set default env vars", func() {
				pod, _ := New(cfg, nodeName, id, serviceIP, nil)

				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(envExecutionID, id))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(envNamespace, namespace))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(envNodeName, nodeName))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(EnvCallbackServiceName, serviceIP))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(EnvCallbackServicePort, "12345"))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(EnvCallbackServiceResultURL, "http://1.1.1.1:12345/report/"+nodeName+"/"+id+"/result"))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(EnvCallbackServiceFileURL, "http://1.1.1.1:12345/report/"+nodeName+"/"+id+"/file"))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(EnvCallbackServiceEventURL, "http://1.1.1.1:12345/report/"+nodeName+"/"+id+"/event"))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar("FOO", "bar"))

				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar(envExecutionID, id))
				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar(envNamespace, namespace))
				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar(envNodeName, nodeName))
				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar(EnvCallbackServiceName, serviceIP))
				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar(EnvCallbackServicePort, "12345"))
				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar(EnvCallbackServiceResultURL, "http://1.1.1.1:12345/report/"+nodeName+"/"+id+"/result"))
				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar(EnvCallbackServiceFileURL, "http://1.1.1.1:12345/report/"+nodeName+"/"+id+"/file"))
				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar(EnvCallbackServiceEventURL, "http://1.1.1.1:12345/report/"+nodeName+"/"+id+"/event"))
				Ω(pod.Spec.InitContainers[0].Env).Should(HaveEnvVar("BAR", "foo"))
			})

			It("should have a correct owner reference", func() {
				ownerID := uuid.New().String()
				ownerName := uuid.New().String()
				pod, _ := New(cfg, nodeName, id, serviceIP, &corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						UID:  ktypes.UID(ownerID),
						Name: ownerName,
					},
				})

				Ω(pod.OwnerReferences).Should(HaveLen(1))
				Ω(string(pod.OwnerReferences[0].UID)).Should(Equal(ownerID))
				Ω(pod.OwnerReferences[0].Name).Should(Equal(ownerName))
			})

			It("should have a correct custom env variables reference", func() {
				pod, _ := New(cfg, nodeName, id, serviceIP, nil, &customEnv{})

				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar(envNamespace, namespace))
				Ω(pod.Spec.Containers[0].Env).Should(HaveEnvVar("CUSTOM", "VALUE"))
			})
		})
	})
})

type customEnv struct{}

func (ce *customEnv) ExtendEnv(_ *config.Config, _ string, _ string, _ string, _ corev1.Container) []corev1.EnvVar {
	return []corev1.EnvVar{{Name: envNamespace, Value: "notMyNamespace"}, {Name: "CUSTOM", Value: "VALUE"}}
}

func HaveEnvVar(name, value string) types.GomegaMatcher {
	return ContainElement(And(WithTransform(getName, Equal(name)), WithTransform(getValue, Equal(value))))
}

func getName(envVar corev1.EnvVar) string {
	return envVar.Name
}

func getValue(envVar corev1.EnvVar) string {
	return envVar.Value
}
