# 源文件夹路径
$sourceFolder = Get-Location
# 后端目标文件夹路径
$destBackend = "C:\cloud_server"
$destFront = "C:\nginx"
$destDeploy = "C:\nssm"

# 检查目标文件夹是否存在，如果不存在则创建
if (!(Test-Path -path $destBackend)) { 
    New-Item -ItemType directory -path $destBackend
}

if (!(Test-Path -path $destFront)) { 
    New-Item -ItemType directory -path $destFront
}

if (!(Test-Path -path $destDeploy)) { 
    New-Item -ItemType directory -path $destDeploy
    Copy-Item -Path $sourceFolder\install_tool\* -Destination $destDeploy -Recurse -Force
    # 获取当前系统环境变量Path
    $currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")

    # 检查C:\nssm是否已经在Path中，如果不在则添加
    #if ($currentPath -notlike "*$destDeploy*") {
    #    # 添加C:\nssm到系统环境变量Path
    #    [Environment]::SetEnvironmentVariable("Path", $currentPath + ";" + $destDeploy, "Machine")
    #}
    # 安装cloud_server_main
    $arg1 = "install", "cloud_server_main", ($destBackend + "\cloud_server_main.exe") -join " "
    Start-Process -NoNewWindow -FilePath $destDeploy\nssm.exe -ArgumentList $arg1
    # 安装nginx
    $arg2 = "install", "nginx", ($destFront + "\nginx.exe"), "-c", ($destFront + "\conf\nginx.conf") -join " "
    Start-Process -NoNewWindow -FilePath $destDeploy\nssm.exe -ArgumentList $arg2
}

# 停止服务
#Start-Process -NoNewWindow -FilePath $destDeploy\nssm.exe -ArgumentList "stop nginx"
#Start-Process -NoNewWindow -FilePath $destDeploy\nssm.exe -ArgumentList "stop cloud_server_main"

# 拷贝文件夹及其包含的文件
Copy-Item -Path $sourceFolder\cloud_server\* -Destination $destBackend -Recurse -Force
Copy-Item -Path $sourceFolder\nginx\* -Destination $destFront -Recurse -Force

# 启动nginx
Start-Process -NoNewWindow -FilePath $destDeploy\nssm.exe -ArgumentList "start nginx"
Start-Process -NoNewWindow -FilePath $destDeploy\nssm.exe -ArgumentList "start cloud_server_main"



