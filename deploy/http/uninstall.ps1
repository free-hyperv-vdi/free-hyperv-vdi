# 后端目标文件夹路径
$destBackend = "C:\cloud_server"
$destFront = "C:\nginx"
$destDeploy = "C:\nssm"

# 停止服务函数
function Stop-ServiceIfExists {
    param(
        [string]$serviceName
    )

    $service = Get-Service -Name $serviceName -ErrorAction SilentlyContinue
    if ($service -ne $null) {
        Stop-Service $serviceName
    }
}

if (Test-Path -path $destDeploy) { 
    # 移除cloud_server_main
	Stop-ServiceIfExists -serviceName "cloud_server_main"
    $arg1 = "remove", "cloud_server_main" -join " "
    Start-Process -NoNewWindow -FilePath $destDeploy\nssm.exe -ArgumentList $arg1
    # 移除nginx
	Stop-ServiceIfExists -serviceName "nginx"
    $arg2 = "remove", "nginx" -join " "
    Start-Process -NoNewWindow -FilePath $destDeploy\nssm.exe -ArgumentList $arg2
	
	Remove-Item -Path $destBackend -Recurse -Force
	Start-Sleep -Seconds 3
	Remove-Item -Path $destFront -Recurse -Force
	Start-Sleep -Seconds 3
	Remove-Item -Path $destDeploy -Recurse -Force
	
	echo "uninstall successfully"
} else {
	echo "uninstall failed."
}





