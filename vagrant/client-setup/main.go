package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	arch := runtime.GOARCH
	os := runtime.GOOS
	fmt.Printf("OS_TYPE: %s\nSYS_ARCH: %s\n", os, arch)

	out, err := exec.Command("cfssl", "version").Output()
	if err != nil {
		cfsslURL := "https://pkg.cfssl.org/R1.2/cfssl_" + os + "-" + arch
		if os == "windows" {
			cfsslURL = "https://pkg.cfssl.org/R1.2/cfssl_" + os + "-" + arch + ".exe"
		}
		fmt.Printf("%s\nDownload and install cfssl\n", err)
		downloadAsFile("/tmp/cfssl", cfsslURL)
		makeFileExecutable("/tmp/cfssl", "cfssl")
	}
	fmt.Println(string(out))

	out, err = exec.Command("cfssljson", "version").Output()
	if err != nil {
		cfsslJSONURL := "https://pkg.cfssl.org/R1.2/cfssljson_" + os + "-" + arch
		if os == "windows" {
			cfsslJSONURL = "https://pkg.cfssl.org/R1.2/cfssljson_" + os + "-" + arch + ".exe"
		}
		fmt.Printf("%s\nDownload and install cfssljson\n", err)
		downloadAsFile("/tmp/cfssljson", cfsslJSONURL)
		makeFileExecutable("/tmp/cfssljson", "cfssljson")
	}
	fmt.Println(string(out))

	out, err = exec.Command("kubectl", "version", "--short").Output()
	if err != nil {
		kubectlVersion := "1.18.0"
		kubectlURL := "https://storage.googleapis.com/kubernetes-release/release/v" + kubectlVersion + "/bin/" + os + "/" + arch + "/kubectl"
		fmt.Printf("%s\nDownload and install kubectl\n", err)
		downloadAsFile("/tmp/kubectl", kubectlURL)
		makeFileExecutable("/tmp/kubectl", "kubectl")
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
	_, err = exec.Command("sudo", "mv", filepath, "/usr/local/bin/"+fileName).Output()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
