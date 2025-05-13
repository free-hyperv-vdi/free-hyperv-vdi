param(
    [Parameter(Mandatory=$true)]
    [string]$VmName,
    [Parameter(Mandatory=$true)]
    [string]$GpuPath
)

try {
    Add-VMGpuPartitionAdapter -VMName $VmName -InstancePath $GpuPath
    Set-VM -GuestControlledCacheTypes $true -VMName $VmName 
    Set-VMGpuPartitionAdapter -VMName $VmName -MinPartitionVRAM 80000000 -MaxPartitionVRAM 100000000 -OptimalPartitionVRAM 100000000 -MinPartitionEncode 80000000 -MaxPartitionEncode 100000000 -OptimalPartitionEncode 100000000 -MinPartitionDecode 80000000 -MaxPartitionDecode 100000000 -OptimalPartitionDecode 100000000 -MinPartitionCompute 80000000 -MaxPartitionCompute 100000000 -OptimalPartitionCompute 100000000 
    Set-VM -LowMemoryMappedIoSpace 1GB -VMName $VmName 
    Set-VM -HighMemoryMappedIoSpace 8GB -VMName $VmName 
}
catch {
    # 捕获异常并输出错误信息到标准错误流
    Write-Error "Error occurred: $_"
    # 抛出异常中止执行
    Throw "Execution stopped due to an error"
}

