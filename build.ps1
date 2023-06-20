$name = "envchecker"
$version = "0.0.1"
$buildTime = Get-Date -Format 'yyyy-MM-dd HH:mm:ss'
$ldFlags = "-s -w -X '$Name/version.version=$version' -X '$Name/version.buildTime=$buildTime'"

$arrOS = "linux", "windows", "android"
$arrARCH = "386", "amd64", "arm", "arm64", "riscv", "riscv64"
Write-Host "=========================================="
Write-Host "   [cmd]:   build.ps1 <os> <arch>"
Write-Host "   [os]:   ", $( $arrOS -join "," )
Write-Host "   [arch]: ", $( $arrARCH -join "," )
Write-Host "=========================================="

$os = $args[0]

if (!$arrOS.Contains($os)) {
    Write-Host "Error: Invalid GOOS" -ForegroundColor:Red;
    exit(-1)
}

$arch = $args[1]
if (!$arrARCH.Contains("$arch")) {
    Write-Host $arch, $args[1]
    Write-Host $arrARCH
    Write-Host "Error: Invalid GOARCH" -ForegroundColor:Red;
    exit(-1)
}

Write-Output("GOOS $( go env GOOS ) => $os")
$env:GOOS = $os
Write-Output("GOARCH $( go env GOARCH ) => $arch")
$env:GOARCH = $arch


$output = ".\app\$os\$name-$os-$arch"
if ($os -eq "linux") {
    go build -ldflags $ldFlags -o $output .\main.go
}
if ($os -eq "windows") {
    $exePath = "$output.exe"
    go build  -ldflags $ldFlags -o $exePath .\main.go
}

Write-Host "build had finished, output:$output" -ForegroundColor:Green

Invoke-Item ".\app\$os\"