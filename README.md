# gke-visibility-tracing

```shell
# Create service account
export APPLICATION=metadata-agent
export APP_SA_NAME=gke-$APPLICATION-sa
gcloud iam service-accounts create $APP_SA_NAME --display-name "GKE $APPLICATION Application Service Account"
export APP_SA_EMAIL=`gcloud iam service-accounts list --format='value(email)' --filter="displayName:$APPLICATION Application Service Account"`

# Bind service account policy
export PROJECT=`gcloud config get-value project`

gcloud projects add-iam-policy-binding $PROJECT --member=serviceAccount:${APP_SA_EMAIL} --role=roles/logging.logWriter
gcloud projects add-iam-policy-binding $PROJECT --member=serviceAccount:${APP_SA_EMAIL} --role=roles/monitoring.metricWriter

# Create service account key and activate it
gcloud iam service-accounts keys create \
    /home/$USER/key.json \
    --iam-account $APP_SA_EMAIL

kubectl create secret generic stackdriver-secret --from-file /home/$USER/key.json
```

REFS
https://github.com/census-instrumentation/opencensus-go
https://github.com/GoogleContainerTools/distroless
https://github.com/GoogleContainerTools/skaffold