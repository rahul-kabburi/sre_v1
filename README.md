# sre_v1  
## System Setup  
1. Clone the repo:  
   ```bash  
   git clone https://github.com/rahul-kabburi/sre_v1.git  
   cd sre_v1  
   ```  

2. Install dependencies (if any):  
   ```bash  
   <your-installation-commands-here>  
   ```  

3. Start the system:  
   ```bash  
   <your-startup-command-here>  
   ```  

## Issue Faced  
- **Error:** ConfigMap 'alertmanager-config' not found  
- **Solution:** Create the ConfigMap with:  
  ```bash  
  kubectl create configmap alertmanager-config --from-file=alertmanager.yml -n monitoring  
  ```  

