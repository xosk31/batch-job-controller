[![Github Build](https://github.com/bakito/batch-job-controller/actions/workflows/build.yml/badge.svg)](https://github.com/bakito/batch-job-controller/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bakito/batch-job-controller)](https://goreportcard.com/report/github.com/bakito/batch-job-controller)
[![Coveralls github](https://img.shields.io/coveralls/github/bakito/batch-job-controller?logo=coveralls)](https://coveralls.io/github/bakito/batch-job-controller?branch=master)
[![GitHub Release](https://img.shields.io/github/release/bakito/batch-job-controller.svg?style=flat)](https://github.com/bakito/batch-job-controller/releases)

# Batch Job Controller

The batch job controller allows executing pods on nodes of a cluster, where the number of concurrent running pods can be
configured. Each pod can report it's results back to the controller to have them exposed as metrics.

## Deployment

The controller expects the following environment variables

| Name | Value                                                                                                                                                                                                                                                          |
| --- |----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| NAMESPACE | The current namespace                                                                                                                                                                                                                                          |
| CONFIG_MAP_NAME | The name of the configmap to read the config from                                                                                                                                                                                                              |
| POD_IP | The IP of the controller Pod. If defined, this IP is used for the callback URL of the job pods.(should be injected via [Downward API](https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/#the-downward-api)) |

## Configuration

The configuration has to be stored in a configmap with the following values

### config.yaml

Controller configuration

```yaml
name: ""                         # name of the controller; will also be used as prefix for the job pods
jobServiceAccount: ""            # service account to be used for the job pods. If empty the default will be used
jobImagePullSecrets:             # pull secrets to be used for the job pods for pulling the image
  - name: secret_name 
jobNodeSelector: { }             # node selector labels to define in which nodes to run the jobs
runOnUnscheduledNodes: true      # if true, jobs are also started on nodes that are unschedulable
cronExpression: "42 3 * * *"     # the cron expression to trigger the job execution
reportDirectory: "/var/www"      # directory to store and serve the reports
reportHistory: 30                # number of execution reports to keep
podPoolSize: 10                  # number of concurrent job pods to run
runOnStartup: true               # if 'true' the jobs are triggered on startup of the controller
startupDelay: 10s                # the delay as duration that is used to start the jobs if runOnStartup is enabled. default is '10s'
callbackServiceName: ""          # name of the controller service
callbackServicePort: 8090        # port of the controller callback api service
custom: { }                       # additional properties that can be used in a custom implementation
latestMetricsLabel: false        # if 'true' each result metric is also created with executionID='latest'
leaderElectionResourceLock: ""   # type of leader election resource lock to be used. ('configmapsleases' (default), 'configmaps', 'endpoints', 'leases', 'endpointsleases')
savePodLog: false                # if enabled, pod logs are saved along other with other job files
metrics:
  prefix: "foo_...."             # prefix for the metrics exposed by the controller
  gauges: # metric gauges that will be exposed by the jobs. The key is uses as suffix for the metrics. 
    test: # suffix of the metric
      help: "help ..."           # help text for the metric
      labels: # list of labels to be used with the metric. node and executionID are automatically added
        - label_a
        - label_b
```

### pod-template.yaml

The template of the pod to be started for each job. When a pod is created it gets enriched by the controller specific
configuration. [pkg\job\job.go](pkg\job\job.go)

## Job Pod

The job pod has the following env variables provided by the controller:

### Environment

| Name | Value |
| --- | --- |
| NAMESPACE | The current namespace |
| NODE_NAME | The name of the node it is running on |
| EXECUTION_ID | The id of the current job execution |
| CALLBACK_SERVICE_NAME | The name/host/ip of the callback service to send the report to |
| CALLBACK_SERVICE_PORT | The port of the callback service to send the report to |
| CALLBACK_SERVICE_RESULT_URL | The full qualified URL of the result callback service  |
| CALLBACK_SERVICE_FILE_URL | The full qualified URL of the file callback service, to send files to the controller |
| CALLBACK_SERVICE_EVENT_URL | The full qualified URL of the event callback service, to create k8s event |

### Callback

The controller exposes by default an endpoint to receive job results. The report is stored locally and metrics of the
reports will be exposed.

#### URL

The report URL is by default: **${CALLBACK_SERVICE_RESULT_URL}**

#### Body

The body of the report contains the metric suffixes that are also defined in the controller config. Each metric has a
decimal value and a map where the key is the label name and value is the value to be used for the metric label.

```json
{
  "test": [
    {
      "value": 1.0,
      "labels": {
        "label_a": "AAA",
        "label_b": "BBB"
      }
    },
    {
      "value": 2.554,
      "labels": {
        "label_a": "AAA2",
        "label_b": "BBB2"
      }
    }
  ]
}
```

Example job script: [helm\batch-job-controller\bin\run.sh](helm\batch-job-controller\bin\run.sh)

### Upload additional files

Additional files can be uploaded.

Use default **'Content-Disposition'** header or the **name** query parameter to define the name of the file. If the name
is not defined an uuid is generated. Each filename is prepended with the node name.

#### URL

The report URL is by default: **${CALLBACK_SERVICE_FILE_URL}**

### Create k8s Events from job pod

k8s Event can be created from each job pod by calling the event endpoint.

The 'reason' should be short and unique; it must be in UpperCamelCase format (starting with a capital letter).

Simple Message:

```json
{
  "warning": false,
  "reason": "TestReason",
  "message": "test message"
}
```

Massage with parameters

```json
{
  "warning": true,
  "reason": "TestReason",
  "message": "test message: %s",
  "args": [
    "a1"
  ]
}
```

#### URL

The event URL is by default: **${CALLBACK_SERVICE_EVENT_URL}**

### Examples

[test-queries.http](./testdata/test-queries.http)
