package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

type tools struct {
	cfssl     string
	cfssljson string
	kubectl   string
}

func main() {
	arch := runtime.GOARCH
	platform := runtime.GOOS

	usrHOMEDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	cfsslURL := "https://pkg.cfssl.org/R1.2/cfssl_" + platform + "-" + arch
	cfsslJSONURL := "https://pkg.cfssl.org/R1.2/cfssljson_" + platform + "-" + arch
	kubectlVersion := "1.18.0"
	kubectlURL := "https://storage.googleapis.com/kubernetes-release/release/v" + kubectlVersion + "/bin/" + platform + "/" + arch + "/kubectl"

	t := tools{
		cfsslURL,
		cfsslJSONURL,
		kubectlURL,
	}

	var wg sync.WaitGroup

	out, err := exec.Command("cfssl", "version").Output()
	if err != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Download and install cfssl:-\n%s\n", t.cfssl)
			downloadAsFile(usrHOMEDir+"/cfssl", t.cfssl)
			makeFileExecutable(usrHOMEDir+"/cfssl", "cfssl")
		}()
	}
	fmt.Println(string(out))

	out, err = exec.Command("cfssljson", "version").Output()
	if err != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Download and install cfssljson:-\n%s\n", t.cfssljson)
			downloadAsFile(usrHOMEDir+"/cfssljson", t.cfssljson)
			makeFileExecutable(usrHOMEDir+"/cfssljson", "cfssljson")
		}()
	}
	fmt.Println(string(out))

	out, err = exec.Command("kubectl", "version", "--client", "--short").Output()
	if err != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Download and install kubectl:-\n%s\n", t.kubectl)
			downloadAsFile(usrHOMEDir+"/kubectl", t.kubectl)
			makeFileExecutable(usrHOMEDir+"/kubectl", "kubectl")
		}()
	}
	wg.Wait()
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
	_, err = exec.Command("sudo", "mv", filepath, installDir).Output()
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	return nil
}
