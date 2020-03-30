package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	arch := runtime.GOARCH
	platform := runtime.GOOS

	usrHOMEDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command("cfssl", "version").Output()
	if err != nil {
		cfsslURL := "https://pkg.cfssl.org/R1.2/cfssl_" + platform + "-" + arch
		fmt.Printf("Download and install cfssl:-\n%s\n", cfsslURL)
		downloadAsFile(usrHOMEDir+"/cfssl", cfsslURL)
		makeFileExecutable(usrHOMEDir+"/cfssl", "cfssl")
	}
	fmt.Println(string(out))

	out, err = exec.Command("cfssljson", "version").Output()
	if err != nil {
		cfsslJSONURL := "https://pkg.cfssl.org/R1.2/cfssljson_" + platform + "-" + arch
		fmt.Printf("Download and install cfssl:-\n%s\n", cfsslJSONURL)
		downloadAsFile(usrHOMEDir+"/cfssljson", cfsslJSONURL)
		makeFileExecutable(usrHOMEDir+"/cfssljson", "cfssljson")
	}
	fmt.Println(string(out))

	out, err = exec.Command("kubectl", "version", "--client").Output()
	if err != nil {
		kubectlVersion := "1.18.0"
		kubectlURL := "https://storage.googleapis.com/kubernetes-release/release/v" + kubectlVersion + "/bin/" + platform + "/" + arch + "/kubectl"
		fmt.Printf("Download and install cfssl:-\n%s\n", kubectlURL)
		downloadAsFile(usrHOMEDir+"/kubectl", kubectlURL)
		makeFileExecutable(usrHOMEDir+"/kubectl", "kubectl")
	}
	fmt.Println(string(out))
}

func downloadAsFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func makeFileExecutable(filepath, fileName string) (err error) {
	err = os.Chmod(filepath, 0111)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}

	installDir := "/usr/local/bin/" + fileName
	if runtime.GOOS == "darwin" {
		installDir = "/usr/local/Cellar/" + fileName
	}
	_, err = exec.Command("sudo", "mv", filepath, installDir).Output()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
