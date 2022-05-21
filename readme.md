## AzureIPs

This is a simple tool that can fetch the JSON file, that Microsoft publishes, of the Azure's public IPs.  It then filters the JSON based on the Service Tag. The default Service Tag is 'AzureCloud'. You can provide the Service Tag to filter on. 

### Default
Service Tag: AzureCloud
Output: stdout

### Usage
Write the IPs for EventHub.NorthEurope to a file called 'eventhub.txt'

`azureips -f eventhub.txt -s EventHub.NorthEurope`

Write the IPs for AzureCosmosDB.NorthCentralUS to the terminal (stdout) 

`azureips -s AzureCosmosDB.NorthCentralUS`