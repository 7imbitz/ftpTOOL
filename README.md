# ftpTOOL

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/7imbitz/ftpTOOL/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://goreportcard.com/badge/github.com/7imbitz/ftpTOOL"><img src="https://goreportcard.com/badge/github.com/7imbitz/ftpTOOL"></a>

</p>

A Simple FTP Enumeration Tool written in golang 

<h1 align="center">

  <img alt="ftpTOOL" src="https://github.com/7imbitz/ftpTOOL/assets/26263598/be835f8d-a68e-4661-afa5-cf86013b88c5" height="650">
  <br>
</h1>


### Installation Instructions

ftpTOOL requires **go1.18** to install successfully. Run the following command to get the repo - 

```bash
go install -v github.com/7imbitz/ftpTOOL/ftpTOOL@latest
```
- TODO
_seems error_

### Workaround Installation

```bash
git clone https://github.com/7imbitz/ftpTOOL.git
cd ftpTOOL
go build ftpTOOL.go
./ftpTOOL
```
_You can move the binary in the `$GOPATH` directory_
```bash
mv ftpTOOL $GOPATH/bin
```

### Current Capability
- Enumeration of msfconsole for 
    - Anonymous account
    - FTP Version
    - Bison FTP Traverse
    - Colorado FTP Traverse
    - Titan FTP xcrc Traverse
- Tested on Kali Linux with go version go1.19.3 linux/amd64
    
 # Notes

- ftpTOOL is still ongoing development (during my free time) , improvement and updates will be done in near future
- ftpTOOL was inspired by [hacktricks](https://book.hacktricks.xyz/network-services-pentesting/pentesting-ftp#hacktricks-automatic-commands) 
- It is currently considered "automate" as user just need to provide IP address and the tool does the work
- Next update will implement modular where user can choose to only dork a/several things rather than selecting all

-----

ftpTOOL is made to ease my work and also to study golang. Community contributions are welcomed. Any issues occur can be [reported](https://github.com/7imbitz/ftpTOOL/issues) 
