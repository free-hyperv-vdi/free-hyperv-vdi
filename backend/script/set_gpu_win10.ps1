param(
    [Parameter(Mandatory=$true)]
    [string]$VmName
)

Add-VMGpuPartitionAdapter -VMName $VmName
Set-VM -GuestControlledCacheTypes $true -VMName $VmName 
Set-VMGpuPartitionAdapter -VMName $VmName -MinPartitionVRAM 80000000 -MaxPartitionVRAM 100000000 -OptimalPartitionVRAM 100000000 -MinPartitionEncode 80000000 -MaxPartitionEncode 100000000 -OptimalPartitionEncode 100000000 -MinPartitionDecode 80000000 -MaxPartitionDecode 100000000 -OptimalPartitionDecode 100000000 -MinPartitionCompute 80000000 -MaxPartitionCompute 100000000 -OptimalPartitionCompute 100000000 
Set-VM -LowMemoryMappedIoSpace 1GB -VMName $VmName 
Set-VM -HighMemoryMappedIoSpace 8GB -VMName $VmName 
