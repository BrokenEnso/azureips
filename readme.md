## AzureIPs

This is a simple tool that can fetch the JSON file, that Microsoft publishes, of the Azure's public IPs.  It then filters the JSON based on the Service Tag. The default Service Tag is 'AzureCloud'. You can provide the Service Tag to filter on. 

### Default
Service Tag: AzureCloud
Output: stdout

> [!CAUTION]
> Each run of this file makes a request to a Microsoft owned/operated web server. This means Microsoft [Terms of Use](https://www.microsoft.com/en-us/legal/terms-of-use) would apply when using this tool. Do not run this in a way that that would be abusive to their systems or services.  

### Usage
Write the IPs for EventHub.NorthEurope to a file called 'eventhub.txt'

`azureips -f eventhub.txt -s EventHub.NorthEurope`

Write the IPs for AzureCosmosDB.NorthCentralUS to the terminal (stdout) 

`azureips -s AzureCosmosDB.NorthCentralUS`