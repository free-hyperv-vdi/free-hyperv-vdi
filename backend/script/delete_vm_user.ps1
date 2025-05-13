param(
    [Parameter(Mandatory=$true)]
    [string]$UserName,
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
    param($UserName)
    
	Remove-LocalUser -Name $UserName

} -ArgumentList $UserName