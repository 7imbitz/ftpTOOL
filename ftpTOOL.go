package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"

	"github.com/projectdiscovery/gologger"

	"github.com/logrusorgru/aurora"
)

//constant variable version & author
const (
	Version = `1.0.0`
	Author  = `7imbitz`
)

//output e.g [CMD]
var (
	msf = aurora.Magenta("CMD")
	res = aurora.Green("OUT")
)

// ftpTOOL logo
var banner = fmt.Sprintf(`
     ______      __________  ____  __ 
    / __/ /_____/_  __/ __ \/ __ \/ / 
   / /_/ __/ __ \/ / / / / / / / / /  
  / __/ /_/ /_/ / / / /_/ / /_/ / /___
 /_/  \__/ .___/_/  \____/\____/_____/
        /_/                           
                                      %s - %s
`, Version, Author)

//showBanner is used to show banner for user
func showBanner() {
	gologger.Print().Msgf("%s", banner)
	gologger.Info().Msgf("ftpTOOL version %s", Version)
	gologger.Info().Label("INF").Msgf("This is a beta version")
	fmt.Println("Consoleless mfs enumeration ftp")
}

//setting up graceful exit
func gracefulExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			gologger.Fatal().Msgf("CTRL+C pressed: Exiting\n")
			os.Exit(1)
		}
	}()
}

//check ftp
func checkFtp(rhost string) {
	// Define the command and its arguments
	cmd := exec.Command("nc", "-nzvw", "5", rhost, "21")

	// Run the command
	err := cmd.Run()

	// Analyze the exit status to determine if the port is open, closed, or timed out
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok && exitErr.ExitCode() == 1 {
			gologger.Info().Label("INF").Msgf("Port is closed")
			gologger.Error().Msgf("No further enumeration done on %s\n", rhost)
			os.Exit(0)
		}
		return // Add a return statement here to exit the function in case of an error
	}

	gologger.Print().Label(res.String()).Msg("Port is open")
}

//msfconsole anon command
func msfCommandAnonymous(rhost string) {
	fmt.Println()
	gologger.Print().Label(res.String()).Msg("Anonymous Account")
	cmd := exec.Command("msfconsole", "-q", "-x", fmt.Sprintf("use auxiliary/scanner/ftp/anonymous; set RHOSTS %s; set RPORT 21; run; exit", rhost))

	// Set the command's stdout and stderr to use the current process's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		gologger.Error().Msgf("Error executing msfconsole: %v\n", err)
	}

}

//msfconsole check version command
func msfCommandVersion(rhost string) {
	fmt.Println()
	gologger.Print().Label(res.String()).Msg("FTP Version")
	cmd := exec.Command("msfconsole", "-q", "-x", fmt.Sprintf("use auxiliary/scanner/ftp/ftp_version; set RHOSTS %s; set RPORT 21; run; exit", rhost))

	// Set the command's stdout and stderr to use the current process's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		gologger.Error().Msgf("Error executing msfconsole: %v\n", err)
	}

}

//msfconsole bison ftp traverse
func msfCommandBison(rhost string) {
	fmt.Println()
	gologger.Print().Label(res.String()).Msg("Bison FTP Traversal")
	cmd := exec.Command("msfconsole", "-q", "-x", fmt.Sprintf("use auxiliary/scanner/ftp/bison_ftp_traversal; set RHOSTS %s; set RPORT 21; run; exit", rhost))

	// Set the command's stdout and stderr to use the current process's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		gologger.Error().Msgf("Error executing msfconsole: %v\n", err)
	}

}

//msfconsole colorado ftp traverse
func msfCommandColorado(rhost string) {
	fmt.Println()
	gologger.Print().Label(res.String()).Msg("Colorado FTP Traversal")
	cmd := exec.Command("msfconsole", "-q", "-x", fmt.Sprintf("use auxiliary/scanner/ftp/colorado_ftp_traversal; set RHOSTS %s; set RPORT 21; run; exit", rhost))

	// Set the command's stdout and stderr to use the current process's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		gologger.Error().Msgf("Error executing msfconsole: %v\n", err)
	}

}

//msfconsole titan ftp xcrc traverse
func msfCommandTitan(rhost string) {
	fmt.Println()
	gologger.Print().Label(res.String()).Msg("Titan FTP xcrc Traversal")
	cmd := exec.Command("msfconsole", "-q", "-x", fmt.Sprintf("use auxiliary/scanner/ftp/titanftp_xcrc_traversal; set RHOSTS %s; set RPORT 21; run; exit", rhost))

	// Set the command's stdout and stderr to use the current process's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		gologger.Error().Msgf("Error executing msfconsole: %v\n", err)
	}

}

//main func
func main() {
	gracefulExit()
	showBanner()
	// Check if the RHOSTS argument is provided
	if len(os.Args) != 2 {
		gologger.Fatal().Msgf("Usage: ftpTOOL <RHOSTS>")
		os.Exit(0)
	}
	//taking rhost from client
	rhost := os.Args[1]

	gologger.Print().Label(msf.String()).Msg("Checking services...")
	checkFtp(rhost)
	gologger.Print().Label(msf.String()).Msg("Running enumeration...")
	msfCommandAnonymous(rhost)
	msfCommandVersion(rhost)
	msfCommandBison(rhost)
	msfCommandColorado(rhost)
	msfCommandTitan(rhost)
	fmt.Println()
	gologger.Print().Label(res.String()).Msg("Happy hacking!")
}
