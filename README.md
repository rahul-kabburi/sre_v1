# sre_v1
## Steps to Set Up the System
1. Deploy the application using Kubernetes.
2. Ensure that Alertmanager is running.
3. Verify the config map using:
   ```
   kubectl get cm alertmanager-config -n monitoring -o yaml
   ```
4. If the config map is missing, create it manually.

## Issue Faced
After deploying, running the following command shows that the config map is missing:
   ```
   kubectl get cm alertmanager-config -n monitoring -o yaml
   Error from server (NotFound): configmaps "alertmanager-config" not found
   ```
