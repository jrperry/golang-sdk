# golang-sdk
iland cloud golang SDK

# Example Usage

client := sdk.NewClient(Username, Password, ClientID, ClientSecret)
virtualMachines := client.GetVirtualMachines()
