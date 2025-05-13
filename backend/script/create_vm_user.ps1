param(
    [Parameter(Mandatory=$true)]
    [string]$UserName,
    [Parameter(Mandatory=$true)]
    [string]$Password,
    [Parameter(Mandatory=$true)]
    [string]$Name,
    [Parameter(Mandatory=$true)]
    [string]$adminUser,
    [Parameter(Mandatory=$true)]
    [string]$adminPwd
)

# 使用安全凭据
$SecurePassword = ConvertTo-SecureString -String $adminPwd -AsPlainText -Force
$Credentials = New-Object System.Management.Automation.PSCredential ($adminUser, $SecurePassword)

# 在远程会话中执行命令
$Session = New-PSSession -VMName $Name -Credential $Credentials
Invoke-Command -Session $Session -ScriptBlock {
    param($UserName, $Password)
    
    # 创建本地用户
    $LocalPassword = ConvertTo-SecureString -String $Password -AsPlainText -Force
    New-LocalUser -Name $UserName -Password $LocalPassword
    
    # 将用户添加到管理员组
    Add-LocalGroupMember -Group "Administrators" -Member $UserName
} -ArgumentList $UserName, $Password